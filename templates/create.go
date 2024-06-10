package templates

import (
	"errors"
	"html/template"
	"path"
)

func Create(files ...string) (*template.Template, error) {
  if len(files) == 0 {
    return nil, errors.New("No template files provided")
  }

  first := files[0]
  name := path.Base(first)

	return template.New(name).ParseFiles(files...)
}
