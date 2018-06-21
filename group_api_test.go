package quatrix

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestGroupApiService_GroupMetadataIdGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	GroupResp := &GroupResp{
		Id:         testStringValue,
		ParentId:   testStringValue,
		ParentName: testStringValue,
		Name:       testStringValue,
		Type_:      testStringValue,
	}

	mu.HandleFunc("/group/metadata/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		GroupRespJson, err := json.Marshal(*GroupResp)
		if err != nil {
			return
		}

		w.Write(GroupRespJson)
	})

	result, response, err := client.GroupApi.GroupMetadataIdGet(context.Background(), testStringValue)
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, *GroupResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestGroupApiService_UserGroupGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	Group := []Group{{
		Id:   testStringValue,
		Name: testStringValue,
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

	result, response, err := client.GroupApi.UserGroupGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, Group) {
		t.Errorf("response not match, obtained %v but expected %v", Group, result)
	}

	deepEqual(t, result, Group)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestGroupApiService_ContactGroupGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	ContactGroupRespItems := []ContactGroupRespItems{{
		Id:       testStringValue,
		Name:     testStringValue,
		Global:   testBoolValue,
		Readonly: testBoolValue,
	}}

	mu.HandleFunc("/contact/group", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		GroupJson, err := json.Marshal(ContactGroupRespItems)
		if err != nil {
			return
		}

		w.Write(GroupJson)
	})

	result, response, err := client.GroupApi.ContactGroupGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, ContactGroupRespItems)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}

func TestGroupApiService_GroupGet(t *testing.T) {
	tearDown := initClient()
	defer tearDown()

	GroupResp := []GroupResp{{
		Id:         testStringValue,
		ParentId:   testStringValue,
		ParentName: testStringValue,
		Name:       testStringValue,
		Type_:      testStringValue,
	}}

	mu.HandleFunc("/group", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		GroupRespJson, err := json.Marshal(GroupResp)
		if err != nil {
			return
		}

		w.Write(GroupRespJson)
	})

	result, response, err := client.GroupApi.GroupGet(context.Background())
	if err != nil {
		t.Error(err)
	}

	deepEqual(t, result, GroupResp)
	checkStatusCode(t, response.StatusCode, http.StatusOK)
}
