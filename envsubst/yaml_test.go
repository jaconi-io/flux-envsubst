package envsubst_test

import (
	"errors"
	"strings"
	"testing"

	. "github.com/jaconi-io/flux-envsubst/v4/envsubst"

	"github.com/stretchr/testify/assert"
)

func TestSplitYAMLInvalid(t *testing.T) {
	yaml := "foo: :"

	err := SplitYAML(strings.NewReader(yaml), func(b []byte) error {
		assert.Fail(t, "unexpected callback invocation")
		return nil
	})
	assert.EqualError(t, err, "yaml: mapping values are not allowed in this context")
}

func TestSplitYAMLCallbackError(t *testing.T) {
	yaml := "foo: bar"

	err := SplitYAML(strings.NewReader(yaml), func(b []byte) error {
		return errors.New("expected")
	})
	assert.EqualError(t, err, "expected")
}

func TestSplitYAML(t *testing.T) {
	yaml := `foo: bar
---
foo: baz
`

	i := 0
	err := SplitYAML(strings.NewReader(yaml), func(b []byte) error {
		switch i {
		case 0:
			assert.Equal(t, "foo: bar\n", string(b))
		case 1:
			assert.Equal(t, "foo: baz\n", string(b))
		default:
			assert.Fail(t, "too many callback invocations")
		}

		i++
		return nil
	})
	assert.NoError(t, err)
}
