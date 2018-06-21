package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestInboxApiService_InboxGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	InboxRespItems := []InboxRespItems{{
		Id:         testStringValue,
		SenderName: testStringValue,
		Activates:  testDigitalValue,
		IsReply:    testBoolValue,
	}}

	mu.HandleFunc("/inbox", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		InboxRespItemsJson, err := json.Marshal(InboxRespItems)
		if err != nil {
			return
		}

		w.Write(InboxRespItemsJson)
	})

	result, response, err := client.InboxApi.InboxGet(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, InboxRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
