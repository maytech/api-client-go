package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestServiceApiService_ServiceGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ServiceResp := []ServiceResp{{
		Id:   testStringValue,
		Key:  testStringValue,
		Base: testBoolValue,
	}}

	mu.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ServiceRespJson, err := json.Marshal(ServiceResp)
		if err != nil {
			return
		}

		w.Write(ServiceRespJson)
	})

	result, response, err := client.ServiceApi.ServiceGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, ServiceResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestServiceApiService_ServiceMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ServiceResp := &ServiceResp{
		Id:   testStringValue,
		Key:  testStringValue,
		Base: testBoolValue,
	}

	mu.HandleFunc("/service/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ServiceRespJson, err := json.Marshal(*ServiceResp)
		if err != nil {
			return
		}

		w.Write(ServiceRespJson)
	})

	result, response, err := client.ServiceApi.ServiceMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ServiceResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
