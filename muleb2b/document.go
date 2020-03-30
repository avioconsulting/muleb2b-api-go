package muleb2b

import (
	"fmt"
	"net/url"
)

type Document struct {
	Id                *string `json:"id"`
	Name              *string `json:"name"`
	Standard          *bool   `json:"isStandard"`
	EdiDocumentTypeId *string `json:"ediDocumentTypeId"`
	SchemaType        *string `json:"schemaType,omitempty"`
	SchemaContent     *string `json:"schemaContent,omitempty"`
	CustomSchemaId    *string `json:"customSchemaId,omitempty"`
}

type createDocumentResponse struct {
	Id *string `json:"id"`
}

func (cli *Client) ListDocuments(partnerId string) (*[]Document, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/documents", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var docs []Document

	_, err = cli.Do(req, &docs)
	if err != nil {
		return nil, err
	}

	return &docs, nil
}

func (cli *Client) GetDocumentById(partnerId, documentId string) (*Document, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/documents/%s", *cli.orgId, *cli.envId, partnerId, documentId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var doc Document

	_, err = cli.Do(req, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (cli *Client) GetDocumentByName(partnerId, name string) (*Document, error) {
	docs, err := cli.ListDocuments(partnerId)
	if err != nil {
		return nil, err
	}

	if docs != nil {
		for _, doc := range *docs {
			if doc.Name != nil && *doc.Name == name {
				return &doc, err
			}
		}
	}

	return nil, nil
}

func (cli *Client) CreateDocument(partnerId string, document *Document) (*string, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/documents", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("POST", u.String(), document)
	if err != nil {
		return nil, err
	}

	var response createDocumentResponse

	_, err = cli.Do(req, &response)
	if err != nil {
		return nil, err
	}

	return response.Id, nil
}

func (cli *Client) UpdateDocument(partnerId string, document *Document) error {
	if document != nil && document.Id != nil {
		rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/documents/%s", *cli.orgId, *cli.envId, partnerId, *document.Id)}
		u := cli.PartnerBaseURL.ResolveReference(rel)

		req, err := cli.NewRequest("PUT", u.String(), document)
		if err != nil {
			return err
		}

		_, err = cli.Do(req, nil)
		if err != nil {
			return err
		}

		return nil

	} else {
		return fmt.Errorf("document id must not be nil")
	}
}

func (cli *Client) DeleteDocument(partnerId string, document *Document) error {
	if document != nil && document.Id != nil {
		return cli.DeleteDocumentById(partnerId, *document.Id)
	} else {
		return fmt.Errorf("document id must not be nil")
	}

}

func (cli *Client) DeleteDocumentById(partnerId, documentId string) error {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/documents/%s", *cli.orgId, *cli.envId, partnerId, documentId)}
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
