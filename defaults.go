package template // import "github.com/orkestr8/template"

import (
	"os"
)

// defaultContextURL returns the default context URL if none is known.
func defaultContextURL() string {
	pwd := "/"
	if wd, err := os.Getwd(); err == nil {
		pwd = wd
	} else {
		pwd = os.Getenv("PWD")
	}
	return "file://localhost" + pwd + "/"
}
