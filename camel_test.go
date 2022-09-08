package camel_test

import (
	"testing"

	"github.com/dtcookie/assert"
	"github.com/dtcookie/camel"
)

func TestUnCamel(t *testing.T) {
	assert := assert.New(t)
	assert.Equals("", camel.Strip(""))
	assert.Equals("a", camel.Strip("a"))
	assert.Equals("a", camel.Strip("A"))
	assert.Equals("a_1", camel.Strip("a1"))
	assert.Equals("a_1", camel.Strip("A1"))
	assert.Equals("1_a", camel.Strip("1a"))
	assert.Equals("1_a", camel.Strip("1A"))
	assert.Equals("1_a", camel.Strip("1a"))
	assert.Equals("1_a", camel.Strip("1A"))
	assert.Equals("a_a", camel.Strip("aA"))
	assert.Equals("aa", camel.Strip("Aa"))
	assert.Equals("aa", camel.Strip("AA"))
	assert.Equals("zero_float_32_p", camel.Strip("zeroFloat32p"))
	assert.Equals("dstring_d", camel.Strip("DStringD"))

}
