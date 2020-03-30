package muleb2b

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

const (
	envResponse = `{
  "data": [
    {
      "id": "9319514f-e3aa-4c42-bc48-6fc232afb1f3",
      "name": "Design",
      "organizationId": "be4f0fba-541b-5f82-b51d-f047b6569645",
      "isProduction": false,
      "type": "design",
      "clientId": "5f4885ee1ee147a49167c37be5e71fa1"
    },
    {
      "id": "3a4d3936-22d9-4d87-a3c0-a8d424bcc032",
      "name": "Sandbox",
      "organizationId": "be4f0fba-541b-5f82-b51d-f047b6569645",
      "isProduction": false,
      "type": "sandbox",
      "clientId": "3e94c12325024413b72baa944a643b9f"
    },
    {
      "id": "3a4d3936-22d9-4d87-a3c0-a8d424bcc033",
      "name": "Production",
      "organizationId": "be4f0fba-541b-5f82-b51d-f047b6569645",
      "isProduction": true,
      "type": "production",
      "clientId": "3e94c12325024413b72baa944a643b9f"
    }
  ],
  "total": 3
}`
)

func TestListEnvironments(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		w.Write([]byte(envResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	envs, err := cli.ListEnvironments()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(*envs))
}

func TestGetEnvironmentByName(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		w.Write([]byte(envResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	env, err := cli.GetEnvironmentByName("Sandbox")

	assert.Nil(t, err)
	assert.NotNil(t, env)
	assert.Equal(t, "3a4d3936-22d9-4d87-a3c0-a8d424bcc032", *(*env).Id)
}

func TestGetEnvironmentByNameEmpty(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		w.Write([]byte(envResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	env, err := cli.GetEnvironmentByName("FakeName")

	assert.Nil(t, err)
	assert.Nil(t, env)
}

func TestGetEnvironmentById(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		w.Write([]byte(envResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	env, err := cli.GetEnvironmentById("3a4d3936-22d9-4d87-a3c0-a8d424bcc032")

	assert.Nil(t, err)
	assert.NotNil(t, env)
	assert.Equal(t, "3a4d3936-22d9-4d87-a3c0-a8d424bcc032", *(*env).Id)
}

func TestGetEnvironmentByIdEmpty(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		w.Write([]byte(envResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	assert.Nil(t, err)

	env, err := cli.GetEnvironmentById("fake-id")

	assert.Nil(t, err)
	assert.Nil(t, env)
}

func TestListEnvironmentsByType(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		w.Write([]byte(envResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	envs, err := cli.ListEnvironmentsByType("design")

	assert.Nil(t, err)
	assert.NotNil(t, envs)
	assert.Equal(t, 1, len(*envs))
	assert.Equal(t, "9319514f-e3aa-4c42-bc48-6fc232afb1f3", *(*envs)[0].Id)
}

func TestListProductionEnvironments(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, strings.Contains(r.URL.String(), "be4f0fba-541b-5f82-b51d-f047b6569645"))
		w.Write([]byte(envResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli, err := NewClient(String("http://devx.anypoint.mulesoft.com/"), String("be4f0fba-541b-5f82-b51d-f047b6569645"), httpClient)

	envs, err := cli.ListProductionEnvironments()

	assert.Nil(t, err)
	assert.NotNil(t, envs)
	assert.Equal(t, 1, len(*envs))
	assert.Equal(t, "3a4d3936-22d9-4d87-a3c0-a8d424bcc033", *(*envs)[0].Id)
}
