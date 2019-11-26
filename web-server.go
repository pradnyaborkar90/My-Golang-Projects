package main
import (
    "fmt"
    "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w,"<h1> Welcome to my Digital Resume </h1><br><p> <b>Visit the following links to know more: </b><ul><li>  /skills </li> </ul> </p>")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, "<h1> The following are my major skills: </h1>")
   fmt.Fprint (w,"<body><ul> <li> <b> GCP</b> </li> <li> <b> AWS </b> </li> <li> <b> Docker </b> </li> <li> <b> ELK stack</b> </li> <li> <b> Google Kubernetes Engine </b> </li> <li> <b> Go <b> </li> <li> <b> Python </b> </li> </ul></body>")

}

func main() {

http.HandleFunc("/",  handlerFunc)
http.HandleFunc("/skills", aboutMe)
http.ListenAndServe(":3000", nil)

}
