package envsubst_test

import (
	"testing"

	. "github.com/jaconi-io/flux-envsubst/envsubst"

	"github.com/stretchr/testify/assert"
)

func TestSplitYAMLInvalid(t *testing.T) {
	yaml := "foo: :"

	out, err := SplitYAML([]byte(yaml))
	assert.EqualError(t, err, "yaml: mapping values are not allowed in this context")
	assert.Len(t, out, 0)
}

func TestSplitYAML(t *testing.T) {
	yaml := `foo: bar
---
foo: baz
`

	out, err := SplitYAML([]byte(yaml))
	assert.NoError(t, err)
	assert.Equal(t, "foo: bar\n", string(out[0]))
	assert.Equal(t, "foo: baz\n", string(out[1]))
}
