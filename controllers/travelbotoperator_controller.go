/*
Copyright 2023.

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

package controllers

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	travelbotv1 "github.com/tcotav/travelbot-operator/api/v1"
	assets "github.com/tcotav/travelbot-operator/assets"
)

// TravelbotOperatorReconciler reconciles a TravelbotOperator object
type TravelbotOperatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=travelbot.ragainsth.com,resources=travelbotoperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=travelbot.ragainsth.com,resources=travelbotoperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=travelbot.ragainsth.com,resources=travelbotoperators/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TravelbotOperator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *TravelbotOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// first get the operator resource object
	operatorCR := &travelbotv1.TravelbotOperator{}
	err := r.Get(ctx, req.NamespacedName, operatorCR)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Operator resource object not found.")
		return ctrl.Result{}, nil
	} else if err != nil {
		logger.Error(err, "Error getting operator resource object")
		return ctrl.Result{}, err
	}

	// then get the deployment object from the cluster
	deployment := &appsv1.Deployment{}
	create := false
	err = r.Get(ctx, req.NamespacedName, deployment)
	// check if the deployment object is not found
	if err != nil && errors.IsNotFound(err) {
		// we don't find it running in the cluster so we set the flag
		// to create further down after creating a new deployment object from the manifest
		create = true
		deployment = assets.GetDeploymentFromFile("manifests/deployment.yaml")
	} else if err != nil {
		logger.Error(err, "Error getting existing Spaceship deployment.")
		return ctrl.Result{}, err
	}

	// set the deployment object's fields to match the operatorCR resource object
	deployment.Namespace = req.Namespace
	deployment.Name = req.Name

	// now we modify the command to match our CR
	command := []string{"/travelbot"}
	if operatorCR.Spec.ShipName != "" {
		command = append(command, fmt.Sprintf("--ship=%s", operatorCR.Spec.ShipName))
	}
	if operatorCR.Spec.ShipSpeed != "" {
		command = append(command, fmt.Sprintf("--speed=%s", operatorCR.Spec.ShipSpeed))
	}
	if operatorCR.Spec.SleepDuration != "" {
		command = append(command, fmt.Sprintf("--sleep=%s", operatorCR.Spec.SleepDuration))
	}
	deployment.Spec.Template.Spec.Containers[0].Command = command

	// now we modify the image to match our CR
	if operatorCR.Spec.DeployImage != "" {
		deployment.Spec.Template.Spec.Containers[0].Image = operatorCR.Spec.DeployImage
	}

	// indicates that the operatorCR resource object should be listed as the OwnerReference of the deployment
	// this is an API field denoting which object owns this object
	ctrl.SetControllerReference(operatorCR, deployment, r.Scheme)

	// perform the appropriate action based on the create flag
	if create {
		err = r.Create(ctx, deployment)
	} else {
		err = r.Update(ctx, deployment)
	}

	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *TravelbotOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&travelbotv1.TravelbotOperator{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
