package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	title := "Gabtec Shared Files"

	shareFolder := os.Getenv("SHARE_FOLDER")
	if shareFolder == "" {
		log.Fatal("SHARE_FOLDER environment variable is not set")
	}

	// Verify if the folder exists
	info, err := os.Stat(shareFolder)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("SHARE_FOLDER does not exist: %s", shareFolder)
		}
		log.Fatalf("Error checking SHARE_FOLDER: %v", err)
	}
	if !info.IsDir() {
		log.Fatalf("SHARE_FOLDER is not a directory: %s", shareFolder)
	}

	// Set up the custom handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fullPath := filepath.Join(shareFolder, filepath.Clean(r.URL.Path))

		// Check if the path exists
		info, err := os.Stat(fullPath)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// If it's a directory, serve a custom listing
		if info.IsDir() {
			files, err := os.ReadDir(fullPath)
			if err != nil {
				http.Error(w, "Error reading directory", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<html><head><title>File Server</title><style>body{font-family:sans-serif;padding:20px;} h1{color:#333;} ul{list-style:none;padding:0;} li{margin:5px 0;} a{text-decoration:none;color:#0066cc;} a:hover{text-decoration:underline;}</style></head><body>")
			fmt.Fprintf(w, "<h1>📁 %s</h1>", title)
			fmt.Fprintf(w, "<p>Serving from: <code>%s</code></p>", shareFolder)
			fmt.Fprintf(w, "<hr>")
			fmt.Fprintf(w, "<ul>")

			// Add a "parent directory" link if not at the root
			if r.URL.Path != "/" {
				fmt.Fprintf(w, "<li><a href=\"..\">.. (Parent Directory)</a></li>")
			}

			for _, file := range files {
				name := file.Name()
				if file.IsDir() {
					name += "/"
				}
				// Ensure the link is relative to the current path
				link := name
				fmt.Fprintf(w, "<li><a href=\"%s\">%s</a></li>", link, name)
			}
			fmt.Fprintf(w, "</ul>")
			fmt.Fprintf(w, "<hr><footer style='font-size:0.8em;color:#666;'>Go File Server: %s</footer>", getVersion())
			fmt.Fprintf(w, "</body></html>")
			return
		}

		// If it's a file, use the default server behavior
		http.ServeFile(w, r, fullPath)
	})

	// Use a default port or allow it to be configured via environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Serving files from %s on port %s...", shareFolder, port)
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
