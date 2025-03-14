package envReader_test

import (
	"testing"
	"os"

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

func TestGetEnvMap(t *testing.T){
	envMap := envReader.GetEnvMap()
	t.Logf("\n %s \n", envMap)
}

func TestSetEnvVars(t *testing.T){
	envReader.SetEnvVars()

	if os.Getenv("Port") != ":8080" {
		t.Error("Unexpected return value at TestSetEnvVars")
	}
}