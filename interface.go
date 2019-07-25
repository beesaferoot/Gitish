// Entry Point to package Gitish
package main
import (
	"Gitish/issues"
	"flag"
	"fmt"
	"log"
	_ "log"
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
var update = flag.String("edit", "", "Edit Issue")
var create = flag.String("create", "" , "Create New Issue in repo")
var issbody = flag.String("b", "","Issue body to be used when creating body")
var list = flag.Bool("view", false,"List all Issues in repo")
var editor = flag.String("editor", "", "specify editor to input issues with")
var editorShort =  flag.String("e", "", "shorthand for --editor argument")
//help flag
var helpStr = `
Gitish utility tool 

commands:
 -repo | --repo [repository name] -- specifies repository name (required)
 -user | --user | -u | --u [login name] -- specifies github login name (required)
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
	if *loginShort != "" && *login == ""{
		*login = *loginShort
	}

	if *editorShort != "" && *editor == ""{
		*editor = *editorShort
	}

	if *helpShort && !*help{
		*help = *helpShort
	}

	// Check of require fields 
	if *repo == "" || *login == "" {
		fmt.Println("required fields missing!!")
		fmt.Print(helpStr)
		os.Exit(1)
	}
	issue.User = new(issues.User)
	issue.User.Login = *login
	issue.Repo = *repo

	switch {
		//Check for view case
		case *list != false: {
			res , err := issues.ViewIssue(issue.Repo, issue.User.Login)
			 if err != nil{
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Total Issues: ", res.TotalCount)
			for index, items := range res.Items{
				fmt.Printf("(%d) %v",index, items.Body)
			}

			os.Exit(1)
		}
		case *update != "":{
			// if editor is specified : use preferred editor
			if ok, err := runedit(editor, &issue); !ok{
				log.Println(err, "Using text from command line!!")
			}
			if *editor == ""{
				issue.Body = *update
			}
			if ok, err := issues.UpdateIssue(&issue); !ok{
				panic(err)
			}
			fmt.Printf("Issue in repo %s updated successfully", issue.Repo)
		}

		case *create != "": {
			// if editor is specified : use preferred editor
			if ok, err := runedit(editor, &issue); !ok{
				log.Println(err, "Using text input from command line!!")
			}
			// set issue title to create
			issue.Title = *create

			// if editor not specified? use issue body in arguments
			if *editor == ""{
				issue.Body = *issbody
			}

			if ok, err := issues.CreateIssue(&issue); !ok{
				panic(err)
			}
			fmt.Printf("Issue in repo %s created successfully", issue.Repo)
		}

		// default action
		default:{
			// output help message
			fmt.Print(helpStr)
		}
	}
 

}

func  runedit(editor *string, issue *issues.Issue) (bool, error) {
	
	if *editor != ""{
		//create temporary file to write issue
		tempfile, err := os.Create("tempfile.txt")
		if err != nil{
			//return  error
			return false, fmt.Errorf("error occured: %v", err)
		}
		tempfile.Close()
		cmd := exec.Command(*editor, tempfile.Name())
		body := make([]byte, 100)
		if err := cmd.Run(); err != nil{
			return false, fmt.Errorf("error occured: %v", err)
		}
		length, err := tempfile.Read(body)
		if length > 0{
			issue.Body = string(body[:length])
		}
		return true, nil
	}else{
		return false, fmt.Errorf("no text editor specified")
	}


}