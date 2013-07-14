// templateに別のtemplateを埋め込む。
package main

import (
	"html/template"
	"net/http"
)

const rootPageTemplate = `
<html>
<head>
<title>Sample - {{.SubTilte}}</title>
</head>
<body>
{{template "header" .}}
{{template "content" .}}
{{template "footer" .}}
</body>
</html>
`

func view(w http.ResponseWriter, r *http.Request) {
	tmpl := *template.Must(template.New("rootPage").Parse(rootPageTemplate))
	tmpl.New("header").Parse(`<h1>header</h1>`)
	tmpl.New("content").Parse(`<div>content here</div>`)
	tmpl.New("footer").Parse(`<div>footer</div>`)
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", view)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
