package git

/*
GitOperation is common git operation functions.
*/
type GitOperation interface {
	// FetchNtCI Get and parse .ntci.yml from git repository
	// If .ntci.yml not exist or parse error, will return error.
	FetchNtCI() (Ntci, error)
}
