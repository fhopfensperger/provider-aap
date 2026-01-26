package core

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/terraform/errors"
	"github.com/hashicorp/terraform-plugin-framework/path"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const coreGroup = "core"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aap_group", func(r *config.Resource) {
		r.ShortGroup = coreGroup
	})
	p.AddResourceConfigurator("aap_host", func(r *config.Resource) {
		r.ShortGroup = coreGroup
	})
	p.AddResourceConfigurator("aap_inventory", func(r *config.Resource) {
		r.ShortGroup = coreGroup
	})
	p.AddResourceConfigurator("aap_job", func(r *config.Resource) {
		r.ShortGroup = coreGroup
		r.TerraformPluginFrameworkIsStateEmptyFn = func(ctx context.Context, tfStateValue tftypes.Value, resourceSchema rschema.Schema) (bool, error) {
			sdkState := tfsdk.State{
				Raw:    tfStateValue,
				Schema: resourceSchema,
			}
			var url string
			if diags := sdkState.GetAttribute(ctx, path.Root("url"), &url); diags.HasError() {
				return false, errors.FrameworkDiagnosticsError("reading url attribute", diags)
			}
			return url == "", nil
		}
	})
	p.AddResourceConfigurator("aap_workflow_job", func(r *config.Resource) {
		r.ShortGroup = coreGroup
		r.Kind = "WorkflowJob"
	})
}
