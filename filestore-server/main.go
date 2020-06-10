package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("failed to start server,err:",err)

	}
}
