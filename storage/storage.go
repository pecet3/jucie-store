package storage

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pecet3/my-api/auth"
)

const ImagesTable = `
create table if not exists images (
	id integer primary key autoincrement,
	url text not null,
	user_id integer default -1,
	created_at timestamp default current_timestamp,
	foreign key (user_id) references user(id) on delete set null
);
`

type storage struct {
	db      *sql.DB
	methods StorageServices
}

func Run(srv *http.ServeMux, db *sql.DB, as *auth.SessionStore) {
	s := &storage{
		db:      db,
		methods: &Services{},
	}
	srv.Handle("POST /upload-image", as.Authorize(s.handleUpload))
	srv.HandleFunc("GET /upload-image", s.serveUploadHTML)
	srv.HandleFunc("/images/", s.serveFileHandler)
}

func (s storage) handleUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("User is uploading a image")
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}

	path, err := s.methods.AddImage(file, header, "")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Saving or compressing a file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", path)
}

func (s storage) serveUploadHTML(w http.ResponseWriter, r *http.Request) {

	htmlFile := "./storage/upload.html"
	file, err := os.ReadFile(htmlFile)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(file)
}

func (s storage) serveFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	log.Println("Accessing to resource:", filePath)
	if _, err := os.Stat("./uploads/" + filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	ext := filepath.Ext(filePath)
	mimeTypes := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
	}

	mime, exists := mimeTypes[ext]
	if !exists {
		http.Error(w, "Invalid image type", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", mime)
	http.ServeFile(w, r, "./uploads/"+filePath)
}
