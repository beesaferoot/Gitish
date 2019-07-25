package issues
import (
	"fmt"
	"encoding/json"
	"bytes"
	"net/http"
)
/*
	Update existing github Issue in specified access repository
*/


func UpdateIssue(issue *Issue) (bool, error) {

	jsonObj, err := json.Marshal(issue)
	if err != nil {
		return false, fmt.Errorf("Error Parsing issue struct to json:%v", err)
	}

	// PUT to issue repo
	queryField := cIssueURL + "/" + issue.User.Login + "/" + issue.Repo + "/issues"
	// Create Client
	Client := &http.Client{}
	req, err  := http.NewRequest("PATCH", queryField, bytes.NewBuffer(jsonObj))
	if err != nil{
		return false, err
	}

	response, err := Client.Do(req)
	if err != nil{
		return false, fmt.Errorf("Failed request to update issue:%v", err)// return Error when failure occurs during PUT request 
	}
	defer response.Body.Close() // Make sure to close response Object

	if response.StatusCode != http.StatusOK{
		// Failed status code
		return false, fmt.Errorf("Post issue request failed with Status: %s", response.Status)
	}
	
	return true, nil

}