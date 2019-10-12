package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"ntci/ci-agent/git"
)

/**
GitLabService handler gitlab request.

url is gitlab repository url.
id is this repository id.
branch is trigger branch name.
*/
type Service struct {
	url        string
	id         int
	branch     string
	name       string
	commit     string
	language   string
	lanversion string
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

	//gitService := new(Service)
	//gitService.url = push.Project.WebURL
	//gitService.id = push.ProjectID
	//gitService.branch = push.Ref
	//gitService.url = drawOffUrl(push)

	//s.url = push.Project.WebURL
	s.id = push.ProjectID
	s.branch = push.Ref
	s.name = push.Project.Name
	s.commit = push.CheckoutSha
	s.url = drawOffUrl(push)

	n, err := git.ParseAndExecuteBuild(s)
	if err != nil {
		logrus.Errorf("Build Error. %s ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
