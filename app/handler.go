package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Result struct {
	A, B int
	Sum  int
}

func AddHandler() http.HandlerFunc {
	tmpl := getTemplate()
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			a, _ := strconv.Atoi(r.FormValue("a"))
			b, _ := strconv.Atoi(r.FormValue("b"))
			tmpl.Execute(w, Result{A: a, B: b, Sum: a + b})
		} else {
			tmpl.Execute(w, nil)
		}
	}
}

func getTemplate() *template.Template {
	return template.Must(template.New("form").Parse(`
		<html>
		<body>
			<h2>Add Two Numbers</h2>
			<form method="POST">
				A: <input name="a" type="number" />
				B: <input name="b" type="number" />
				<input type="submit" value="Add" />
			</form>
			{{if .}}
				<h3>Result: {{.A}} + {{.B}} = {{.Sum}}</h3>
			{{end}}
		</body>
		</html>
	`))
}
