package files

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeFormData(t *testing.T) {
	data, contentType, err := makeFormData("example.txt", []byte("example"))
	assert.Nil(t, err, "should be nil")
	log.Println(contentType)
	log.Println(string(data))
}
