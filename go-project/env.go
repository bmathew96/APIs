package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"strings"

	"github.com/kelseyhightower/envconfig"
)

/*
	These environment variables must be set in order for getEnvConfig() to work
	export MYAPI_DBUSERNAME=
	export MYAPI_DBPASSWORD=
	export MYAPI_DBHOSTNAME=
	export MYAPI_DBPORT=3306
	export MYAPI_DEBUG=true
	export MYAPI_APIKEYPUBLIC=""
	export MYAPI_APIKEYPRIVATE=""
*/

// EnvConfig contains the Enviornment Variables we are using
type EnvConfig struct {
	DBUsername    string
	DBPassword    string
	DBHostname    string
	APIKeyPublic  string `envconfig:"APIKeyPublic"`
	APIKeyPrivate string `envconfig:"APIKeyPrivate"`
	DBPort        int
	Debug         bool
}

func getEnvConfig() (config EnvConfig, err error) {

	// parse out environment variables
	err = envconfig.Process("myapi", &config)

	return
}

func getKeyHash() (hash string, timeStamp string) {
	timeStamp = strings.Replace(time.Now().String(), " ", "", -1)

	//creates md5 hash
	h := md5.New()
	io.WriteString(h, timeStamp)
	io.WriteString(h, config.APIKeyPrivate)
	io.WriteString(h, config.APIKeyPublic)
	hash = fmt.Sprintf("%x", h.Sum(nil))

	return
}
