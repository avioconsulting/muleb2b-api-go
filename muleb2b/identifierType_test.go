package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	identifierTestUrl      = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/identifierTypes"
	identifierTestResponse = `[
    {
        "id": "65715d63-b628-4eb0-913d-cfdb55057c15",
        "name": "AS2",
        "label": "AS2 Identity",
        "description": "",
        "readonly": true,
        "environment_id": "00000000-0000-0000-0000-000000000000",
        "qualifiers": [
            {
                "id": "25c1bc8a-801f-4947-a2a6-7721ef971460",
                "identifierTypeId": "65715d63-b628-4eb0-913d-cfdb55057c15",
                "code": "AS2",
                "label": "AS2 Identity",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            }
        ]
    },
    {
        "id": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
        "name": "X12-ISA",
        "label": "X12 - ISA",
        "description": "",
        "readonly": true,
        "environment_id": "00000000-0000-0000-0000-000000000000",
        "qualifiers": [
            {
                "id": "dacbb21c-f452-434e-985f-c8c588ebf2f2",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "01",
                "label": "01 (Duns Number)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "26432f80-b58b-4c96-aac6-58af5d5580fc",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "12",
                "label": "12 (Phone Number)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "2e92b8d1-0510-4d59-ac3b-953d08127f78",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "14",
                "label": "14 (Duns Number and Suffix Number)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "fc397111-9590-4fea-a238-5853d575e04e",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "ZZ",
                "label": "ZZ (Mutually Defined)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "888bbdfc-4693-4267-a0d4-22a059fde2e4",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "02",
                "label": "02 (SCAC - Standard Carrier Alpha Code)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "4f583fac-07f0-439a-9596-4459b9b07f99",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "03",
                "label": "03 (FMC - Federal Maritime Commission)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "34ec0e15-b57f-4f57-b70e-3b25e55ee200",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "04",
                "label": "04 (IATA - International Air Transport Association)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "912a6767-ae6b-45b1-a552-689057d1ad29",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "09",
                "label": "09 (X.121)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "05a7c47a-24f3-4f90-ac6e-7dde0e083690",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "10",
                "label": "10 (DoD - Department of Defense Activity Address Code)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "2cab1b3a-3128-4e41-9e75-85d5bc9dd350",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "11",
                "label": "11 (DEA - Drug Enforcement Administration)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "94cfbc8c-efbf-444a-a452-98fed4932dbd",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "13",
                "label": "13 (UCS Code)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "b606f1cc-c717-4389-9800-556993e3a943",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "16",
                "label": "16 (Duns Number With 4-Character Suffix)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "ba360720-ac58-4823-a324-649c78d15223",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "19",
                "label": "19 (EDI Council of Australia Communications ID)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "2d3cd43f-e6f8-4967-9083-7b0b4f2122e4",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "20",
                "label": "20 - (HIN - Health Industry Number)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "e6b315d7-6335-4e53-95de-996ddbfcc95c",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "30",
                "label": "30 (U.S. Federal Tax Identification Number)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "e5e79171-07cd-413b-a9db-37f276ab0dfa",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "NR",
                "label": "NR (NRMA - National Retail Merchants Association)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "447b44ac-977d-4dca-a49a-59bf56e65c8d",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "SN",
                "label": "SN (Standard Address Number)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            },
            {
                "id": "789fb88c-c353-46f6-ba7e-2f9784e5e504",
                "identifierTypeId": "ce7e0612-8913-4bac-a99b-7f7c6c8f456f",
                "code": "08",
                "label": "08 (UCC EDI Communications ID)",
                "segmentIdentifier": null,
                "descritpion": null,
                "environmentId": "00000000-0000-0000-0000-000000000000"
            }
        ]
    }
]`
)

func TestClient_ListIdentifierTypes(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(identifierTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.ListIdentifierTypes()

	assert.Nil(t, err)
	assert.NotNil(t, identifiers)
	assert.Equal(t, 2, len(identifiers))
}

func TestClient_GetIdentifierTypesByName(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(identifierTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.GetIdentifierTypesByName("AS2")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(identifiers))
	assert.NotNil(t, (*identifiers[0]).Id)
	assert.Equal(t, "65715d63-b628-4eb0-913d-cfdb55057c15", *(*identifiers[0]).Id)
}

func TestClient_GetIdentifierTypesByLabel(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(identifierTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.GetIdentifierTypesByLabel("AS2 Identity")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(identifiers))
	assert.NotNil(t, (*identifiers[0]).Id)
	assert.Equal(t, "65715d63-b628-4eb0-913d-cfdb55057c15", *(*identifiers[0]).Id)
}

func TestClient_GetIdentifierTypesById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(identifierTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.GetIdentifierTypesById("65715d63-b628-4eb0-913d-cfdb55057c15")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(identifiers))
	assert.NotNil(t, (*identifiers[0]).Name)
	assert.Equal(t, "AS2", *(*identifiers[0]).Name)
}

func TestIdentifier_GetIdentifierQualifierByCode(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(identifierTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.GetIdentifierTypesByName("X12-ISA")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(identifiers))
	assert.Equal(t, "ce7e0612-8913-4bac-a99b-7f7c6c8f456f", *(*identifiers[0]).Id)
	assert.Equal(t, 18, len((*identifiers[0]).Qualifiers))

	qualifiers, err := identifiers[0].GetIdentifierTypeQualifiersByCode("30")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(qualifiers))
	assert.NotNil(t, qualifiers[0])
	assert.NotNil(t, *qualifiers[0].Id)
	assert.Equal(t, "e6b315d7-6335-4e53-95de-996ddbfcc95c", *(*qualifiers[0]).Id)
}

func TestIdentifier_GetIdentifierQualifierByLabel(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(identifierTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.GetIdentifierTypesByName("X12-ISA")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(identifiers))
	assert.Equal(t, "ce7e0612-8913-4bac-a99b-7f7c6c8f456f", *(*identifiers[0]).Id)
	assert.Equal(t, 18, len((*identifiers[0]).Qualifiers))

	qualifiers, err := identifiers[0].GetIdentifierTypeQualifiersByLabel("30 (U.S. Federal Tax Identification Number)")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(qualifiers))
	assert.NotNil(t, qualifiers[0])
	assert.NotNil(t, *qualifiers[0].Id)
	assert.Equal(t, "e6b315d7-6335-4e53-95de-996ddbfcc95c", *(*qualifiers[0]).Id)
}

func TestIdentifier_GetIdentifierQualifierById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, identifierTestUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(identifierTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	assert.Nil(t, err)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	identifiers, err := cli.GetIdentifierTypesByName("X12-ISA")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(identifiers))
	assert.Equal(t, "ce7e0612-8913-4bac-a99b-7f7c6c8f456f", *(*identifiers[0]).Id)
	assert.Equal(t, 18, len((*identifiers[0]).Qualifiers))

	qualifiers, err := identifiers[0].GetIdentifierTypeQualifiersById("e6b315d7-6335-4e53-95de-996ddbfcc95c")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(qualifiers))
	assert.NotNil(t, qualifiers[0])
	assert.NotNil(t, *qualifiers[0].Code)
	assert.Equal(t, "30", *(*qualifiers[0]).Code)
}
