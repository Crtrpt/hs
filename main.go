package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	cd, _ := os.Getwd()
	var addr = flag.String("addr", ":8088", "监听的ip和地址")
	var root = flag.String("root", cd, "root目录")
	var index = flag.String("index", "index.html", "index文件")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()
		pathSeg := strings.Split(path, "/")
		filePath := path
		if len(pathSeg) == 0 || pathSeg[len(pathSeg)-1] == "" {
			filePath = path + *index
		}
		file, err := os.Open(*root + filePath)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(404)
			fmt.Fprint(w, "file not found, path:"+filePath)
		} else {
			_, err := io.Copy(w, file)
			if err != nil {
				fmt.Errorf("file copy error", err)
			}
		}
		fmt.Printf("[%s] %v \r\n", time.Now().Format("2006-01-02 15:04:05"), filePath)
	})
	fmt.Printf("addr: %v \r\n", *addr)
	fmt.Printf("root: %v \r\n", *root)
	fmt.Printf("index: %v \r\n", *index)

	// listen to port
	http.ListenAndServe(*addr, nil)
}
