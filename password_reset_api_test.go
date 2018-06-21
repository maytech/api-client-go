package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestPasswordResetApiService_ResetPasswordRequestPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ResetPasswordRequestResp := &ResetPasswordRequestResp{
		Email: []string{testStringValue},
	}

	mu.HandleFunc("/reset-password/request", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ResetPasswordRequestRespJson, err := json.Marshal(*ResetPasswordRequestResp)
		if err != nil {
			return
		}

		w.Write(ResetPasswordRequestRespJson)
	})

	body := ResetPasswordRequestReq{
		Email: []string{testStringValue},
	}

	result, response, err := client.PasswordResetApi.ResetPasswordRequestPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ResetPasswordRequestResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPasswordResetApiService_ResetPasswordResetIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/reset-password/reset/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	body := ResetPasswordResetReq{
		Password: testStringValue,
	}

	result, response, err := client.PasswordResetApi.ResetPasswordResetIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ResetPasswordResetReq{}

	_, _, errIncor := client.PasswordResetApi.ResetPasswordResetIdPost(context.Background(), testStringValue, bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestPasswordResetApiService_ResetPasswordMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ResetPasswordMetadataResp := &ResetPasswordMetadataResp{
		Id:               testStringValue,
		MultipleAccounts: testBoolValue,
	}

	mu.HandleFunc("/reset-password/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ResetPasswordMetadataRespJson, err := json.Marshal(*ResetPasswordMetadataResp)
		if err != nil {
			return
		}

		w.Write(ResetPasswordMetadataRespJson)
	})

	result, response, err := client.PasswordResetApi.ResetPasswordMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ResetPasswordMetadataResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
