package quatrix

import (
	"net/http"
	"testing"

	"encoding/json"

	"golang.org/x/net/context"
)

func TestCheckApiService_CheckPingGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	CheckPingResp := &CheckPingResp{
		Response: testStringValue,
	}

	mu.HandleFunc("/check/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		CheckPingRespJson, err := json.Marshal(*CheckPingResp)
		if err != nil {
			return
		}

		w.Write(CheckPingRespJson)
	})

	result, response, err := client.CheckApi.CheckPingGet(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *CheckPingResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
