package envsubst

import (
	"io"

	"gopkg.in/yaml.v3"
)

// SplitYAML into separate documents.
func SplitYAML(in io.Reader, callback func([]byte) error) error {
	decoder := yaml.NewDecoder(in)

	for {
		var document interface{}
		err := decoder.Decode(&document)
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		rawDocument, err := yaml.Marshal(document)
		if err != nil {
			return err
		}

		err = callback(rawDocument)
		if err != nil {
			return err
		}
	}

	return nil
}
