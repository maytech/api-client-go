package quatrix

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestAccountApiService_AccountLogoGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/account/logo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	response, err := client.AccountApi.AccountLogoGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestAccountApiService_AccountRolesGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AccountRolesRespItems := []AccountRolesRespItems{{
		Id:    testStringValue,
		Email: testStringValue,
		Name:  testStringValue}}

	mu.HandleFunc("/account/roles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		AccountRolesRespItemsJson, err := json.Marshal(AccountRolesRespItems)
		if err != nil {
			return
		}

		w.Write(AccountRolesRespItemsJson)
	})

	result, response, err := client.AccountApi.AccountRolesGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, AccountRolesRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestAccountApiService_AccountInfoGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AccountInfoResp := &AccountInfoResp{
		Id:             testStringValue,
		Key:            testStringValue,
		Base:           testBoolValue,
		ServerLocation: testStringValue}

	mu.HandleFunc("/account/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		AccountInfoRespJson, err := json.Marshal(*AccountInfoResp)
		if err != nil {
			return
		}

		w.Write(AccountInfoRespJson)
	})

	result, response, err := client.AccountApi.AccountInfoGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *AccountInfoResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestAccountApiService_AccountMetadataGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AccountMetadataResp := &AccountMetadataResp{
		Status:      testStringValue,
		Title:       testStringValue,
		Language:    testStringValue,
		DefaultLogo: testBoolValue}

	mu.HandleFunc("/account/metadata", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		AccountMetadataRespJson, err := json.Marshal(*AccountMetadataResp)
		if err != nil {
			return
		}

		w.Write(AccountMetadataRespJson)
	})

	result, response, err := client.AccountApi.AccountMetadataGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *AccountMetadataResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestAccountApiService_AccountListGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AccountListRespItems := []AccountListRespItems{{
		Id:       testStringValue,
		Hostname: testStringValue,
		Plan:     testStringValue,
		Status:   testStringValue,
		Current:  testBoolValue,
	}}

	mu.HandleFunc("/account/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		AccountListRespItemsJson, err := json.Marshal(AccountListRespItems)
		if err != nil {
			return
		}

		w.Write(AccountListRespItemsJson)
	})

	result, response, err := client.AccountApi.AccountListGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, AccountListRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
