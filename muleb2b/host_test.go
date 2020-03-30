package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

const (
	okHostResponse = `{
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
  "dependencyCounts": null,
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
    "HTTP"
  ],
  "standards": [],
  "usedInDeployments": {
    "inboundCount": 0,
    "outboundCount": 1
  }
}`
)

func TestGetHostDetails(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032"))
		w.Write([]byte(okHostResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	c.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")
	err = c.GetHostDetails()

	assert.Nil(t, err)
	assert.Equal(t, "9e39423e-da65-429a-b735-941e4b3fd350", *c.hostPartnerId)
	assert.Equal(t, "Test-2", *c.hostPartnerName)
}

func TestGetHostDetailsEnvNotSet(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032"))
		w.Write([]byte(okHostResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	c, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	err = c.GetHostDetails()

	assert.NotNil(t, err)
}
