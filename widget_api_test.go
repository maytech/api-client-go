package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestWidgetApiService_WidgetUploadLinkIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	WidgetUploadLinkResp := &WidgetUploadLinkResp{
		FileId:   testStringValue,
		FilePath: testStringValue,
		ParentId: testStringValue,
		RoleId:   testStringValue,
	}

	mu.HandleFunc("/widget/upload-link/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		WidgetUploadLinkRespJson, err := json.Marshal(*WidgetUploadLinkResp)
		if err != nil {
			return
		}

		w.Write(WidgetUploadLinkRespJson)
	})

	result, response, err := client.WidgetApi.WidgetUploadLinkIdPost(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *WidgetUploadLinkResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestWidgetApiService_WidgetFinalizeUploadIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	WidgetFinalizeUploadResp := &WidgetFinalizeUploadResp{
		Id:       testStringValue,
		ParentId: testStringValue,
		Size:     testDigitalValue,
	}

	mu.HandleFunc("/widget/finalize-upload/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		WidgetFinalizeUploadRespJson, err := json.Marshal(*WidgetFinalizeUploadResp)
		if err != nil {
			return
		}

		w.Write(WidgetFinalizeUploadRespJson)
	})

	result, response, err := client.WidgetApi.WidgetFinalizeUploadIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *WidgetFinalizeUploadResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestWidgetApiService_WidgetMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/widget/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.WidgetApi.WidgetMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
