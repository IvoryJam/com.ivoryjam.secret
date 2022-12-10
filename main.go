package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./resources"))
	http.Handle("/", fs)
	http.HandleFunc("/loadGifts", getTable)
	http.ListenAndServe(":3000", nil)
}

func getTable(w http.ResponseWriter, r *http.Request) {
	bullShitURL := "https://docs.google.com/spreadsheets/d/e/2PACX-1vRHzMzusRPzI3NbkGJv6Urtr0qmlWbzuq5wujhrgxD8sPrviyYGZjOcFbrkeP_ZgQNYSzQsMNfiqpDm/pub?gid=0&single=true&output=tsv"
	bullShitBody, err := http.Get(bullShitURL)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer bullShitBody.Body.Close()

	b, err := ioutil.ReadAll(bullShitBody.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write(b)
}
