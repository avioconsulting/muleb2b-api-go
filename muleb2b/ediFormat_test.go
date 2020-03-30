package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	ediListResp = `[
	{
		"id": "25c1bc8a-801f-4337-a2a6-7721ef971460",
		"formatType": "X12",
		"description": "X12 format type",
		"label": "X12"
	},
	{
		"id": "eac1bc8a-801f-4337-12a6-7721ef971460",
		"formatType": "JSON",
		"description": "Json format type",
		"label": "JSON"
	}
]`
	ediFormatsUrl        = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/ediFormats"
	ediFormatVersiosnUrl = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/ediFormats/X12/ediFormatVersions"

	ediFormatVersionsResp = `[
    {
        "id": "7f9d1dcc-6262-43b8-8990-26a4f2787aec",
        "formatType": "X12",
        "version": "v003010",
        "description": "X12 v.3010",
        "label": "3010"
    },
    {
        "id": "095c7c8d-5279-4fb9-98a6-80fc80878619",
        "formatType": "X12",
        "version": "v003020",
        "description": "X12 v.3020",
        "label": "3020"
    },
    {
        "id": "65d08bcd-6880-4130-9184-6f6361db1d7c",
        "formatType": "X12",
        "version": "v003030",
        "description": "X12 v.3030",
        "label": "3030"
    },
    {
        "id": "bd10f1d5-7572-41cf-8f24-e3ef353fb89c",
        "formatType": "X12",
        "version": "v003040",
        "description": "X12 v.3040",
        "label": "3040"
    }
]
`
)

func TestClient_ListEdiFormats(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatsUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediListResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	formats, err := cli.ListEdiFormats()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*formats))
}

func TestClient_GetEdiFormatByLabel(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatsUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediListResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	format, err := cli.GetEdiFormatByLabel("JSON")

	assert.Nil(t, err)
	assert.NotNil(t, format)
	assert.NotNil(t, format.Id)
	assert.Equal(t, "eac1bc8a-801f-4337-12a6-7721ef971460", *format.Id)
}

func TestClient_GetEdiFormatByFormat(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatsUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediListResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	format, err := cli.GetEdiFormatByFormat("JSON")

	assert.Nil(t, err)
	assert.NotNil(t, format)
	assert.NotNil(t, format.Id)
	assert.Equal(t, "eac1bc8a-801f-4337-12a6-7721ef971460", *format.Id)
}

func TestClient_GetEdiFormatById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatsUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediListResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	format, err := cli.GetEdiFormatById("eac1bc8a-801f-4337-12a6-7721ef971460")

	assert.Nil(t, err)
	assert.NotNil(t, format)
	assert.NotNil(t, format.FormatType)
	assert.Equal(t, "JSON", *format.FormatType)
}

func TestClient_ListEdiFormatVersions(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatVersiosnUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediFormatVersionsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	versions, err := cli.ListEdiFormatVersions("X12")

	assert.Nil(t, err)
	assert.Equal(t, 4, len(*versions))
}

func TestClient_GetEdiFormatVersionByLabel(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatVersiosnUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediFormatVersionsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	version, err := cli.GetEdiFormatVersionByLabel("X12", "3010")

	assert.Nil(t, err)
	assert.NotNil(t, version)
	assert.NotNil(t, version.Id)
	assert.Equal(t, "7f9d1dcc-6262-43b8-8990-26a4f2787aec", *version.Id)
}

func TestClient_GetEdiFormatVersionByVersion(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatVersiosnUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediFormatVersionsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	version, err := cli.GetEdiFormatVersionByVersion("X12", "v003030")

	assert.Nil(t, err)
	assert.NotNil(t, version)
	assert.NotNil(t, version.Id)
	assert.Equal(t, "65d08bcd-6880-4130-9184-6f6361db1d7c", *version.Id)
}

func TestClient_GetEdiFormatVersionById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ediFormatVersiosnUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ediFormatVersionsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	version, err := cli.GetEdiFormatVersionById("X12", "bd10f1d5-7572-41cf-8f24-e3ef353fb89c")

	assert.Nil(t, err)
	assert.NotNil(t, version)
	assert.NotNil(t, version.Label)
	assert.Equal(t, "3040", *version.Label)
}
