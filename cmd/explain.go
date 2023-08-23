/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/kunxl-gg/no-issues/constants"
	"github.com/kunxl-gg/no-issues/helpers"
	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {

		var (
			owner string = args[0]
			repo string = args[1]
			issueNumber string = args[2]
		)

		// get the issue
		issue := getIssue(owner, repo, issueNumber)

		// making a request to the chatGPT API
		ans := getAns(issueBody)


	},
}

func getAns(issueBody string){

	// making a url 
	url := constants.OPENAI_BASE_URL + "/v1/engines/davinci/completions"

	// writing the prompt
	prompt := "Q: " + issueBody + "\nA:"

	// making a request
	req, err := http.NewRequest("POST", url, nil)
	helpers.CheckError(err)

	// adding headers
	req.Header.Set("Authorization", "Bearer " + constants.OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")

	// making a client
	client := &http.Client{}
	res, err := client.Do(req)
	helpers.CheckError(err)

	// closing the response body
	defer res.Body.Close()

	// reading the response
	body, err := ioutil.ReadAll(res.Body)
	helpers.CheckError(err)

	// printing the response
	fmt.Println(string(body))
	

}

func getIssue(owner string, repo string, issueNumber string) string {

	// making a url
	url := constants.GITHUB_API_BASE_URL + "/repos/" + owner + "/" + repo + "/issues/" + issueNumber

	// making a response
	req, err := http.NewRequest("GET", url, nil)
	helpers.CheckError(err)

	// adding headers
	req.Header.Set("Authorization", "Bearer " + constants.GH_API_KEY)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	// making a client 
	client := &http.Client{}
	res, err := client.Do(req)
	helpers.CheckError(err)

	// closing the response body
	defer res.Body.Close()

	// reading the response
	body, err := ioutil.ReadAll(res.Body)
	helpers.CheckError(err)

	return string(body)

}

func init() {
	rootCmd.AddCommand(explainCmd)
}
