package issues
/*
	View issues for specified access github repository 
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func ViewIssue(repo string, login string) (*IssueList, error) {

	// fetch issues from repo
	queryField := cIssueURL + "/" + login + "/" + repo + "/issues"
	response, err := http.Get(queryField)
	if err != nil {
		return nil, fmt.Errorf("Error Occured trying to fetch issues: %v", err)
	}

	defer response.Body.Close() // Make sure to close response Object

	if response.StatusCode != http.StatusOK{
		// Failed status code
		return nil, fmt.Errorf("Get issue request failed with Status: %s", response.Status)
	}

	var result IssueList
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil{
		return nil, err
	}

	return &result, nil
}