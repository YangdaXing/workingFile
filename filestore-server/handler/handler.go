package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter,r *http.Request){
	 if r.Method == "GET"{
	 	//返回上传html
	 	data,err := ioutil.ReadFile("./static/view/index.html")
		 if err != nil {
			 io.WriteString(w,"err")
			 return
		 }
		 io.WriteString(w,string(data))
	 }else if r.Method== "POST" {
	 	//接收文件流及存储到本地目录
		 file,head,err := r.FormFile("file")
		 if err != nil {
			 fmt.Println("Failed to get data,Error:",err)
			 return
		 }
		 defer file.Close()
		 newFile , err := os.Create("~/"+head.Filename)
		 if err != nil {
			 fmt.Println("Failed to create file,err",err)
			 return
		 }
		 defer newFile.Close()
		 _,err = io.Copy(newFile,file)
		 if err != nil {
			 fmt.Println("Failed to save data into file,err :",err)
			 return
		 }
		 http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	 }
}

func UploadSucHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Upload finished!")
}