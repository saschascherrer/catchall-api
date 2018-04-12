package main

import "net/http"
import "log"
import "io/ioutil"
import "flag"

func main() {
	var socket string
	flag.StringVar(&socket, "l", ":8080", "Socket to listen to")

	log.Printf("Starting catchall-api on %s \n", socket)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		b, err := ioutil.ReadAll(r.Body)
		if ( err != nil) {
			log.Println("400: Unreadable Body")
		}
		log.Printf("%s\n", b)
		log.Println("-----");
	})
	log.Fatal(http.ListenAndServe(socket, nil))
}
