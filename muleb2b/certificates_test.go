package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	baseCertificatesTestUrl       = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/partners/9e39423e-da65-429a-b735-941e4b3fd350/certificates"
	listCertificatesTestResponse  = `[{"id":"89e81a66-9fd9-476c-9390-df15c4c694d1","certificateType":"PEM","name":"Certificate-1","csmSecret":{"secretId":"998c2ebb-f5bf-4026-a4c1-c935d08edff9","secretGroupId":"b0c91b69-27c8-4a36-831e-b5e5f18e2616","grantResponse":{"path":"certificates/998c2ebb-f5bf-4026-a4c1-c935d08edff9","accessGrant":""}},"authority":"tst","serialNumber":"F414286E8C8E5108","startDate":"2020-02-27T18:42:11Z","expires":"2021-02-26T18:42:11Z","partnerId":"9e39423e-da65-429a-b735-941e4b3fd350","usedInEndpoints":0,"flowDependencyCount":0,"usedInAs2":true,"runtimeStatus":null},{"id":"3f18450e-d8ea-4088-beb6-3591bce7dbb4","certificateType":"PEM","name":"Certificate-2","csmSecret":{"secretId":"91ea8c05-6474-4b7a-8e5b-4bda0efe66f1","secretGroupId":"b0c91b69-27c8-4a36-831e-b5e5f18e2616","grantResponse":{"path":"certificates/91ea8c05-6474-4b7a-8e5b-4bda0efe66f1","accessGrant":""}},"authority":"ts","serialNumber":"DD5D816942935AEF","startDate":"2020-02-27T18:43:55Z","expires":"2021-02-26T18:43:55Z","partnerId":"9e39423e-da65-429a-b735-941e4b3fd350","usedInEndpoints":0,"flowDependencyCount":0,"usedInAs2":false,"runtimeStatus":null}]`
	singleCertificateTestUrl      = "/partnermanager/partners/api/v1/organizations/be4f0fba-541b-5f82-b51d-f047b6569645/environments/3a4d3936-22d9-4d87-a3c0-a8d424bcc032/partners/9e39423e-da65-429a-b735-941e4b3fd350/certificates/89e81a66-9fd9-476c-9390-df15c4c694d1"
	getCertificateTestResponse    = `{"id":"89e81a66-9fd9-476c-9390-df15c4c694d1","certificateType":"PEM","name":"Certificate-1","csmSecret":{"secretId":"998c2ebb-f5bf-4026-a4c1-c935d08edff9","secretGroupId":"b0c91b69-27c8-4a36-831e-b5e5f18e2616","grantResponse":{"path":"certificates/998c2ebb-f5bf-4026-a4c1-c935d08edff9","accessGrant":""}},"authority":"tst","serialNumber":"F414286E8C8E5108","startDate":"2020-02-27T18:42:11Z","expires":"2021-02-26T18:42:11Z","partnerId":"9e39423e-da65-429a-b735-941e4b3fd350","usedInEndpoints":0,"flowDependencyCount":0,"usedInAs2":true,"runtimeStatus":null}`
	createCertificateTestResponse = `{"id":"3f18450e-d8ea-4088-beb6-3591bce7dbb4"}`
	demoCertContent               = `-----BEGIN CERTIFICATE-----
MIIFSjCCAzICCQD0FChujI5RCDANBgkqhkiG9w0BAQsFADBnMQswCQYDVQQGEwJ1
czELMAkGA1UECAwCbW4xDTALBgNVBAcMBG1wbHMxDDAKBgNVBAoMA3RzdDEMMAoG
A1UECwwDdHN0MQwwCgYDVQQDDAN0c3QxEjAQBgkqhkiG9w0BCQEWA3N0YTAeFw0y
MDAyMjcxODQyMTFaFw0yMTAyMjYxODQyMTFaMGcxCzAJBgNVBAYTAnVzMQswCQYD
VQQIDAJtbjENMAsGA1UEBwwEbXBsczEMMAoGA1UECgwDdHN0MQwwCgYDVQQLDAN0
c3QxDDAKBgNVBAMMA3RzdDESMBAGCSqGSIb3DQEJARYDc3RhMIICIjANBgkqhkiG
9w0BAQEFAAOCAg8AMIICCgKCAgEAsUVj2jmsxeVlR4IRfg8WxjOdKMataqIsEzu6
asx4gX/CrCdrGgJjuYiNSL9cxCuKH+cZ0WJev4a9d3DjIos4T4Vtv6clcV1/fmM7
r/pjZY1+6mYcTZExSA/kdIoKS1RGBLqFM6U9fZy7dH4/VUsgnQ2rVe/RShjaokso
MyKOAcy8qa+pzEHQZkhbgJpzhxA5D9eL2xqmC6fRj0cvZdOoXzvlLAaDwekWWLmy
eA54vM7I1r8BqtSOwWaYQDJZvqnOLNYtgHGr9L+o4Ate9rHTt6agYOMguXSmHbuZ
w9UxmPjqBNa26moq0vBTPUkO//8XU3x3pjyRMhBy9+rkGvC85Fv5iUbqjKMUL7E8
gAZ2UTiYyzhYBLaQ291+hpAICxLvTLmnxUIpIVJZFbvtLknwXktKA4Qppi8/zPo1
N8QGL1Ly9NSfO9Xr4RNlYQ6vEdazDuN8CXFe9W9tRrD9Y0XYpkkOvicbrvw4UNRC
hWjfCoa+UzucmfqYRP03TTXO89sFaG+2I6S+P84sD3qQ84XbG42LFjRKaESyfNsq
vosmW2dvS8rZfTK+3HmhCWCcBYrGyCHR3R1FgXC18Dq3qXhWr7/g1waIN7WhtBgP
irC6/t3ftfVsGDafzzmrvQ/y0TFt2KkiIu0cZPcpMlw6CjVJRF6uNRMJCCWsBOGX
fFi/p1cCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEAf3ScwbiZMuBEGBFlvAdwROPe
B13k3A7ou0VClcslxO8FAkCtH9GOO9XU/iMDz1FUBtM59TlADo8tW1cLIfZ1bQ5z
Vs2WOF18eGPkTbx+uJZ9yUhK7q3ha0QaBop6CU/krT4t/OIHSw1YR2MClmEOOKuE
srieswPG0YU4w6U6GdfUv9i3Tzxht+C8ETUAaxFeOkVKPxqx0g8oqKZXK1BwZIau
BNRVOMWDUy81EkfdsZ63UwAJcUkxSfkdyoMMKcFIkz2bAozbi8xSCw6PQ02T+FYG
PgNNPcr5h4KuSLRVP5rTHiP0zVYkJ56LBaQ8V7EEcRquUCT+E7T/JTJw5dKR2bt3
GKjq93O/rMH//W+JSMDFd1YiqWPCe+hPxJtApA+z/39E31OmnK29yIKMTteRAEOV
GRB+S387t0DcmFrZkg81Fufwg9Zx5ZyYkar/ePR3h9lZEaXOM9VFVlC9UTrfdF84
sQjP3jyeddplRXqADT8x+4iRCjaWo5cQZBu+WPjeY9sSIuscfsk+PS0eG+5+LekK
AnQpD6g2QPpGY2x/3moAnS9IAnxhi3SxxwfLdOrhKMIoIuZaWex/Tbp79U4wblSa
HTUSfNOJpiUGKNGZ4wVNoa6Qki5fH1pWlGE+HTcizBYW/Sv5mHgNIbjOTiPsVvZd
axlm6QIk1IMrzP7+XQo=
-----END CERTIFICATE-----`
)

