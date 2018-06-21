package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestMFAApiService_UserSetMfaPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileRemoveMfaResp := &ProfileRemoveMfaResp{
		Status: testStringValue,
	}

	mu.HandleFunc("/user/set-mfa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileRemoveMfaRespJson, err := json.Marshal(*ProfileRemoveMfaResp)
		if err != nil {
			return
		}

		w.Write(ProfileRemoveMfaRespJson)
	})

	body := UserSetMfaReq{
		Ids:      []string{testStringValue},
		AuthType: testStringValue,
		Code:     testStringValue,
	}

	result, response, err := client.MFAApi.UserSetMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileRemoveMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestMFAApiService_ProfileSetMfaPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileSetMfaResp := &ProfileSetMfaResp{
		SessionIds: testStringValue,
	}

	mu.HandleFunc("/profile/set-mfa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileSetMfaRespJson, err := json.Marshal(*ProfileSetMfaResp)
		if err != nil {
			return
		}

		w.Write(ProfileSetMfaRespJson)
	})

	body := ProfileSetMfaReq{
		AuthType: testStringValue,
		Code:     testStringValue,
	}

	result, response, err := client.MFAApi.ProfileSetMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileSetMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ProfileSetMfaReq{}

	_, _, errIncor := client.MFAApi.ProfileSetMfaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestMFAApiService_UserRemoveMfaPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileRemoveMfaResp := &ProfileRemoveMfaResp{
		Status: testStringValue,
	}

	mu.HandleFunc("/user/remove-mfa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileRemoveMfaRespJson, err := json.Marshal(*ProfileRemoveMfaResp)
		if err != nil {
			return
		}

		w.Write(ProfileRemoveMfaRespJson)
	})

	body := UserRemoveMfaReq{
		AuthType: testStringValue,
		Ids:      []string{testStringValue},
	}

	result, response, err := client.MFAApi.UserRemoveMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileRemoveMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UserRemoveMfaReq{}

	_, _, errIncor := client.MFAApi.UserRemoveMfaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestMFAApiService_SessionLoginPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SessionLoginResp := &SessionLoginResp{
		SessionId: testStringValue,
	}

	mu.HandleFunc("/session/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SessionLoginRespJson, err := json.Marshal(*SessionLoginResp)
		if err != nil {
			return
		}

		w.Write(SessionLoginRespJson)
	})

	body := SessionLoginPostResp{
		AuthType: testStringValue,
		Code:     testStringValue,
	}

	result, response, err := client.MFAApi.SessionLoginPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SessionLoginResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := SessionLoginPostResp{}

	_, _, errIncor := client.MFAApi.SessionLoginPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestMFAApiService_ProfileRemoveMfaPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileRemoveMfaResp := &ProfileRemoveMfaResp{
		Status: testStringValue,
	}

	mu.HandleFunc("/profile/remove-mfa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileRemoveMfaRespJson, err := json.Marshal(*ProfileRemoveMfaResp)
		if err != nil {
			return
		}

		w.Write(ProfileRemoveMfaRespJson)
	})

	body := ProfileRemoveMfaReq{
		AuthType: testStringValue,
	}

	result, response, err := client.MFAApi.ProfileRemoveMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileRemoveMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ProfileRemoveMfaReq{}

	_, _, errIncor := client.MFAApi.ProfileRemoveMfaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}
