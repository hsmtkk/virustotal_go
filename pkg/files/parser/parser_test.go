package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/virustotal_go/pkg/files/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	bs, err := ioutil.ReadFile("../example_response.json")
	assert.Nil(t, err, "should be nil")
	ps := parser.New()
	want := "NjY0MjRlOTFjMDIyYTkyNWM0NjU2NWQzYWNlMzFmZmI6MTQ3NTA0ODI3Nw=="
	got, err := ps.ParseResponse(string(bs))
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should match")
}
