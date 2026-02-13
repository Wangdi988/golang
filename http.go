// package main

// import (
//     "fmt"
//     "net/http"
// )


// // An http.ResponseWriter which is where you write your text/html response to.
// // An http.Request which contains all information about this HTTP request including things like the URL or header fields.

// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//     })

//     http.ListenAndServe(":80", nil)
// }



package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", nil)
}