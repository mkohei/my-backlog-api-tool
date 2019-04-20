// プロジェクト一覧の取得
package backlog

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/mkohei/my-backlog-api-tool/config"
)

type Project struct {
	ID         int    `json:"id"`
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
}

const PROJECTS_URL = "/api/v2/projects"

func GetProject(conf config.Config, projectKey string) (Project, error) {
	var project Project

	values := url.Values{}
	values.Add("apiKey", conf.APIKey)

	url := conf.SpaceURL + PROJECTS_URL + "/" + projectKey + "?" + values.Encode()

	body, err := Get(url)
	if err != nil {
		fmt.Println(err)
		return project, err
	}

	json.Unmarshal(body, &project)
	return project, nil
}

func GetProjects(conf config.Config) ([]Project, error) {
	values := url.Values{}
	values.Add("apiKey", conf.APIKey)

	url := conf.SpaceURL + PROJECTS_URL + "?" + values.Encode()
	body, err := Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var projects []Project
	json.Unmarshal(body, &projects)
	return projects, nil
}
