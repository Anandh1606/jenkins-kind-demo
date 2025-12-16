package main
import (
    "fmt"
    "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello Kubernetes!")
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":11000", nil)
}
