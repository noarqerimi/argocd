package pulumi-argocd

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/noarqerimi/pulumi-argocd//tree/main/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/oboukili/terraform-provider-argocd/tree/master/argocd"
)
const (
	mainPkg = "pulumi-argocd"
	mainMod = "index"
)
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(pulumi-argocd.Provider())
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "pulumi-argocd",
		DisplayName: "",
		Publisher: "Pulumi",
		LogoURL: "",
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing pulumi-argocd cloud resources.",
		Keywords:   []string{"pulumi", "pulumi-argocd", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/noarqerimi/pulumi-argocd",
		GitHubOrg: "",
		Config:    map[string]*tfbridge.SchemaInfo{
		},
		PreConfigureCallback: preConfigureCallback,
		Resources:            map[string]*tfbridge.ResourceInfo{
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// For more information, please reference: https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("pulumi-argocd_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliasing()
	prov.SetAutonaming(255, "-")

	return prov
}
