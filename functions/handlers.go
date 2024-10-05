package ascii

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Page struct {
	Title        string
	Art          string
	Text         string
	Banner       string
	MessageError string
	Code         int
	HttpErr      string
}

var art = &Page{Title: "index.html"}

// render templates
func renderTemplate(w http.ResponseWriter, p *Page) {
	temp, err := template.ParseFiles("./templates/" + p.Title)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Oops!!! Internal Server Error")
		return
	}
	err = temp.Execute(w, p)
	art = &Page{Title: "index.html"}
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Oops!!! Internal Server Error")
		return
	}
}

// Home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed, "Oops!! Method not allowed")
		return
	}

	if r.URL.Path != "/" {
		if strings.ContainsRune(r.URL.Path[1:], '/') {
			http.Redirect(w, r, "/notFound", http.StatusFound)
			return
		}
		ErrorHandler(w, http.StatusNotFound, " Oops!!Page Not Found")
		return
	}

	if art.Art == "" {
		os.Remove("output/file.txt")
	}

	renderTemplate(w, art)
	art.MessageError = ""
}

// Fonction qui gère la création et l'affichage de l'ASCII Art
func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed, "Oops!! Method not allowed")
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if len(text) > 3000 || text == "" || banner == "" {
		ErrorHandler(w, http.StatusBadRequest, "Oops!! Bad Request")
		return
	}
	data, err := Art(text, banner)
	if err != nil {
		if err.Error() == "not printable" {
			art.MessageError = "*your text is not printable"
		} else if err.Error() == "not a banner" {
			art.MessageError = "*This banner is not valide"
		} else if err.Error() == "this banner is not exists" {
			ErrorHandler(w, http.StatusBadRequest, "Oops!! Bad Request")
			return
		} else {
			ErrorHandler(w, http.StatusNotFound, "Oops!!Page Not Found")
			return
		}
	}
	art.Text = text
	art.Banner = banner
	art.Art = string(data)

	// download a file
	download := r.FormValue("download")
	if download == "download" {
		os.MkdirAll("output", 0o777)
		err = os.WriteFile("./output/file.txt", data, 0o666)
		if err != nil {
			fmt.Println(err.Error())
		}
		art.Text = ""
		art.Banner = ""
		art.Art = ""
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
		http.ServeFile(w, r, "output/file.txt")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// Fonction qui gère les erreurs HTTP et affiche les pages d'erreur correspondantes
func ErrorHandler(w http.ResponseWriter, status int, message string) {
	art.Code = status
	art.HttpErr = message
	w.WriteHeader(status)
	art.Title = "error.html"
	renderTemplate(w, art)
	art.Title = "index.html"
}
