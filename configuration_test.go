package quatrix

import (
	"testing"
)

var testConfig = &Configuration{
	BasePath:      "https://api.quatrix.it/api/1.0",
	DefaultHeader: make(map[string]string),
	UserAgent:     "Swagger-Codegen/1.0.0/go",
}

func TestNewConfiguration(t *testing.T) {

	deepEqual(t, NewConfiguration(), testConfig)
}

func TestConfiguration_AddDefaultHeader(t *testing.T) {

	newCfg := &Configuration{
		BasePath:      "https://api.quatrix.it/api/1.0",
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		DefaultHeader: map[string]string{"testHeader": "testValue"},
	}

	testConfig.AddDefaultHeader("testHeader", "testValue")

	deepEqual(t, newCfg, testConfig)
}
