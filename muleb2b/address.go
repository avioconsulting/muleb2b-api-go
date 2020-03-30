package muleb2b

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Id         *string `json:"id,omitempty"`
	Addr1      *string `json:"addressLine_1"`
	Addr2      *string `json:"addressLine_2,omitempty"`
	City       *string `json:"city"`
	State      *string `json:"state"`
	Country    *string `json:"country"`
	PostalCode *string `json:"postalCode"`
}

// TODO: This needs to be updated after the Public API is updated to make Addresses available
func (cli *Client) GetPartnerAddress(partnerId string) (*Address, error) {
	partner, err := cli.GetPartnerProfile(partnerId)
	if err != nil {
		return nil, err
	}
	if partner != nil && len(partner.Addresses) > 0 {
		return partner.Addresses[0], nil
	} else {
		return nil, fmt.Errorf("partner is nil")
	}
}

func (cli *Client) UpdatePartnerAddress(partnerId string, address *Address) error {
	partner, err := cli.GetPartnerProfile(partnerId)
	if err != nil {
		return err
	}
	if len(partner.Addresses) > 0 {
		address.Id = partner.Addresses[0].Id
		partner.Addresses[0] = address
	} else {
		var addresses []*Address
		addresses = append(addresses, address)
		partner.Addresses = addresses
	}
	err = cli.UpdatePartnerProfile(partner)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Client) DeletePartnerAddress(partnerId string) error {
	if cli.envId == nil {
		return fmt.Errorf("environment must be set before Partner can be updated")
	}

	address, err := cli.GetPartnerAddress(partnerId)
	if err != nil {
		return err
	}

	empty := String("")
	address.Addr1 = empty
	address.Addr2 = empty
	address.City = empty
	address.State = empty
	address.Country = empty
	address.PostalCode = empty

	err = cli.UpdatePartnerAddress(partnerId, address)
	return err
}

func (address *Address) Empty() bool {
	if address == nil {
		return true
	}
	if address.Addr1 != nil && *address.Addr1 != "" {
		return false
	}
	if address.Addr2 != nil && *address.Addr2 != "" {
		return false
	}
	if address.City != nil && *address.City != "" {
		return false
	}
	if address.State != nil && *address.State != "" {
		return false
	}
	if address.PostalCode != nil && *address.PostalCode != "" {
		return false
	}
	if address.Country != nil && *address.Country != "" {
		return false
	}
	return true
}

func (address *Address) String() string {
	j, _ := json.Marshal(address)
	return string(j)
}
