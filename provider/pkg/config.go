package pkg

import (
	"context"
	"net/url"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	ApiToken string `pulumi:"apiToken,optional" provider:"secret"`
	Client   serviceClient.Client
	BaseUrl  string
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.ApiToken, "Your API token")
	a.SetDefault(&c.ApiToken, "", "PROVIDER_APITOKEN")
	a.Describe(&c.BaseUrl, "Url of the API")
}

// var _ = (infer.Annotated)((*Config)(nil))

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx context.Context) error {
	// httpClient := http.Client{
	// 	Timeout: 60 * time.Second,
	// }

	baseUrl, err := url.Parse(c.BaseUrl)
	if err != nil {
		return err
	}

	c.Client, err := serviceClient.NewClientWithResponses(baseUrl.String())

	c.Client = client
}
