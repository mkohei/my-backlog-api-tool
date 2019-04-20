// 課題一覧の取得
package backlog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/mkohei/my-backlog-api-tool/config"
)

const ISSUES_URL = "/api/v2/issues"

type Issue struct {
	ID          int    `json:"id"`
	ProjectID   int    `json:"projectId"`
	IssueKey    string `json:"issueKey"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetIssue(conf config.Config, issueKey string) (Issue, error) {
	var issue Issue

	values := url.Values{}
	values.Add("apiKey", conf.APIKey)

	url := conf.SpaceURL + ISSUES_URL + "/" + issueKey + "?" + values.Encode()
	body, err := Get(url)
	if err != nil {
		fmt.Println(err)
		return issue, err
	}

	json.Unmarshal(body, &issue)
	return issue, nil
}

func GetIssues(conf config.Config, params map[string]string) ([]Issue, error) {
	values := url.Values{}
	values.Add("apiKey", conf.APIKey)
	for key, val := range params {
		values.Add(key, val)
	}

	url := conf.SpaceURL + ISSUES_URL + "?" + values.Encode()
	body, err := Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var issues []Issue
	json.Unmarshal(body, &issues)
	return issues, nil
}

func SearchIssueKeys(str string, projectKey string) []string {
	rep := regexp.MustCompile(projectKey + `-[\d]+`)
	keys := rep.FindAllString(str, -1)
	return keys
}

func DispIssue(issue Issue, layer int, showStatus bool) {
	buf := strings.Repeat("    ", layer)
	if layer != 0 {
		buf += "->"
	}
	status := ""
	if showStatus {
		status = "[ " + issue.Status.Name + " ]"
	}
	fmt.Println(buf, status, issue.IssueKey, issue.Summary)
}

func DispIssueNotCompleted(issue Issue, layout int, showStatus bool) {
	completeID := 4
	if completeID != issue.Status.ID {
		DispIssue(issue, layout, showStatus)
	}
}
