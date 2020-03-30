package muleb2b

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

const (
	listEndpointsResponse = `[
    {
        "id": "84f24b8b-b8ae-4c44-ae76-30673e657807",
        "name": "HTTP target 1",
        "description": "",
        "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
        "endpointTypeId": "aa1fd35b-50af-47fe-91bb-48a7ed4ab685",
        "partnerId": "9e39423e-da65-429a-b735-941e4b3fd350",
        "hostCertificateId": null,
        "partnerCertificateId": null,
        "usedInDeployments": 0,
        "flowDependencyCount": 1,
        "usedInConfigurations": 0,
        "deployment": null,
        "endpointRole": "SEND",
        "endpointType": "http",
        "config": {
            "protocol": "HTTP",
            "path": "/",
            "allowedMethods": "POST",
            "persistentConnections": false,
            "connectionIdleTimeout": 30000,
            "responseTimeout": 15000,
            "allowStorageApiUIConnection": false,
            "tlsContext": null,
            "configName": "http",
            "serverAddress": "test.com",
            "serverPort": 80,
            "authMode": {
                "authType": "NONE",
                "csmSecret": null
            }
        },
        "createdAt": "2020-01-16T18:29:25Z",
        "createdBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "updatedAt": "2020-01-16T18:29:25Z",
        "updatedBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "isComplete": null,
        "dependenciesInfo": [],
        "partner": {
            "id": "9e39423e-da65-429a-b735-941e4b3fd350",
            "name": "Test-2",
            "description": null,
            "websiteUrl": null,
            "status": {
                "id": "3fb16466-aa97-42a2-bf7f-592d741440c7",
                "startDate": null,
                "endDate": null,
                "status": "ACTIVE"
            },
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": true,
            "dependencyCounts": null
        },
        "hostCertificate": null,
        "partnerCertificate": null
    },
    {
        "id": "e6420269-6fb7-4daf-bdeb-43e8acc39046",
        "name": "HTTP target 2",
        "description": "",
        "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
        "endpointTypeId": "aa1fd35b-50af-47fe-91bb-48a7ed4ab685",
        "partnerId": "9e39423e-da65-429a-b735-941e4b3fd350",
        "hostCertificateId": null,
        "partnerCertificateId": null,
        "usedInDeployments": 0,
        "flowDependencyCount": 1,
        "usedInConfigurations": 0,
        "deployment": null,
        "endpointRole": "SEND",
        "endpointType": "http",
        "config": {
            "protocol": "HTTP",
            "path": "/",
            "allowedMethods": "POST",
            "persistentConnections": false,
            "connectionIdleTimeout": 30000,
            "responseTimeout": 15000,
            "allowStorageApiUIConnection": false,
            "tlsContext": null,
            "configName": "http",
            "serverAddress": "blah.com",
            "serverPort": 80,
            "authMode": {
                "authType": "NONE",
                "csmSecret": null
            }
        },
        "createdAt": "2020-01-16T19:07:49Z",
        "createdBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "updatedAt": "2020-01-16T19:07:49Z",
        "updatedBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "isComplete": null,
        "dependenciesInfo": [],
        "partner": {
            "id": "9e39423e-da65-429a-b735-941e4b3fd350",
            "name": "Test-2",
            "description": null,
            "websiteUrl": null,
            "status": {
                "id": "3fb16466-aa97-42a2-bf7f-592d741440c7",
                "startDate": null,
                "endDate": null,
                "status": "ACTIVE"
            },
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": true,
            "dependencyCounts": null
        },
        "hostCertificate": null,
        "partnerCertificate": null
    }
]`

	endpointsForPartnerResponse = `[
    {
        "id": "95d9e857-1810-4757-a1d3-57593fd235b7",
        "name": "HTTP send-ack 1",
        "description": "something",
        "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
        "endpointTypeId": "aa1fd35b-50af-47fe-91bb-48a7ed4ab685",
        "partnerId": "5d6db9ed-e7ec-4c0a-9ff5-f9c3ab56d8ea",
        "hostCertificateId": null,
        "partnerCertificateId": null,
        "usedInDeployments": 0,
        "flowDependencyCount": 0,
        "usedInConfigurations": 0,
        "deployment": null,
        "endpointRole": "RECEIVE_ACK",
        "endpointType": "http",
        "config": {
            "protocol": "HTTP",
            "path": "/",
            "allowedMethods": "POST",
            "persistentConnections": false,
            "connectionIdleTimeout": 30000,
            "responseTimeout": 15000,
            "allowStorageApiUIConnection": false,
            "tlsContext": null,
            "configName": "http",
            "serverAddress": "app.test.com",
            "serverPort": 80,
            "authMode": {
                "apiKey": "2e795218-eb8a-4e43-8a01-4e636edb1db6",
                "httpHeaderName": "X-API-Key",
                "authType": "API_KEY",
                "csmSecret": {
                    "secretId": "2e795218-eb8a-4e43-8a01-4e636edb1db6",
                    "secretGroupId": "b0c91b69-27c8-4a36-831e-b5e5f18e2616",
                    "grantResponse": {
                        "path": "sharedSecrets/2e795218-eb8a-4e43-8a01-4e636edb1db6",
                        "accessGrant": ""
                    }
                }
            }
        },
        "createdAt": "2019-11-21T20:02:35Z",
        "createdBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "updatedAt": "2019-11-21T20:02:35Z",
        "updatedBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "isComplete": null,
        "dependenciesInfo": [],
        "partner": {
            "id": "5d6db9ed-e7ec-4c0a-9ff5-f9c3ab56d8ea",
            "name": "Sample",
            "description": "Some description",
            "websiteUrl": "http://sample.com",
            "status": {
                "id": "1799e1e7-ae3d-4353-bff0-bb02c721566d",
                "startDate": null,
                "endDate": null,
                "status": "ACTIVE"
            },
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": false,
            "dependencyCounts": null
        }
    }
]`

	getEndpointResponse = `{
    "id": "84f24b8b-b8ae-4c44-ae76-30673e657807",
    "name": "HTTP target 1",
    "description": "",
    "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
    "endpointTypeId": "aa1fd35b-50af-47fe-91bb-48a7ed4ab685",
    "partnerId": "9e39423e-da65-429a-b735-941e4b3fd350",
    "hostCertificateId": null,
    "partnerCertificateId": null,
    "usedInDeployments": 0,
    "flowDependencyCount": 1,
    "usedInConfigurations": 0,
    "deployment": null,
    "endpointRole": "SEND",
    "endpointType": "http",
    "config": {
        "protocol": "HTTP",
        "path": "/",
        "allowedMethods": "POST",
        "persistentConnections": false,
        "connectionIdleTimeout": 30000,
        "responseTimeout": 15000,
        "allowStorageApiUIConnection": false,
        "tlsContext": null,
        "configName": "http",
        "serverAddress": "test.com",
        "serverPort": 80,
        "authMode": {
            "authType": "NONE",
            "csmSecret": null
        }
    },
    "createdAt": "2020-01-16T18:29:25Z",
    "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
    "updatedAt": "2020-01-16T18:29:25Z",
    "updatedBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
    "isComplete": true,
    "dependenciesInfo": [],
    "hostCertificate": null,
    "partnerCertificate": null
}`
)

