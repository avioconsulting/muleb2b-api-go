package muleb2b

import (
	"fmt"
	"net/url"
)

type environmentResponse struct {
	Data  *[]Environment `json:"data"`
	Total *int           `json:"total"`
}

type Environment struct {
	Id           *string `json:"id"`
	Name         *string `json:"name"`
	OrgId        *string `json:"organizationId"`
	IsProduction *bool   `json:"isProduction"`
	Type         *string `json:"type"`
	ClientId     *string `json:"clientId"`
}

func (cli *Client) ListEnvironments() (*[]Environment, error) {
	rel := &url.URL{Path: fmt.Sprintf("accounts/api/organizations/%s/environments", *cli.orgId)}
	u := cli.BaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var envResp environmentResponse

	_, err = cli.Do(req, &envResp)
	if err != nil {
		return nil, err
	}

	envList := *envResp.Data

	return &envList, nil
}

func (cli *Client) GetEnvironmentByName(name string) (*Environment, error) {
	envs, err := cli.ListEnvironments()

	if err != nil {
		return nil, err
	}

	for _, env := range *envs {
		if *env.Name == name {
			return &env, nil
		}
	}

	return nil, nil
}

func (cli *Client) GetEnvironmentById(id string) (*Environment, error) {
	envs, err := cli.ListEnvironments()

	if err != nil {
		return nil, err
	}

	for _, env := range *envs {
		if *env.Id == id {
			return &env, nil
		}
	}

	return nil, nil
}

func (cli *Client) ListEnvironmentsByType(envType string) (*[]Environment, error) {
	envs, err := cli.ListEnvironments()

	if err != nil {
		return nil, err
	}

	var envList []Environment

	for _, env := range *envs {
		if *env.Type == envType {
			envList = append(envList, env)
		}
	}

	return &envList, nil
}

func (cli *Client) ListProductionEnvironments() (*[]Environment, error) {
	envs, err := cli.ListEnvironments()

	if err != nil {
		return nil, err
	}

	var envList []Environment

	for _, env := range *envs {
		if *env.IsProduction {
			envList = append(envList, env)
		}
	}

	return &envList, nil
}

func (cli *Client) SetEnvironment(id string) {
	cli.envId = &id
}
