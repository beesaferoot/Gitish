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

//TODO: fix bugs
func CreateIssue(issue *Issue) (bool, error) {

	queryFields := issue.User.Login + "/" + issue.Repo + "/issues" // :owner/repo/issue
	URL := cIssueURL + "/" + queryFields
	fmt.Println("URL - ", URL)

	jsonObj, err := json.Marshal(*issue)
	if err != nil {
		return false, fmt.Errorf("Error Occured: %v", err)
	}

	// Create New Client
	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonObj))
	if err != nil {
		// Error occured
		return false, fmt.Errorf("Error Occured: %v", err)
	}

	req.SetBasicAuth(issue.User.Login, issue.User.Password)
	req.Header.Set("Content-Type", "application/json; charset: utf-8")
	fmt.Println(req.Body)

	response, err := client.Do(req)

	if err != nil {
		// Error occured
		return false, fmt.Errorf("Error Occured: %v", err)
	}

	defer response.Body.Close() // Make sure to close response Object

	// Check of Status Code is set to  Created
	if response.StatusCode != http.StatusCreated {
		return false, fmt.Errorf("Post issue request failed with Status: %s", response.Status)
	}

	return true, nil
}
