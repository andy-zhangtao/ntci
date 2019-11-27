package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"ntci/ci-agent/dataBus"
	"ntci/ci-agent/git"
	"ntci/ci-agent/store"
)

/**
GitLabService handler gitlab request.

url is gitlab repository url.
id is this repository id.
branch is trigger branch name.
*/
type Service struct {
	url    string
	webURL string
	// id : project id
	id int
	// jid : db idx
	jid        int
	branch     string
	name       string
	commit     string
	language   string
	lanversion string
	user       string
	sha        string
	message    string
	namespace  string
	buidScript string
	dockerfile string
}

func (s Service) GitCallBack(w http.ResponseWriter, r *http.Request) {
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

	commits := len(push.Commits)
	s.id = push.ProjectID
	s.branch = drawOffBranch(push)
	s.name = converName(strings.ToLower(push.Project.Name))
	s.commit = push.CheckoutSha
	s.webURL = push.Project.HTTPURL
	s.url = drawOffUrl(push)
	s.dockerfile = push.Dockerfile

	if commits == 0 {
		s.user = push.UserUsername
		s.message = ""
	} else {
		s.user = push.Commits[commits-1].Author.Email
		if s.user == "" {
			s.user = push.UserUsername
		}
		s.message = push.Commits[commits-1].Message
	}

	if push.CheckoutSha != "" {
		s.sha = push.CheckoutSha[:12]
	}

	s.namespace = push.Project.PathWithNamespace

	bus := dataBus.GetBus()

	id, err := bus.Pb.AddNewBuild(store.Build{
		Name:      s.name,
		Branch:    s.branch,
		Status:    store.BuildReady,
		Git:       s.url,
		Timestamp: time.Now().Local(),
		User:      s.user,
		Sha:       s.sha,
		Message:   s.message,
		Namespace: s.namespace,
	})
	if err != nil {
		logrus.Errorf("Add New Build Error. %s ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.jid = id

	//buildScript := ""
	if push.Build != "" {
		s.buidScript = push.Build
	}

	n, err := git.ParseAndExecuteBuild(s)
	logrus.Debugf("ntct.yml: %v", n)

	if err != nil {
		logrus.Errorf("Build Error. %s ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	nt, err := yaml.Marshal(n)
	//buildScript = string(nt)

	err = bus.Pb.AddNtci(s.user, s.name, s.branch, string(nt))
	if err != nil {
		logrus.Errorf("Save Configure Error. %s ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("%s-%d", s.name, s.jid)))
	return
}

// converName conver '_' to '-'
func converName(name string) string {
	idx := []string{
		"_",
	}

	for _, i := range idx {
		name = strings.Replace(name, i, "-", -1)
	}

	return name
}

/*
drawOffUrl

Get gitlab url from web url.  Since web url format is: http://[domain/ip][:port]/[namespace]/name.

So use split web url, and return the first element.

*/
func drawOffUrl(p pushEvent) string {
	//end := ""
	//if p.Project.Namespace != "" {
	//	end = fmt.Sprintf("%s/%s", p.Project.Namespace, p.Project.Name)
	//} else {
	//	end = fmt.Sprintf("%s", p.Project.Name)
	//}

	s := strings.Split(p.Project.WebURL, p.Project.PathWithNamespace)

	return s[0]
}

func drawOffBranch(p pushEvent) string {
	branch := "master"

	if p.Ref == "refs/heads/master" {
		return branch
	}

	if strings.HasPrefix(p.Ref, "refs/heads/") {
		branch = strings.Split(p.Ref, "refs/heads/")[1]
	}

	return branch
}
