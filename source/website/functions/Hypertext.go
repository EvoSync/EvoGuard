package functions

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/valyala/fasttemplate"
)

/*
	Hypertext.go will implement the remote rendering
	interface for the html files found within the
	file server service.
*/

// WriteHypertext will attempt to write the hypertext contents to the node
func WriteHypertext(response http.ResponseWriter, elements map[string]string, file ...string) error {
	contents, err := os.ReadFile(filepath.Join(file...))
	if err != nil {
		return err
	}

	// ExportLast will attempt to continue within the hypertext transcript
	return ExportLast(response.Write([]byte(fasttemplate.ExecuteFuncString(string(contents), "[[", "]]", func(w io.Writer, tag string) (int, error) {
		if len(tag) <= 0 || tag[0] != '$' {
			return 0, nil
		}

		content, ok := elements[tag[1:]]
		if !ok || len(content) == 0 {
			return 0, nil
		}

		return w.Write([]byte(content))
	}))))
}
