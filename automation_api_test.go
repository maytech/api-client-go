package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestAutomationApiService_AutomationCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AutomationResp := &AutomationResp{
		Id:      testStringValue,
		OwnerId: testStringValue,
		FileId:  testStringValue,
		Created: testDigitalValue,
	}

	mu.HandleFunc("/automation/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		AutomationRespJson, err := json.Marshal(*AutomationResp)
		if err != nil {
			return
		}

		w.Write(AutomationRespJson)
	})

	body := AutomationCreateReq{
		FileId: testStringValue,
		Action: testStringValue,
		Status: testStringValue,
	}

	result, response, err := client.AutomationApi.AutomationCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *AutomationResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := AutomationCreateReq{}

	_, _, errIncor := client.AutomationApi.AutomationCreatePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestAutomationApiService_AutomationMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AutomationResp := &AutomationResp{
		Id:      testStringValue,
		OwnerId: testStringValue,
		FileId:  testStringValue,
		Created: testDigitalValue,
	}

	mu.HandleFunc("/automation/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		AutomationRespJson, err := json.Marshal(*AutomationResp)
		if err != nil {
			return
		}

		w.Write(AutomationRespJson)
	})

	result, response, err := client.AutomationApi.AutomationMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *AutomationResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

}

func TestAutomationApiService_AutomationEditPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AutomationResp := []AutomationResp{{
		Id:      testStringValue,
		OwnerId: testStringValue,
		FileId:  testStringValue,
		Created: testDigitalValue,
	}}

	mu.HandleFunc("/automation/edit/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		AutomationRespJson, err := json.Marshal(AutomationResp)
		if err != nil {
			return
		}

		w.Write(AutomationRespJson)
	})

	body := AutomationEditReq{
		Ids:    []string{testStringValue},
		Action: testStringValue,
	}

	result, response, err := client.AutomationApi.AutomationEditPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, AutomationResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := AutomationEditReq{}

	_, _, errIncor := client.AutomationApi.AutomationEditPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestAutomationApiService_AutomationDeletePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AutomationDeleteResp := &AutomationDeleteResp{
		Ids: []string{testStringValue},
	}

	mu.HandleFunc("/automation/delete", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		AutomationDeleteRespJson, err := json.Marshal(*AutomationDeleteResp)
		if err != nil {
			return
		}

		w.Write(AutomationDeleteRespJson)
	})

	body := IdsReq{
		Ids: []string{testStringValue},
	}

	result, response, err := client.AutomationApi.AutomationDeletePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *AutomationDeleteResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := IdsReq{}

	_, _, errIncor := client.AutomationApi.AutomationDeletePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestAutomationApiService_AutomationGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	AutomationResp := []AutomationResp{{
		Id:      testStringValue,
		OwnerId: testStringValue,
		FileId:  testStringValue,
		Created: testDigitalValue,
	}}

	mu.HandleFunc("/automation", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		AutomationRespJson, err := json.Marshal(AutomationResp)
		if err != nil {
			return
		}

		w.Write(AutomationRespJson)
	})

	result, response, err := client.AutomationApi.AutomationGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, AutomationResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
