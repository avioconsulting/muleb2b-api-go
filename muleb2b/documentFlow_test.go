package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	getDocFlowsExpectedUrl = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/documentflows/summarylist"
	docFlowsExpectedUrl    = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/documentflows"
	singleDocFlowExpected  = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/documentflows/a23f0680-f6ff-4047-b631-714c4cc79421"
	listDocFlowsResp       = `[
    {
        "id": "028ccd8d-8ea7-42f7-ace9-5844a3f4331e",
        "name": "Untitled message flow",
        "direction": "INBOUND",
        "partnerFrom": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2020-02-20T16:13:29Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2020-02-20T16:13:29Z",
            "updatedBy": "",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "b60ec9cd-2e40-402f-a90c-1e67e75a8b31",
            "name": "SampleABCD",
            "description": null,
            "websiteUrl": null,
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": false,
            "dependencyCounts": null
        },
        "partnerTo": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2019-11-20T21:23:13Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2019-11-20T21:23:13Z",
            "updatedBy": "",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "9e39423e-da65-429a-b735-941e4b3fd350",
            "name": "Test-2",
            "description": null,
            "websiteUrl": null,
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": true,
            "dependencyCounts": null
        },
        "createdAt": "2020-02-20T16:13:30Z",
        "createdBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "updatedAt": "2020-02-20T16:15:09Z",
        "updatedBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "deploymentSummary": [],
        "lastDeployTime": null,
        "lastDeployBy": null,
        "lastDeployStatus": null
    },
    {
        "id": "a23f0680-f6ff-4047-b631-714c4cc79421",
        "name": "X12-3010-819",
        "direction": "INBOUND",
        "partnerFrom": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2019-11-21T22:22:44Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2019-11-21T22:22:44Z",
            "updatedBy": "",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "ae5e69af-9de2-4355-8502-f2a330326ad2",
            "name": "Sample2",
            "description": null,
            "websiteUrl": null,
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": false,
            "dependencyCounts": null
        },
        "partnerTo": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2019-11-20T21:23:13Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2019-11-20T21:23:13Z",
            "updatedBy": "",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "9e39423e-da65-429a-b735-941e4b3fd350",
            "name": "Test-2",
            "description": null,
            "websiteUrl": null,
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": true,
            "dependencyCounts": null
        },
        "createdAt": "2020-01-16T18:41:44Z",
        "createdBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "updatedAt": "2020-02-17T20:02:27Z",
        "updatedBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "deploymentSummary": [],
        "lastDeployTime": null,
        "lastDeployBy": null,
        "lastDeployStatus": null
    },
    {
        "id": "8dcc492a-7da4-46c8-b56e-eea5ae70748f",
        "name": "Untitled message flow",
        "direction": "INBOUND",
        "partnerFrom": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2019-11-21T18:27:34Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2019-11-21T18:53:40Z",
            "updatedBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "5d6db9ed-e7ec-4c0a-9ff5-f9c3ab56d8ea",
            "name": "Sample",
            "description": "Some description",
            "websiteUrl": "http://sample.com",
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": false,
            "dependencyCounts": null
        },
        "partnerTo": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2019-11-20T21:23:13Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2019-11-20T21:23:13Z",
            "updatedBy": "",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "9e39423e-da65-429a-b735-941e4b3fd350",
            "name": "Test-2",
            "description": null,
            "websiteUrl": null,
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": true,
            "dependencyCounts": null
        },
        "createdAt": "2020-01-16T22:28:08Z",
        "createdBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "updatedAt": "2020-01-16T22:28:08Z",
        "updatedBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "deploymentSummary": [],
        "lastDeployTime": null,
        "lastDeployBy": null,
        "lastDeployStatus": null
    },
    {
        "id": "b594b3d2-2317-4617-a70c-cb190beab0c9",
        "name": "X12-3010-810",
        "direction": "INBOUND",
        "partnerFrom": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2019-11-21T18:27:34Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2019-11-21T18:53:40Z",
            "updatedBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "5d6db9ed-e7ec-4c0a-9ff5-f9c3ab56d8ea",
            "name": "Sample",
            "description": "Some description",
            "websiteUrl": "http://sample.com",
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": false,
            "dependencyCounts": null
        },
        "partnerTo": {
            "logo": null,
            "contacts": [],
            "identifiers": [],
            "addresses": [],
            "createdAt": "2019-11-20T21:23:13Z",
            "createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "updatedAt": "2019-11-20T21:23:13Z",
            "updatedBy": "",
            "protocols": [],
            "standards": [],
            "usedInDeployments": null,
            "id": "9e39423e-da65-429a-b735-941e4b3fd350",
            "name": "Test-2",
            "description": null,
            "websiteUrl": null,
            "status": null,
            "environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
            "hostFlag": true,
            "dependencyCounts": null
        },
        "createdAt": "2019-11-21T20:37:37Z",
        "createdBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "updatedAt": "2020-01-16T18:31:59Z",
        "updatedBy": {
            "userId": "383c750f-40c7-4da0-8f0e-5d0be135532e",
            "firstName": "test",
            "lastName": "test",
            "userName": "test"
        },
        "deploymentSummary": [],
        "lastDeployTime": null,
        "lastDeployBy": null,
        "lastDeployStatus": null
    }
]`
	createDocFlowResp = `{"id":"test"}`
	updateDocFlowResp = `{
	"id": "a23f0680-f6ff-4047-b631-714c4cc79421",
	"name": "X12-3010-810",
	"direction": "INBOUND",
	"partnerFromId": "ae5e69af-9de2-4355-8502-f2a330326ad2",
	"partnerToId": "9e39423e-da65-429a-b735-941e4b3fd350",
	"createdAt": "2020-01-16T18:41:44Z",
	"createdBy": "383c750f-40c7-4da0-8f0e-5d0be135532e",
	"modifierAt": null,
	"modifiedBy": null,
	"configurations": [
		{
			"id": "2d4286d2-2c0f-4d7f-9b2e-db079d128248",
			"documentFlowId": "a23f0680-f6ff-4047-b631-714c4cc79421",
			"environmentId": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
			"status": "DRAFT",
			"version": 4,
			"preProcessingEndpointId": null,
			"receivingEndpointId": "4218d516-bcd7-4a9c-89ea-fe73e6538e5f",
			"receivingAckEndpointId": null,
			"targetEndpointId": "e6420269-6fb7-4daf-bdeb-43e8acc39046",
			"sourceDocTypeId": "dc25e9c2-1366-4fbf-a6b8-072f8d843736",
			"targetDocTypeId": "8d42a6ec-e49c-4f32-88a6-8d06ae1cbdc0",
			"documentMapping": [
				{
					"id": "584bc4f4-bcc5-4b4a-b763-b6416d9c2065",
					"mappingType": "DWL_FILE",
					"mappingContent": null,
					"mappingSourceRef": "anything_to_json.dwl"
				}
			],
			"receivingAckConfig": null
		}
	],
	"last_deploy_status": null,
	"last_deploy_message": null
}`
)

