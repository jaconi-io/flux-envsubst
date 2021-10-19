package envsubst_test

import "os"

func tmpEnv(key, value string) func() {
	old, set := os.LookupEnv(key)
	err := os.Setenv(key, value)
	if err != nil {
		panic(err)
	}

	return func() {
		if set {
			err = os.Setenv(key, old)
		} else {
			err = os.Unsetenv(key)
		}

		if err != nil {
			panic(err)
		}
	}
}
