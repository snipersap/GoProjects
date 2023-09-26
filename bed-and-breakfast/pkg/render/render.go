package render

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"path/filepath"
)

// Create a template cache w/o manual input

// RenderTemplate takes the name of the template as input and writes the parsed
// output from templates and layouts to the response writer
func RenderTemplate(w http.ResponseWriter, tName string) error {

	//create the Template Cache
	tc, err := createTemplateCache()
	if err != nil {
		slog.Error("could not load cache from disk!")
		return err
	}

	// get requested template from cache
	t, ok := tc[tName]
	if !ok {
		slog.Error("could not load template from cache")
		return err
	}

	//Using the buffer to write to response instead of directly writing it
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		slog.Error("could not write parsed templates to buffer", "error", err.Error())
		return err
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		slog.Error("could not write to response", "error", err.Error())
		return err
	}
	return nil
}

// createTemplateCache reads the templates and layouts from the templates directory
// and returns the template cache after parsing
func createTemplateCache() (map[string]*template.Template, error) {

	// init the cache
	tmplCache := map[string]*template.Template{}

	// Read the template files before reading the layout files
	pages, err := pageTemplates()
	if err != nil {
		slog.Error(err.Error())
		return tmplCache, err
	}

	//Parse the template files into page templates
	for _, page := range pages {
		tName := filepath.Base(page)
		tSet, err := template.New(tName).ParseFiles(page)
		if err != nil {
			slog.Error(err.Error())
			return tmplCache, err
		}

		// Parse the page templates with the help of layout files into pages
		err = parseLayouts(tSet)
		if err != nil {
			slog.Error(err.Error())
			return tmplCache, err
		}

		// Set the template into the cache
		tmplCache[tName] = tSet
	}

	log.Println(tmplCache)
	return tmplCache, nil
}

// parseLayouts parses the layouts with the page and returns an error if any
func parseLayouts(tSet *template.Template) error {
	layoutPattern := "./templates/*.layout.tmpl"
	layouts, err := filepath.Glob(layoutPattern)
	if err != nil {
		slog.Error("bad pattern for layout", "search query", layoutPattern)
		return err
	} else if len(layouts) < 1 {
		err = errors.New("could not read layouts from disk")
		slog.Error(err.Error(), "search query", layoutPattern)
		return err
	}

	_, err = tSet.ParseGlob(layoutPattern)
	if err != nil {
		slog.Error("could not parse layouts with pages", "layout query", layoutPattern)
		return err
	}

	return nil
}

// pageTemplates returns the template files from disk
func pageTemplates() ([]string, error) {
	pattern := "./templates/*.page.tmpl"
	pages, err := filepath.Glob(pattern)
	if err != nil {
		slog.Error("bad pattern of page template", "search query", pattern)
		return nil, err
	} else if len(pages) < 1 {
		err = errors.New("could not load templates from disk")
		slog.Error(err.Error(), "search query", pattern)
		return nil, err
	} else {
		log.Println(pages)
		return pages, nil
	}
}
