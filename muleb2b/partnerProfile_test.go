package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	partnerProfileTestUrl     = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/partnerprofiles/67d0c3e9-6542-437b-9042-7eee278914f0"
	getPartnerProfileResponse = `{"logo":null,"contacts":[{"id":"2bf070b0-b767-4f0c-bb40-79b197b9184a","name":"John Smith","email":"john.smith@example.com","phoneNumber":"9441231234","status":"ACTIVE","contactType":{"id":"50615d27-0356-47fc-bd1e-440c992cd98e","name":"Business","label":"Business Contact","description":"Business Contact"}},{"id":"8ccd553c-596d-4bd8-b0fd-5aaf6ba17a8b","name":"Jim Beam","email":"jim.beam@example.com","phoneNumber":"666123123","status":"ACTIVE","contactType":{"id":"020f4c28-a0c2-4e70-b25d-8ab68f1a2020","name":"Technical","label":"Technical Contact","description":"Technical Contact"}},{"id":"3eee37e1-e675-4047-b7e9-047e84eb7927","name":"Bob","email":"bob@example.com","phoneNumber":"1234123456","status":"ACTIVE","contactType":{"id":"50615d27-0356-47fc-bd1e-440c992cd98e","name":"Business","label":"Business Contact","description":"Business Contact"}}],"identifiers":[{"id":"9467e501-06a6-4587-bd0a-f568283a38ea","identifierTypeQualifierId":"3ca16584-94b7-4d97-b57c-e5d6600bd441","status":"ACTIVE","qualifierLabel":"DUNS Number","typeLabel":"DUNS","code":"DUNS","value":"asdf"}],"addresses":[{"id":"0512296a-39a5-4470-8535-ee500f108202","addressLine_1":"123 Main ST","addressLine_2":"Suite 1","city":"Minneapolis","state":"MN","country":"US","postalCode":"55401"}],"createdAt":"2020-03-20T17:59:29Z","createdBy":{"userId":"383c750f-40c7-4da0-8f0e-5d0be135532e","firstName":"test","lastName":"test","userName":"test"},"updatedAt":"2020-03-20T20:58:57Z","updatedBy":{"userId":"383c750f-40c7-4da0-8f0e-5d0be135532e","firstName":"test","lastName":"test","userName":"test"},"protocols":[],"standards":[],"usedInDeployments":{"inboundCount":0,"outboundCount":0},"id":"67d0c3e9-6542-437b-9042-7eee278914f0","name":"Sample","description":null,"websiteUrl":null,"status":{"id":"bbc0557f-d43d-4e2b-9f38-895ccbc2b063","startDate":null,"endDate":null,"status":"ACTIVE"},"environmentId":"3a4d3936-22d9-4d87-a3c0-a8d424bcc032","hostFlag":false,"dependencyCounts":null}`
)

func TestGetPartnerProfile(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, partnerProfileTestUrl, r.URL.String())
		assert.True(t, r.Method == "GET")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getPartnerProfileResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")
	assert.Nil(t, err)

	profile, err := cli.GetPartnerProfile("67d0c3e9-6542-437b-9042-7eee278914f0")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(profile.Contacts))
}

func TestUpdatePartnerProfile(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, partnerProfileTestUrl, r.URL.String())
		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(getPartnerProfileResponse))
		} else if r.Method == "PATCH" {
			w.WriteHeader(http.StatusNoContent)
		}

	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")
	assert.Nil(t, err)

	profile, err := cli.GetPartnerProfile("67d0c3e9-6542-437b-9042-7eee278914f0")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(profile.Contacts))

	profile.WebsiteUrl = String("test.com")

	err = cli.UpdatePartnerProfile(profile)
	assert.Nil(t, err)
}
