package main

import(
    "fmt"
    "html"
    "net/http"
    "log"
)

func main() { 
    const port = ":3333"
    // home endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World from endpoint %q", html.EscapeString(r.URL.Path))
    })

    // create the listener in here
    fmt.Printf("Listening in %s", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

