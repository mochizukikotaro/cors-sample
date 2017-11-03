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

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	ping := Ping{http.StatusOK, "root"}
	res, _ := json.Marshal(ping)
	w.Header().Set("Content-Type", "application/json")

	// これを返せば、GET はできる、POST もできる
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	// Cookie を取得したい場合はワイルドカードは使えない
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9999")

	// これを追加すると POST with application/json が可能となる
	w.Header().Set("Access-Control-Allow-Headers","Content-Type, X-Hoge")

	// これを追加すると Cookie が取得できる
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Cookie
	cookie, _ := r.Cookie("hoge")
	if cookie != nil {
		v := cookie.Value
		fmt.Println(v)
	} else {
		fmt.Println("Cookie が取得できなかった")
	}

	w.Write(res)
}


func main() {
	var httpServer http.Server
	http.HandleFunc("/", rootHandler)
	httpServer.Addr = ":8888"
	log.Println(httpServer.ListenAndServe())
}