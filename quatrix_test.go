package quatrix

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

const (
	testStringValue   = "testStringValue"
	testDigitalValue  = 12
	testBoolValue     = true
	validateIncorrect = "incorrect validation of required parameters"
)

var (
	mu     *mux.Router
	server *httptest.Server
	client *APIClient
)

func initClient() func() {
	mu = mux.NewRouter()
	server = httptest.NewServer(mu)

	client = NewAPIClient(&Configuration{
		BasePath:      server.URL,
		DefaultHeader: make(map[string]string),
	})

	return func() {
		server.Close()
	}
}

func deepEqual(t *testing.T, x, y interface{}) {
	t.Helper()

	if !reflect.DeepEqual(x, y) {
		t.Errorf("response not match, obtained %v but expected %v", y, x)
	}
}

func checkStatusCode(t *testing.T, x, y int) {
	t.Helper()

	if x != y {
		t.Errorf("status code not match, obtained %d but expected %d", x, y)
	}
}
