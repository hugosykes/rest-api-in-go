package main

import (
	//"github.com/gorilla/mux"
	"log"
	"net/http"
	"html"
	//"fmt"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", Index)
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func Index(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080"+html.EscapeString(r.URL.Path), bytes.NewBuffer([]byte(html.EscapeString(r.URL.Path))))

	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Fprint(w,"response Status: ", resp.Status)
	fmt.Fprint(w,"response Headers:", resp.Header)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body is:", string(body))
}
