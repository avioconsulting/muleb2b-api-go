package muleb2b

import (
	"fmt"
	"net/url"
)

// TODO: This is only needed until Address and Contact are available in the Public API
type PartnerProfile struct {
	Id          *string    `json:"id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	WebsiteUrl  *string    `json:"websiteUrl"`
	Logo        *string    `json:"logo"`
	Addresses   []*Address `json:"addresses"`
	Contacts    []*Contact `json:"contacts"`
}

func (cli *Client) GetPartnerProfile(partnerId string) (*PartnerProfile, error) {
	if cli.envId == nil {
		return nil, fmt.Errorf("environment must be set before Partner can be read")
	} else if partnerId == "" {
		return nil, fmt.Errorf("partner must exist before it can be read. id is nil or empty")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partnerprofiles/%s", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var partner PartnerProfile
	_, err = cli.Do(req, &partner)
	if err != nil {
		return nil, err
	}

	return &partner, nil
	return nil, nil
}

func (cli *Client) UpdatePartnerProfile(partner *PartnerProfile) error {
	if cli.envId == nil {
		return fmt.Errorf("environment must be set before Partner can be updated")
	} else if partner.Id == nil {
		return fmt.Errorf("partner must exist before it is updated. Partner.Id is nil")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partnerprofiles/%s", *cli.orgId, *cli.envId, *partner.Id)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("PATCH", u.String(), partner)
	if err != nil {
		return err
	}

	_, err = cli.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
