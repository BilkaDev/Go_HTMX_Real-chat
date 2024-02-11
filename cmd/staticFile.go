package main

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/a-h/templ"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/auth"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/home"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/layout"
)

func GenerateStaticFiles(rootPath string) {
	if err := os.Mkdir(rootPath, 0755); err != nil {
		if !os.IsExist(err) {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}

	createNewPage("index.html", rootPath, home.Show(), layout.Base())
	createNewPage("login", rootPath, auth.Show("login"), layout.Base())
	createNewPage("signup", rootPath, auth.Show("signup"), layout.Base())
}

func writeToFile(f *os.File, component templ.Component, layout templ.Component) error {
	return layout.Render(templ.WithChildren(context.Background(), component), f)
}

func createNewPage(name string, rootPath string, component templ.Component, layout templ.Component) {
	// Create   page.
	f, err := os.Create(path.Join(rootPath, name+""))
	if err != nil {
		if !os.IsExist(err) {
			log.Fatalf("failed to create output file: %v", err)
		}
	}
	// Write it out.
	err = writeToFile(f, component, layout)
	if err != nil {
		log.Fatalf("failed to write %v page: %v", name, err)
	}
}
