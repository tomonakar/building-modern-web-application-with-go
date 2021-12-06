package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w, tmpl)
	if err != nil {
		fmt.Println("Error getting template cache", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter, tmpl string) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// filapath.Globでパターンにマッチするファイルを取得
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		// filepath.Baseでファイルパスの最後の部分を取得（つまりファイル名）
		name := filepath.Base(page)

		// template.Newは、指定した名前のHTMLテンプレートを割り当てる
		// Funcsはテンプレートに関数を割り当てる
		// ParseFilesは、指定したファイルを読み込んで、テンプレートをパースする
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// ParseGlobは、パターンによって特定されたファイルを読み込んで、テンプレートをパースする
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}
	return myCache, nil
}

// var functions = template.FuncMap{}

// // RenderTemplate renders a template
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	_, err := RenderTemplateTest(w)
// 	if err != nil {
// 		fmt.Println("Error getting template cache", err)
// 	}

// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

// 	err = parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 	}
// }

// func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

// 	myCache := map[string]*template.Template{}

// 	pages, err := filepath.Glob("./templates/*.page.tmpl")
// 	if err != nil {
// 		return myCache, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)
// 		fmt.Println("Page is currently", page)
// 		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
// 		if err != nil {
// 			return myCache, err
// 		}

// 		matches, err := filepath.Glob("./templates/*.layout.tmpl")
// 		if err != nil {
// 			return myCache, err
// 		}

// 		if len(matches) > 0 {
// 			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
// 			if err != nil {
// 				return myCache, err
// 			}
// 		}

// 		myCache[name] = ts
// 	}

// 	return myCache, nil
// }
