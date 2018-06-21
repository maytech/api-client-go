package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestProfileApiService_ProfileSetPasswordPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileSetPasswordResp := &ProfileSetPasswordResp{
		Email: testStringValue,
	}

	mu.HandleFunc("/profile/set-password", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileSetPasswordRespJson, err := json.Marshal(*ProfileSetPasswordResp)
		if err != nil {
			return
		}

		w.Write(ProfileSetPasswordRespJson)
	})

	body := ProfileSetPasswordReq{
		Password:    testStringValue,
		NewPassword: testStringValue,
	}

	result, response, err := client.ProfileApi.ProfileSetPasswordPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileSetPasswordResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ProfileSetPasswordReq{}

	_, _, errIncor := client.ProfileApi.ProfileSetPasswordPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestProfileApiService_ProfileGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileResp := &ProfileResp{
		Id:         testStringValue,
		Name:       testStringValue,
		Email:      testStringValue,
		SuperAdmin: testStringValue,
		HomeId:     testStringValue,
		Operations: testDigitalValue,
	}

	mu.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileRespJson, err := json.Marshal(*ProfileResp)
		if err != nil {
			return
		}

		w.Write(ProfileRespJson)
	})

	result, response, err := client.ProfileApi.ProfileGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestProfileApiService_ProfileSetMfaPost(t *testing.T) {
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

	result, response, err := client.ProfileApi.ProfileSetMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileSetMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ProfileSetMfaReq{}

	_, _, errIncor := client.ProfileApi.ProfileSetMfaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestProfileApiService_ProfileInfoGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileInfoResp := &ProfileInfoResp{
		UserUsed:  testDigitalValue,
		UserLimit: testDigitalValue,
		AccLimit:  testDigitalValue,
		AccUsed:   testDigitalValue,
	}

	mu.HandleFunc("/profile/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileSetMfaRespJson, err := json.Marshal(*ProfileInfoResp)
		if err != nil {
			return
		}

		w.Write(ProfileSetMfaRespJson)
	})

	result, response, err := client.ProfileApi.ProfileInfoGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileInfoResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestProfileApiService_ProfileSetPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileSetResp := &ProfileSetResp{
		Name:             testStringValue,
		Email:            testStringValue,
		Language:         testStringValue,
		MessageSignature: testStringValue,
	}

	mu.HandleFunc("/profile/set", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileSetRespJson, err := json.Marshal(*ProfileSetResp)
		if err != nil {
			return
		}

		w.Write(ProfileSetRespJson)
	})

	result, response, err := client.ProfileApi.ProfileSetPost(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileSetResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestProfileApiService_Profile2faGenerateGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/profile/2fa/generate", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.ProfileApi.Profile2faGenerateGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestProfileApiService_ProfileRemoveMfaPost(t *testing.T) {
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

	result, response, err := client.ProfileApi.ProfileRemoveMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileRemoveMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ProfileRemoveMfaReq{}

	_, _, errIncor := client.ProfileApi.ProfileRemoveMfaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}
