package main

import (
	"fmt"
	"net/http"
)

const html = `<html>
	<body>
		<h1>Hello, world</h1>
	</body>
</html>`

func main() {
	fmt.Println("starting up...")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
