// Entry Point to package Gitish
package main

import (
	"Gitish/issues"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

/*
	Algo:
		- get required flags
		- update Issue struct
		- send request specified by flags
*/

// declare flags
var repo = flag.String("repo", "", "specify github repo")
var login = flag.String("user", "", "specify github login name")
var loginShort = flag.String("u", "", "short form for -user [name]| --user [name]")
var password = flag.String("p", "", " github password for API permission verification")
var update = flag.String("edit", "", "Edit Issue")
var create = flag.String("create", "", "Create New Issue in repo")
var issbody = flag.String("b", "", "Issue body to be used when creating body")
var list = flag.Bool("view", false, "List all Issues in repo")
var editor = flag.String("editor", "", "specify editor to input issues with")
var editorShort = flag.String("e", "", "shorthand for --editor argument")

//help flag
var helpStr = `
Gitish utility tool 

commands:
 -repo | --repo [repository name] -- specifies repository name (required)
 -user | --user | -u | --u [login name] -- specifies github login name (required)
 --p | -p [password] --  github password for API permission verification (required)
 -edit | --edit IssueString..  -- Edit issue  
 -create [title]| --create [Issue title String][-b Issue Body or use -e to specify preferred editor].. -- Create new issue
 -view | --view  -- List all issues in repository
-editor | --editor | -e | --e -- specify preferred editor to input Issue  

`
var help = flag.Bool("help", false, "display help message")
var helpShort = flag.Bool("h", false, "display help message")

// declare issue
var issue issues.Issue

func main() {

	flag.Parse()

	/*
		TODO: 1.Write Logic to Initialize Issue with provided fields
			  2. couple package issues functions
	*/
	if *loginShort != "" && *login == "" {
		*login = *loginShort
	}

	if *editorShort != "" && *editor == "" {
		*editor = *editorShort
	}

	if *helpShort && !*help {
		*help = *helpShort
	}

	if *help {
		fmt.Print(helpStr)
		os.Exit(1)
	}
	fmt.Println(*issbody)
	// Check for required fields
	if *repo == "" || *login == "" || *password == "" {
		fmt.Println("required fields missing!!")
		fmt.Print(helpStr)
		os.Exit(1)
	}
	issue.User = new(issues.User)
	issue.User.Login = *login
	issue.User.Password = *password
	issue.Repo = *repo

	switch {
	//Check for view case
	case *list != false:
		{
			res, err := issues.ViewIssue(issue.Repo, issue.User.Login)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Total Issues: ", res.TotalCount)
			for index, items := range res.Items {
				fmt.Printf("(%d) %v\n", index, items.Body)
			}

			os.Exit(1)
		}
	case *update != "":
		{
			// if editor is specified : use preferred editor
			if ok, err := runedit(editor, &issue); !ok {
				log.Println(err, "Using text from command line!!")
			}
			if *editor == "" {
				issue.Body = *update
			}
			if ok, err := issues.UpdateIssue(&issue); !ok {
				panic(err)
			}
			fmt.Printf("Issue in repo %s updated successfully\n", issue.Repo)
		}

	case *create != "":
		{
			// if editor is specified : use preferred editor
			if ok, err := runedit(editor, &issue); !ok {
				log.Println(err, "Using text input from command line!!")
			}
			// set issue title to create
			issue.Title = *create

			// if editor not specified? use issue body in arguments
			if *editor == "" {
				issue.Body = *issbody
			}

			if ok, err := issues.CreateIssue(&issue); !ok {
				panic(err)
			}
			fmt.Printf("Issue in repo %s created successfully\n", issue.Repo)
		}

	// default action
	default:
		{
			// output help message
			fmt.Print(`No operation specified. check out -h for help`)
			//fmt.Print(helpStr)
		}
	}

}

//TODO: Work on edit functionality!
func runedit(editor *string, issue *issues.Issue) (bool, error) {

	if *editor != "" {
		//create temporary file to write issue
		tempfile, err := os.Create("tempfile.txt")
		if err != nil {
			//return  error
			return false, fmt.Errorf("error occured: %v\n", err)
		}
		cmd := exec.Command(*editor, tempfile.Name())
		body := make([]byte, 300)
		// start command process
		if err := cmd.Start(); err != nil {
			return false, fmt.Errorf("error occured: %v\n", err)
		}
		// wait for command process to exit
		if err := cmd.Wait(); err != nil {
			return false, fmt.Errorf("error occured: %v\n", err)
		}

		keyinput := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter any key to continue...")
		keyinput.Scan()

		// read length of temporary file
		length, err := tempfile.Read(body)
		fmt.Println("length of tempfile read: ", length)
		if length > 0 {
			// assign issue body
			issue.Body = string(body[:length])
		}

		// delete tempfile
		if err := tempfile.Close(); err != nil {
			return false, fmt.Errorf("Error Occured: %v\n", err)
		}
		defer os.Remove(tempfile.Name()) // make sure to delete tempfile.txt after being used...
		return true, nil
	} else {
		return false, fmt.Errorf("no text editor specified")
	}

}
