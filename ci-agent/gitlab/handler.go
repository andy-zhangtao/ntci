package gitlab

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

/**
GitLabService handler gitlab request.
*/
type Service struct{}

func (s *Service) GitCallBack(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Read Git Lab Request Error. %s ", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.Debugf("gitlab request data. %s ", string(data))
}
