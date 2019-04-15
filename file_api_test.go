package quatrix

import (
	"net/http"
	"reflect"
	"testing"

	"encoding/json"

	"golang.org/x/net/context"
)

func TestFileApiService_FileInfoIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileInfoResp := &FileInfoResp{
		Id:      testStringValue,
		Created: testDigitalValue,
		Name:    testStringValue,
		Type_:   testStringValue,
	}

	mu.HandleFunc("/file/info/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileInfoRespJson, err := json.Marshal(*FileInfoResp)
		if err != nil {
			return
		}

		w.Write(FileInfoRespJson)
	})

	result, response, err := client.FileApi.FileInfoIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, *FileInfoResp) {
		t.Errorf("response not match, obtained %v but expected %v", *FileInfoResp, result)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("status code not match, obtained %d but expected %d", response.StatusCode, http.StatusOK)
	}
}

func TestFileApiService_FileMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileMetadataGetResp := &FileMetadataGetResp{
		Id:      testStringValue,
		Gid:     testDigitalValue,
		Uid:     testDigitalValue,
		Created: testDigitalValue,
		Name:    testStringValue,
	}

	mu.HandleFunc("/file/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileMetadataGetRespJson, err := json.Marshal(*FileMetadataGetResp)
		if err != nil {
			return
		}

		w.Write(FileMetadataGetRespJson)
	})

	result, response, err := client.FileApi.FileMetadataIdGet(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, *FileMetadataGetResp) {
		t.Errorf("response not match, obtained %v but expected %v", *FileMetadataGetResp, result)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("status code not match, obtained %d but expected %d", response.StatusCode, http.StatusOK)
	}
}

