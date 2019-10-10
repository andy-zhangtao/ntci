package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"ntci/ci-agent/dataBus"
	"ntci/ci-agent/git"
)

/**
GitLabService handler gitlab request.

url is gitlab repository url.
id is this repository id.
branch is trigger branch name.
*/
type Service struct {
	url    string
	id     int
	branch string
}

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
	gitService.branch = push.Ref
	gitService.url = drawOffUrl(push)

	n, err := git.ParseProject(gitService)
	if err != nil {
		logrus.Errorf("Parse .ntci.yml Error. %s ", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Debugf("ntct.yml: %v", n)
}

/*
drawOffUrl

Get gitlab url from web url.  Since web url format is: http://[domain/ip][:port]/[namespace]/name.

So use split web url, and return the first element.

*/
func drawOffUrl(p pushEvent) string {
	end := ""
	if p.Project.Namespace != "" {
		end = fmt.Sprintf("%s/%s", p.Project.Namespace, p.Project.Name)
	} else {
		end = fmt.Sprintf("%s", p.Project.Name)
	}

	s := strings.Split(p.Project.WebURL, end)

	return s[0]
}
