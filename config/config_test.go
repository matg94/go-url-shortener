package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	testConfigString := "base_url: \"http://test\"\n" +
		"redis:\n" +
		"  MaxIdle: 123\n" +
		"  Username: \"test_user\""

	err := ioutil.WriteFile("./testFile.yaml", []byte(testConfigString), 0644)

	if err != nil {
		t.Logf("Failed to write test config: %s", err)
		t.Fail()
	}

	data, err := ReadConfigFile("./testFile.yaml")

	if err != nil {
		t.Logf("Failed to read test config: %s", err)
		t.Fail()
	}
	if string(data) != testConfigString {
		t.Log("Expected data read from file to match data defined in code")
		t.Fail()
	}
	os.Remove("./testFile.yaml")
}

func TestParseConfig(t *testing.T) {
	testParsedConfig := &AppConfig{
		BaseUrl: "http://test",
		Redis: RedisConfig{
			MaxIdle: 123,
			User:    "test_user",
		},
	}
	testConfigString := "base_url: \"http://test\"\n" +
		"redis:\n" +
		"  MaxIdle: 123\n" +
		"  Username: \"test_user\""

	parsedConfig, err := ParseYamlConfig([]byte(testConfigString))

	if err != nil {
		t.Logf("Failed to parse test config: %s", err)
		t.Fail()
	}

	if *testParsedConfig != *parsedConfig {
		t.Logf("Expected test config and parsed config to match")
		t.Fail()
	}
}
