package texttemplatedemo

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
)

type person struct {
	Name    string
	Surname string
	// Surname template.JS // import html/template required
}

const st = `<!DOCTYPE html>
<html> "github.com/Masterminds/sprig"
   <head>
      <title>Test</title>
   </head>
   <body>
      <script>
        var {{.Name}} = {{.Surname}};
      </script>
      <div>{{.Name}}, {{.Surname}}</div>
      <p>{{.Name}}, {{.Surname}}</p>
      <DP>{{.Name}}, {{.Surname}}</DP>
      <DP a={{.Name}}>{{.Surname}}</DP>
   </body>
</html>
`

var content1 = bookmarkHTMLContent{
	ID1:    "1",
	Href1:  "www.google.com",
	Title1: "Google",
	Bookmarks: []bookmark{
		{
			AddDate: time.Now(),
			ID:      "2",
			Href:    "www.amazon.com",
			Title:   "Amazon",
		},
		{
			AddDate: time.Now(),
			ID:      "3",
			Href:    "www.ebay.com",
			Title:   "eBay",
		},
		{
			AddDate: time.Now(),
			ID:      "",
			Href:    "www.booking.com",
			Title:   "Booking",
		},
	},
}

func readTemplateFromStringWriteToStdout() {
	t, e := template.New("web_page_template").Parse(st)
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}
	t.Execute(os.Stdout, &person{"Bojan", "Komazec"})
}

type bookmark struct {
	AddDate time.Time
	ID      string
	Href    string
	Title   string
}

type bookmarkHTMLContent struct {
	ID1       string
	Href1     string
	Title1    string
	Bookmarks []bookmark
}

