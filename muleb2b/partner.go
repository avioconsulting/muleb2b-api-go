package muleb2b

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

type Partner struct {
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name"`
	Description   *string `json:"description,omitempty"`
	WebsiteUrl    *string `json:"websiteUrl,omitempty"`
	Status        *Status `json:"status,omitempty"`
	EnvironmentId *string `json:"environmentId"`
	HostFlag      *bool   `json:"hostFlag,omitempty"`
}

type Status struct {
	Id        *string    `json:"id"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
	Status    *string    `json:"status"`
}

type creationResponse struct {
	Id *string `json:"id"`
}

// Create a Partner
func (cli *Client) CreatePartner(partner *Partner) (*string, error) {

	if cli.envId == nil {
		return nil, errors.New("Environment must be set before Partner can be created")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("POST", u.String(), partner)
	if err != nil {
		return nil, err
	}

	var resp creationResponse

	_, err = cli.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Id, nil
}

// Update a Partner
func (cli *Client) UpdatePartner(partner *Partner) error {
	if cli.envId == nil {
		return errors.New("Environment must be set before Partner can be updated")
	} else if partner.Id == nil {
		return errors.New("Partner must exist before it is updated. Partner.Id is nil")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s", *cli.orgId, *cli.envId, *partner.Id)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("PUT", u.String(), partner)
	if err != nil {
		return err
	}

	_, err = cli.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Delete a partner using its ID
func (cli *Client) DeletePartnerById(id *string) error {
	if cli.envId == nil {
		return errors.New("Environment must be set before Partner can be updated")
	} else if id == nil || *id == "" {
		return errors.New("Partner must exist before it is deleted. ID is nil or empty")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s", *cli.orgId, *cli.envId, *id)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}

	_, err = cli.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Delete a partner given a Partner containing at least Id
func (cli *Client) DeletePartner(partner *Partner) error {
	if partner != nil {
		return cli.DeletePartnerById(partner.Id)
	}
	return errors.New("Partner is nil")
}

// Retrieve a Partner using its ID
func (cli *Client) GetPartner(id string) (*Partner, error) {
	if cli.envId == nil {
		return nil, errors.New("Environment must be set before Partner can be read")
	} else if id == "" {
		return nil, errors.New("Partner must exist before it can be read. id is nil or empty")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s", *cli.orgId, *cli.envId, id)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var partner Partner
	_, err = cli.Do(req, &partner)
	if err != nil {
		return nil, err
	}

	return &partner, nil
}

func (cli *Client) ListPartners() (*[]Partner, error) {
	if cli.envId == nil {
		return nil, errors.New("Environment must be set before Partner can be read")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var partners []Partner
	_, err = cli.Do(req, &partners)
	if err != nil {
		return nil, err
	}

	return &partners, nil
}

// TODO: Change this once the bug for hostFlag is fixed
func (cli *Client) GetHostPartner() (*Partner, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partnerprofiles/host", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var partner Partner
	_, err = cli.Do(req, &partner)
	if err != nil {
		return nil, err
	}
	return &partner, nil
}

func (cli *Client) GetPartnerByName(name string) (*Partner, error) {
	if cli.envId == nil {
		return nil, errors.New("Environment must be set before Partner can be read")
	}

	partners, err := cli.ListPartners()
	if err != nil {
		return nil, err
	}
	if partners != nil {
		for _, partner := range *partners {
			if partner.Name != nil && *partner.Name == name {
				return &partner, nil
			}
		}
	}

	return nil, nil
}

func (p Partner) String() string {
	j, _ := json.Marshal(p)
	return string(j)
}
