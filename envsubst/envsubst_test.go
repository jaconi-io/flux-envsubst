package envsubst_test

import (
	"testing"

	. "github.com/jaconi-io/flux-envsubst/envsubst"

	"github.com/stretchr/testify/assert"
)

func TestStrictEvalMissingVariable(t *testing.T) {
	out, err := StrictEval("Hello ${foo}!", func(s string) string {
		return ""
	})

	assert.EqualError(t, err, "variable \"foo\" is not set")
	assert.Equal(t, "", out)
}

func TestStrictEval(t *testing.T) {
	out, err := StrictEval("Hello ${foo}!", func(s string) string {
		return "bar"
	})

	assert.NoError(t, err)
	assert.Equal(t, "Hello bar!", out)
}

func TestStrictEvalEnvMissingVariable(t *testing.T) {
	out, err := StrictEvalEnv("Hello ${foo}!")

	assert.EqualError(t, err, "variable \"foo\" is not set")
	assert.Equal(t, "", out)
}

func TestStrictEvalEnv(t *testing.T) {
	reset := tmpEnv("foo", "bar")
	defer reset()

	out, err := StrictEvalEnv("Hello ${foo}!")

	assert.NoError(t, err)
	assert.Equal(t, "Hello bar!", out)
}
