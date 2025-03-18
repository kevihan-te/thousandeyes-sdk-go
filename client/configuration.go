package client

import (
	"context"
	"fmt"
	"golang.org/x/mod/modfile"
	"log"
	"net/http"
	"os"
	"strings"
)

var sdkVersion string

func SDKVersion() string {
	return sdkVersion
}

func init() {
	sdkVersion = getCurrentModuleVersion()
}

func getCurrentModuleVersion() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(modFile.Module.Mod.Path, "/")
	return parts[len(parts)-1]
}

// Configuration stores the configuration of the API client
type Configuration struct {
	UserAgent  string `json:"userAgent,omitempty"`
	Debug      bool   `json:"debug,omitempty"`
	ServerURL  string
	AuthToken  string
	HTTPClient *http.Client
	Context    context.Context
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		Debug:     false,
		ServerURL: "https://api.thousandeyes.com/v7",
	}
	return cfg
}

func (c *Configuration) GenerateUserAgent() string {
	var sdkUserAgent = fmt.Sprintf("ThousandEyesSDK-Go/%s", SDKVersion())
	if c.UserAgent == "" {
		return sdkUserAgent
	}
	return sdkUserAgent + " " + c.UserAgent
}

func (c *Configuration) WithAuthToken(authToken string) *Configuration {
	c.AuthToken = authToken
	return c
}

func (c *Configuration) WithServerUrl(serverUrl string) *Configuration {
	c.ServerURL = serverUrl
	return c
}
