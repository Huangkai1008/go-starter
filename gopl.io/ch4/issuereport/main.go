package main

import (
	"go-starter/gopl.io/ch4/github"
	"html/template"
	"log"
	"os"
	"time"
)

const temp = `{{.TotalCount}} issues
{{range .Items}}---------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title| printf "%.64s"}}
Age: {{.CreatedAt|daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
