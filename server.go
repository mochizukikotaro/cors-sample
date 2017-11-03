package main

import (
	"net/http"
	"fmt"
	"net/http/httputil"
	"log"
	"encoding/json"
)

type Ping struct {
	Status int `json:"status"`
	Result string `json:"result"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet) {
		http.Error(w, "GET 以外です", http.StatusInternalServerError)
		return
	}

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	ping := Ping{http.StatusOK, "root"}
	res, _ := json.Marshal(ping)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}


func main() {
	var httpServer http.Server
	http.HandleFunc("/", rootHandler)
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}