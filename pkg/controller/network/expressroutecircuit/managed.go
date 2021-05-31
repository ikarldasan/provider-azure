/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package subnet

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
)

// Error strings.
const (
// errNotExpressRouteCircuit    = "managed resource is not an Express Route Circuit"
// errCreateExpressRouteCircuit = "cannot create Express Route Circuit"
// errUpdateExpressRouteCircuit = "cannot update Express Route Circuit"
// errGetExpressRouteCircuit    = "cannot get Express Route Circuit"
// errDeleteExpressRouteCircuit = "cannot delete Express Route Circuit"
)

// Setup adds a controller that reconciles Subnets.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter) error {
	name := managed.ControllerName(v1alpha3.ExpressRouteCircuitGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewDefaultManagedRateLimiter(rl),
		}).
		For(&v1alpha3.ExpressRouteCircuit{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha3.ExpressRouteCircuitGroupVersionKind),
			managed.WithConnectionPublishers(),
			managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

// type connecter struct {
// 	client client.Client
// }

// func (c *connecter) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
// creds, auth, err := azureclients.GetAuthInfo(ctx, c.client, mg)
// if err != nil {
// 	return nil, err
// }
// cl := azurenetwork.NewSubnetsClient(creds[azureclients.CredentialsKeySubscriptionID])
// cl.Authorizer = auth
// return &external{client: cl}, nil
// return &external{}, nil
// }

// type external struct{ client networkapi.SubnetsClientAPI }

// func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
// 	// s, ok := mg.(*v1alpha3.Subnet)
// 	// if !ok {
// 	// 	return managed.ExternalObservation{}, errors.New(errNotExpressRouteCircuit)
// 	// }

// 	// az, err := e.client.Get(ctx, s.Spec.ResourceGroupName, s.Spec.VirtualNetworkName, meta.GetExternalName(s), "")
// 	// if azureclients.IsNotFound(err) {
// 	// 	return managed.ExternalObservation{ResourceExists: false}, nil
// 	// }
// 	// if err != nil {
// 	// 	return managed.ExternalObservation{}, errors.Wrap(err, errGetSubnet)
// 	// }

// 	// network.UpdateSubnetStatusFromAzure(s, az)
// 	// s.SetConditions(xpv1.Available())

// 	o := managed.ExternalObservation{
// 		ResourceExists:    true,
// 		ConnectionDetails: managed.ConnectionDetails{},
// 	}

// 	// return o, nil
// 	return o, nil
// }

// func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
// 	// s, ok := mg.(*v1alpha3.Subnet)
// 	// if !ok {
// 	// 	return managed.ExternalCreation{}, errors.New(errNotSubnet)
// 	// }

// 	// s.Status.SetConditions(xpv1.Creating())

// 	// snet := network.NewSubnetParameters(s)
// 	// if _, err := e.client.CreateOrUpdate(ctx, s.Spec.ResourceGroupName, s.Spec.VirtualNetworkName, meta.GetExternalName(s), snet); err != nil {
// 	// 	return managed.ExternalCreation{}, errors.Wrap(err, errCreateSubnet)
// 	// }

// 	return managed.ExternalCreation{}, nil
// }

// func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
// 	// s, ok := mg.(*v1alpha3.Subnet)
// 	// if !ok {
// 	// 	return managed.ExternalUpdate{}, errors.New(errNotSubnet)
// 	// }

// 	// az, err := e.client.Get(ctx, s.Spec.ResourceGroupName, s.Spec.VirtualNetworkName, meta.GetExternalName(s), "")
// 	// if err != nil {
// 	// 	return managed.ExternalUpdate{}, errors.Wrap(err, errGetSubnet)
// 	// }

// 	// if network.SubnetNeedsUpdate(s, az) {
// 	// 	snet := network.NewSubnetParameters(s)
// 	// 	if _, err := e.client.CreateOrUpdate(ctx, s.Spec.ResourceGroupName, s.Spec.VirtualNetworkName, meta.GetExternalName(s), snet); err != nil {
// 	// 		return managed.ExternalUpdate{}, errors.Wrap(err, errUpdateSubnet)
// 	// 	}
// 	// }
// 	return managed.ExternalUpdate{}, nil
// }

// func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
// 	// s, ok := mg.(*v1alpha3.Subnet)
// 	// if !ok {
// 	// 	return errors.New(errNotSubnet)
// 	// }

// 	// mg.SetConditions(xpv1.Deleting())

// 	// _, err := e.client.Delete(ctx, s.Spec.ResourceGroupName, s.Spec.VirtualNetworkName, meta.GetExternalName(s))
// 	// return errors.Wrap(resource.Ignore(azureclients.IsNotFound, err), errDeleteSubnet)
// 	return errors.Wrap(resource.Ignore(azureclients.IsNotFound, nil), errDeleteExpressRouteCircuit)
// }
