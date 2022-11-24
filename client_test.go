package atomicasset

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_SendError(t *testing.T) {
	client := New("http://0.0.0.0:8080")

	_, err := client.send("GET", "/", nil)

	assert.EqualError(t, err, "Get \"http://0.0.0.0:8080/\": dial tcp 0.0.0.0:8080: connect: connection refused")
}

func TestClient_SendEncodeParametersFail(t *testing.T) {
	client := Client{}

	_, err := client.send("GET", "/", "a string")

	assert.EqualError(t, err, "expects struct input, got string")
}

func TestClient_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		payload := `{
          "success": false,
          "message": "Some internal error"
        }`

		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(500)
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetHealth()

	assert.EqualError(t, err, "API Error: Some internal error")
}

func TestClient_APIErrorEmptyPayload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(404)
		_, err := res.Write([]byte(`{}`))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	health, err := client.GetHealth()

	assert.NoError(t, err)
	assert.Equal(t, 404, health.HTTPStatusCode)
}

func TestClient_ErrorNoPayload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(200)
		_, err := res.Write([]byte{})
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetHealth()

	assert.EqualError(t, err, "unexpected end of JSON input")
}

func TestClient_HostHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "my-custom-host", req.Host)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(200)
		_, err := res.Write([]byte{})
		assert.NoError(t, err)
	}))

	client := New(srv.URL)
	client.Host = "my-custom-host"

	_, err := client.send("GET", "/", nil)
	assert.NoError(t, err)
}

func TestClient_InvalidContentType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "some-type")
	}))

	client := New(srv.URL)

	_, err := client.send("GET", "/", nil)

	assert.EqualError(t, err, "invalid content-type 'some-type', expected 'application/json'")
}
