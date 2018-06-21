package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestSiteSettingsApiService_SettingsAuthMethodsGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SettingsAuthMethodsRespItems := []SettingsAuthMethodsRespItems{{
		Rank: testDigitalValue,
		Id:   testStringValue,
		Key:  testStringValue,
	}}

	mu.HandleFunc("/settings/auth-methods", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SettingsAuthMethodsRespItemsJson, err := json.Marshal(SettingsAuthMethodsRespItems)
		if err != nil {
			return
		}

		w.Write(SettingsAuthMethodsRespItemsJson)
	})

	result, response, err := client.SiteSettingsApi.SettingsAuthMethodsGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, SettingsAuthMethodsRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestSiteSettingsApiService_SettingsGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SettingsResp := &SettingsResp{
		Title: testStringValue,
		Bcc:   []string{testStringValue},
	}

	mu.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SettingsRespJson, err := json.Marshal(*SettingsResp)
		if err != nil {
			return
		}

		w.Write(SettingsRespJson)
	})

	result, response, err := client.SiteSettingsApi.SettingsGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SettingsResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestSiteSettingsApiService_SettingsUploadLogoLinkGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SettingsUploadLogoLinkResp := &SettingsUploadLogoLinkResp{
		UploadKey: testStringValue,
	}

	mu.HandleFunc("/settings/upload-logo-link", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SettingsUploadLogoLinkRespJson, err := json.Marshal(*SettingsUploadLogoLinkResp)
		if err != nil {
			return
		}

		w.Write(SettingsUploadLogoLinkRespJson)
	})

	result, response, err := client.SiteSettingsApi.SettingsUploadLogoLinkGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SettingsUploadLogoLinkResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestSiteSettingsApiService_SettingsSetPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	SettingsResp := &SettingsResp{
		Title: testStringValue,
		Bcc:   []string{testStringValue},
	}

	mu.HandleFunc("/settings/set", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		SettingsRespJson, err := json.Marshal(*SettingsResp)
		if err != nil {
			return
		}

		w.Write(SettingsRespJson)
	})

	body := SettingsSetReq{}

	result, response, err := client.SiteSettingsApi.SettingsSetPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *SettingsResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
