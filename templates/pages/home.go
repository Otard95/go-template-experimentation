package pages

import (
	"io"
	"testing/teplates/templates"
	"testing/teplates/templates/components"
)

type HomeProps struct {
	Header components.HeaderProps
}

func Home(wr io.Writer, props HomeProps) error {
	tmpl, err := templates.Create(templates.PageTemlplate, components.HeaderTemlplate, "templates/pages/home.html")
	if err != nil {
		return err
	}

	return tmpl.Execute(wr, props)
}