func TestListPartnerCertificates(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, baseCertificatesTestUrl, r.URL.String())
		assert.True(t, r.Method == "GET")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(listCertificatesTestResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")
	assert.Nil(t, err)

	certificates, err := cli.ListPartnerCertificates("9e39423e-da65-429a-b735-941e4b3fd350")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(certificates))
}

func TestGetPartnerCertificate(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, singleCertificateTestUrl, r.URL.String())
		assert.True(t, r.Method == "GET")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(getCertificateTestResponse))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")
	assert.Nil(t, err)

	certificate, err := cli.GetPartnerCertificate("9e39423e-da65-429a-b735-941e4b3fd350", "89e81a66-9fd9-476c-9390-df15c4c694d1")
	assert.Nil(t, err)
	assert.NotNil(t, certificate)
	assert.NotNil(t, certificate.Name)
	assert.Equal(t, "Certificate-1", *certificate.Name)
}

func TestDeletePartnerCertificate(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, singleCertificateTestUrl, r.URL.String())
		assert.Equal(t, "DELETE", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")
	assert.Nil(t, err)

	err = cli.DeletePartnerCertificate("9e39423e-da65-429a-b735-941e4b3fd350", "89e81a66-9fd9-476c-9390-df15c4c694d1")
	assert.Nil(t, err)
}

func TestCreatePartnerCertificate(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, baseCertificatesTestUrl, r.URL.String())
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(createCertificateTestResponse))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)
	cli.SetEnvironment("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")
	assert.Nil(t, err)

	id, err := cli.CreatePartnerCertificate("9e39423e-da65-429a-b735-941e4b3fd350", demoCertContent, "Certificate-1", "PEM")
	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, "3f18450e-d8ea-4088-beb6-3591bce7dbb4", *id)
}
