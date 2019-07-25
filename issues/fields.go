
/*
	Data Fields for parsing returned issues
*/

package issues
import "time"


// URL for creating issues
const (
	cIssueURL = "https://api.github.com/repos"
)

type IssueList struct{
	TotalCount int `json:"total_count"`
	Items []Issue
}


type Issue struct {
	Number int 
	Title string 
	State string 
	Label *Label
	User *User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	HTMLURL string `json:"html_url"`
	Repo string 
	Sort string 
	Body string 
}

type Label struct {
	Name string 
	Description string 
}
type User struct {
	Login string 
	HTMLURL string `json:"html_url"`
}
