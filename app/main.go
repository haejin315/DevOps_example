package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Result struct {
	A, B int
	Sum  int
}

func main() {
	tmpl := template.Must(template.New("form").Parse(`
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			a, _ := strconv.Atoi(r.FormValue("a"))
			b, _ := strconv.Atoi(r.FormValue("b"))
			tmpl.Execute(w, Result{A: a, B: b, Sum: a + b})
		} else {
			tmpl.Execute(w, nil)
		}
	})

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
