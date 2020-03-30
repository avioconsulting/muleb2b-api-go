package muleb2b

import (
	"fmt"
	"net/url"
)

type IdentifierType struct {
	Id            *string                    `json:"id"`
	Name          *string                    `json:"name"`
	Label         *string                    `json:"label"`
	Description   *string                    `json:"description"`
	Readonly      *bool                      `json:"readonly"`
	EnvironmentId *string                    `json:"environment_id"`
	Qualifiers    []*IdentifierTypeQualifier `json:"qualifiers"`
}

type IdentifierTypeQualifier struct {
	Id                *string `json:"id"`
	IdentifierTypeId  *string `json:"identifierTypeId"`
	Code              *string `json:"code"`
	Label             *string `json:"label"`
	SegmentIdentifier *string `json:"segmentIdentifier"`
	Description       *string `json:"descritpion"`
	EnvironmentId     *string `json:"environmentId"`
}

func (cli *Client) ListIdentifierTypes() ([]*IdentifierType, error) {
	if cli.envId != nil {
		rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/identifierTypes", *cli.orgId, *cli.envId)}
		u := cli.PartnerBaseURL.ResolveReference(rel)

		req, err := cli.NewRequest("GET", u.String(), nil)
		if err != nil {
			return nil, err
		}

		var identifierTypeList []*IdentifierType

		_, err = cli.Do(req, &identifierTypeList)
		if err != nil {
			return nil, err
		}

		return identifierTypeList, nil
	} else {
		return nil, fmt.Errorf("envId is not set")
	}
}

func (cli *Client) GetIdentifierTypesByName(name string) ([]*IdentifierType, error) {
	identifierTypes, err := cli.ListIdentifierTypes()

	if err != nil {
		return nil, err
	}

	var identifierTypeList []*IdentifierType
	for _, identifierType := range identifierTypes {
		if identifierType.Name != nil && *identifierType.Name == name {
			identifierTypeList = append(identifierTypeList, identifierType)
		}
	}
	return identifierTypeList, nil
}

func (cli *Client) GetIdentifierTypesByLabel(label string) ([]*IdentifierType, error) {
	identifierTypes, err := cli.ListIdentifierTypes()

	if err != nil {
		return nil, err
	}

	var identifierTypeList []*IdentifierType
	for _, identifierType := range identifierTypes {
		if identifierType.Label != nil && *identifierType.Label == label {
			identifierTypeList = append(identifierTypeList, identifierType)
		}
	}
	return identifierTypeList, nil
}

func (cli *Client) GetIdentifierTypesById(id string) ([]*IdentifierType, error) {
	identifierTypes, err := cli.ListIdentifierTypes()

	if err != nil {
		return nil, err
	}

	var identifierTypeList []*IdentifierType
	for _, identifierType := range identifierTypes {
		if identifierType.Id != nil && *identifierType.Id == id {
			identifierTypeList = append(identifierTypeList, identifierType)
		}
	}
	return identifierTypeList, nil
}

func (identifier *IdentifierType) GetIdentifierTypeQualifiersByCode(code string) ([]*IdentifierTypeQualifier, error) {
	var qualifierList []*IdentifierTypeQualifier
	for _, iq := range identifier.Qualifiers {
		if iq.Code != nil && *iq.Code == code {
			qualifierList = append(qualifierList, iq)
		}
	}
	return qualifierList, nil
}

func (identifier *IdentifierType) GetIdentifierTypeQualifiersByLabel(label string) ([]*IdentifierTypeQualifier, error) {
	var qualifierList []*IdentifierTypeQualifier
	for _, iq := range identifier.Qualifiers {
		if iq.Label != nil && *iq.Label == label {
			qualifierList = append(qualifierList, iq)
		}
	}
	return qualifierList, nil
}

func (identifier *IdentifierType) GetIdentifierTypeQualifiersById(id string) ([]*IdentifierTypeQualifier, error) {
	var qualifierList []*IdentifierTypeQualifier
	for _, iq := range identifier.Qualifiers {
		if iq.Id != nil && *iq.Id == id {
			qualifierList = append(qualifierList, iq)
		}
	}
	return qualifierList, nil
}
