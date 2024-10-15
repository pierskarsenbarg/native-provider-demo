package pkg

import (
	serviceClient "github.com/pierskarsenbarg/native-provider-demo/client"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	ApiToken string `pulumi:"apiToken,optional" provider:"secret"`
	Client   serviceClient.Client
}

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.ApiToken, "Your Turso API token")
	a.SetDefault(&c.ApiToken, "", "PROVIDER_APITOKEN")
}

var _ = (infer.Annotated)((*Config)(nil))
