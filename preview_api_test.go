package quatrix

import (
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestPreviewApiService_PreviewVideoIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/preview/video/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.PreviewApi.PreviewVideoIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPreviewApiService_PreviewDocumentPdfIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/preview/document_pdf/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.PreviewApi.PreviewDocumentPdfIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPreviewApiService_PreviewImageIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/preview/image/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.PreviewApi.PreviewImageIdGet(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
