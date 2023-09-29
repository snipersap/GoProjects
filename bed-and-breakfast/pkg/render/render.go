package render

// Create a template cache w/o manual input
import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/config"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/models"
)

// RenderTemplate takes the name of the template as input and writes the parsed
// output from templates and layouts to the response writer
func RenderTemplate(w http.ResponseWriter, tName string, td *models.TemplateData) error {

	//get parsed template
	t, err := getParsedTemplate(tName)
	if err != nil {
		return err
	}

	// Placeholder for adding some default data before writing the response
	td = addDefaultData(td)

	//Using the buffer to write to response instead of directly writing it
	buf := new(bytes.Buffer)
	err = t.Execute(buf, td)
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

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func getParsedTemplate(tName string) (*template.Template, error) {
	var tc map[string]*template.Template
	var err error

	//Pull from cache?
	if isCacheAllowed() {
		// load template cache from app config
		tc = appConfigInRender().TemplateCache
		if tc == nil {
			err := errors.New("could not load cache from app config")
			slog.Error(err.Error())
			return nil, err
		}
	} else {
		//create the Template Cache
		tc, err = GetTemplateCache()
		if err != nil {
			log.Fatalln("cannot load template cache. Exiting app:", err.Error())
		}
	}

	// get requested template from cache
	t, ok := tc[tName]
	if !ok {
		err := errors.New("could not load template from cache")
		slog.Error(err.Error())
		return nil, err
	}

	return t, nil
}

// isCacheAllowed returns UseCache from App config
func isCacheAllowed() bool {
	return appConfigInRender().UseCache
}

// GetTemplateCache reads the templates and layouts from the templates directory
// and returns the template cache after parsing
func GetTemplateCache() (map[string]*template.Template, error) {

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

	// log.Println(tmplCache)
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
		// log.Println(pages)
		return pages, nil
	}
}

// Populate the app config

// app is local to package render and stores the app config
var app *config.AppConfig

// SetAppConfig sets the reference of the application config
// locally in the render package
func SetAppConfig(a *config.AppConfig) {
	app = a
}

// appConfigInRender returns the reference of the local app config in render package
func appConfigInRender() *config.AppConfig {
	return app
}
