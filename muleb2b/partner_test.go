package muleb2b

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

const (
	partnerCreationResponse = `{"id":"680b6a5f-df8f-4f3a-8277-a651e609a804"}`
	getPartnerResponse      = `{
  "id": "test-id",
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
  "dependencyCounts": {
    "messageFlowDeployment": 0,
    "messageFlowConfig": 1,
    "identifiers": 2,
    "endpoints": 1,
    "certificates": 0
  }
	}`
	getHostPartnerResponse = `{
    "logo": null,
    "contacts": [],
    "identifiers": [
        {
            "id": "56e8756b-3d33-4ce7-948b-292b88473f60",
            "identifierTypeQualifierId": "26432f80-b58b-4c96-aac6-58af5d5580fc",
            "status": "ACTIVE",
            "qualifierLabel": "12 (Phone Number)",
            "typeLabel": "X12 - ISA",
            "code": "12",
            "value": "6546546543"
        },
        {
            "id": "9afdb370-4873-4547-9013-74a32ce28340",
            "identifierTypeQualifierId": "25c1bc8a-801f-4947-a2a6-7721ef971460",
            "status": "ACTIVE",
            "qualifierLabel": "AS2 Identity",
            "typeLabel": "AS2 Identity",
            "code": "AS2",
            "value": "abc"
        },
        {
            "id": "95b33674-71f5-4df4-a77d-9dac1d27da55",
            "identifierTypeQualifierId": "25c1bc8a-801f-4947-a2a6-7721ef971460",
            "status": "ACTIVE",
            "qualifierLabel": "AS2 Identity",
            "typeLabel": "AS2 Identity",
            "code": "AS2",
            "value": "def"
        },
        {
            "id": "661c7324-53e5-4320-bdc8-e1e1dc660c24",
            "identifierTypeQualifierId": "25c1bc8a-801f-4947-a2a6-7721ef971460",
            "status": "ACTIVE",
            "qualifierLabel": "AS2 Identity",
            "typeLabel": "AS2 Identity",
            "code": "AS2",
            "value": "flyingmonkey"
        }
    ],
    "addresses": [],
    "createdAt": "2019-11-20T21:23:13Z",
    "createdBy": {
        "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
        "firstName": "test",
        "lastName": "test",
        "userName": "test"
    },
    "updatedAt": "2019-11-20T21:23:13Z",
    "updatedBy": "",
    "protocols": [
        "HTTP",
        "SFTP"
    ],
    "standards": [
        "JSON"
    ],
    "usedInDeployments": {
        "inboundCount": 0,
        "outboundCount": 3
    },
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
}`
	getHostPartnerExpectedUri   = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/partnerprofiles/host"
	getPartnerByNameExpectedUri = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/partners"
	getPartnerByNameResponse    = `[
    {
        "id": "ae5e69af-9de2-4355-8502-f2a330326ad2",
        "name": "Sample2",
        "description": null,
        "websiteUrl": null,
        "status": {
            "id": "0d0117c6-bef9-4842-be8d-ad02b18f01a6",
            "startDate": null,
            "endDate": null,
            "status": "ACTIVE"
        },
        "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
        "hostFlag": false,
        "dependencyCounts": null
    },
    {
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
    },
    {
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
    }
]`
)

func TestCreatePartner(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID

		var p Partner
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&p)

		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, "Test1234", *p.Name)

		w.Write([]byte(partnerCreationResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	partner := Partner{
		Name:          String("Test1234"),
		EnvironmentId: String("3a4d3936-22d9-4d87-a3c0-a8d424bcc032"),
	}

	id, err := c.CreatePartner(&partner)

	assert.Nil(t, err)
	assert.Equal(t, "680b6a5f-df8f-4f3a-8277-a651e609a804", *id)
}

func TestUpdatePartner(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID

		var p Partner
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&p)

		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, "test-id", *p.Id)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	partner := Partner{
		Id:            String("test-id"),
		Name:          String("Test1234"),
		Description:   String("This is a test"),
		WebsiteUrl:    String("test.com"),
		EnvironmentId: String("3a4d3936-22d9-4d87-a3c0-a8d424bcc032"),
	}

	err = c.UpdatePartner(&partner)

	assert.Nil(t, err)
}

func TestDeletePartnerSuccess(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "test-id"))
		assert.True(t, r.Method == "DELETE")
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)
	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	partner := Partner{
		Id: String("test-id"),
	}
	err = c.DeletePartner(&partner)

	assert.Nil(t, err)
}

func TestDeletePartnerByIdSuccess(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "test-id"))
		assert.True(t, r.Method == "DELETE")
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)
	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	err = c.DeletePartnerById(String("test-id"))

	assert.Nil(t, err)
}

func TestGetPartnerSuccess(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "test-id"))
		assert.True(t, r.Method == "GET")
		w.Write([]byte(getPartnerResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)
	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	partner, err := c.GetPartner("test-id")

	assert.Nil(t, err)
	assert.NotNil(t, partner)
	assert.Equal(t, "Sample", *(*partner).Name)
}

func TestGetHostPartner(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, getHostPartnerExpectedUri, r.URL.String())
		assert.True(t, r.Method == "GET")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getHostPartnerResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)
	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	partner, err := c.GetHostPartner()

	assert.Nil(t, err)
	assert.NotNil(t, partner)
	assert.Equal(t, "Test-2", *(*partner).Name)
	assert.Equal(t, "9e39423e-da65-429a-b735-941e4b3fd350", *(*partner).Id)
}

func TestGetPartnerByName(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, getPartnerByNameExpectedUri, r.URL.String())
		assert.True(t, r.Method == "GET")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getPartnerByNameResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)
	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	partner, err := c.GetPartnerByName("Sample2")

	assert.Nil(t, err)
	assert.NotNil(t, partner)
	assert.Equal(t, "ae5e69af-9de2-4355-8502-f2a330326ad2", *(*partner).Id)
}
