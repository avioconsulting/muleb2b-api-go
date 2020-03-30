package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

const (
	listDocumentsResponse = `[
    {
        "id": "dc25e9c2-1366-4fbf-a6b8-072f8d843736",
        "name": "X12-3010-810",
        "isStandard": true,
        "ediDocumentTypeId": "6117a01c-d661-4517-80a5-5a6fe08d833c",
        "baseType": {
            "id": "6117a01c-d661-4517-80a5-5a6fe08d833c",
            "documentName": "810",
            "label": "810",
            "schemaPath": "/schemas/x12/3010/810.esl",
            "description": "",
            "formatType": "X12",
            "version": "v003010",
            "versionLabel": "3010",
            "ediFormatVersionId": "7f9d1dcc-6262-43b8-8990-26a4f2787aec"
        }
    },
    {
        "customSchemaId": "7be3fd9a-4ecb-4435-b57a-b457e37020c8",
        "schemaType": "customSchemaType",
        "schemaContent": "aaaaaaaa",
        "id": "95ec0efa-f149-4dd5-a58d-3d7e5a6212e4",
        "name": "X12-3010-819",
        "isStandard": false,
        "ediDocumentTypeId": "f4190af3-a1ee-4aa7-8a1c-67d385b57c90",
        "baseType": {
            "id": "f4190af3-a1ee-4aa7-8a1c-67d385b57c90",
            "documentName": "819",
            "label": "819",
            "schemaPath": "/schemas/x12/3010/819.esl",
            "description": "",
            "formatType": "X12",
            "version": "v003010",
            "versionLabel": "3010",
            "ediFormatVersionId": "7f9d1dcc-6262-43b8-8990-26a4f2787aec"
        }
    }
]`
	getDocumentByIdResponse = `{
    "id": "af130dd2-09ee-48d9-bcc9-0347af6872b1",
    "name": "MySampleDocument",
    "isStandard": true,
    "ediDocumentTypeId": "6117a01c-d661-4517-80a5-5a6fe08d833c",
    "baseType": {
        "id": "6117a01c-d661-4517-80a5-5a6fe08d833c",
        "documentName": "810",
        "label": "810",
        "schemaPath": "/schemas/x12/3010/810.esl",
        "description": "",
        "formatType": "X12",
        "version": "v003010",
        "versionLabel": "3010",
        "ediFormatVersionId": "7f9d1dcc-6262-43b8-8990-26a4f2787aec"
    }
}`
	createDocResponse = `{
    "id": "af130dd2-09ee-48d9-bcc9-0347af6872b1"
}`
)

func TestClient_ListDocuments(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(listDocumentsResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	docs, err := cli.ListDocuments("ae5e69af-9de2-4355-8502-f2a330326ad2")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*docs))
}

func TestClient_GetDocumentByName(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(listDocumentsResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	doc, err := cli.GetDocumentByName("ae5e69af-9de2-4355-8502-f2a330326ad2", "X12-3010-810")

	assert.Nil(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, "dc25e9c2-1366-4fbf-a6b8-072f8d843736", *doc.Id)
}

func TestClient_GetDocumentById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID
		assert.True(t, strings.Contains(r.URL.String(), "af130dd2-09ee-48d9-bcc9-0347af6872b1")) // Document ID

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getDocumentByIdResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	doc, err := cli.GetDocumentById("ae5e69af-9de2-4355-8502-f2a330326ad2", "af130dd2-09ee-48d9-bcc9-0347af6872b1")

	assert.Nil(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, "af130dd2-09ee-48d9-bcc9-0347af6872b1", *doc.Id)
}

func TestClient_CreateDocument(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(createDocResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	doc := Document{
		Name:              String("Test"),
		Standard:          Boolean(true),
		EdiDocumentTypeId: String("6117a01c-d661-4517-80a5-5a6fe08d833c"),
	}

	id, err := cli.CreateDocument("ae5e69af-9de2-4355-8502-f2a330326ad2", &doc)

	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, "af130dd2-09ee-48d9-bcc9-0347af6872b1", *id)
}

func TestClient_UpdateDocument(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID
		assert.True(t, strings.Contains(r.URL.String(), "af130dd2-09ee-48d9-bcc9-0347af6872b1")) // Document ID

		w.WriteHeader(http.StatusNoContent)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	doc := Document{
		Id:                String("af130dd2-09ee-48d9-bcc9-0347af6872b1"),
		Name:              String("Test"),
		Standard:          Boolean(true),
		EdiDocumentTypeId: String("6117a01c-d661-4517-80a5-5a6fe08d833c"),
	}

	err = cli.UpdateDocument("ae5e69af-9de2-4355-8502-f2a330326ad2", &doc)

	assert.Nil(t, err)
}

func TestClient_DeleteDocument(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645")) // Org ID
		assert.True(t, strings.Contains(r.URL.String(), "3a4d3936-22d9-4d87-a3c0-a8d424bcc032")) // Environment ID
		assert.True(t, strings.Contains(r.URL.String(), "ae5e69af-9de2-4355-8502-f2a330326ad2")) // Partner ID
		assert.True(t, strings.Contains(r.URL.String(), "af130dd2-09ee-48d9-bcc9-0347af6872b1")) // Document ID

		w.WriteHeader(http.StatusNoContent)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	doc := Document{
		Id:                String("af130dd2-09ee-48d9-bcc9-0347af6872b1"),
		Name:              String("Test"),
		Standard:          Boolean(true),
		EdiDocumentTypeId: String("6117a01c-d661-4517-80a5-5a6fe08d833c"),
	}

	err = cli.DeleteDocument("ae5e69af-9de2-4355-8502-f2a330326ad2", &doc)

	assert.Nil(t, err)
}
