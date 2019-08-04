// Package backlog provides backlog api
package backlog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"

	"github.com/mkohei/my-backlog-api-tool/config"
)

// IssuesURL is issues endpoint
const IssuesURL = "/api/v2/issues"

// Issue show response issue
type Issue struct {
	ID            int      `json:"id"`
	ProjectID     int      `json:"projectId"`
	IssueKey      string   `json:"issueKey"`
	Summary       string   `json:"summary"`
	Description   string   `json:"description"`
	Status        Status   `json:"status"`
	Assignee      Assignee `json:"assignee"`
	ParentIssueID int      `json:"parentIssueId"`
}

// Status show response status
type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Assignee show response asignee
type Assignee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MailAddress string `json:"mailAddress"`
}

// GetIssue implements to get issue by IssueKey
func GetIssue(conf config.Config, issueKey string) (Issue, error) {
	var issue Issue

	values := url.Values{}
	values.Add("apiKey", conf.APIKey)

	url := makeIssueURL(conf, issueKey) + "?" + values.Encode()
	body, err := Get(url)
	if err != nil {
		fmt.Println(err)
		return issue, err
	}

	json.Unmarshal(body, &issue)
	return issue, nil
}

// GetIssues implements to get issue by params
func GetIssues(conf config.Config, params map[string]string) ([]Issue, error) {
	values := url.Values{}
	values.Add("apiKey", conf.APIKey)
	for key, val := range params {
		values.Add(key, val)
	}

	url := conf.SpaceURL + IssuesURL + "?" + values.Encode()
	body, err := Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var issues []Issue
	json.Unmarshal(body, &issues)
	return issues, nil
}

// SearchIssueKeys provides to search IssueKey in text
func SearchIssueKeys(str string, projectKey string) []string {
	rep := regexp.MustCompile(projectKey + `-[\d]+`)
	keys := rep.FindAllString(str, -1)
	return keys
}

func makeIssueURL(conf config.Config, issueKey string) string {
	return conf.SpaceURL + IssuesURL + "/" + issueKey
}

// MakeViewURL provides backlog issue view url
func MakeViewURL(conf config.Config, issueKey string) string {
	return conf.SpaceURL + "/view/" + issueKey
}
