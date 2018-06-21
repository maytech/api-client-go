package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestSSHKeyApiService_SshKeyEditPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SshKeyResp := &SshKeyResp{
		Id:          testStringValue,
		Caption:     testStringValue,
		Fingerprint: testStringValue,
	}

	mu.HandleFunc("/ssh-key/edit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SshKeyRespJson, err := json.Marshal(*SshKeyResp)
		if err != nil {
			return
		}

		w.Write(SshKeyRespJson)
	})

	body := SshKeyEditReq{
		Id:      testStringValue,
		Caption: testStringValue,
		Status:  testStringValue,
		Expires: testDigitalValue,
	}

	result, response, err := client.SSHKeyApi.SshKeyEditPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SshKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := SshKeyEditReq{}

	_, _, errIncor := client.SSHKeyApi.SshKeyEditPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestSSHKeyApiService_SshKeyGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SshKeyResp := []SshKeyResp{{
		Id:          testStringValue,
		Caption:     testStringValue,
		Fingerprint: testStringValue,
	}}

	mu.HandleFunc("/ssh-key", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SshKeyRespJson, err := json.Marshal(SshKeyResp)
		if err != nil {
			return
		}

		w.Write(SshKeyRespJson)
	})

	result, response, err := client.SSHKeyApi.SshKeyGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, SshKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestSSHKeyApiService_SshKeyDeletePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/ssh-key/delete", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	body := SshKeyDeleteReq{
		Id: testStringValue,
	}

	result, response, err := client.SSHKeyApi.SshKeyDeletePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := SshKeyDeleteReq{}

	_, _, errIncor := client.SSHKeyApi.SshKeyDeletePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestSSHKeyApiService_SshKeyCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SshKeyResp := &SshKeyResp{
		Id:          testStringValue,
		Caption:     testStringValue,
		Fingerprint: testStringValue,
	}

	mu.HandleFunc("/ssh-key/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SshKeyRespJson, err := json.Marshal(*SshKeyResp)
		if err != nil {
			return
		}

		w.Write(SshKeyRespJson)
	})

	body := SshKeyCreateReq{
		Caption: testStringValue,
		Key:     testStringValue,
	}

	result, response, err := client.SSHKeyApi.SshKeyCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SshKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := SshKeyCreateReq{}

	_, _, errIncor := client.SSHKeyApi.SshKeyCreatePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestSSHKeyApiService_SshKeyDeleteIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/ssh-key/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.SSHKeyApi.SshKeyDeleteIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestSSHKeyApiService_SshKeyMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SshKeyResp := &SshKeyResp{
		Id:          testStringValue,
		Caption:     testStringValue,
		Fingerprint: testStringValue,
	}

	mu.HandleFunc("/ssh-key/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SshKeyRespJson, err := json.Marshal(*SshKeyResp)
		if err != nil {
			return
		}

		w.Write(SshKeyRespJson)
	})

	result, response, err := client.SSHKeyApi.SshKeyMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SshKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
