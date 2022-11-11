package envsubst

import (
	"fmt"
	"os"

	"github.com/drone/envsubst/v2"
)

// StrictEval is similar to Eval of github.com/drone/envsubst, but logs a warning when a variable is missing.
func StrictEval(in string, mapping func(string) string) (out string, err error) {
	out, err = envsubst.Eval(in, func(s string) string {
		v := mapping(s)
		if v == "" {
			fmt.Fprintf(os.Stderr, "variable %q might be missing\n", s)
		}

		return v
	})

	return out, err
}

// StrictEvalEnv is similar to EvalEnv of github.com/drone/envsubst, but logs a warning when a variable is missing.
func StrictEvalEnv(s string) (string, error) {
	return StrictEval(s, os.Getenv)
}
