package envsubst

import (
	"context"
	"fmt"

	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/yaml"
)

func SubstituteVariables(ctx context.Context, res *resource.Resource) (*resource.Resource, error) {
	resData, err := res.AsYAML()
	if err != nil {
		return nil, err
	}

	// Disable variable substitution, if annotation / label is present.
	key := "kustomize.toolkit.fluxcd.io/substitute"
	if res.GetLabels()[key] == "disabled" || res.GetAnnotations()[key] == "disabled" {
		return res, nil
	}

	output, err := StrictEvalEnv(string(resData))
	if err != nil {
		return nil, fmt.Errorf("variable substitution failed for %s: %w", id(res), err)
	}

	jsonData, err := yaml.YAMLToJSON([]byte(output))
	if err != nil {
		return nil, fmt.Errorf("conversion to JSON failed for %s: %w", id(res), err)
	}

	err = res.UnmarshalJSON(jsonData)
	if err != nil {
		return nil, fmt.Errorf("conversion from JSON failed for %s: %w", id(res), err)
	}

	return res, nil
}

func id(res *resource.Resource) string {
	kind := res.GetKind()
	namespace := res.GetNamespace()
	name := res.GetName()
	return fmt.Sprintf("%s %s/%s", kind, namespace, name)
}
