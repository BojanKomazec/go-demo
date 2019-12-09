package htmltemplatedemo

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"os"
)

type person struct {
	Name    string
	Surname string
	// Surname template.JS // import html/template required
}

const st = `<!DOCTYPE html>
<html>
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

var c = bookmarkHTMLContent{
	ID1:    "1",
	Href1:  "www.google.com",
	Title1: "Google",
	Bookmarks: []bookmark{
		{
			ID:    "2",
			Href:  "www.amazon.com",
			Title: "Amazon",
		},
		{
			ID:    "3",
			Href:  "www.ebay.com",
			Title: "eBay",
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
	ID    string
	Href  string
	Title string
}

type bookmarkHTMLContent struct {
	ID1       string
	Href1     string
	Title1    string
	Bookmarks []bookmark
}

func readTemplateFromFileWriteToStdout() {
	t, e := template.ParseFiles("./internal/pkg/templatedemo/netscape_bookmarks_template_01.html")
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	t.Execute(os.Stdout, &c)
}

func readTemplateFromFileExecuteInBufferWriteToStdout() {
	t, e := template.ParseFiles("./internal/pkg/templatedemo/netscape_bookmarks_template_01.html")
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		return
	}

	var buf bytes.Buffer
	if e := t.Execute(&buf, &c); e != nil {
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
	fmt.Println("htmltemplatedemo.demoRangeArray()")
	t := template.Must(template.New("").Parse(stringTemplate))
	t.Execute(os.Stdout, []string{"first", "second", "third"})
	fmt.Println("~htmltemplatedemo.demoRangeArray()")
	fmt.Println()
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nhtmltemplatedemo.ShowDemo()\n\n")
	readTemplateFromStringWriteToStdout()
	readTemplateFromFileWriteToStdout()
	readTemplateFromFileExecuteInBufferWriteToStdout()
	showSpaceTrimming()
	showUsingVariablesInTemplate()
	demoPrintfInTemplate()
	demoRangeArray(st5)
	demoRangeArray(st6)
	demoRangeArray(st7)
	demoRangeArray(st8)
	fmt.Printf("\n\n~htmltemplatedemo.ShowDemo()\n\n")
}
