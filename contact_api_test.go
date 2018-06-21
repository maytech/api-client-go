package quatrix

import (
	"net/http"
	"testing"

	"encoding/json"

	"golang.org/x/net/context"
)

func TestContactApiService_ContactDeletePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ContactDeleteRespItems := []ContactDeleteRespItems{
		{Id: testStringValue},
	}

	mu.HandleFunc("/contact/delete", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ContactDeleteRespItemsJson, err := json.Marshal(ContactDeleteRespItems)
		if err != nil {
			return
		}

		w.Write(ContactDeleteRespItemsJson)
	})

	body := IdsReq{
		Ids: []string{testStringValue},
	}

	result, response, err := client.ContactApi.ContactDeletePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, ContactDeleteRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := IdsReq{}

	_, _, errIncor := client.ContactApi.ContactDeletePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestContactApiService_ContactUpgradeIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserResp := &UserResp{
		Id:         testStringValue,
		ParentId:   testStringValue,
		HomeId:     testStringValue,
		HomeName:   testStringValue,
		SuperAdmin: testStringValue,
	}

	mu.HandleFunc("/contact/upgrade/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserRespJson, err := json.Marshal(*UserResp)
		if err != nil {
			return
		}

		w.Write(UserRespJson)
	})

	result, response, err := client.ContactApi.ContactUpgradeIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *UserResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestContactApiService_ContactPgpKeyIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyResp := &PgpKeyResp{
		Id:      testStringValue,
		Name:    testStringValue,
		Created: testDigitalValue,
	}

	mu.HandleFunc("/contact/pgp-key/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRespJson, err := json.Marshal(*PgpKeyResp)
		if err != nil {
			return
		}

		w.Write(PgpKeyRespJson)
	})

	result, response, err := client.ContactApi.ContactPgpKeyIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PgpKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestContactApiService_ContactMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ContactResp := &ContactResp{
		Id:    testStringValue,
		Name:  testStringValue,
		Email: testStringValue,
		Type_: testStringValue,
	}

	mu.HandleFunc("/contact/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ContactRespJson, err := json.Marshal(*ContactResp)
		if err != nil {
			return
		}

		w.Write(ContactRespJson)
	})

	result, response, err := client.ContactApi.ContactMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ContactResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestContactApiService_ContactGroupGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ContactGroupRespItems := []ContactGroupRespItems{{
		Id:       testStringValue,
		Name:     testStringValue,
		Global:   testBoolValue,
		Readonly: testBoolValue,
	}}

	mu.HandleFunc("/contact/group", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ContactGroupRespItemsJson, err := json.Marshal(ContactGroupRespItems)
		if err != nil {
			return
		}

		w.Write(ContactGroupRespItemsJson)
	})

	result, response, err := client.ContactApi.ContactGroupGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, ContactGroupRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestContactApiService_ContactGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ContactResp := []ContactResp{{
		Id:    testStringValue,
		Name:  testStringValue,
		Email: testStringValue,
		Type_: testStringValue,
	}}

	mu.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ContactRespJson, err := json.Marshal(ContactResp)
		if err != nil {
			return
		}

		w.Write(ContactRespJson)
	})

	result, response, err := client.ContactApi.ContactGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, ContactResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestContactApiService_ContactCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ContactResp := &ContactResp{
		Id:    testStringValue,
		Name:  testStringValue,
		Email: testStringValue,
		Type_: testStringValue,
	}

	mu.HandleFunc("/contact/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ContactRespJson, err := json.Marshal(*ContactResp)
		if err != nil {
			return
		}

		w.Write(ContactRespJson)
	})

	body := ContactCreateReq{
		Name:  testStringValue,
		Email: testStringValue,
	}

	result, response, err := client.ContactApi.ContactCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ContactResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := ContactCreateReq{}

	_, _, errIncor := client.ContactApi.ContactCreatePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestContactApiService_ContactEditIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ContactResp := &ContactResp{
		Id:    testStringValue,
		Name:  testStringValue,
		Email: testStringValue,
		Type_: testStringValue,
	}

	mu.HandleFunc("/contact/edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ContactRespJson, err := json.Marshal(*ContactResp)
		if err != nil {
			return
		}

		w.Write(ContactRespJson)
	})

	result, response, err := client.ContactApi.ContactEditIdPost(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ContactResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
