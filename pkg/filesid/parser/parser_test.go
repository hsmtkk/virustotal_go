package parser_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/virustotal_go/pkg/filesid/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	bs, err := ioutil.ReadFile("../example_response.json")
	assert.Nil(t, err, "should be nil")
	ps := parser.New()
	stats, err := ps.ParseResponse(string(bs))
	assert.Nil(t, err, "should be nil")
	assert.Zero(t, stats.HarmLess, "should be zero")
	assert.Equal(t, 59, stats.Undetected, "should match")
	assert.Zero(t, stats.Suspicious, "should be zero")
	assert.Zero(t, stats.Malicious, "should be zero")
}
