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
	url := "https://apis.choreo.dev/projects/1.0.0/graphql"
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

type Owner struct {
	ID        string `json:"id"`
	IdpID     string `json:"idpId"`
	CreatedAt string `json:"createdAt"`
}

type Organization struct {
	ID     string `json:"id"`
	UUID   string `json:"uuid"`
	Handle string `json:"handle"`
	Name   string `json:"name"`
	Owner  Owner  `json:"owner"`
}

func GetOrganizations(token string) ([]Organization, error) {
	url := "https://apis.choreo.dev/orgs/1.0.0/orgs"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
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

	var organizations []Organization
	if err := json.NewDecoder(resp.Body).Decode(&organizations); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return organizations, nil
}

type Component struct {
	ProjectID   string `json:"projectId"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Handler     string `json:"handler"`
	DisplayName string `json:"displayName"`
	DisplayType string `json:"displayType"`
}

type ComponentsResponse struct {
	Data struct {
		Components []Component `json:"components"`
	} `json:"data"`
}

func GetComponents(orgHandler string, projectID string, token string) ([]Component, error) {
	url := "https://apis.choreo.dev/projects/1.0.0/graphql"
	query := fmt.Sprintf(`{"query":"query{ components(orgHandler: \"%s\", projectId: \"%s\"){\n          projectId, \n          id,\n          name,\n          status,\n          handler, \n          displayName,\n          displayType,\n        } \n      }"}`, orgHandler, projectID)

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

	var componentsResponse ComponentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&componentsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return componentsResponse.Data.Components, nil
}

type Environment struct {
	ID                   string `json:"id"`
	CreatedAt            string `json:"created_at"`
	OrganizationID       int    `json:"organization_id"`
	OrganizationUUID     string `json:"organization_uuid"`
	EnvName              string `json:"env_name"`
	Region               string `json:"region"`
	ChoreoEnv            string `json:"choreo_env"`
	ClusterID            string `json:"cluster_id"`
	DockerCredentialUUID string `json:"docker_credential_uuid"`
	ExternalApimEnvName  string `json:"external_apim_env_name"`
	InternalApimEnvName  string `json:"internal_apim_env_name"`
	SandboxApimEnvName   string `json:"sandbox_apim_env_name"`
	Critical             bool   `json:"critical"`
	DnsPrefix            string `json:"dns_prefix"`
	PdpWebAppDnsPrefix   string `json:"pdp_web_app_dns_prefix"`
	DeletionStatus       string `json:"deletion_status"`
	Sandbox              bool   `json:"sandbox"`
}

type EnvironmentsResponse struct {
	Data []Environment `json:"data"`
}

func GetEnvironments(orgID string, token string) ([]Environment, error) {
	url := fmt.Sprintf("https://apis.choreo.dev/devops/1.0.0/api/v1/organizations/%s/environment-templates", orgID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
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

	var environmentsResponse EnvironmentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&environmentsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return environmentsResponse.Data, nil
}
