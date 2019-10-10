package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"ntci/ci-agent/dataBus"
	"ntci/ci-agent/git"
)

/**
GitLabService handler gitlab request.

url is gitlab repository url.
id is this repository id.
*/
type Service struct {
	url string
	id  int
}

func (s *Service) FetchNtCI() (n git.Ntci, err error) {
	queryURL := fmt.Sprintf("%s/api/v4/projects/%d/repository/blobs/.ntci.yml/raw", s.url, s.id)
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

func (s *Service) GitCallBack(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Read Git Lab Request Error. %s ", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Debugf("gitlab request data. %s ", string(data))

	var push pushEvent

	err = json.Unmarshal(data, &push)
	if err != nil {
		logrus.Errorf("Unmarshal Git Lab Push Event Error. %s ", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if push.ObjectKind != "push" && push.EventName != "push" {
		logrus.Errorf("Wrong Event. Kind: %s Event: %s ", push.ObjectKind, push.EventName)
		//w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	gitService := new(Service)
	gitService.url = push.Project.WebURL
	gitService.id = push.ProjectID
	n, err := git.ParseProject(gitService)
	if err != nil {
		logrus.Errorf("Parse .ntci.yml Error. %s ", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	logrus.Debugf("ntct.yml: %v", n)
}
