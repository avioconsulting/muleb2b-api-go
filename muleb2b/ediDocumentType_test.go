package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	ediDocTypesResp = `[
    {
        "id": "6117a01c-d661-4517-80a5-5a6fe08d833c",
        "documentName": "810",
        "label": "810",
        "schemaPath": "/schemas/x12/3010/810.esl",
        "description": "",
        "formatType": null,
        "version": null,
        "versionLabel": null,
        "ediFormatVersionId": "7f9d1dcc-6262-43b8-8990-26a4f2787aec"
    },
    {
        "id": "f4190af3-a1ee-4aa7-8a1c-67d385b57c90",
        "documentName": "819",
        "label": "819",
        "schemaPath": "/schemas/x12/3010/819.esl",
        "description": "",
        "formatType": null,
        "version": null,
        "versionLabel": null,
        "ediFormatVersionId": "7f9d1dcc-6262-43b8-8990-26a4f2787aec"
    },
    {
        "id": "0c42aebe-113c-4691-88b2-036555f86dd6",
        "documentName": "820",
        "label": "820",
        "schemaPath": "/schemas/x12/3010/820.esl",
        "description": "",
        "formatType": null,
        "version": null,
        "versionLabel": null,
        "ediFormatVersionId": "7f9d1dcc-6262-43b8-8990-26a4f2787aec"
    }
]`
	ediDocTypesUrl = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/ediFormats/X12/ediFormatVersions/7f9d1dcc-6262-43b8-8990-26a4f2787aec/ediDocumentTypes"
)

func TestClient_ListEdiDocumentTypes(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediDocTypesUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediDocTypesResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	docTypes, err := cli.ListEdiDocumentTypes("X12", "7f9d1dcc-6262-43b8-8990-26a4f2787aec")

	assert.Nil(t, err)
	assert.Equal(t, 3, len(*docTypes))
}

func TestClient_GetEdiDocumentTypeByLabel(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediDocTypesUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediDocTypesResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	docType, err := cli.GetEdiDocumentTypeByLabel("X12", "7f9d1dcc-6262-43b8-8990-26a4f2787aec", "810")

	assert.Nil(t, err)
	assert.NotNil(t, docType)
	assert.NotNil(t, docType.Id)
	assert.Equal(t, "6117a01c-d661-4517-80a5-5a6fe08d833c", *docType.Id)

}

func TestClient_GetEdiDocumentTypeByName(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediDocTypesUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediDocTypesResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	docType, err := cli.GetEdiDocumentTypeByName("X12", "7f9d1dcc-6262-43b8-8990-26a4f2787aec", "810")

	assert.Nil(t, err)
	assert.NotNil(t, docType)
	assert.NotNil(t, docType.Id)
	assert.Equal(t, "6117a01c-d661-4517-80a5-5a6fe08d833c", *docType.Id)

}

func TestClient_GetEdiDocumentTypeById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediDocTypesUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediDocTypesResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	docType, err := cli.GetEdiDocumentTypeById("X12", "7f9d1dcc-6262-43b8-8990-26a4f2787aec", "f4190af3-a1ee-4aa7-8a1c-67d385b57c90")

	assert.Nil(t, err)
	assert.NotNil(t, docType)
	assert.NotNil(t, docType.Label)
	assert.Equal(t, "819", *docType.Label)

}
