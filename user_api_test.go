package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestUserApiService_UserSetMfaPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileRemoveMfaResp := &ProfileRemoveMfaResp{
		Status: testStringValue,
	}

	mu.HandleFunc("/user/set-mfa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileRemoveMfaRespJson, err := json.Marshal(*ProfileRemoveMfaResp)
		if err != nil {
			return
		}

		w.Write(ProfileRemoveMfaRespJson)
	})

	body := UserSetMfaReq{
		Ids:      []string{testStringValue},
		AuthType: testStringValue,
		Code:     testStringValue,
	}

	result, response, err := client.UserApi.UserSetMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileRemoveMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UserSetMfaReq{}

	_, _, errIncor := client.UserApi.UserSetMfaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestUserApiService_UserDeletePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	JobResp := &JobResp{
		JobId: testStringValue,
	}

	mu.HandleFunc("/user/delete", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		JobRespJson, err := json.Marshal(*JobResp)
		if err != nil {
			return
		}

		w.Write(JobRespJson)
	})

	body := UserDeleteReq{
		Ids: []string{testStringValue},
	}

	result, response, err := client.UserApi.UserDeletePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *JobResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UserDeleteReq{}

	_, _, errIncor := client.UserApi.UserDeletePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestUserApiService_UserGroupGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	Group := []Group{{
		Id: testStringValue,
	}}

	mu.HandleFunc("/user/group", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		GroupJson, err := json.Marshal(Group)
		if err != nil {
			return
		}

		w.Write(GroupJson)
	})

	result, response, err := client.UserApi.UserGroupGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, Group)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestUserApiService_UserGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserResp := []UserResp{{
		Id:         testStringValue,
		ParentId:   testStringValue,
		HomeId:     testStringValue,
		HomeName:   testStringValue,
		SuperAdmin: testStringValue,
		Name:       testStringValue,
	}}

	mu.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserRespJson, err := json.Marshal(UserResp)
		if err != nil {
			return
		}

		w.Write(UserRespJson)
	})

	result, response, err := client.UserApi.UserGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, UserResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestUserApiService_UserCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserResp := &UserResp{
		Id:       testStringValue,
		ParentId: testStringValue,
		HomeName: testStringValue,
		HomeId:   testStringValue,
	}

	mu.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserRespJson, err := json.Marshal(*UserResp)
		if err != nil {
			return
		}

		w.Write(UserRespJson)
	})

	body := UserCreateReq{
		Name:  testStringValue,
		Email: testStringValue,
	}

	result, response, err := client.UserApi.UserCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *UserResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UserCreateReq{}

	_, _, errIncor := client.UserApi.UserCreatePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestUserApiService_UserRemoveMfaPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ProfileRemoveMfaResp := &ProfileRemoveMfaResp{
		Status: testStringValue,
	}

	mu.HandleFunc("/user/remove-mfa", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ProfileRemoveMfaRespJson, err := json.Marshal(*ProfileRemoveMfaResp)
		if err != nil {
			return
		}

		w.Write(ProfileRemoveMfaRespJson)
	})

	body := UserRemoveMfaReq{
		Ids:      []string{testStringValue},
		AuthType: testStringValue,
	}

	result, response, err := client.UserApi.UserRemoveMfaPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *ProfileRemoveMfaResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UserRemoveMfaReq{}

	_, _, errIncor := client.UserApi.UserRemoveMfaPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestUserApiService_UserMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserResp := &UserResp{
		Id:       testStringValue,
		ParentId: testStringValue,
		HomeName: testStringValue,
		HomeId:   testStringValue,
	}

	mu.HandleFunc("/user/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserRespJson, err := json.Marshal(*UserResp)
		if err != nil {
			return
		}

		w.Write(UserRespJson)
	})

	result, response, err := client.UserApi.UserMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *UserResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestUserApiService_UserPgpKeyIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PgpKeyResp := &PgpKeyResp{
		Id:      testStringValue,
		Name:    testStringValue,
		Created: testDigitalValue,
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

	result, response, err := client.UserApi.UserPgpKeyIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PgpKeyResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestUserApiService_UserEditPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserResp := []UserResp{{
		Id:       testStringValue,
		ParentId: testStringValue,
		HomeName: testStringValue,
		HomeId:   testStringValue,
	}}

	mu.HandleFunc("/user/edit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserRespJson, err := json.Marshal(UserResp)
		if err != nil {
			return
		}

		w.Write(UserRespJson)
	})

	body := UserEditReq{
		Ids: []string{testStringValue},
	}

	result, response, err := client.UserApi.UserEditPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, UserResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := UserEditReq{}

	_, _, errIncor := client.UserApi.UserEditPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestUserApiService_UserSignupPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	mu.HandleFunc("/user/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	body := UserSignupReq{
		Email:    testStringValue,
		Password: testStringValue,
	}

	response, err := client.UserApi.UserSignupPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
