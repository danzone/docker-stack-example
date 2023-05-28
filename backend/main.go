package main
 
import (
   "fmt"
   "log"
   "net/http"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Println(r.URL.RawQuery)
   fmt.Fprintf(w, `
+++++++++++++++++++++++
_________ _________
| _______| | _______|
| | | |
| | | |
| | | |
| |_______ | |_______
|_________| |_________|
 
++++++++++++++++++++++++
 
------------------------------------------
Hey, there!!! This is Daniele Gagliardi. :)
------------------------------------------
`)
}
 
func main() {
   http.HandleFunc("/", handler)
   log.Fatal(http.ListenAndServe(":80", nil))
}
