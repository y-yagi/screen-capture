package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	log.Print("Listening on http://localhost:8888/")
	exec.Command("xdg-open", "http://localhost:8888/").Run()
	log.Printf("Running command and waiting for it to finish...")
	http.ListenAndServe(":8888", nil)
}

//go:embed static
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		return http.FS(os.DirFS("static"))
	}

	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
