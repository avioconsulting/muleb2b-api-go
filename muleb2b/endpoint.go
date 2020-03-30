package muleb2b

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// These should eventually be retrieved from a B2b service since they could change
const sftpTypeId string = "3bcc65e5-040b-47eb-8fc7-27c89225f1bc"
const httpTypeId string = "aa1fd35b-50af-47fe-91bb-48a7ed4ab685"

type Endpoint struct {
	ID                   *string         `json:"id,omitempty"`
	Name                 *string         `json:"name"`
	Description          *string         `json:"description"`
	EnvironmentID        *string         `json:"environmentId"`
	EndpointTypeID       *string         `json:"endpointTypeId"`
	PartnerID            *string         `json:"partnerId"`
	HostCertificateID    *string         `json:"hostCertificateId,omitempty"`
	PartnerCertificateID *string         `json:"partnerCertificateId,omitempty"`
	UsedInDeployments    *int64          `json:"usedInDeployments,omitempty"`
	FlowDependencyCount  *int64          `json:"flowDependencyCount,omitempty"`
	UsedInConfigurations *int64          `json:"usedInConfigurations,omitempty"`
	Deployment           *interface{}    `json:"deployment,omitempty"`
	EndpointRole         *string         `json:"endpointRole"` // SEND, STORAGE_API, RECEIVE_ACK, RECEIVE
	EndpointType         *string         `json:"endpointType"`
	Config               *EndpointConfig `json:"config"`
	CreatedAt            *string         `json:"createdAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	IsComplete         *bool          `json:"isComplete,omitempty"`
	DependenciesInfo   *[]interface{} `json:"dependenciesInfo,omitempty"`
	HostCertificate    *interface{}   `json:"hostCertificate,omitempty"`
	PartnerCertificate *interface{}   `json:"partnerCertificate,omitempty"`
}

type EndpointConfig struct {
	FileAge                     *int64      `json:"fileAge,omitempty"`
	MovedPath                   *string     `json:"movedPath,omitempty"`
	SizeCheckWaitTime           *int        `json:"sizeCheckWaitTime,omitempty"`
	PollingFrequency            *int        `json:"pollingFrequency,omitempty"`
	Path                        *string     `json:"path,omitempty"`
	ServerAddress               *string     `json:"serverAddress,omitempty"`
	ServerPort                  *int        `json:"serverPort,omitempty"`
	ConfigName                  *string     `json:"configName"`
	AuthMode                    *AuthMode   `json:"authMode,omitempty"`
	Protocol                    *string     `json:"protocol,omitempty"`
	AllowedMethods              *string     `json:"allowedMethods,omitempty"`
	PersistentConnections       *bool       `json:"persistentConnections,omitempty"`
	ConnectionIdleTimeout       *int        `json:"connectionIdleTimeout,omitempty"`
	ResponseTimeout             *int        `json:"responseTimeout,omitempty"`
	AllowStorageApiUIConnection *bool       `json:"allowStorageApiUIConnection,omitempty"`
	TlsContext                  *TlsContext `json:"tlsContext,omitempty"`
}

type TlsContext struct {
	Insecure        *bool `json:"insecure"`
	NeedCertificate *bool `json:"needCertificate"`
}

type User struct {
	ID        *string `json:"userId"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	UserName  *string `json:"userName"`
}

// AuthMode Types
//   NONE
//   BASIC (+username, +password)
//   API_KEY (+apiKey, +httpHeaderName)
//   CLIENT_CREDENTIALS (+clientId, +clientSecret, +clientIdHeader, +clientSecretHeader)
//   OAUTH_TOKEN (+tokenUrl, +clientId, +clientSecret)
type AuthMode struct {
	AuthType           *string `json:"authType"`
	Username           *string `json:"username,omitempty"`
	Password           *string `json:"password,omitempty"`
	HttpHeaderName     *string `json:"httpHeaderName,omitempty"`
	ApiKey             *string `json:"apiKey,omitempty"`
	ClientId           *string `json:"clientId,omitempty"`
	ClientSecret       *string `json:"clientSecret,omitempty"`
	ClientIdHeader     *string `json:"clientIdHeader,omitempty"`
	ClientSecretHeader *string `json:"clientSecretHeader,omitempty"`
	TokenUrl           *string `json:"tokenUrl,omitempty"`
}

func (cli *Client) ListEndpoints() (*[]Endpoint, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/endpoints", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var endpoints []Endpoint

	_, err = cli.Do(req, &endpoints)
	if err != nil {
		return nil, err
	}
	return &endpoints, nil
}

func (cli *Client) GetEndpointsForPartner(partnerId string) (*[]Endpoint, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/endpoints?partnerId=%s", *cli.orgId, *cli.envId, partnerId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var endpoints []Endpoint

	_, err = cli.Do(req, &endpoints)
	if err != nil {
		return nil, err
	}
	return &endpoints, nil
}

func (cli *Client) GetEndpoint(id string) (*Endpoint, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/endpoints/%s", *cli.orgId, *cli.envId, id)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	req, err := cli.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var endpoint Endpoint

	resp, err := cli.Do(req, &endpoint)
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("endpoint (%s) not found", id)
	}
	if err != nil {
		return nil, err
	}
	return &endpoint, nil
}

func (cli *Client) CreateEndpoint(endpoint Endpoint) (*string, error) {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/endpoints", *cli.orgId, *cli.envId)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	if *endpoint.EndpointType == "sftp" && (endpoint.EndpointTypeID == nil || *endpoint.EndpointTypeID == "") {
		endpoint.EndpointTypeID = String(sftpTypeId)
	} else if *endpoint.EndpointType == "http" && (endpoint.EndpointTypeID == nil || *endpoint.EndpointTypeID == "") {
		endpoint.EndpointTypeID = String(httpTypeId)
	}

	fmt.Sprintf("Create Endpoint: %s\n", endpoint.String())

	req, err := cli.NewRequest("POST", u.String(), endpoint)
	if err != nil {
		return nil, err
	}

	var id string

	resp, err := cli.Do(req, &id)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("endpoint not created, status code is %d\nrequest is %s", resp.StatusCode, endpoint.String())
	}

	return &id, nil
}

func (cli *Client) UpdateEndpoint(endpoint Endpoint) error {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/endpoints/%s", *cli.orgId, *cli.envId, *endpoint.ID)}
	u := cli.PartnerBaseURL.ResolveReference(rel)

	if *endpoint.EndpointType == "sftp" && (endpoint.EndpointTypeID == nil || *endpoint.EndpointTypeID == "") {
		endpoint.EndpointTypeID = String(sftpTypeId)
	} else if *endpoint.EndpointType == "http" && (endpoint.EndpointTypeID == nil || *endpoint.EndpointTypeID == "") {
		endpoint.EndpointTypeID = String(httpTypeId)
	}

	fmt.Sprintf("Update Endpoint: %s\n", endpoint.String())

	req, err := cli.NewRequest("PUT", u.String(), endpoint)
	if err != nil {
		return err
	}

	resp, err := cli.Do(req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code %d\nEndpoint JSON is %s", resp.StatusCode, endpoint.String())
	}
	return nil
}

func (cli *Client) DeleteEndpoint(id string) error {
	rel := &url.URL{Path: fmt.Sprintf("organizations/%s/environments/%s/endpoints/%s", *cli.orgId, *cli.envId, id)}
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

func (e Endpoint) String() string {
	j, _ := json.Marshal(e)
	return string(j)
}
