package main
import (
	"fmt"
//	"log"
	"net/http"
)
type Hello struct{}
//这样写参数好记录这个修改，如果写到一行那么修改了好发现
func (h *Hello) ServeHTTP(
	w http.ResponseWriter, 
	r *http.Request) {
		fmt.Fprint(w, "hello!")
}

	
func main(){
//	var h Hello
//	err := http.ListenAndServe("localhost:4000", h)
//	if err != nil {
//		log.Fatal(err)
//	}
	http.Handle("/", &Hello{})
	http.ListenAndServe(":12345", nil)
}
