package muleb2b

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DocumentFlow struct {
	Id             *string                      `json:"id,omitempty"`
	Name           *string                      `json:"name"`
	Direction      *string                      `json:"direction,omitempty"`
	PartnerFromId  *string                      `json:"partnerFromId,omitempty"`
	PartnerToId    *string                      `json:"partnerToId,omitempty"`
	Configurations []*DocumentFlowConfiguration `json:"configurations,omitempty"`
}

type DocumentFlowConfiguration struct {
	Id                      *string    `json:"id,omitempty"`
	DocumentFlowId          *string    `json:"documentFlowId,omitempty"`
	EnvironmentId           *string    `json:"environmentId"`
	Status                  *string    `json:"status"`
	Version                 *int       `json:"version"`
	PreProcessingEndpointId *string    `json:"preProcessingEndpointId"`
	ReceivingEndpointId     *string    `json:"receivingEndpointId"`
	ReceivingAckEndpointId  *string    `json:"receivingAckEndpointId"`
	TargetEndpointId        *string    `json:"targetEndpointId"`
	SourceDocTypeId         *string    `json:"sourceDocTypeId"`
	TargetDocTypeId         *string    `json:"targetDocTypeId"`
	ReceivingAckConfig      *string    `json:"receivingAckConfig"`
	DocumentMapping         []*Mapping `json:"documentMapping,omitempty"`
}

type createDocumentFlowResponse struct {
	Id *string `json:"id"`
}

func (cli *Client) ListDocumentFlows() (*[]DocumentFlow, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/summarylist", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var docflows []DocumentFlow

	_, err = cli.Do(req, &docflows)
	if err != nil {
		return nil, err
	}

	return &docflows, nil
}

func (cli *Client) GetDocumentFlowById(flowId string) (*DocumentFlow, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/%s", *cli.orgId, *cli.envId, flowId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var docFlow DocumentFlow

	_, err = cli.Do(req, &docFlow)
	if err != nil {
		return nil, err
	}

	return &docFlow, nil
}

func (cli *Client) GetDocumentFlowByName(name string) (*DocumentFlow, error) {
	docFlows, err := cli.ListDocumentFlows()
	if err != nil {
		return nil, err
	}

	if docFlows != nil {
		for _, doc := range *docFlows {
			if doc.Name != nil && *doc.Name == name {
				return &doc, err
			}
		}
	}
	return nil, nil
}

func (cli *Client) CreateDocumentFlow(flow *DocumentFlow) (*string, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("POST", u.String(), flow)
	if err != nil {
		return nil, err
	}

	var response createDocumentFlowResponse

	resp, err := cli.Do(req, &response)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create document flow")
	}

	return response.Id, nil
}

func (cli *Client) UpdateDocumentFlow(flow *DocumentFlow) (*DocumentFlow, error) {
	if flow != nil {
		rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/%s", *cli.orgId, *cli.envId, *(*flow).Id)}
		u := cli.PartnerBaseURL.ResolveReference(rel)

		req, err := cli.NewRequest("PATCH", u.String(), flow)
		if err != nil {
			return nil, err
		}

		var docFlow DocumentFlow

		resp, err := cli.Do(req, &docFlow)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("received invalid status (%d) in response", resp.StatusCode)
		}
		return &docFlow, nil
	} else {
		return nil, fmt.Errorf("DocumentFlow input is nil")
	}
}

func (cli *Client) DeleteDocumentFlow(id string) error {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/documentflows/%s", *cli.orgId, *cli.envId, id)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := cli.Do(req, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code for delete (%d)", resp.StatusCode)
	}

	return nil
}

func (docFlow DocumentFlow) String() string {
	j, _ := json.Marshal(docFlow)
	return string(j)
}
