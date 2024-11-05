package pkg

import (
	"context"
	"net/http"
	"time"

	apiClient "github.com/pierskarsenbarg/native-provider-demo/provider/internal"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	ApiToken string `pulumi:"apiToken,optional" provider:"secret"`
	Client   apiClient.Client
	BaseUrl  string
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.ApiToken, "Your API token")
	a.SetDefault(&c.ApiToken, "", "PROVIDER_APITOKEN")
	a.Describe(&c.BaseUrl, "Url of the API")
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx context.Context) error {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	client, err := apiClient.NewClient(&httpClient, c.ApiToken, "")
	if err != nil {
		return err
	}

	c.Client = *client

	return nil
}
