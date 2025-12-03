// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	group "github.com/crossplane-contrib/provider-aap/internal/controller/namespaced/core/group"
	host "github.com/crossplane-contrib/provider-aap/internal/controller/namespaced/core/host"
	inventory "github.com/crossplane-contrib/provider-aap/internal/controller/namespaced/core/inventory"
	job "github.com/crossplane-contrib/provider-aap/internal/controller/namespaced/core/job"
	workflowjob "github.com/crossplane-contrib/provider-aap/internal/controller/namespaced/core/workflowjob"
	providerconfig "github.com/crossplane-contrib/provider-aap/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		host.Setup,
		inventory.Setup,
		job.Setup,
		workflowjob.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.SetupGated,
		host.SetupGated,
		inventory.SetupGated,
		job.SetupGated,
		workflowjob.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
