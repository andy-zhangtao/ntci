package message

type DeployMsg struct {
	// Project Name
	Name string `json:"name"`
	// Build Branch
	Branch string `json:"branch"`
	// Git Project URL
	Git string `json:"git"`
	// Commit author
	Owner string `json:"owner"`
	// Docker Image. Options
	Img string `json:"img"`
}
