package main

import(
	"fmt"
	"net/http"
)

func pippoFunc(w http.ResponseWriter, r*http.Request){
	w.Write([]byte(`  
		<!DOCTYPER HTML>
		ciao



		`))
}


func main(){
	http.HandleFunc("/pippo", pippoFunc)

	fmt.Println("Serving: http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}