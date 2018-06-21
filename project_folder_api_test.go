package quatrix

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func TestProjectFolderApiService_ProjectFolderCreatePost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PfCreateResp := &PfCreateResp{
		Id:   testStringValue,
		Name: testStringValue,
	}

	mu.HandleFunc("/project-folder/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PfCreateRespJson, err := json.Marshal(*PfCreateResp)
		if err != nil {
			return
		}

		w.Write(PfCreateRespJson)
	})

	body := PfcreateReq{
		Created:          testDigitalValue,
		Name:             testStringValue,
		FileId:           testStringValue,
		UsersPermissions: []UserPermissionReq{{}},
		Notify:           testBoolValue,
	}

	result, response, err := client.ProjectFolderApi.ProjectFolderCreatePost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PfCreateResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := PfcreateReq{}

	_, _, errIncor := client.ProjectFolderApi.ProjectFolderCreatePost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestProjectFolderApiService_ProjectFolderUsersIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserPermissionResp := []UserPermissionResp{{
		UserId:     testStringValue,
		Operations: testDigitalValue,
		Owns:       testBoolValue,
	}}

	mu.HandleFunc("/project-folder/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserPermissionRespJson, err := json.Marshal(UserPermissionResp)
		if err != nil {
			return
		}

		w.Write(UserPermissionRespJson)
	})

	result, response, err := client.ProjectFolderApi.ProjectFolderUsersIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, UserPermissionResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestProjectFolderApiService_ProjectFolderMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PfMetadataResp := &PfMetadataResp{
		FolderName: testStringValue,
		Created:    testDigitalValue,
	}

	mu.HandleFunc("/project-folder/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PfMetadataRespJson, err := json.Marshal(*PfMetadataResp)
		if err != nil {
			return
		}

		w.Write(PfMetadataRespJson)
	})

	result, response, err := client.ProjectFolderApi.ProjectFolderMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *PfMetadataResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestProjectFolderApiService_ProjectFolderAddUsersIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserPermissionResp := []UserPermissionResp{{
		UserId:     testStringValue,
		Operations: testDigitalValue,
		Owns:       testBoolValue,
	}}

	mu.HandleFunc("/project-folder/add-users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserPermissionRespJson, err := json.Marshal(UserPermissionResp)
		if err != nil {
			return
		}

		w.Write(UserPermissionRespJson)
	})

	body := PfaddUsersReq{
		UsersPermissions: []UserPermissionReq{{
			UserId:     testStringValue,
			Notify:     testBoolValue,
			Operations: testDigitalValue,
		}},
	}

	result, response, err := client.ProjectFolderApi.ProjectFolderAddUsersIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, UserPermissionResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := PfaddUsersReq{}

	_, _, errIncor := client.ProjectFolderApi.ProjectFolderAddUsersIdPost(context.Background(), testStringValue, bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestProjectFolderApiService_ProjectFolderDeleteUsersIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	PfDeleteUsersRespItems := []PfDeleteUsersRespItems{{
		UserId: testStringValue,
	}}

	mu.HandleFunc("/project-folder/delete-users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		PfDeleteUsersRespItemsJson, err := json.Marshal(PfDeleteUsersRespItems)
		if err != nil {
			return
		}

		w.Write(PfDeleteUsersRespItemsJson)
	})

	body := PfdeleteUsersReq{
		Users: []ProjectfolderdeleteusersUsers{{UserId:testStringValue}},
	}

	result, response, err := client.ProjectFolderApi.ProjectFolderDeleteUsersPost(context.Background(), body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, PfDeleteUsersRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := PfdeleteUsersReq{}

	_, _, errIncor := client.ProjectFolderApi.ProjectFolderDeleteUsersPost(context.Background(), bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestProjectFolderApiService_ProjectFolderEditUsersIdPost(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	UserPermissionResp := []UserPermissionResp{{
		UserId:     testStringValue,
		Operations: testDigitalValue,
		Owns:       testBoolValue,
	}}

	mu.HandleFunc("/project-folder/edit-users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		UserPermissionRespJson, err := json.Marshal(UserPermissionResp)
		if err != nil {
			return
		}

		w.Write(UserPermissionRespJson)
	})

	body := PfeditUsersReq{
		UsersPermissions: []UserPermissionReq{{
			UserId: testStringValue,
		}},
	}

	result, response, err := client.ProjectFolderApi.ProjectFolderEditUsersIdPost(context.Background(), testStringValue, body)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, UserPermissionResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)

	// Incorrect call
	bodyIncorrect := PfeditUsersReq{}

	_, _, errIncor := client.ProjectFolderApi.ProjectFolderEditUsersIdPost(context.Background(), testStringValue, bodyIncorrect)
	if errIncor == nil {
		t.Error(validateIncorrect)
	}
}

func TestProjectFolderApiService_ProjectFolderDeleteIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	IdResp := &IdResp{
		Id: testStringValue,
	}

	mu.HandleFunc("/project-folder/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		IdRespJson, err := json.Marshal(*IdResp)
		if err != nil {
			return
		}

		w.Write(IdRespJson)
	})

	result, response, err := client.ProjectFolderApi.ProjectFolderDeleteIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *IdResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
