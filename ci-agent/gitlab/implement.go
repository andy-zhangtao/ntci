package gitlab

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"ntci/ci-agent/dataBus"
	"ntci/ci-agent/git"
	build_rpc_v1 "ntci/ci-grpc/build"
)

func (s *Service) FetchNtCI() (n git.Ntci, err error) {
	queryURL := fmt.Sprintf("%sapi/v4/projects/%d/repository/files/.ntci.yml/raw?ref=%s", s.url, s.id, s.branch)
	logrus.Debugf("Fetch .ntci.yml request: %s", queryURL)

	reqest, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		return
	}

	reqest.Header.Add("PRIVATE-TOKEN", dataBus.GetBus().Access.Gitlab.Token)
	logrus.Debugf("Fetch .ntci.yml Token: %s", dataBus.GetBus().Access.Gitlab.Token)
	client := &http.Client{}

	response, _ := client.Do(reqest)
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &n)
	if err != nil {
		return
	}

	return
}

func (s *Service) VerifyNtci(ntci git.Ntci) bool {
	bus := dataBus.GetBus()

	lanuage := ntci.Language
	tag := "latest"
	if strings.Contains(ntci.Language, ":") {
		lanuage = strings.Split(ntci.Language, ":")[0]
	}

	if _, ok := bus.LanguageRuntime[lanuage]; !ok {
		return false
	}

	l := bus.LanguageRuntime[lanuage]

	if _, ok := l[tag]; !ok {
		return false
	}

	return true
}

func (s *Service) InvokeBuildService(ntci git.Ntci) (err error) {

	bus := dataBus.GetBus()

	conn, err := grpc.Dial(bus.Build[bus.BuildMode].Addr, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := build_rpc_v1.NewBuildServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Run(ctx, &build_rpc_v1.Request{
		Name:   s.name,
		Id:     s.commit,
		Branch: s.branch,
		Url:    s.url,
	})

	if err != nil {
		logrus.Errorf("Invoke Build Service Error.  %v", err)
		return err
	}

	if r.Code != GRPC_SUCC {
		logrus.Errorf("Invoke Build Service Failed.  %d, %s", r.Code, r.Message)
		return errors.New("Invoke Build Service Failed ")
	}

	logrus.Infof("Invoke Build Service Success: %d", r.Code)
	return
}
