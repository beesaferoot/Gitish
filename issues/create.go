
/*
	Create Issues 
*/
package issues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateIssue(issue *Issue) (bool, error){

	queryFields := issue.User.Login + "/" + issue.Repo + "/issues" // :owner/repo/issue
	URL := cIssueURL + "/" + queryFields
	fmt.Println("URL - ", URL)

	jsonObj, err := json.Marshal(*issue)
	if err != nil{
		return false, fmt.Errorf("Error Occured: %v", err)
	}
	response, err := http.Post(URL, "application/json; charset=utf8", bytes.NewBuffer(jsonObj))
	
	if err != nil {
		// Error occured 
		return false, fmt.Errorf("Error Occured: %v", err)
	}

	defer response.Body.Close() // Make sure to close response Object

	// Check of Status Code is Ok
	if response.StatusCode != http.StatusOK{
		return false, fmt.Errorf("Post issue request failed with Status: %s", response.Status)
	}

	return true, nil
}