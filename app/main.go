package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
