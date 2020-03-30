package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	ptnrIdentifiersTestUrl         = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/partners/6514a87c-ae99-45da-9c6e-29114dc19040/identifiers"
	ptnrSingleIdentifierTestUrl    = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/partners/6514a87c-ae99-45da-9c6e-29114dc19040/identifiers/6ba3a211-a1c0-4a19-9b74-31e2312939dc"
	ptnrIdentifierListTestResponse = `[
    {
        "id": "681ea553-5527-4131-a866-48a32a8fd5bf",
        "identifierTypeQualifierId": "25c1bc8a-801f-4947-a2a6-7721ef971460",
        "status": "ACTIVE",
        "qualifierLabel": "AS2 Identity",
        "typeLabel": "AS2 Identity",
        "code": "AS2",
        "value": "ghgg"
    },
    {
        "id": "6ba3a211-a1c0-4a19-9b74-31e2312939dc",
        "identifierTypeQualifierId": "25c1bc8a-801f-4947-a2a6-7721ef971460",
        "status": "ACTIVE",
        "qualifierLabel": "AS2 Identity",
        "typeLabel": "AS2 Identity",
        "code": "AS2",
        "value": "TestingTesting123"
    },
    {
        "id": "6ba3a211-a1c0-4a19-9b74-31e2312939db",
        "identifierTypeQualifierId": "25c1bc8a-801f-4947-a2a6-7721ef971460",
        "status": "ACTIVE",
        "qualifierLabel": "AS2 Identity",
        "typeLabel": "AS2 Identity",
        "code": "AS2",
        "value": "TestingTesting123"
    }
	]`
	ptnrIdentifierGetResponse = `{
		"id": "681ea553-5527-4131-a866-48a32a8fd5bf",
		"identifierTypeQualifierId": "25c1bc8a-801f-4947-a2a6-7721ef971460",
		"status": "ACTIVE",
		"qualifierLabel": "AS2 Identity",
		"typeLabel": "AS2 Identity",
		"code": "AS2",
		"value": "ghgg"
	}`
	ptnrCreationResponse = "{}"
)

func TestClient_ListIdentifiers(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ptnrIdentifiersTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ptnrIdentifierListTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.ListPartnerIdentifiers("6514a87c-ae99-45da-9c6e-29114dc19040")

	assert.Nil(t, err)
	assert.NotNil(t, identifiers)
	assert.Equal(t, 3, len(identifiers))
}

func TestClient_GetPartnerIdentifierById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ptnrSingleIdentifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ptnrIdentifierGetResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifier, err := cli.GetPartnerIdentifierById("6514a87c-ae99-45da-9c6e-29114dc19040", "6ba3a211-a1c0-4a19-9b74-31e2312939dc")

	assert.Nil(t, err)
	assert.NotNil(t, identifier)
	assert.NotNil(t, identifier.Value)
	assert.Equal(t, "ghgg", *identifier.Value)
}

func TestClient_GetPartnerIdentifiersByValue(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ptnrIdentifiersTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ptnrIdentifierListTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.GetPartnerIdentifiersByValue("6514a87c-ae99-45da-9c6e-29114dc19040", "TestingTesting123")

	assert.Nil(t, err)
	assert.NotNil(t, identifiers)
	assert.Equal(t, 2, len(identifiers))
}

func TestClient_CreatePartnerIdentifier(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ptnrIdentifiersTestUrl, r.URL.String())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(ptnrCreationResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifier := Identifier{
		IdentifierTypeQualifierId: String("25c1bc8a-801f-4947-a2a6-7721ef971460"),
		Status:                    String("ACTIVE"),
		Value:                     String("TestingTesting1234"),
	}

	err = cli.CreatePartnerIdentifier("6514a87c-ae99-45da-9c6e-29114dc19040", &identifier)

	assert.Nil(t, err)
}

func TestClient_DeletePartnerIdentifier(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ptnrSingleIdentifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(ptnrCreationResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	cli.DeletePartnerIdentifier("6514a87c-ae99-45da-9c6e-29114dc19040", "6ba3a211-a1c0-4a19-9b74-31e2312939dc")

	assert.Nil(t, err)
}
