package git

import "errors"

/*
ParseProject

Parse Project via parse .ntci.yml file.
*/
func ParseProject(g GitOperation) (n Ntci, err error) {
	n, err = g.FetchNtCI()
	if err != nil {
		return
	}

	if !g.VerifyNtci(n) {
		err = errors.New("Invalid ntci configure ")
		return
	}

	return
}
