package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestLanguageApiService_LanguagesDefaultGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	LanguagesDefaultResp := &LanguagesDefaultResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/languages/default", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		LanguagesDefaultRespJson, err := json.Marshal(*LanguagesDefaultResp)
		if err != nil {
			return
		}

		w.Write(LanguagesDefaultRespJson)
	})

	result, response, err := client.LanguageApi.LanguagesDefaultGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *LanguagesDefaultResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestLanguageApiService_LanguagesGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	LanguagesRespItems := []LanguagesRespItems{{
		Id:       testStringValue,
		Name:     testStringValue,
		Filename: testStringValue,
	}}

	mu.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		LanguagesRespItemsJson, err := json.Marshal(LanguagesRespItems)
		if err != nil {
			return
		}

		w.Write(LanguagesRespItemsJson)
	})

	result, response, err := client.LanguageApi.LanguagesGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, LanguagesRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
