/*
 * @Date: 2022.02.14 15:37
 * @Description: Omit
 * @LastEditors: Rustle Karl
 * @LastEditTime: 2022.02.14 15:37
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"golang.org/x/net/webdav"
)

func main() {
	var port, directory string

	flag.StringVar(&port, "P", "7654", "http port")
	flag.StringVar(&directory, "D", ".", "service directory")

	flag.Parse()

	hostPort := ":" + port

	srv := &webdav.Handler{
		FileSystem: webdav.Dir(directory),
		LockSystem: webdav.NewMemLS(),
		Logger: func(request *http.Request, err error) {
			log.Printf(
				"%s %s %s\n",
				request.RemoteAddr,
				request.Method,
				request.URL,
			)
			if err != nil {
				log.Printf("%s\n", err)
				req, _ := httputil.DumpRequest(request, true)
				log.Printf("%s\n", req)
			}
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && ListDirectory(srv.FileSystem, w, r.URL.Path) {
			log.Printf("%s %s\n", r.Method, r.URL.Path)
			return
		}
		srv.ServeHTTP(w, r)
	})

	log.Printf("initialized: port='%s', directory='%s'\n", port, directory)

	_ = http.ListenAndServe(hostPort, nil)
}

func ListDirectory(fs webdav.FileSystem, w http.ResponseWriter, uri string) bool {
	if file, err := fs.OpenFile(context.Background(), uri, os.O_RDONLY, 0); err == nil {
		defer file.Close()

		if fi, er := file.Stat(); er != nil || fi == nil || !fi.IsDir() {
			return false
		}

		items, er := file.Readdir(-1)
		if er != nil {
			return false
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		_, _ = fmt.Fprintf(w, "<pre>\n")
		for i := range items {
			n := items[i].Name()
			if items[i].IsDir() {
				n += "/"
			}
			_, _ = fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", n, n)
		}
		_, _ = fmt.Fprintf(w, "</pre>\n")

		return true
	}

	return false
}
