package core

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aap_group", func(r *config.Resource) {
		r.ShortGroup = "core"
	})
	p.AddResourceConfigurator("aap_host", func(r *config.Resource) {
		r.ShortGroup = "core"
	})
	p.AddResourceConfigurator("aap_inventory", func(r *config.Resource) {
		r.ShortGroup = "core"
	})
	p.AddResourceConfigurator("aap_job", func(r *config.Resource) {
		r.ShortGroup = "core"
	})
	p.AddResourceConfigurator("aap_workflow_job", func(r *config.Resource) {
		r.ShortGroup = "core"
		r.Kind = "WorkflowJob"
	})
}
