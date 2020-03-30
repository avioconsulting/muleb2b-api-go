package muleb2b

import (
	"fmt"
	"net/url"
)

type EdiDocumentType struct {
	Id           *string `json:"id"`
	DocumentName *string `json:"documentName"`
	Label        *string `json:"label"`
	SchemaPath   *string `json:"schemaPath"`
	Description  *string `json:"description"`
}

func (cli *Client) ListEdiDocumentTypes(formatType, formatVersionId string) (*[]EdiDocumentType, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/ediFormats/%s/ediFormatVersions/%s/ediDocumentTypes", *cli.orgId, *cli.envId, formatType, formatVersionId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var docTypes []EdiDocumentType

	_, err = cli.Do(req, &docTypes)
	if err != nil {
		return nil, err
	}

	return &docTypes, nil
}

func (cli *Client) GetEdiDocumentTypeByLabel(formatType, formatVersionId, label string) (*EdiDocumentType, error) {
	docTypes, err := cli.ListEdiDocumentTypes(formatType, formatVersionId)

	if err != nil {
		return nil, err
	}

	if docTypes != nil {
		for _, docType := range *docTypes {
			if *docType.Label == label {
				return &docType, nil
			}
		}
	}
	return nil, nil
}

func (cli *Client) GetEdiDocumentTypeByName(formatType, formatVersionId, name string) (*EdiDocumentType, error) {
	docTypes, err := cli.ListEdiDocumentTypes(formatType, formatVersionId)

	if err != nil {
		return nil, err
	}

	if docTypes != nil {
		for _, docType := range *docTypes {
			if *docType.DocumentName == name {
				return &docType, nil
			}
		}
	}
	return nil, nil
}

func (cli *Client) GetEdiDocumentTypeById(formatType, formatVersionId, id string) (*EdiDocumentType, error) {
	docTypes, err := cli.ListEdiDocumentTypes(formatType, formatVersionId)

	if err != nil {
		return nil, err
	}

	if docTypes != nil {
		for _, docType := range *docTypes {
			if *docType.Id == id {
				return &docType, nil
			}
		}
	}
	return nil, nil
}