func TestFileApiService_FileMovePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdsResp := &IdsResp{
		Ids: []string{testStringValue},
	}

	mu.HandleFunc("/file/move", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdsRespJson, err := json.Marshal(*IdsResp)
		if err != nil {
			return
		}

		w.Write(IdsRespJson)
	})

	body := CopyMoveFilesReq{
		Ids:     []string{testStringValue},
		Target:  testStringValue,
		Resolve: testBoolValue,
	}

	result, response, err := client.FileApi.FileMovePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdsResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := CopyMoveFilesReq{}

	_, _, errIncor := client.FileApi.FileMovePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestFileApiService_FileDownloadLinkPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdsResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/file/download-link", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdsRespJson, err := json.Marshal(*IdsResp)
		if err != nil {
			return
		}

		w.Write(IdsRespJson)
	})

	body := IdsReq{
		Ids: []string{testStringValue},
	}

	result, response, err := client.FileApi.FileDownloadLinkPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdsResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := IdsReq{}

	_, _, errIncor := client.FileApi.FileDownloadLinkPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestFileApiService_FileDiffIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileDiffResp := &FileDiffResp{
		From: testDigitalValue,
		To:   testDigitalValue,
	}

	mu.HandleFunc("/file/diff/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileDiffRespJson, err := json.Marshal(*FileDiffResp)
		if err != nil {
			return
		}

		w.Write(FileDiffRespJson)
	})

	result, response, err := client.FileApi.FileDiffIdGet(context.Background(), testStringValue, testDigitalValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileDiffResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FilePreviewIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FilePreviewResp := &FilePreviewResp{
		Type_:     testStringValue,
		Uri:	   testStringValue,
	}

	mu.HandleFunc("/file/preview/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FilePreviewRespJson, err := json.Marshal(*FilePreviewResp)
		if err != nil {
			return
		}

		w.Write(FilePreviewRespJson)
	})

	result, response, err := client.FileApi.FilePreviewIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FilePreviewResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileCopyPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	JobResp := &JobResp{
		JobId: testStringValue,
	}

	mu.HandleFunc("/file/copy", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		JobRespJson, err := json.Marshal(*JobResp)
		if err != nil {
			return
		}

		w.Write(JobRespJson)
	})

	body := CopyMoveFilesReq{
		Ids:     []string{testStringValue},
		Target:  testStringValue,
		Resolve: testBoolValue,
	}

	result, response, err := client.FileApi.FileCopyPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *JobResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileMakedirPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileResp := &FileResp{
		Id:       testStringValue,
		Name:     testStringValue,
		ParentId: testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/file/makedir", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileRespJson, err := json.Marshal(*FileResp)
		if err != nil {
			return
		}

		w.Write(FileRespJson)
	})

	body := MakeDirReq{
		Target:  testStringValue,
		Name:    testStringValue,
		Resolve: testBoolValue,
	}

	result, response, err := client.FileApi.FileMakedirPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileDownloadIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/file/download/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.FileApi.FileDownloadIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileDeletePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdsResp := &IdsResp{
		Ids: []string{testStringValue},
	}

	mu.HandleFunc("/file/delete", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdsRespJson, err := json.Marshal(*IdsResp)
		if err != nil {
			return
		}

		w.Write(IdsRespJson)
	})

	body := IdsReq{
		Ids: []string{testStringValue},
	}

	result, response, err := client.FileApi.FileDeletePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdsResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileAddTagIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileTagResp := &FileTagResp{
		Id:     testStringValue,
		FileId: testStringValue,
		Type_:  testStringValue,
		Value:  testStringValue,
	}

	mu.HandleFunc("/file/add-tag/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileTagRespJson, err := json.Marshal(*FileTagResp)
		if err != nil {
			return
		}

		w.Write(FileTagRespJson)
	})

	body := FileAddTagReq{
		Type_: testStringValue,
		Value: testStringValue,
	}

	result, response, err := client.FileApi.FileAddTagIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileTagResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileSearchPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileResp := []FileResp{{
		Id:       testStringValue,
		Name:     testStringValue,
		ParentId: testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}}

	mu.HandleFunc("/file/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileRespJson, err := json.Marshal(FileResp)
		if err != nil {
			return
		}

		w.Write(FileRespJson)
	})

	body := SearchReq{}

	result, response, err := client.FileApi.FileSearchPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, FileResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileRenameIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileRenameResp := &FileRenameResp{
		Id:       testStringValue,
		Name:     testStringValue,
		ParentId: testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/file/rename/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileRespJson, err := json.Marshal(*FileRenameResp)
		if err != nil {
			return
		}

		w.Write(FileRespJson)
	})

	body := FileRenameReq{
		Name:    testStringValue,
		Resolve: testBoolValue,
	}

	result, response, err := client.FileApi.FileRenameIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileRenameResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileMetadataPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileMetadataPostResp := &FileMetadataPostResp{
		Id:       testStringValue,
		Name:     testStringValue,
		ParentId: testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/file/metadata", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileMetadataPostRespJson, err := json.Marshal(*FileMetadataPostResp)
		if err != nil {
			return
		}

		w.Write(FileMetadataPostRespJson)
	})

	body := FileMetadataPostReq{
		Id:    testStringValue,
		Mtime: testDigitalValue,
	}

	result, response, err := client.FileApi.FileMetadataPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileMetadataPostResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := FileMetadataPostReq{}

	_, _, errIncor := client.FileApi.FileMetadataPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestFileApiService_FileMetadataGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileMetadataGetResp := &FileMetadataGetResp{
		Id:       testStringValue,
		Name:     testStringValue,
		ParentId: testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue}

	mu.HandleFunc("/file/metadata", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileMetadataGetRespJson, err := json.Marshal(*FileMetadataGetResp)
		if err != nil {
			return
		}

		w.Write(FileMetadataGetRespJson)
	})

	result, response, err := client.FileApi.FileMetadataGet(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileMetadataGetResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileTagsIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	FileTagResp := []FileTagResp{{
		Id:     testStringValue,
		FileId: testStringValue,
		Type_:  testStringValue,
		Value:  testStringValue,
	}}

	mu.HandleFunc("/file/tags/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		FileTagRespJson, err := json.Marshal(FileTagResp)
		if err != nil {
			return
		}

		w.Write(FileTagRespJson)
	})

	result, response, err := client.FileApi.FileTagsIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, FileTagResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestFileApiService_FileModifyPost(t *testing.T) {
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
		Id:       testStringValue,
		Truncate: testDigitalValue,
		Name:     testStringValue,
	}

	result, response, err := client.FileApi.FileModifyPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *FileModifyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := FileModifyReq{}

	_, _, errIncor := client.FileApi.FileModifyPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}
