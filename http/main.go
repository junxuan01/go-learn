package main

import (
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	txt, err := readfile()
	if err != nil {
		w.Write([]byte("没有找到文件"))
	}
	w.Write(txt)
}

func readfile() ([]byte, error) {
	// 读取文件内容
	data, err := os.ReadFile("test.txt")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9090", nil)

}
