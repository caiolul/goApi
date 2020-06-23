package main

import (
    "fmt"
    "net/http"
)

func main(){
    http.HandleFunc("/api/wellcome/", func(res http.ResponseWriter, req *http.Request){
     fmt.Fprintf(res, "hello word")
    })
    var port = ":8000"

    fmt.Println("Sever runing it http://localhost"+port+"/api/wellcome/")

    http.ListenAndServe(port, nil)
}
