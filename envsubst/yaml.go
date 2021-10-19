package envsubst

import (
	"bytes"
	"io"

	"gopkg.in/yaml.v3"
)

// SplitYAML into separate documents.
func SplitYAML(raw []byte) ([][]byte, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(raw))

	var rawDocuments [][]byte
	for {
		var document interface{}
		err := decoder.Decode(&document)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		rawDocument, err := yaml.Marshal(document)
		if err != nil {
			return nil, err
		}

		rawDocuments = append(rawDocuments, rawDocument)
	}

	return rawDocuments, nil
}
