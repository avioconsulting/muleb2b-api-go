package muleb2b

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
)

type Certificate struct {
	Id              *string `json:"id"`
	CertificateType *string `json:"certificateType,omitempty"`
	Name            *string `json:"name,omitempty"`
	Authority       *string `json:"authority,omitempty"`
	SerialNumber    *string `json:"serialNumber,omitempty"`
}

func (cli *Client) CreatePartnerCertificate(partnerId, certContent, certificateName, certificateType string) (*string, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/certificates", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("certName", certificateName)
	writer.WriteField("certType", certificateType)
	caCertHeader := make(textproto.MIMEHeader)
	caCertHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="certFile"; filename="cert.crt"`))
	caCertHeader.Set("Content-Type", "application/x-x509-ca-cert")
	certPart, err := writer.CreatePart(caCertHeader)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(certPart, strings.NewReader(certContent))
	if err != nil {
		return nil, err
	}
	writer.Close()
	request, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Accept", "application/json")
	if cli.accessToken != nil {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *cli.accessToken))
	}

	var cert Certificate
	_, err = cli.Do(request, &cert)
	if err != nil {
		return nil, err
	}

	if cert.Id != nil {
		return cert.Id, nil
	} else {
		return nil, fmt.Errorf("certificate ID is nil")
	}
}

func (cli *Client) ListPartnerCertificates(partnerId string) ([]*Certificate, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/certificates", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var certificates []*Certificate
	_, err = cli.Do(req, &certificates)
	if err != nil {
		return nil, err
	}

	return certificates, nil

}

func (cli *Client) GetPartnerCertificate(partnerId, certificateId string) (*Certificate, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/certificates/%s", *cli.orgId, *cli.envId, partnerId, certificateId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var certificate Certificate
	_, err = cli.Do(req, &certificate)
	if err != nil {
		return nil, err
	}

	return &certificate, nil
}

func (cli *Client) DeletePartnerCertificate(partnerId, certificateId string) error {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/partners/%s/certificates/%s", *cli.orgId, *cli.envId, partnerId, certificateId)}
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
