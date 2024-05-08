package templates

import "html/template"

func Parse(names ...string) (*template.Template, error) {
	patterns := []string{"partial/*.tmpl"}
	for _, name := range names {
		patterns = append(patterns, name+".tmpl")
	}
	return template.New("").ParseFS(FS, patterns...)
}

func MustParse(names ...string) *template.Template {
	tpl, err := Parse(names...)
	if err != nil {
		panic(err)
	}
	return tpl
}
