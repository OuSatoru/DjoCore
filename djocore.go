package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/OuSatoru/djocore/setIE"
)

func main() {
	setIE.SetProxy("http:127.0.0.1:8888;https:127.0.0.1:8888")
	if setIE.GetProxy() == "127.0.0.1:8888" {
		fmt.Println("True.")
	}
	http.HandleFunc("/echo", visitPort)
	http.HandleFunc("/", holdAddr)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func visitPort(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>Djocore Echo Service</h1>\n")
	fmt.Fprintf(w, "%s %s %s<br><br>\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]: %q<br>\n", k, v)
		//fmt.Printf("Header[%q]: %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q<br>\n", r.Host)
	//fmt.Printf("Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q<br>\n", r.RemoteAddr)
	//fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q<br>\n", k, v)
	}
}

func holdAddr(w http.ResponseWriter, r *http.Request)  {

}