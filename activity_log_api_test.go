package quatrix

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestActivityLogApiService_TrackingActivityGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	TrackingActivityRespItems := []TrackingActivityRespItems{{
		Id:         testStringValue,
		Email:      testStringValue,
		UserName:   testStringValue,
		UserStatus: testStringValue,
	}}

	mu.HandleFunc("/tracking/activity", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		TrackingActivityRespItemsJson, err := json.Marshal(TrackingActivityRespItems)
		if err != nil {
			return
		}

		w.Write(TrackingActivityRespItemsJson)
	})

	result, response, err := client.ActivityLogApi.TrackingActivityGet(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, TrackingActivityRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestActivityLogApiService_TrackingFilesIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	TrackingFilesRespItems := []TrackingFilesRespItems{{
		Id:        testStringValue,
		Name:      testStringValue,
		Size:      testDigitalValue,
		Type_:     testStringValue,
		Downloads: testDigitalValue,
	}}

	mu.HandleFunc("/tracking/files/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		TrackingFilesRespItemsJson, err := json.Marshal(TrackingFilesRespItems)
		if err != nil {
			return
		}

		w.Write(TrackingFilesRespItemsJson)
	})

	result, response, err := client.ActivityLogApi.TrackingFilesIdGet(context.Background(), "testID")
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, TrackingFilesRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestActivityLogApiService_TrackingCsvGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	TrackingCsvRespItems := []TrackingCsvRespItems{{
		Time:       testStringValue,
		Event:      testStringValue,
		Action:     testStringValue,
		ClientIPs:  testStringValue,
		UserName:   testStringValue,
		Email:      testStringValue,
		UserStatus: testStringValue,
		Protocol:   testStringValue,
	}}

	mu.HandleFunc("/tracking/csv", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		TrackingCsvRespItemsJson, err := json.Marshal(TrackingCsvRespItems)
		if err != nil {
			return
		}

		w.Write(TrackingCsvRespItemsJson)
	})

	result, response, err := client.ActivityLogApi.TrackingCsvGet(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, TrackingCsvRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestActivityLogApiService_TrackingDownloadsIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	TrackingDownloadsRespItems := []TrackingDownloadsRespItems{{
		Ip:        testStringValue,
		Name:      testStringValue,
		Size:      testDigitalValue,
		Duration:  testDigitalValue,
		Time:      testDigitalValue,
		Recipient: testStringValue,
	}}

	mu.HandleFunc("/tracking/downloads/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		TrackingDownloadsRespItemsJson, err := json.Marshal(TrackingDownloadsRespItems)
		if err != nil {
			return
		}

		w.Write(TrackingDownloadsRespItemsJson)
	})

	result, response, err := client.ActivityLogApi.TrackingDownloadsIdGet(context.Background(), "testID")
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, TrackingDownloadsRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
