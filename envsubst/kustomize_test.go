package envsubst_test

import (
	"context"
	"testing"

	. "github.com/jaconi-io/flux-envsubst/envsubst"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/yaml"

	"github.com/stretchr/testify/assert"
)

func TestSubstituteVariables(t *testing.T) {
	reset := tmpEnv("foo", "bar")
	defer reset()

	raw := []byte("foo: ${foo}")
	res := &resource.Resource{}
	err := yaml.Unmarshal(raw, res)
	assert.NoError(t, err)

	out, err := SubstituteVariables(context.TODO(), res)
	assert.NoError(t, err)

	rawOut, err := out.AsYAML()
	assert.NoError(t, err)
	assert.Equal(t, "foo: bar\n", string(rawOut))
}

func TestSubstituteVariables_MissingVariable(t *testing.T) {
	raw := []byte("foo: ${foo}")
	res := &resource.Resource{}
	err := yaml.Unmarshal(raw, res)
	assert.NoError(t, err)

	out, err := SubstituteVariables(context.TODO(), res)
	assert.NoError(t, err)

	rawOut, err := out.AsYAML()
	assert.NoError(t, err)
	assert.Equal(t, "foo: null\n", string(rawOut))
}

func TestSubstituteVariables_WithDefault(t *testing.T) {
	reset := tmpEnv("foo", "bar")
	defer reset()

	raw := []byte("foo: ${foo:-baz}")
	res := &resource.Resource{}
	err := yaml.Unmarshal(raw, res)
	assert.NoError(t, err)

	out, err := SubstituteVariables(context.TODO(), res)
	assert.NoError(t, err)

	rawOut, err := out.AsYAML()
	assert.NoError(t, err)
	assert.Equal(t, "foo: bar\n", string(rawOut))
}

func TestSubstituteVariables_MissingVariableWithDefault(t *testing.T) {
	raw := []byte("foo: ${foo:-bar}")
	res := &resource.Resource{}
	err := yaml.Unmarshal(raw, res)
	assert.NoError(t, err)

	out, err := SubstituteVariables(context.TODO(), res)
	assert.NoError(t, err)

	rawOut, err := out.AsYAML()
	assert.NoError(t, err)
	assert.Equal(t, "foo: bar\n", string(rawOut))
}

func TestSubstituteVariablesAnnotation(t *testing.T) {
	reset := tmpEnv("foo", "bar")
	defer reset()

	raw := []byte(`foo: ${foo}
metadata:
  annotations:
    kustomize.toolkit.fluxcd.io/substitute: disabled
`)
	res := &resource.Resource{}
	err := yaml.Unmarshal(raw, res)
	assert.NoError(t, err)

	out, err := SubstituteVariables(context.TODO(), res)
	assert.NoError(t, err)

	rawOut, err := out.AsYAML()
	assert.NoError(t, err)
	assert.Equal(t, string(raw), string(rawOut))
}
