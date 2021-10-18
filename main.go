package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func ifErr(err error) {
	if err != nil {
		panic(err)
	} else {
		return
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func getImageNames(files []os.FileInfo) []string {
	ImageNames := make([]string, 0)
	for _, file := range files {
		ImageNames = append(ImageNames, file.Name())
	}

	return ImageNames
}

type routeInfo struct {
	ImageLocation string
	Route         string
	ImageNames    []string
	RouteName     string
}

func main() {
	// todo change the http Dir to be programatially fetched through package 'os'
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("/home/ansverma/OLD/Anshul/Documents/Devendra/try1/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/fashion", http.StatusSeeOther)
	})
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	handlers := []func(http.ResponseWriter, *http.Request){}
	routesInfo := []routeInfo{{Route: "/fashion", RouteName: "Fashion"}, {Route: "/celebrity", RouteName: "Celebrity"}, {Route: "/product", RouteName: "Product"}, {Route: "/boudior", RouteName: "Boudior"}, {Route: "/wedding", RouteName: "Wedding"}, {Route: "/editorial", RouteName: "Editorial"}, {Route: "/conceptual", RouteName: "Conceptual"}}
	for _, routeInfo := range routesInfo {

		routeInfo.ImageLocation = fmt.Sprintf("./img%s", routeInfo.Route)
		files, err := ioutil.ReadDir(routeInfo.ImageLocation)
		ifErr(err)
		routeInfo.ImageNames = getImageNames(files)
		route := routeInfo
		handler := func(w http.ResponseWriter, r *http.Request) {
			err = tpl.ExecuteTemplate(w, "index.html", route)
			ifErr(err)
		}
		handlers = append(handlers, handler)
		/*
			http.HandleFunc("fashion", func(w http.ResponseWriter, r *http.Request) {
						fmt.Println(routeInfo.Route)
						err = tpl.ExecuteTemplate(w, "index.html", routeInfo)
						ifErr(err)
			})
		*/
	}

	for i, h := range handlers {
		http.HandleFunc(routesInfo[i].Route, h)
		i++
	}
	err := http.ListenAndServe(":8000", nil)
	ifErr(err)
}

func about(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles("about.html")
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles("contact.html")
	tpl.ExecuteTemplate(w, "contact.html", nil)
}
