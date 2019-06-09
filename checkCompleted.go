// The seedgen program ...
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mkohei/my-backlog-api-tool/backlog"
	"github.com/mkohei/my-backlog-api-tool/config"
)

func errorExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	targetIssueKey := getTargetIssueKey()

	conf, err := config.LoadConfig()
	errorExit(err)

	// 課題情報の取得
	targetIssue, err := backlog.GetIssue(conf, targetIssueKey)
	errorExit(err)
	displayIssue(conf, targetIssue, 0, false)

	params := map[string]string{}

	// 対象課題の子課題
	params["parentIssueId[]"] = strconv.Itoa(targetIssue.ID)
	childIssues, err := backlog.GetIssues(conf, params)
	errorExit(err)
	for _, childIssue := range childIssues {
		displayIssueNotCompleted(conf, childIssue, 1, true)
	}

	// 詳細にある IssueKey を取得
	projectKey := strings.Split(targetIssue.IssueKey, "-")[0]
	issueKeys := backlog.SearchIssueKeys(targetIssue.Description, projectKey)
	sort.Strings(issueKeys)

	for _, issueKey := range issueKeys {
		issue, err := backlog.GetIssue(conf, issueKey)
		errorExit(err)
		displayIssue(conf, issue, 1, true)

		// 詳細にある Issue の子課題
		params["parentIssueId[]"] = strconv.Itoa(issue.ID)
		childIssues, err := backlog.GetIssues(conf, params)
		errorExit(err)
		for _, childIssue := range childIssues {
			displayIssueNotCompleted(conf, childIssue, 2, true)
		}
	}
}

func getTargetIssueKey() string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please set IssueKey.\n$go run checkCompleted.go [IssueKey]")
		os.Exit(1)
	}
	return args[0]
}

func displayIssue(conf config.Config, issue backlog.Issue, layer int, showStatus bool) {
	// buffer
	buf := strings.Repeat("	", layer)
	buf1 := buf
	buf2 := buf + "	"
	if layer != 0 {
		buf1 += "->"
	}
	// status
	status := ""
	if showStatus {
		status = "[ " + issue.Status.Name + " ]"
	}
	// assignee
	assignee := " ( " + issue.Assignee.Name + " ) "
	// url
	url := backlog.MakeViewURL(conf, issue.IssueKey)
	// display
	//fmt.Println(buf1, status, issue.IssueKey, issue.Summary)
	//fmt.Print(buf2, assignee, url, "\n\n")
	fmt.Println(buf1, status, assignee, issue.IssueKey, issue.Summary, url, buf2)
}

func displayIssueNotCompleted(conf config.Config, issue backlog.Issue, layer int, showStatus bool) {
	completeID := 4
	if completeID != issue.Status.ID {
		displayIssue(conf, issue, layer, showStatus)
	}
}
