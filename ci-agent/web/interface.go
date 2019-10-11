package web

import "net/http"

type gitRequestCallBack interface {
	/**
	Git Repository Http Call Back Function. This Function Can Receive Git Pull Request.
	Since git hub and git lab send different body, so this function should extract a interface.
	*/
	GitCallBack(http.ResponseWriter, *http.Request)
}
