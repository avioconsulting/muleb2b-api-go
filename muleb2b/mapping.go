package muleb2b

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	MappingContentPrefix = "data:application/octet-stream;base64,"
)

type Mapping struct {
	Id               *string `json:"id"`
	MappingType      *string `json:"mappingType"`
	MappingContent   *string `json:"mappingContent"`
	MappingSourceRef *string `json:"mappingSourceRef"`
}

func (cli *Client) ListMappings(docFlowId string) (*[]Mapping, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/%s/mappings", *cli.orgId, *cli.envId, docFlowId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var mappings []Mapping

	_, err = cli.Do(req, &mappings)
	if err != nil {
		return nil, err
	}
	return &mappings, nil
}

func (cli *Client) GetMappingById(docFlowId, mappingId string) (*Mapping, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/%s/mappings/%s", *cli.orgId, *cli.envId, docFlowId, mappingId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var mapping Mapping

	_, err = cli.Do(req, &mapping)
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

func (cli *Client) CreateMapping(docFlowId string, mapping *Mapping) error {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/%s/mappings", *cli.orgId, *cli.envId, docFlowId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("POST", u.String(), mapping)
	if err != nil {
		return err
	}

	resp, err := cli.Do(req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error creating mapping for document flow (%s)", docFlowId)
	}

	return nil
}

func (cli *Client) UpdateMapping(docFlowId string, mapping *Mapping) error {
	if mapping != nil {
		rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/%s/mappings/%s", *cli.orgId, *cli.envId, docFlowId, *(*mapping).Id)}
		u := cli.PartnerBaseURL.ResolveReference(rel)

		req, err := cli.NewRequest("PUT", u.String(), mapping)
		if err != nil {
			return err
		}

		resp, err := cli.Do(req, nil)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusNoContent {
			return fmt.Errorf("error creating mapping for document flow (%s)", docFlowId)
		}

		return nil
	} else {
		return fmt.Errorf("mapping is nil")
	}
}
