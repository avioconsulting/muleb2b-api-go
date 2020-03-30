package muleb2b

import (
	"encoding/json"
	"fmt"
	"net/url"
)

/*
 * QualifierLabel, TypeLabel, and Code are all identified by IdentifierTypeQualifierId, and thus, not necessary
 */
type Identifier struct {
	Id                        *string `json:"id,omitempty"`
	IdentifierTypeQualifierId *string `json:"identifierTypeQualifierId"`
	Status                    *string `json:"status,omitempty"`
	Value *string `json:"value"`
}

func (cli *Client) CreatePartnerIdentifier(partnerId string, identifier *Identifier) error {
	if cli.envId != nil {
		if identifier != nil {
			rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/identifiers", *cli.orgId, *cli.envId, partnerId)}
			u := cli.PartnerBaseURL.ResolveReference(rel)

			req, err := cli.NewRequest("POST", u.String(), identifier)
			if err != nil {
				return err
			}

			_, err = cli.Do(req, nil)
			if err != nil {
				return err
			}
			return nil
		} else {
			return fmt.Errorf("cannot accept a nil Identifier")
		}
	} else {
		return fmt.Errorf("envId is not set")
	}
}

func (cli *Client) ListPartnerIdentifiers(partnerId string) ([]*Identifier, error) {
	if cli.envId != nil {
		rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/identifiers", *cli.orgId, *cli.envId, partnerId)}
		u := cli.PartnerBaseURL.ResolveReference(rel)

		req, err := cli.NewRequest("GET", u.String(), nil)

		if err != nil {
			return nil, err
		}

		var identifierList []*Identifier

		_, err = cli.Do(req, &identifierList)
		if err != nil {
			return nil, err
		}

		return identifierList, nil
	} else {
		return nil, fmt.Errorf("envId is not set")
	}
}

func (cli *Client) GetPartnerIdentifiersByValue(partnerId, value string) ([]*Identifier, error) {
	identifiers, err := cli.ListPartnerIdentifiers(partnerId)
	if err != nil {
		return nil, err
	}

	if len(identifiers) > 0 {
		var identifierList []*Identifier
		for _, identifier := range identifiers {
			if identifier.Value != nil && *identifier.Value == value {
				identifierList = append(identifierList, identifier)
			}
		}
		return identifierList, nil
	} else {
		return nil, nil
	}
}

func (cli *Client) GetPartnerIdentifierById(partnerId, identifierId string) (*Identifier, error) {
	if cli.envId != nil {
		rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/identifiers/%s", *cli.orgId, *cli.envId, partnerId, identifierId)}
		u := cli.PartnerBaseURL.ResolveReference(rel)

		req, err := cli.NewRequest("GET", u.String(), nil)

		if err != nil {
			return nil, err
		}

		var identifier *Identifier

		_, err = cli.Do(req, &identifier)
		if err != nil {
			return nil, err
		}

		return identifier, nil
	} else {
		return nil, fmt.Errorf("envId is not set")
	}
}

func (cli *Client) GetPartnerIdentifierByQualifierIdAndValue(partnerId, identifierTypeQualifierId, value string) (*Identifier, error) {
	identifiers, err := cli.ListPartnerIdentifiers(partnerId)
	if err != nil {
		return nil, err
	}

	if len(identifiers) > 0 {
		for _, identifier := range identifiers {
			if identifier.IdentifierTypeQualifierId != nil && *identifier.IdentifierTypeQualifierId == identifierTypeQualifierId &&
				identifier.Value != nil && *identifier.Value == value {
				return identifier, nil
			}
		}
	}
	return nil, nil

}

func (cli *Client) DeletePartnerIdentifier(partnerId, identifierId string) error {
	if cli.envId != nil {
		rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/identifiers/%s", *cli.orgId, *cli.envId, partnerId, identifierId)}
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
	} else {
		return fmt.Errorf("envId is not set")
	}
}

func (i *Identifier) QualifierIdAndValueEqual(o *Identifier) bool {
	if i == nil && o == nil {
		return true
	} else if i != nil && o != nil {
		if i.IdentifierTypeQualifierId != nil &&
			o.IdentifierTypeQualifierId != nil &&
			*(*i).IdentifierTypeQualifierId == *(*o).IdentifierTypeQualifierId &&
			i.Value != nil &&
			o.Value != nil &&
			*(*i).Value == *(*o).Value {
			return true
		}
	}
	return false
}

func (i *Identifier) String() string {
	j, _ := json.Marshal(i)
	return string(j)
}
