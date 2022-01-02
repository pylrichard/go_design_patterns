package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrganization(t *testing.T) {
	o := NewOrganization().Count()
	assert.Equal(t, 20, o)
}