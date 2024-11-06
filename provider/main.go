package main

import (
	"fmt"
	"os"

	"github.com/pierskarsenbarg/native-provider-demo/provider/pkg"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version = "0.0.1"

const Name string = "nativeProvider"

func main() {
	err := p.RunProvider(Name, Version, provider())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func provider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: Name,
			Description: "Demo native provider",
			LanguageMap: map[string]any{
				"go": gogen.GoPackageInfo{
					ModulePath:     "github.com/pierskarsenbarg/native-provider-demo/provider/sdk",
					ImportBasePath: "github.com/pierskarsenbarg/native-provider-demo/provider/sdk/go/nativeProvider",
				},
				"nodejs": nodejsgen.NodePackageInfo{
					PackageName: "@pierskarsenbarg/nativeProvider",
					Dependencies: map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
					DevDependencies: map[string]string{
						"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
						"@types/mime": "^2.0.0",
					},
				},
				"csharp": dotnetgen.CSharpPackageInfo{
					RootNamespace: "MyNamespace",
					PackageReferences: map[string]string{
						"Pulumi": "3.*",
					},
				},
				"python": pythongen.PackageInfo{
					Requires: map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
					PackageName: "pierskarsenbarg_pulumi_nativeProvider",
				},
			},
			PluginDownloadURL: "github://api.github.com/pierskarsenbarg/native-provider-demo",
			Publisher:         "Piers Karsenbarg",
		},
		Resources: []infer.InferredResource{
			infer.Resource[*pkg.Organisation, pkg.OrganisationArgs, pkg.OrganisationState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"pkg": "index", // required because the folder with everything in is "pkg"
		},
		Functions: []infer.InferredFunction{
			infer.Function[*pkg.GetOrganisation, pkg.GetOrganisationArgs, pkg.OrganisationState](),
		},
		Config: infer.Config[*pkg.Config](),
	})
}