func TestClient_ListEndpoints(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID

		w.Write([]byte(listEndpointsResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	endpoints, err := cli.ListEndpoints()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*endpoints))
}

func TestClient_GetEndpointsForPartner(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "5d6db9ed-e7ec-4c0a-9ff5-f9c3ab56d8ea")) // Partner ID

		w.Write([]byte(endpointsForPartnerResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	endpoints, err := cli.GetEndpointsForPartner("5d6db9ed-e7ec-4c0a-9ff5-f9c3ab56d8ea")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*endpoints))
}

func TestClient_GetEndpoint(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "84f24b8b-b8ae-4c44-ae76-30673e657807")) // Endpoint ID

		w.Write([]byte(getEndpointResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	endpoint, err := cli.GetEndpoint("84f24b8b-b8ae-4c44-ae76-30673e657807")

	assert.Nil(t, err)
	assert.NotNil(t, endpoint.ID)
	assert.Equal(t, "84f24b8b-b8ae-4c44-ae76-30673e657807", *endpoint.ID)
	assert.Equal(t, "SEND", *endpoint.EndpointRole)
}

func TestClient_CreateEndpoint(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`"84f24b8b-b8ae-4c44-ae76-30673e657807"`))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	authMode := AuthMode{
		AuthType: String("NONE"),
	}
	conf := EndpointConfig{
		Path:          String("/"),
		ServerAddress: String("test.com"),
		ServerPort:    Integer(80),
		ConfigName:    String("http"),
		AuthMode:      &authMode,
	}

	endpoint := Endpoint{
		Name:           String("HTTP Target"),
		Description:    String(""),
		EnvironmentID:  cli.envId,
		EndpointTypeID: String(httpTypeId),
		PartnerID:      String("5d6db9ed-e7ec-4c0a-9ff5-f9c3ab56d8ea"),
		EndpointRole:   String("SEND"),
		EndpointType:   String("http"),
		Config:         &conf,
	}

	id, err := cli.CreateEndpoint(endpoint)

	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, "84f24b8b-b8ae-4c44-ae76-30673e657807", *id)
}

func TestClient_UpdateEndpoint(t *testing.T) {
	description := "Test description"
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "84f24b8b-b8ae-4c44-ae76-30673e657807")) // Endpoint ID
		if r.Method == "GET" {
			w.Write([]byte(getEndpointResponse))
		} else {
			var endpoint Endpoint
			defer r.Body.Close()
			err := json.NewDecoder(r.Body).Decode(&endpoint)
			assert.Nil(t, err)
			assert.Equal(t, description, *endpoint.Description)
		}
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	endpoint, err := cli.GetEndpoint("84f24b8b-b8ae-4c44-ae76-30673e657807")

	assert.Nil(t, err)
	assert.NotNil(t, endpoint.ID)

	*endpoint.Description = description

	err = cli.UpdateEndpoint(*endpoint)
}

func TestClient_DeleteEndpoint(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "84f24b8b-b8ae-4c44-ae76-30673e657807")) // Endpoint ID
		assert.Equal(t, "DELETE", r.Method)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	err = cli.DeleteEndpoint("84f24b8b-b8ae-4c44-ae76-30673e657807")

	assert.Nil(t, err)
}
