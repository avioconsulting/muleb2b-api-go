package muleb2b

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	okResponse = `{
		"access_token": "76e2b486-55b0-4980-8257-15620d8ca9a0",
		"token_type": "bearer",
		"redirectUrl": "/home/"
	}`
)

func TestClientLogin(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req authRequest
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&req)

		assert.Nil(t, err)

		assert.Equal(t, "user", *req.Username)
		assert.Equal(t, "pass", *req.Password)
		w.Write([]byte(okResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	err = cli.Login("user", "pass")

	assert.Nil(t, err)
	assert.Equal(t, "76e2b486-55b0-4980-8257-15620d8ca9a0", *cli.accessToken)
}
