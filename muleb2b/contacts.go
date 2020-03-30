package muleb2b

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Contact struct {
	Id          *string      `json:"id,omitempty"`
	Name        *string      `json:"name"`
	Email       *string      `json:"email"`
	Phone       *string      `json:"phoneNumber"`
	Status      *string      `json:"status,omitempty"`
	ContactType *ContactType `json:"contactType"`
}

type ContactType struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Label       *string `json:"label"`
}

// The Contact Types endpoint isn't available yet
func GetBusinessContactType() *ContactType {
	return &ContactType{
		Id:          String("50615d27-0356-47fc-bd1e-440c992cd98e"),
		Name:        String("Business"),
		Description: String("Business contact"),
		Label:       String("Business contact"),
	}
}

func GetTechnicalContactType() *ContactType {
	return &ContactType{
		Id:          String("020f4c28-a0c2-4e70-b25d-8ab68f1a2020"),
		Name:        String("Technical"),
		Description: String("Technical contact"),
		Label:       String("Technical contact"),
	}
}

func GetOtherContactType() *ContactType {
	return &ContactType{
		Id:          String("248ed98f-dcc4-403f-ad3d-dfaacd2ca3f0"),
		Name:        String("Other"),
		Description: String("Other contact"),
		Label:       String("Other contact"),
	}
}

// TODO: This needs to be updated once Partner Contacts are available in the Public API
func (cli *Client) GetPartnerContacts(partnerId string) ([]*Contact, error) {
	partner, err := cli.GetPartnerProfile(partnerId)
	if err != nil {
		return nil, err
	}
	if partner != nil {
		return partner.Contacts, nil
	} else {
		return nil, fmt.Errorf("partner is nil")
	}
}

func (cli *Client) GetPartnerContactById(partnerId, contactId string) (*Contact, error) {
	contacts, err := cli.GetPartnerContacts(partnerId)
	if err != nil {
		return nil, err
	}
	for _, contact := range contacts {
		if contact.Id != nil && *contact.Id == contactId {
			return contact, nil
		}
	}
	return nil, nil
}

func (cli *Client) UpdatePartnerContacts(partnerId string, contacts []*Contact) error {
	partner, err := cli.GetPartnerProfile(partnerId)
	if err != nil {
		return err
	}
	partner.Contacts = contacts
	err = cli.UpdatePartnerProfile(partner)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Client) DeletePartnerContact(partnerId, contactId string) error {
	if cli.envId == nil {
		return fmt.Errorf("environment must be set before Partner can be updated")
	}

	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/contacts/%s", *cli.orgId, *cli.envId, partnerId, contactId)}
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

func (contact *Contact) String() string {
	j, _ := json.Marshal(contact)
	return string(j)
}
