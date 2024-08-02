package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type TemplateData struct {
	Username string
}

var scriptTemplate = ` 
#!/usr/bin/env bash

wget https://github.com/{{ .Username }}/.dotfiles/archive/refs/heads/main.zip -O {{ .Username }}-dotfiles.zip
unzip {{ .Username }}-dotfiles.zip 
(cd .dotfiles-main && /bin/bash install.sh -f)
rm -rf {{ .Username }}-dotfiles.zip .dotfiles-main
`

func serveScript(w http.ResponseWriter, r *http.Request) {
	strippedPath := strings.ReplaceAll(r.URL.Path, "/", "")
	data := TemplateData{strippedPath}
	fmt.Printf("%v", data)
	tmpl := template.Must(template.New("script").Parse(scriptTemplate))
	tmpl.Execute(w, data)
}

func main() {
	fmt.Println("Starting backend")
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveScript)
	mux.HandleFunc("/index.html", serveScript)

	err := http.ListenAndServe(":3000", mux)

	fmt.Printf("Stopped due to an error: %s", err)
}
