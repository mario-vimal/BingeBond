package gql_schema

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

func String() (string, error) {
	buf := bytes.Buffer{}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	files := []string{dir + "/gql_schema/schema.graphql", dir + "/gql_schema/query.graphql", dir + "/gql_schema/mutation.graphql"}

	root := dir + "/gql_schema/types"
	filesDir, errDir := ioutil.ReadDir(root)
	if errDir != nil {
		return "", err
	}

	for _, f := range filesDir {
		files = append(files, root+"/"+f.Name())
	}

	for _, path := range files {
		b, err := ioutil.ReadFile(path)

		if err != nil {
			return "", err
		}

		buf.Write(b)

		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String(), nil

}
