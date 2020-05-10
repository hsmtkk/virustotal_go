package files

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeFormData(t *testing.T) {
	data, err := makeFormData([]byte("example"))
	assert.Nil(t, err, "should be nil")
	log.Println(string(data))
}
