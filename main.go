package main

import (
	"fmt"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<script src="https://unpkg.com/vue@3/dist/vue.global.js"></script>

	<div id="app">{{ message }}</div>
	
	<script>
	  const { createApp } = Vue
	
	  createApp({
		data() {
		  return {
			message: 'Hello Vue!'
		  }
		}
	  }).mount('#app')
	</script>
	`)
}
func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(MyHandler))
	mux.Handle("/api/iot")

	func() error {
		server := &http.Server{Addr: ":8181", Handler: http.Handler(mux)}
		return server.ListenAndServe()
	}()
}
