package gitlab

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"ntci/ci-agent/dataBus"
	"ntci/ci-agent/git"
	"ntci/ci-agent/rpc/builder"
	"ntci/ci-agent/store"
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

	logrus.Debugf(".ntci.yml content: %s", string(data))

	err = yaml.Unmarshal(data, &n)
	if err != nil {
		return
	}

	return
}

func (s *Service) VerifyNtci(ntci git.Ntci) bool {

	s.lanversion = "latest"
	if strings.Contains(ntci.Language, ":") {
		s.language = strings.Split(ntci.Language, ":")[0]
		s.lanversion = strings.Split(ntci.Language, ":")[1]
	} else {
		s.language = ntci.Language
	}

	return true
}

func (s *Service) InvokeBuildService(ntci git.Ntci) (err error) {

	bus := dataBus.GetBus()

	if s.language != "" {
		err = bus.Pb.UpdateBuildLanguage(s.language, s.lanversion, s.id, s.language, s.lanversion)
		if err != nil {
			return
		}
	}

	env, err := bus.Pb.GetCommonEnv()
	if err != nil {
		logrus.Error(err)
	}

	r, err := builder.InvokeBuilderServiceRun(&build_rpc_v1.Request{
		Name:       s.name,
		Id:         int32(s.jid),
		Branch:     s.branch,
		Url:        s.webURL,
		Language:   s.language,
		Lanversion: s.lanversion,
		User:       s.user,
		Sha:        s.sha,
		Message:    s.message,
		Env:        env,
	})

	if err != nil {
		bus.Pb.UpdataBuildStatus(int32(store.BuildFailed), s.jid, s.name, s.user)
		logrus.Errorf("Invoke Build Service Error.  %v", err)
		return err
	}

	if r.Code != GRPC_SUCC {
		bus.Pb.UpdataBuildStatus(int32(store.BuildFailed), s.jid, s.name, s.user)
		logrus.Errorf("Invoke Build Service Failed.  %d, %s", r.Code, r.Message)
		return errors.New("Invoke Build Service Failed ")
	}

	bus.Pb.UpdataBuildStatus(int32(store.BuildEnv), s.jid, s.name, s.user)
	logrus.Infof("Invoke Build Service Success: %d", r.Code)
	return
}
