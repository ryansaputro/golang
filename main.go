package main

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/ryansaputro/golang/api"
)

const FSPATH = "./build/"

func main() {
	fs := http.FileServer(http.Dir(FSPATH))

	api.Run()
	http.HandleFunc("/login_view", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If the requested file exists then return if; otherwise return index.html (fileserver default page)
		if r.URL.Path != "/" {
			fullPath := FSPATH + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
			_, err := os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					panic(err)
				}
				// Requested file does not exist so we return the default (resolves to index.html)
				r.URL.Path = "/"
			}
		}
		fs.ServeHTTP(w, r)
	})
}
