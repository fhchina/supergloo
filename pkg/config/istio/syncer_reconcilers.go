package istio

import (
	"context"

	"github.com/solo-io/supergloo/pkg/config/utils"

	v1 "github.com/solo-io/supergloo/pkg/api/v1"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	policyv1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"
	networkingv1alpha3 "github.com/solo-io/supergloo/pkg/api/external/istio/networking/v1alpha3"
	rbacv1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/rbac/v1alpha1"
	"github.com/solo-io/supergloo/pkg/translator/istio"
)

type Reconcilers interface {
	ReconcileAll(ctx context.Context, config *istio.MeshConfig) error
}

type istioReconcilers struct {
	ownerLabels map[string]string

	rbacConfigReconciler         rbacv1alpha1.RbacConfigReconciler
	serviceRoleReconciler        rbacv1alpha1.ServiceRoleReconciler
	serviceRoleBindingReconciler rbacv1alpha1.ServiceRoleBindingReconciler
	meshPolicyReconciler         policyv1alpha1.MeshPolicyReconciler
	destinationRuleReconciler    networkingv1alpha3.DestinationRuleReconciler
	virtualServiceReconciler     networkingv1alpha3.VirtualServiceReconciler
	tlsSecretReconciler          v1.TlsSecretReconciler
}

func NewIstioReconcilers(ownerLabels map[string]string,
	rbacConfigReconciler rbacv1alpha1.RbacConfigReconciler,
	serviceRoleReconciler rbacv1alpha1.ServiceRoleReconciler,
	serviceRoleBindingReconciler rbacv1alpha1.ServiceRoleBindingReconciler,
	meshPolicyReconciler policyv1alpha1.MeshPolicyReconciler,
	destinationRuleReconciler networkingv1alpha3.DestinationRuleReconciler,
	virtualServiceReconciler networkingv1alpha3.VirtualServiceReconciler,
	tlsSecretReconciler v1.TlsSecretReconciler) Reconcilers {
	return &istioReconcilers{
		ownerLabels:                  ownerLabels,
		rbacConfigReconciler:         rbacConfigReconciler,
		serviceRoleReconciler:        serviceRoleReconciler,
		serviceRoleBindingReconciler: serviceRoleBindingReconciler,
		meshPolicyReconciler:         meshPolicyReconciler,
		destinationRuleReconciler:    destinationRuleReconciler,
		virtualServiceReconciler:     virtualServiceReconciler,
		tlsSecretReconciler:          tlsSecretReconciler,
	}
}

func (s *istioReconcilers) ReconcileAll(ctx context.Context, config *istio.MeshConfig) error {
	logger := contextutils.LoggerFrom(ctx)

	// this list should always either be empty or contain the global mesh policy
	var meshPoliciesToReconcile policyv1alpha1.MeshPolicyList
	if config.MeshPolicy != nil {
		logger.Infof("MeshPolicy: %v", config.MeshPolicy.Metadata.Name)
		utils.SetLabels(s.ownerLabels, config.MeshPolicy)
		meshPoliciesToReconcile = append(meshPoliciesToReconcile, config.MeshPolicy)
	}
	if err := s.meshPolicyReconciler.Reconcile(
		"",
		meshPoliciesToReconcile, // mesh policy is a singleton
		nil,
		clients.ListOpts{
			Ctx:      ctx,
			Selector: nil, // allows overwriting a user-created mesh policy
		},
	); err != nil {
		return errors.Wrapf(err, "reconciling default mesh policy")
	}

	// this list should always either be empty or contain the global rbac config
	var rbacConfigsToReconcile rbacv1alpha1.RbacConfigList
	if config.RbacConfig != nil {
		logger.Infof("RbacConfig: %v", config.RbacConfig.Metadata.Name)
		utils.SetLabels(s.ownerLabels, config.RbacConfig)
		rbacConfigsToReconcile = append(rbacConfigsToReconcile, config.RbacConfig)
	}
	if err := s.rbacConfigReconciler.Reconcile(
		"",
		rbacConfigsToReconcile, // rbac config is a singleton
		nil,
		clients.ListOpts{
			Ctx:      ctx,
			Selector: nil, // allows overwriting a user-created rbac config
		},
	); err != nil {
		return errors.Wrapf(err, "reconciling default rbac config")
	}

	// this list should always either be empty or contain the global cacerts root cert secret
	var tlsSecretsToReconcile v1.TlsSecretList
	if config.RootCert != nil {
		logger.Infof("RootCert: %v", config.RootCert.Metadata.Name)
		utils.SetLabels(s.ownerLabels, config.RootCert)
		tlsSecretsToReconcile = append(tlsSecretsToReconcile, config.RootCert)
	}
	if err := s.tlsSecretReconciler.Reconcile(
		"",
		tlsSecretsToReconcile, // root cert is a singleton
		nil,
		clients.ListOpts{
			Ctx:      ctx,
			Selector: s.ownerLabels,
		},
	); err != nil {
		return errors.Wrapf(err, "reconciling cacerts root cert")
	}

	logger.Infof("DestinationRules: %v", config.DestinationRules.Names())
	utils.SetLabels(s.ownerLabels, config.DestinationRules.AsResources()...)
	if err := s.destinationRuleReconciler.Reconcile(
		"",
		config.DestinationRules,
		nil,
		clients.ListOpts{
			Ctx:      ctx,
			Selector: s.ownerLabels,
		},
	); err != nil {
		return errors.Wrapf(err, "reconciling destination rules")
	}

	logger.Infof("VirtualServices: %v", config.VirtualServices.Names())
	utils.SetLabels(s.ownerLabels, config.VirtualServices.AsResources()...)
	if err := s.virtualServiceReconciler.Reconcile(
		"",
		config.VirtualServices,
		nil,
		clients.ListOpts{
			Ctx:      ctx,
			Selector: s.ownerLabels,
		},
	); err != nil {
		return errors.Wrapf(err, "reconciling virtual services")
	}

	logger.Infof("ServiceRoles: %v", config.ServiceRoles.Names())
	utils.SetLabels(s.ownerLabels, config.ServiceRoles.AsResources()...)
	if err := s.serviceRoleReconciler.Reconcile(
		"",
		config.ServiceRoles,
		nil,
		clients.ListOpts{
			Ctx:      ctx,
			Selector: s.ownerLabels,
		},
	); err != nil {
		return errors.Wrapf(err, "reconciling service roles")
	}

	logger.Infof("ServiceRoleBindings: %v", config.ServiceRoleBindings.Names())
	utils.SetLabels(s.ownerLabels, config.ServiceRoleBindings.AsResources()...)
	if err := s.serviceRoleBindingReconciler.Reconcile(
		"",
		config.ServiceRoleBindings,
		nil,
		clients.ListOpts{
			Ctx:      ctx,
			Selector: s.ownerLabels,
		},
	); err != nil {
		return errors.Wrapf(err, "reconciling service role bindings")
	}

	return nil
}
