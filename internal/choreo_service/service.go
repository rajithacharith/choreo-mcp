package choreoservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Project struct {
	ID                        string   `json:"id"`
	OrgID                     int      `json:"orgId"`
	Name                      string   `json:"name"`
	Version                   string   `json:"version"`
	CreatedDate               string   `json:"createdDate"`
	Handler                   string   `json:"handler"`
	Region                    string   `json:"region"`
	Description               string   `json:"description"`
	DefaultDeploymentPipeline string   `json:"defaultDeploymentPipelineId"`
	DeploymentPipelineIDs     []string `json:"deploymentPipelineIds"`
	Type                      string   `json:"type"`
	GitProvider               *string  `json:"gitProvider"`
	GitOrganization           *string  `json:"gitOrganization"`
	Repository                *string  `json:"repository"`
	Branch                    *string  `json:"branch"`
	SecretRef                 *string  `json:"secretRef"`
	UpdatedAt                 string   `json:"updatedAt"`
}

type ProjectsResponse struct {
	Data struct {
		Projects []Project `json:"projects"`
	} `json:"data"`
}

func GetProjects(orgID string, token string) ([]Project, error) {
	url := "https://apis.preview-dv.choreo.dev/projects/1.0.0/graphql"
	query := fmt.Sprintf(`{"query":"query{projects(orgId: %s){id, orgId, name, version, createdDate, handler, region, description, defaultDeploymentPipelineId, deploymentPipelineIds, type, gitProvider, gitOrganization, repository, branch, secretRef, updatedAt}}"}`, orgID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var projectsResponse ProjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&projectsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return projectsResponse.Data.Projects, nil
}
