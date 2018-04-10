// file http_test1.go
package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("%s\n", r.PostFormValue("test2"))
	fmt.Printf("%s\n", r.PostFormValue("pic"))
	fmt.Println("PATH: ", r.URL.Path)
	fmt.Println("SCHEME: ", r.URL.Scheme)
	fmt.Println("METHOD: ", r.Method)
	fmt.Println()

	fmt.Fprintf(w, "<h1>Index Page</h1><form action='/uploadfile' method='post' enctype='multipart/form-data'><input type='text' id='test2' name='test2' placeholder='input anything' /><br/><input type='file' name='pic'/><input type='submit' value='Submit'/></form>")

	w.Write([]byte("hello"))
}

func UpLoadFile(w http.ResponseWriter, r *http.Request){
	//获取文件内容 要这样获取
	file, head, err := r.FormFile("pic")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//创建文件
	fW, err := os.Create("./upload/" + head.Filename)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return
	}
	w.Write([]byte("SUCCESS"))
}

func main() {
	http.HandleFunc("/test", HandleIndex)
	http.HandleFunc("/uploadfile",UpLoadFile)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
