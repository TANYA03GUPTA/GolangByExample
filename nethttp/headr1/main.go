//w.header().add()
//add key val pair ->res header 
//convert key to canonical form
//w.header.set()
//replace any existing values 
package main
import(
	"encoding/json"
	"net/http"
)
func main(){
	handler := http.HandlerFunc(handlereq)
	http.Handle("/example",handler)
	http.ListenAndServe(":8075", nil)
}

func handlereq(w http.ResponseWriter, r *http.request){
	w.Header().Set("content-type", "application/json")
	w.Header().Add("foo", "bar1")
	w.Header().Add("foo", "bar2")

	resp := make(map[string]string)
	resp["message"] = "success"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp) 
}