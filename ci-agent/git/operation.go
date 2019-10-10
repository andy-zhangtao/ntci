package git

/*
ParseProject

Parse Project via parse .ntci.yml file.
*/
func ParseProject(g GitOperation) (n Ntci, err error) {
	n, err = g.FetchNtCI()
	if err != nil {
		return
	}

	return
}
