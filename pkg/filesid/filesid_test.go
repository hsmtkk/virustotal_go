package filesid_test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"testing"

	"github.com/hsmtkk/virustotal_go/pkg/apikey"
	"github.com/hsmtkk/virustotal_go/pkg/filesid"
	"github.com/stretchr/testify/assert"
)

const hashSample = "f4d76f4ad2977077b00035901b614d04a1fd5e5dec9d22309279304c8da56865"

func TestReal(t *testing.T) {
	apiKey, err := apikey.New().LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	api := filesid.New(apiKey)
	resp, err := api.RetrieveInformation(hashSample)
	assert.Nil(t, err, "should be nil")
	log.Println(resp)
}

func TestLocal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := httputil.DumpRequest(r, true)
		assert.Nil(t, err, "should be nil")
		log.Println(string(bs))
		src, err := os.Open("example_response.json")
		assert.Nil(t, err, "should be nil")
		io.Copy(w, src)
	}))
	defer ts.Close()

	api := filesid.NewForTest(ts.Client(), ts.URL)
	_, err := api.RetrieveInformation(hashSample)
	assert.Nil(t, err, "should be nil")
}
