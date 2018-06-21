package quatrix

import (
	"net/http"
	"testing"

	"encoding/json"

	"golang.org/x/net/context"
)

func TestBillingApiService_BillingUpgradePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	BillingUpgradeResp := &BillingUpgradeResp{
		Status:           testStringValue,
		UpgradeInvoiceId: testStringValue,
		ErrorText:        testStringValue,
	}

	mu.HandleFunc("/billing/upgrade", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		BillingUpgradeRespJson, err := json.Marshal(*BillingUpgradeResp)
		if err != nil {
			return
		}

		w.Write(BillingUpgradeRespJson)
	})

	body := BillingUpgradeReq{
		Currency:      testStringValue,
		RenewalPeriod: testStringValue,
	}

	result, response, err := client.BillingApi.BillingUpgradePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *BillingUpgradeResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := BillingUpgradeReq{}

	_, _, errIncor := client.BillingApi.BillingUpgradePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}
