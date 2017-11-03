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

	//if (r.Method != http.MethodGet) {
	//	http.Error(w, "GET 以外です", http.StatusInternalServerError)
	//	return
	//}

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	ping := Ping{http.StatusOK, "root"}
	res, _ := json.Marshal(ping)
	w.Header().Set("Content-Type", "application/json")

	// ひとつめの学び
	// これを返せば、GET はできる、POST もできる
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// これを追加すると POST with application/json が可能となる
	w.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Content-Type, X-CORS-Sample")


	// Cookie
	// 1
	cookie, err := r.Cookie("hoge")

	if err != nil {
		fmt.Println("hoge が取得できなかった")
		fmt.Println("Cookie", err)
	}
	// 2
	v := cookie.Value
	fmt.Println(v)

	w.Write(res)

}


func main() {
	var httpServer http.Server
	http.HandleFunc("/", rootHandler)
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}