func TestClient_ListDocumentFlows(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, getDocFlowsExpectedUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(listDocFlowsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	docFlows, err := cli.ListDocumentFlows()

	assert.Nil(t, err)
	assert.Equal(t, 4, len(*docFlows))
}

func TestClient_GetDocumentFlowByName(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, getDocFlowsExpectedUrl, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(listDocFlowsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	docFlow, err := cli.GetDocumentFlowByName("X12-3010-819")

	assert.Nil(t, err)
	assert.NotNil(t, docFlow)
	assert.Equal(t, "a23f0680-f6ff-4047-b631-714c4cc79421", *(*docFlow).Id)
}

func TestClient_CreateDocumentFlow(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, docFlowsExpectedUrl, r.URL.String())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(createDocFlowResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	flow := DocumentFlow{
		Name:          String("X12-Test"),
		Direction:     String("INBOUND"),
		PartnerFromId: String("partnerFrom"),
		PartnerToId:   String("partnerTo"),
	}

	flowId, err := cli.CreateDocumentFlow(&flow)
	assert.Nil(t, err)
	assert.NotNil(t, flowId)
	assert.Equal(t, "test", *flowId)
}

func TestClient_UpdateDocumentFlow(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, singleDocFlowExpected, r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(updateDocFlowResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	docFlowConfig := DocumentFlowConfiguration{

		Id:                  String("2d4286d2-2c0f-4d7f-9b2e-db079d128248"),
		DocumentFlowId:      String("test"),
		EnvironmentId:       cli.envId,
		Version:             Integer(3),
		ReceivingEndpointId: String("tests"),
		TargetEndpointId:    String("asdf"),
	}

	flow := DocumentFlow{
		Id:            String("a23f0680-f6ff-4047-b631-714c4cc79421"),
		Name:          String("X12-Test"),
		Direction:     String("INBOUND"),
		PartnerFromId: String("partnerFrom"),
		PartnerToId:   String("partnerTo"),
		Configurations: []*DocumentFlowConfiguration{
			&docFlowConfig,
		},
	}

	updatedFlow, err := cli.UpdateDocumentFlow(&flow)
	assert.Nil(t, err)
	assert.NotNil(t, updatedFlow)
	assert.Equal(t, "2d4286d2-2c0f-4d7f-9b2e-db079d128248", *updatedFlow.Configurations[0].Id)
}

func TestClient_DeleteDocumentFlow(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, singleDocFlowExpected, r.URL.String())
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(listDocFlowsResp))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)

	err = cli.DeleteDocumentFlow("a23f0680-f6ff-4047-b631-714c4cc79421")

	assert.Nil(t, err)
}
