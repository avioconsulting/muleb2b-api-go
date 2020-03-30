package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	listMappingsExpectedUrl  = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/documentflows/b594b3d2-2317-4617-a70c-cb190beab0c9/mappings"
	singleMappingExpectedUrl = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/documentflows/b594b3d2-2317-4617-a70c-cb190beab0c9/mappings/c6b67e55-a8cc-48f0-a460-be58c14d7c0e"
	listMappingsResp         = `[
    {
        "id": "584bc4f4-bcc5-4b4a-b763-b6416d9c2065",
        "mappingType": "DWL_FILE",
        "mappingContent": "data:application/octet-stream;base64,JWR3IDIuMA0Kb3V0cHV0IGFwcGxpY2F0aW9uL2pzb24NCi0tLSANCnBheWxvYWQ=",
        "mappingSourceRef": "anything_to_json.dwl"
    }
]`
	getMappingResp = `{
    "id": "584bc4f4-bcc5-4b4a-b763-b6416d9c2065",
    "mappingType": "DWL_FILE",
    "mappingContent": "data:application/octet-stream;base64,JWR3IDIuMA0Kb3V0cHV0IGFwcGxpY2F0aW9uL2pzb24NCi0tLSANCnBheWxvYWQ=",
    "mappingSourceRef": "anything_to_json.dwl"
}`
)

func TestClient_ListMappings(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, listMappingsExpectedUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(listMappingsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	mappings, err := cli.ListMappings("b594b3d2-2317-4617-a70c-cb190beab0c9")

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*mappings))
}

func TestClient_GetMapping(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, singleMappingExpectedUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getMappingResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	mapping, err := cli.GetMappingById("b594b3d2-2317-4617-a70c-cb190beab0c9", "c6b67e55-a8cc-48f0-a460-be58c14d7c0e")

	assert.Nil(t, err)
	assert.Equal(t, "anything_to_json.dwl", *mapping.MappingSourceRef)
}

func TestClient_CreateMapping(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, listMappingsExpectedUrl, r.URL.String())
		w.WriteHeader(http.StatusCreated)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	mapping := Mapping{
		MappingType:      String("DWL_FILE"),
		MappingContent:   String("data:application/octet-stream;base64,JWR3IDIuMA0Kb3V0cHV0IGFwcGxpY2F0aW9uL2pzb24NCi0tLSANCnBheWxvYWQ="),
		MappingSourceRef: String("anything_to_json.dwl"),
	}

	cli.CreateMapping("b594b3d2-2317-4617-a70c-cb190beab0c9", &mapping)

	assert.Nil(t, err)
}

func TestClient_UpdateMapping(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, singleMappingExpectedUrl, r.URL.String())
		w.WriteHeader(http.StatusCreated)
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	mapping := Mapping{
		Id:               String("c6b67e55-a8cc-48f0-a460-be58c14d7c0e"),
		MappingType:      String("DWL_FILE"),
		MappingContent:   String("data:application/octet-stream;base64,JWR3IDIuMA0Kb3V0cHV0IGFwcGxpY2F0aW9uL2pzb24NCi0tLSANCnBheWxvYWQ="),
		MappingSourceRef: String("anything_to_json.dwl"),
	}

	cli.UpdateMapping("b594b3d2-2317-4617-a70c-cb190beab0c9", &mapping)

	assert.Nil(t, err)
}
