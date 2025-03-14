package envReader_test

import (
	"testing"

	envReader "github.com/udan-jayanith/envReader"
)

func TestLoadEnv(t *testing.T) {
	envReader.LoadEnv()
}

func TestGet(t *testing.T) {
	port := envReader.Get("Port")
	t.Logf("\n %s \n", port)
	if port != ":8080"{
		t.Error("Unexpected value at TestGet ", port)
	}
}
