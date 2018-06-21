package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestWOPIApiService_FileWopiTokenIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/file/wopi-token/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.WOPIApi.FileWopiTokenIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
