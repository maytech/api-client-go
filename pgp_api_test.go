package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestPGPApiService_PgpKeyDeleteIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/pgp-key/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ContactRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(ContactRespJson)
	})

	result, response, err := client.PGPApi.PgpKeyDeleteIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPGPApiService_PgpKeyRecipientsPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyRecipientsRespItems := []PgpKeyRecipientsRespItems{{
		Id:       testStringValue,
		Name:     testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
		Public:   testStringValue,
	}}

	mu.HandleFunc("/pgp-key/recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRecipientsRespItemsJson, err := json.Marshal(PgpKeyRecipientsRespItems)
		if err != nil {
			return
		}

		w.Write(PgpKeyRecipientsRespItemsJson)
	})

	body := IdsReq{
		Ids: []string{testStringValue},
	}

	result, response, err := client.PGPApi.PgpKeyRecipientsPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, PgpKeyRecipientsRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := IdsReq{}

	_, _, errIncor := client.PGPApi.PgpKeyRecipientsPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestPGPApiService_KeyRequestRespondIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	JobResp := &JobResp{
		JobId: testStringValue,
	}

	mu.HandleFunc("/key-request/respond/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		JobRespJson, err := json.Marshal(*JobResp)
		if err != nil {
			return
		}

		w.Write(JobRespJson)
	})

	body := KeyRequestRespondReq{
		Name:    testStringValue,
		Private: testStringValue,
		Public:  testStringValue,
	}

	result, response, err := client.PGPApi.KeyRequestRespondIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *JobResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := KeyRequestRespondReq{}

	_, _, errIncor := client.PGPApi.KeyRequestRespondIdPost(context.Background(), testStringValue, bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestPGPApiService_PgpKeyEditIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyResp := &PgpKeyResp{
		Id:       testStringValue,
		Name:     testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/pgp-key/edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRespJson, err := json.Marshal(*PgpKeyResp)
		if err != nil {
			return
		}

		w.Write(PgpKeyRespJson)
	})

	result, response, err := client.PGPApi.PgpKeyEditIdPost(context.Background(), testStringValue, nil)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PgpKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPGPApiService_PgpKeyMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyResp := &PgpKeyResp{
		Id:       testStringValue,
		Name:     testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/pgp-key/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRespJson, err := json.Marshal(*PgpKeyResp)
		if err != nil {
			return
		}

		w.Write(PgpKeyRespJson)
	})

	result, response, err := client.PGPApi.PgpKeyMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PgpKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPGPApiService_ContactPgpKeyIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyResp := &PgpKeyResp{
		Id:       testStringValue,
		Name:     testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
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

	result, response, err := client.PGPApi.ContactPgpKeyIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PgpKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPGPApiService_UserPgpKeyIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyResp := &PgpKeyResp{
		Id:       testStringValue,
		Name:     testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/user/pgp-key/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRespJson, err := json.Marshal(*PgpKeyResp)
		if err != nil {
			return
		}

		w.Write(PgpKeyRespJson)
	})

	result, response, err := client.PGPApi.UserPgpKeyIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PgpKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPGPApiService_PgpKeyRequestIdsGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	JobResp := &JobResp{
		JobId: testStringValue,
	}

	mu.HandleFunc("/pgp-key/request/{ids[]}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRespJson, err := json.Marshal(*JobResp)
		if err != nil {
			return
		}

		w.Write(PgpKeyRespJson)
	})

	result, response, err := client.PGPApi.PgpKeyRequestIdsGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *JobResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPGPApiService_PgpKeyCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyResp := &PgpKeyResp{
		Id:       testStringValue,
		Name:     testStringValue,
		Created:  testDigitalValue,
		Modified: testDigitalValue,
	}

	mu.HandleFunc("/pgp-key/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRespJson, err := json.Marshal(*PgpKeyResp)
		if err != nil {
			return
		}

		w.Write(PgpKeyRespJson)
	})

	body := PgpCreateReq{
		Name:    testStringValue,
		Public:  testStringValue,
		Private: testStringValue,
	}

	result, response, err := client.PGPApi.PgpKeyCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PgpKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestPGPApiService_KeyRequestMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	KeyRequestMetadataResp := &KeyRequestMetadataResp{
		Id:          testStringValue,
		SenderEmail: testStringValue,
		SenderName:  testStringValue,
		UserName:    testStringValue,
		UserEmail:   testStringValue,
	}

	mu.HandleFunc("/key-request/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PgpKeyRespJson, err := json.Marshal(*KeyRequestMetadataResp)
		if err != nil {
			return
		}

		w.Write(PgpKeyRespJson)
	})

	result, response, err := client.PGPApi.KeyRequestMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *KeyRequestMetadataResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
