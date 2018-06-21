package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestUploadApiService_UploadLinkPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileModifyResp := &FileModifyResp{
		Name:      testStringValue,
		FileId:    testStringValue,
		ParentId:  testStringValue,
		UploadKey: testStringValue,
	}

	mu.HandleFunc("/upload/link", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileModifyRespJson, err := json.Marshal(*FileModifyResp)
		if err != nil {
			return
		}

		w.Write(FileModifyRespJson)
	})

	body := UploadLinkReq{
		Name:     testStringValue,
		ParentId: testStringValue,
	}

	result, response, err := client.UploadApi.UploadLinkPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileModifyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UploadLinkReq{}

	_, _, errIncor := client.UploadApi.UploadLinkPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestUploadApiService_UploadFinalizeIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UploadFinalizeResp := &UploadFinalizeResp{
		Id:       testStringValue,
		ParentId: testStringValue,
		Size:     testDigitalValue,
	}

	mu.HandleFunc("/upload/finalize/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UploadFinalizeRespJson, err := json.Marshal(*UploadFinalizeResp)
		if err != nil {
			return
		}

		w.Write(UploadFinalizeRespJson)
	})

	result, response, err := client.UploadApi.UploadFinalizeIdGet(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *UploadFinalizeResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestUploadApiService_SettingsUploadLogoLinkGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SettingsUploadLogoLinkResp := &SettingsUploadLogoLinkResp{
		UploadKey: testStringValue,
	}

	mu.HandleFunc("/settings/upload-logo-link", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SettingsUploadLogoLinkRespJson, err := json.Marshal(*SettingsUploadLogoLinkResp)
		if err != nil {
			return
		}

		w.Write(SettingsUploadLogoLinkRespJson)
	})

	result, response, err := client.UploadApi.SettingsUploadLogoLinkGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SettingsUploadLogoLinkResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestUploadApiService_FileModifyPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileModifyResp := &FileModifyResp{
		Name:      testStringValue,
		FileId:    testStringValue,
		ParentId:  testStringValue,
		UploadKey: testStringValue,
	}

	mu.HandleFunc("/file/modify", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileModifyRespJson, err := json.Marshal(*FileModifyResp)
		if err != nil {
			return
		}

		w.Write(FileModifyRespJson)
	})

	body := FileModifyReq{
		Id:   testStringValue,
		Name: testStringValue,
	}

	result, response, err := client.UploadApi.FileModifyPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileModifyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := FileModifyReq{}

	_, _, errIncor := client.UploadApi.FileModifyPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}
