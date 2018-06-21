package quatrix

import (
	"net/http"
	"testing"

	"encoding/json"

	"golang.org/x/net/context"
)

func TestAuthApiService_SessionUnblockCaptchaPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SessionUnblockCaptchaResp := &SessionUnblockCaptchaResp{
		Valid: testBoolValue}

	mu.HandleFunc("/session/unblock-captcha", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SessionUnblockCaptchaRespJson, err := json.Marshal(*SessionUnblockCaptchaResp)
		if err != nil {
			return
		}

		w.Write(SessionUnblockCaptchaRespJson)
	})

	body := UnblockCaptchaReq{
		Token: testStringValue,
	}

	result, response, err := client.AuthApi.SessionUnblockCaptchaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SessionUnblockCaptchaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UnblockCaptchaReq{}

	_, _, errIncor := client.AuthApi.SessionUnblockCaptchaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestAuthApiService_SessionLogoutGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/session/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.AuthApi.SessionLogoutGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestAuthApiService_SessionKeepaliveGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/session/keepalive", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.AuthApi.SessionKeepaliveGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestAuthApiService_SessionLoginPost(t *testing.T) {
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

	result, response, err := client.AuthApi.SessionLoginPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SessionLoginResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := SessionLoginPostResp{}

	_, _, errIncor := client.AuthApi.SessionLoginPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestAuthApiService_SessionLoginGet(t *testing.T) {
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

	result, response, err := client.AuthApi.SessionLoginGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SessionLoginResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
