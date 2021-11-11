package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/fantasticrabbit/ClickupCLI/internal"
	"github.com/fantasticrabbit/ClickupCLI/mocks"
	"github.com/stretchr/testify/assert"
)

func init() {
	internal.Client = &mocks.MockClient{}
}

func TestGetClientResponseSuccess(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "admin"}}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	request := internal.getJSON("http://localhost/testpath")
	assert.NotNil(t, request)
	assert.EqualValues(t, "Test Name", request.Name)
}
