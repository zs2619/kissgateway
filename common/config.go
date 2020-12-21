package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ProxyServer struct {
	Name string   `json:"name"`
	Addr []string `json:"addrList"`
}

type WebConfigType struct {
	LogLevel            string        `json:"logLevel"`
	ServerPort          int           `json:"serverPort"`
	ServerType          string        `json:"serverType"`
	ProxyServerConfList []ProxyServer `json:"proxyServerList"`
}

func getConfigRaw(config string) ([]byte, error) {
	configPath := os.Getenv("GATEWAY_CONFIGPATH")
	filePath := filepath.Join(configPath, config)
	raw, err := ioutil.ReadFile(filePath)
	return raw, err
}

var WebConfig WebConfigType

func LoadWebConfig(config string) (*WebConfigType, error) {
	rawFile, err := getConfigRaw(config)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawFile, &WebConfig)
	if err != nil {
		return nil, err
	}
	return &WebConfig, nil
}
