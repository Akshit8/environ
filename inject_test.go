package environ

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupEnv() {
	os.Setenv("ENV_HOST", "localhost")
	os.Setenv("ENV_PORT", "8080")
	os.Setenv("ENV_FREQUENCY", "10.12")
}

func cleanEnv() {
	os.Unsetenv("ENV_HOST")
	os.Unsetenv("ENV_PORT")
	os.Unsetenv("ENV_FREQUENCY")
}

func TestOk(t *testing.T) {
	setupEnv()
	defer func() {
		cleanEnv()
		parsers = map[reflect.Type]reflect.Value{}
	}()

	type testConfig struct {
		Host  string `environ:"ENV_HOST"`
		Port  string `environ:"ENV_PORT"`
		Freq  string `environ:"ENV_FREQUENCY"`
		Token string `environ:"ENV_TOKEN,default_token"`
	}

	var config testConfig

	err := Inject(&config)
	assert.NoError(t, err)

	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "10.12", config.Freq)
	assert.Equal(t, "default_token", config.Token)
}

func TestOkWithParser(t *testing.T) {
	setupEnv()
	defer func() {
		cleanEnv()
		parsers = map[reflect.Type]reflect.Value{}
	}()

	type testConfig struct {
		Host  string  `environ:"ENV_HOST"`
		Port  int     `environ:"ENV_PORT"`
		Freq  float64 `environ:"ENV_FREQUENCY"`
		Token string  `environ:"ENV_TOKEN,default_token"`
	}

	intParser := func(s string) (int, error) {
		return strconv.Atoi(s)
	}

	floatParser := func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}

	UseParser(intParser)
	UseParser(floatParser)

	var config testConfig

	err := Inject(&config)
	assert.NoError(t, err)

	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, 10.12, config.Freq)
	assert.Equal(t, "default_token", config.Token)
}

func TestPassStructValue(t *testing.T) {
	defer func() {
		cleanEnv()
		parsers = map[reflect.Type]reflect.Value{}
	}()

	err := Inject(struct{}{})
	assert.NotNil(t, err)
}

func TestNoDefaultVar(t *testing.T) {
	defer func() {
		cleanEnv()
		parsers = map[reflect.Type]reflect.Value{}
	}()

	type testConfig struct {
		Host string `environ:"ENV_HOST"`
	}

	var config testConfig

	err := Inject(&config)
	assert.NotNil(t, err)
}

func TestNoParserFound(t *testing.T) {
	setupEnv()
	defer func() {
		cleanEnv()
		parsers = map[reflect.Type]reflect.Value{}
	}()

	type testConfig struct {
		Host string `environ:"ENV_HOST"`
		Port int    `environ:"ENV_PORT"`
	}

	var config testConfig

	err := Inject(&config)
	assert.NotNil(t, err)
}

func TestParserError(t *testing.T) {
	setupEnv()
	defer func() {
		cleanEnv()
		parsers = map[reflect.Type]reflect.Value{}
	}()

	type testConfig struct {
		Host string `environ:"ENV_HOST"`
		Port int    `environ:"ENV_PORT"`
	}

	intParser := func(s string) (int, error) {
		return 0, errors.New("error")
	}

	UseParser(intParser)

	var config testConfig

	err := Inject(&config)
	assert.NotNil(t, err)
}

func TestRandomTags(t *testing.T) {
	setupEnv()
	defer func() {
		cleanEnv()
		parsers = map[reflect.Type]reflect.Value{}
	}()

	type testConfig struct {
		Host   string `environ:"ENV_HOST"`
		Port   string `environ:"ENV_PORT"`
		Random string `environ:"ENV_RANDOM,default_random,random_tag"`
		Key    string
	}

	var config testConfig

	err := Inject(&config)
	assert.NoError(t, err)

	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "", config.Random)
	assert.Equal(t, "", config.Key)
}
