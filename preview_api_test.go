package quatrix

import (
	"testing"
	"net/http"
)

func TestPreviewApiService_PreviewIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/preview/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		
		// ... add another response parameters 
	})
	
	// ... write quatrix here
}

func TestPreviewApiService_FilePreviewIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/file/preview/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		
		// ... add another response parameters 
	})
	
	// ... write quatrix here
}

