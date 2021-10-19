package envsubst

import (
	"fmt"
	"os"

	"github.com/drone/envsubst"
)

type MissingVariableError struct {
	Variable string
}

func (m MissingVariableError) Error() string {
	return fmt.Sprintf("variable %q is not set", m.Variable)
}

// StrictEval is similar to Eval of github.com/drone/envsubst, but returns an error when a variable is missing.
func StrictEval(in string, mapping func(string) string) (out string, err error) {
	// Recover panics caused by missing variables.
	defer func() {
		rec := recover()
		recErr, ok := rec.(MissingVariableError)
		if ok {
			out = ""
			err = recErr
		}
	}()

	out, err = envsubst.Eval(in, func(s string) string {
		v := mapping(s)
		if v == "" {
			panic(MissingVariableError{Variable: s})
		}

		return v
	})

	return out, err
}

// StrictEvalEnv is similar to EvalEnv of github.com/drone/envsubst, but returns an error when a variable is missing.
func StrictEvalEnv(s string) (string, error) {
	return StrictEval(s, os.Getenv)
}
