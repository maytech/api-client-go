package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestShareApiService_ShareRequestPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ShareRequestResp := &ShareRequestResp{
		Id:    testStringValue,
		Url:   testStringValue,
		JobId: testStringValue,
	}

	mu.HandleFunc("/share/request", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ShareRequestRespJson, err := json.Marshal(*ShareRequestResp)
		if err != nil {
			return
		}

		w.Write(ShareRequestRespJson)
	})

	body := ShareRequestReq{
		ReturnPgpEncrypted: testBoolValue,
	}

	result, response, err := client.ShareApi.ShareRequestPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ShareRequestResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareSendRequestIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	JobResp := &JobResp{
		JobId: testStringValue,
	}

	mu.HandleFunc("/share/send-request/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		JobRespJson, err := json.Marshal(*JobResp)
		if err != nil {
			return
		}

		w.Write(JobRespJson)
	})

	result, response, err := client.ShareApi.ShareSendRequestIdPost(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *JobResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_QuicklinkRevokeIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/quicklink/revoke/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.ShareApi.QuicklinkRevokeIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_FilesReturnUploadLinkIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FilesReturnUploadLinkResp := &FilesReturnUploadLinkResp{
		FileId:    testStringValue,
		UploadKey: testStringValue,
		ParentId:  testStringValue,
	}

	mu.HandleFunc("/files-return/upload-link/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FilesReturnUploadLinkRespJson, err := json.Marshal(*FilesReturnUploadLinkResp)
		if err != nil {
			return
		}

		w.Write(FilesReturnUploadLinkRespJson)
	})

	body := FilesReturnUploadLinkReq{
		ParentId: testStringValue,
		FileSize: testDigitalValue,
	}

	result, response, err := client.ShareApi.FilesReturnUploadLinkIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FilesReturnUploadLinkResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := FilesReturnUploadLinkReq{}

	_, _, errIncor := client.ShareApi.FilesReturnUploadLinkIdPost(context.Background(), testStringValue, bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestShareApiService_QuicklinkLoginPinPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/quicklink/login-pin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	body := QuicklinkLoginPinReq{
		Id:  testStringValue,
		Pin: testStringValue,
	}

	response, err := client.ShareApi.QuicklinkLoginPinPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_TrackingIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	TrackingIdRespItems := []TrackingIdRespItems{{
		Id:      testStringValue,
		Subject: testStringValue,
		Type_:   testStringValue,
	}}

	mu.HandleFunc("/tracking/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		TrackingIdRespItemsJson, err := json.Marshal(TrackingIdRespItems)
		if err != nil {
			return
		}

		w.Write(TrackingIdRespItemsJson)
	})

	result, response, err := client.ShareApi.TrackingIdGet(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, TrackingIdRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ShareCreateResp := &ShareCreateResp{
		Id:        testStringValue,
		Activates: testDigitalValue,
	}

	mu.HandleFunc("/share/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ShareCreateRespJson, err := json.Marshal(*ShareCreateResp)
		if err != nil {
			return
		}

		w.Write(ShareCreateRespJson)
	})

	body := ShareCreateReq{
		FolderId: testStringValue,
		Files:    []string{testStringValue},
		Contacts: []string{testStringValue},
	}

	result, response, err := client.ShareApi.ShareCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ShareCreateResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ShareCreateReq{}

	_, _, errIncor := client.ShareApi.ShareCreatePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestShareApiService_SharePreviewIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/share/preview/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.ShareApi.SharePreviewIdGet(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareDownloadIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/share/download/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.ShareApi.ShareDownloadIdGet(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareDownloadLinkIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/share/download-link/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.ShareApi.ShareDownloadLinkIdPost(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareDownloadLinkIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/share/download-link/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.ShareApi.ShareDownloadLinkIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_QuicklinkCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	QuicklinkCreateResp := &QuicklinkCreateResp{
		Id:        testStringValue,
		Expires:   testDigitalValue,
		Activates: testDigitalValue,
	}

	mu.HandleFunc("/quicklink/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		QuicklinkCreateRespJson, err := json.Marshal(*QuicklinkCreateResp)
		if err != nil {
			return
		}

		w.Write(QuicklinkCreateRespJson)
	})

	body := QuicklinkCreateReq{
		Subject: testStringValue,
		Files:   []string{testStringValue},
		Pin:     testStringValue,
	}

	result, response, err := client.ShareApi.QuicklinkCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *QuicklinkCreateResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := QuicklinkCreateReq{}

	_, _, errIncor := client.ShareApi.QuicklinkCreatePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestShareApiService_ShareRevokeIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/share/revoke/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.ShareApi.ShareRevokeIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_TrackingGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	TrackingRespItems := []TrackingRespItems{{
		Id:      testStringValue,
		Subject: testStringValue,
		Type_:   testStringValue,
	}}

	mu.HandleFunc("/tracking/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		TrackingRespItemsJson, err := json.Marshal(TrackingRespItems)
		if err != nil {
			return
		}

		w.Write(TrackingRespItemsJson)
	})

	result, response, err := client.ShareApi.TrackingGet(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, TrackingRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareDownloadInfoIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ShareDownloadInfoResp := &ShareDownloadInfoResp{
		Id:        testStringValue,
		Message:   testStringValue,
		Subject:   testStringValue,
		Activates: testDigitalValue,
	}

	mu.HandleFunc("/share/download-info/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ShareDownloadInfoRespJson, err := json.Marshal(*ShareDownloadInfoResp)
		if err != nil {
			return
		}

		w.Write(ShareDownloadInfoRespJson)
	})

	result, response, err := client.ShareApi.ShareDownloadInfoIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ShareDownloadInfoResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareRecipientsGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ShareRecipientsResp := &ShareRecipientsResp{
		Id:    testStringValue,
		Name:  testStringValue,
		Email: testStringValue,
	}

	mu.HandleFunc("/share/recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ShareRecipientsRespJson, err := json.Marshal(*ShareRecipientsResp)
		if err != nil {
			return
		}

		w.Write(ShareRecipientsRespJson)
	})

	result, response, err := client.ShareApi.ShareRecipientsGet(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ShareRecipientsResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_FilesReturnSendPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	JobResp := &JobResp{
		JobId: testStringValue,
	}

	mu.HandleFunc("/files-return/send", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		JobRespJson, err := json.Marshal(*JobResp)
		if err != nil {
			return
		}

		w.Write(JobRespJson)
	})

	body := FilesReturnSendReq{
		Id:      testStringValue,
		Files:   []string{testStringValue},
		Message: testStringValue,
	}

	result, response, err := client.ShareApi.FilesReturnSendPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *JobResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := FilesReturnSendReq{}

	_, _, errIncor := client.ShareApi.FilesReturnSendPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestShareApiService_FilesReturnMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FilesReturnMetadataResp := &FilesReturnMetadataResp{
		Id:          testStringValue,
		Subject:     testStringValue,
		SenderName:  testStringValue,
		SenderEmail: testStringValue,
		SenderKey:   testStringValue,
	}

	mu.HandleFunc("/files-return/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FilesReturnMetadataRespJson, err := json.Marshal(*FilesReturnMetadataResp)
		if err != nil {
			return
		}

		w.Write(FilesReturnMetadataRespJson)
	})

	result, response, err := client.ShareApi.FilesReturnMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FilesReturnMetadataResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_FilesReturnMakedirIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FilesReturnMakedirResps := &FilesReturnMakedirResps{
		Id:       testStringValue,
		ParentId: testStringValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/files-return/makedir/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FilesReturnMakedirRespsJson, err := json.Marshal(*FilesReturnMakedirResps)
		if err != nil {
			return
		}

		w.Write(FilesReturnMakedirRespsJson)
	})

	body := FilesReturnMakedirReq{
		Name:     testStringValue,
		ParentId: testStringValue,
	}

	result, response, err := client.ShareApi.FilesReturnMakedirIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FilesReturnMakedirResps)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := FilesReturnMakedirReq{}

	_, _, errIncor := client.ShareApi.FilesReturnMakedirIdPost(context.Background(), testStringValue, bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestShareApiService_ShareLoginPinPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/share/login-pin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	body := ShareLoginPinReq{
		Id:  testStringValue,
		Pin: testStringValue,
	}

	response, err := client.ShareApi.ShareLoginPinPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestShareApiService_ShareFilesIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ShareFilesRespItems := []ShareFilesRespItems{{
		Id:   testStringValue,
		Name: testStringValue,
	}}

	mu.HandleFunc("/share/files/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ShareFilesRespItemsJson, err := json.Marshal(ShareFilesRespItems)
		if err != nil {
			return
		}

		w.Write(ShareFilesRespItemsJson)
	})

	result, response, err := client.ShareApi.ShareFilesIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, ShareFilesRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
