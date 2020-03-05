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

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\ntexttemplatedemo.ShowDemo()\n\n")
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
	fmt.Printf("\n\n~texttemplatedemo.ShowDemo()\n\n")
}
