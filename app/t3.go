package main

import (
    "html/template"
    "os"

    "github.com/Masterminds/sprig/v3"
)

type Alert struct {
    URL string
}

func main() {
    tmpl, err := template.New("test").Funcs(sprig.FuncMap()).Parse(`
{{- $x := regexFind "earliest=(.+?)&" .URL | replace "earliest=" "" | replace "&" "" -}}
{{ regexReplaceAll "earliest=(.+?)&" .URL (list "earliest=" ((now).UTC.Unix | toString) "&" | join "") }}
{{ .URL }}`)
    if err != nil {
        panic(err)
    }
    err = tmpl.Execute(os.Stdout, &Alert{URL: "https://search.com?earliest=-15m&latest=now"})
    if err != nil {
        panic(err)
    }
}
