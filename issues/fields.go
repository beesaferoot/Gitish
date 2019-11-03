/*
	Data Fields for parsing returned issues
*/

package issues

import "time"

// URL for creating issues
const (
	cIssueURL = "https://api.github.com/repos"
)

type IssueList struct {
	TotalCount int `json:"total_count"`
	Items      []Issue
}

type Issue struct {
	Number    int
	Title     string  `json:"title"`
	State     string  `json:"state"`
	Labels    []Label `json:"labels"`
	User      *User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	HTMLURL   string    `json:"html_url"`
	Repo      string    `json:"repo"`
	Assignees []string  `json:"assignees"`
	Sort      string
	Body      string `json:"body"`
}

type Label struct {
	Name        string
	Description string
}

type User struct {
	Login    string
	Password string
	HTMLURL  string `json:"html_url"`
}
