package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	coreCluster "github.com/crossplane-contrib/provider-aap/config/cluster/core"
	coreNamespaced "github.com/crossplane-contrib/provider-aap/config/namespaced/core"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
)

const (
	resourcePrefix = "aap"
	modulePath     = "github.com/crossplane-contrib/provider-aap"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider(fwProvider fwprovider.Provider) *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("aap.crossplane.io"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithIncludeList(CLIReconciledResourceList()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		ujconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		ujconfig.WithDefaultResourceOptions(ResourceConfigurator()),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		coreCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced(fwProvider fwprovider.Provider) *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("aap.m.crossplane.io"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithIncludeList(CLIReconciledResourceList()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		ujconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		ujconfig.WithDefaultResourceOptions(ResourceConfigurator()),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}),
	)
	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		coreNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
