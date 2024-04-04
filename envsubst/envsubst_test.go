package envsubst_test

import (
	"testing"

	. "github.com/jaconi-io/flux-envsubst/v4/envsubst"

	"github.com/stretchr/testify/assert"
)

func TestStrictEvalMissingVariable(t *testing.T) {
	out, err := StrictEval("Hello ${foo}!", func(s string) string {
		return ""
	})

	assert.NoError(t, err)
	assert.Equal(t, "Hello !", out)
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

	assert.NoError(t, err)
	assert.Equal(t, "Hello !", out)
}

func TestStrictEvalEnvMissingVariableWithDefault(t *testing.T) {
	out, err := StrictEvalEnv("Hello ${foo:=bar}!")

	assert.NoError(t, err)
	assert.Equal(t, "Hello bar!", out)
}

func TestStrictEvalEnv(t *testing.T) {
	reset := tmpEnv("foo", "bar")
	defer reset()

	out, err := StrictEvalEnv("Hello ${foo}!")

	assert.NoError(t, err)
	assert.Equal(t, "Hello bar!", out)
}
