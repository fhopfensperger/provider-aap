package config

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/upjet/v2/pkg/config"
)

// When constructing the state, we want to set the default values using omittedFieldWithDefaultValue,
// to avoid when observing an already existing resources, that an update is triggered
type omittedFieldWithDefaultValue struct {
	FieldName    string
	DefaultValue any
}

// TerraformPluginFrameworkExternalNameConfigs contains all external name configurations for this
// provider.
var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"aap_group":     FrameworkResourceWithComputedIdentifierWithOmittedField("url", "/api/v2/groups/abcdef/"),
	"aap_host":      FrameworkResourceWithComputedIdentifierWithOmittedField("url", "/api/v2/hosts/abcdef/"),
	"aap_inventory": FrameworkResourceWithComputedIdentifierWithOmittedField("url", "/api/v2/inventories/abcdef/"),
	"aap_job": FrameworkResourceWithComputedIdentifierWithOmittedField("url", "/api/v2/jobs/abcdef/",
		omittedFieldWithDefaultValue{"wait_for_completion", false},
		omittedFieldWithDefaultValue{"wait_for_completion_timeout_seconds", 120},
	),
	"aap_workflow_job": FrameworkResourceWithComputedIdentifierWithOmittedField("url", "/api/v2/jobs/abcdef/",
		omittedFieldWithDefaultValue{"wait_for_completion", false},
		omittedFieldWithDefaultValue{"wait_for_completion_timeout_seconds", 120},
	)}

func FrameworkResourceWithComputedIdentifierWithOmittedField(identifier, placeholder string, omittedFieldWithDefaultValues ...omittedFieldWithDefaultValue) config.ExternalName {
	en := config.NewExternalNameFrom(config.IdentifierFromProvider,
		config.WithSetIdentifierArgumentsFn(func(fn config.SetIdentifierArgumentsFn, base map[string]any, externalName string) {
			if id, ok := base[identifier]; !ok || id == placeholder {
				if externalName == "" {
					base[identifier] = placeholder
				} else {
					base[identifier] = externalName
				}
				for _, omittedField := range omittedFieldWithDefaultValues {
					if omittedField.DefaultValue != nil {
						base[omittedField.FieldName] = omittedField.DefaultValue
					}
				}
			}
		}),
		config.WithGetExternalNameFn(func(fn config.GetExternalNameFn, tfState map[string]any) (string, error) {
			if id, ok := tfState[identifier]; ok {
				idStr := fmt.Sprintf("%v", id)
				if len(idStr) > 0 {
					return idStr, nil
				}
			}
			return "", errors.Errorf("cannot find attribute %q in tfstate", identifier)
		}),
	)
	// en.TFPluginFrameworkOptions.ComputedIdentifierAttributes = append([]string{identifier}, omittedFields...)
	en.TFPluginFrameworkOptions.ComputedIdentifierAttributes = []string{identifier}

	omittedFields := []string{}
	for _, f := range omittedFieldWithDefaultValues {
		omittedFields = append(omittedFields, f.FieldName)
	}
	en.OmittedFields = omittedFields

	return en
}

func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// If an external name is configured for multiple architectures,
		// Terraform Plugin Framework takes precedence over Terraform
		// Plugin SDKv2, which takes precedence over CLI architecture.
		e, configured := TerraformPluginFrameworkExternalNameConfigs[r.Name]
		if !configured {
			return
		}
		r.Version = "v1alpha1"
		r.ExternalName = e
	}
}

func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}

func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
