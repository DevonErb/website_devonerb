// Sends the template to the client
func serveAbout(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/about.html"))
    tmpl.Execute(w, nil)
}