// This is a conversion function that converts item's value into string so can be used for filling in the HTML template.
func timeToUnixTimeString(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

const templateFilePath string = "./internal/pkg/texttemplatedemo/netscape_bookmarks_template_01.html"

// Template has to be created with New() before calling Funcs() on it.
//
// Template name passed to New() has to match the name of template file.
// If that's not the case, we'll get the error like in this example:
// 		// Error: template: bookmark: "bookmark" is an incomplete or empty template
// 		template.New("bookmark").ParseFiles("some_file.html")
//
// https://stackoverflow.com/questions/24837883/golang-templates-minus-function
// https://stackoverflow.com/questions/49043292/error-template-is-an-incomplete-or-empty-template
// https://stackoverflow.com/questions/17843311/template-and-custom-function-panic-function-not-defined
// Funcs() must be called on template before parsing (calling ParseFiles()).
func readTemplateFromFileWriteToStdout() {
	templateFileName := filepath.Base(templateFilePath)
	t, e := template.New(templateFileName).Funcs(template.FuncMap{"timeToUnixTimeString": timeToUnixTimeString}).ParseFiles(templateFilePath)
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	err := t.Execute(os.Stdout, &content1)
	if err != nil {
		fmt.Println("template.Execute() failed with error:", err.Error())
	}
}

func readTemplateFromFileExecuteInBufferWriteToStdout() {
	templateFileName := filepath.Base(templateFilePath)
	t, e := template.New(templateFileName).Funcs(template.FuncMap{"timeToUnixTimeString": timeToUnixTimeString}).ParseFiles(templateFilePath)
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	var buf bytes.Buffer
	if e := t.Execute(&buf, &content1); e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	fmt.Println("Buffer content: ", buf.String())

	fb, e := format.Source(buf.Bytes())
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	os.Stdout.Write(fb)
}

const st2 = `Greeting:
{{- if eq .TimeOfDay "morning" -}}
Good morning!
{{- else if eq .TimeOfDay "evening" -}}
Good evening!
{{- end -}}`

// https://golang.org/pkg/text/template/#hdr-Text_and_spaces
func showSpaceTrimming() {
	t := template.Must(template.New("").Parse(st2))
	t.Execute(os.Stdout, map[string]interface{}{"TimeOfDay": "evening"})
	fmt.Println()
	t.Execute(os.Stdout, map[string]interface{}{"TimeOfDay": "morning"})
	fmt.Println()
}

const st3 = `Greeting:
{{- $Name:=.Name}}
{{- if eq .TimeOfDay "morning" -}}
Good morning, {{$Name}}!
{{- else if eq .TimeOfDay "evening" -}}
Good evening, {{$Name}}!
{{- end -}}`

// https://golang.org/pkg/text/template/#hdr-Text_and_spaces
func showUsingVariablesInTemplate() {
	t := template.Must(template.New("").Parse(st3))
	t.Execute(os.Stdout, map[string]interface{}{"TimeOfDay": "evening", "Name": "Bob"})
	fmt.Println()
	t.Execute(os.Stdout, map[string]interface{}{"TimeOfDay": "morning", "Name": "Alice"})
	fmt.Println()
}

const st4 = `abcd{{printf "efgh\n"}}`

func demoPrintfInTemplate() {
	t := template.Must(template.New("").Parse(st4))
	t.Execute(os.Stdout, struct{}{})
	fmt.Println()
}

const st5 = `{{- range $value := .}}
{{- $value}}
{{end -}}`

const st6 = `{{- range $value := . -}}
<p>{{$value}}</p>
{{end -}}`

const st7 = `<div>
	{{range $value := . -}}
    <p>{{$value}}</p>
	{{end}}
</div>
`

const st8 = `<div>
	{{- range $value := .}}
	{{$value}}
	{{- end}}
</div>
`

func demoRangeArray(stringTemplate string) {
	fmt.Println("texttemplatedemo.demoRangeArray()")
	t := template.Must(template.New("").Parse(stringTemplate))
	t.Execute(os.Stdout, []string{"first", "second", "third"})
	fmt.Println("~texttemplatedemo.demoRangeArray()")
	fmt.Println()
}

// DirNode struct
type DirNode struct {
	Name  string
	Files []string
	Nodes []DirNode
}

// Solution for recursive templates:
// 		https://gist.github.com/mxlje/8e6279a90dc8f79f65fa8c855e1d7a79
//
// Links to read in order to understand solution for getting indentation of nodes created by recursive template:
// 		https://github.com/Masterminds/sprig
// 		https://github.com/Masterminds/sprig/issues/57
// 		https://stackoverflow.com/questions/43821989/how-to-indent-content-of-included-template
// 		https://github.com/helm/helm/blob/8648ccf5d35d682dcd5f7a9c2082f0aaf071e817/pkg/engine/engine.go#L148-L154
func recursiveTemplateDemo() {
	// I want to use sprig package so can do: {{ template "dirnode" . | indent 4 }}.
	// cannot use sprig.FuncMap() (type "html/template".FuncMap) as type "text/template".FuncMap in argument to "text/template".New("root").Funcs
	// tmpl, err := template.New("root").Funcs(sprig.FuncMap()).Parse(`

	tmpl := template.New("root")

    var funcMap template.FuncMap = map[string]interface{}{}
    // copied from: https://github.com/helm/helm/blob/8648ccf5d35d682dcd5f7a9c2082f0aaf071e817/pkg/engine/engine.go#L147-L154
    funcMap["include"] = func(name string, data interface{}) (string, error) {
        buf := bytes.NewBuffer(nil)
        if err := tmpl.ExecuteTemplate(buf, name, data); err != nil {
            return "", err
        }
        return buf.String(), nil
	}

    tmpl = tmpl.Funcs(sprig.TxtFuncMap()).Funcs(funcMap)

	var err error
	tmpl, err = tmpl.Parse(
		`{{ define "dirnode" }}
			<DIRNODE>
				Name = {{ .Name -}}
				{{ if gt (len .Files) 0 }}
					{{- range .Files }}
				File = {{ . }}
					{{- end }}
				{{- end }}
				{{- if gt (len .Nodes) 0 }}
				{{- range .Nodes }}
					{{- include "dirnode" . | indent 8 }}
				{{- end }}
				{{- end }}
			</DIRNODE>
		{{- end }}
		{{ include "dirnode" . | indent 4 }}
	`)

	if err != nil {
		panic(err)
    }

	data := DirNode{
		Name: "Root",
		Files: []string{
			"root_file1",
			"root_file2",
		},
		Nodes: []DirNode{
			DirNode{
				Name: "Users",
				Files: []string{
					"root_file1",
					"root_file2",
				},
				Nodes: []DirNode{
					DirNode{
						Name: "Alice",
						Files: []string{
							"alice_file1",
							"alice_file2",
						},
						Nodes: []DirNode{},
					},
					DirNode{
						Name: "Bob",
						Files: []string{
							"bob_file1",
							"bob_file2",
						},
						Nodes: []DirNode{
							DirNode{
								Name: "Books",
								Files: []string{
									"bobs_book1",
									"bobs_book2",
								},
								Nodes: []DirNode{},
							},
						},
					},
				},
			},
			DirNode{
				Name: "Libraries",
				Files: []string{
					"libraries_file1",
					"libraries_file2",
				},
				Nodes: []DirNode{
					DirNode{
						Name: "LibA",
						Files: []string{
							"libA_file1",
							"libA_file2",
						},
						Nodes: []DirNode{},
					},
				},
			},
		},
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\ntexttemplatedemo.ShowDemo()\n\n")
	// readTemplateFromStringWriteToStdout()
	// readTemplateFromFileWriteToStdout()
	// readTemplateFromFileExecuteInBufferWriteToStdout()
	// showSpaceTrimming()
	// showUsingVariablesInTemplate()
	// demoPrintfInTemplate()
	// demoRangeArray(st5)
	// demoRangeArray(st6)
	// demoRangeArray(st7)
	// demoRangeArray(st8)
	recursiveTemplateDemo()
	fmt.Printf("\n\n~texttemplatedemo.ShowDemo()\n\n")
}
