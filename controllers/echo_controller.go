/*


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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testappv1 "example/api/v1"
)

// EchoReconciler reconciles a Echo object
type EchoReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=testapp.my.domain,resources=echoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=testapp.my.domain,resources=echoes/status,verbs=get;update;patch

func (r *EchoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("echo", req.NamespacedName)

	var echo testappv1.Echo
	if err := r.Get(ctx, req.NamespacedName, &echo); err != nil {
		log.Error(err, "unable to fetch Echo")

		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.V(1).Info("Get echo successfully", "echo entity", fmt.Sprintf("%v", echo))

	resource := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": echo.Spec.APIVersion,
			"kind":       echo.Spec.Kind,
		},
	}

	nn := types.NamespacedName{
		Name:      echo.Spec.Name,
		Namespace: echo.Spec.Namespace,
	}

	if err := r.Get(ctx, nn, resource); err != nil {
		log.Error(err, "unable to fetch specified resource")

		return ctrl.Result{}, err
	}

	field, found, err := unstructured.NestedFieldCopy(resource.Object, strings.Split(echo.Spec.RefPath, ".")...)
	if err != nil {
		log.Error(err, "failed to get nested field")

		return ctrl.Result{}, err
	}

	if !found {
		log.Error(fmt.Errorf("refpath not exist"), "failed to get refpath")

		return ctrl.Result{}, err
	}

	fieldData, err := json.Marshal(field)
	if err != nil {
		log.Error(err, "marshal field failed")

		return ctrl.Result{}, err
	}

	log.V(1).Info("Get RefPath successfully", "RefPath", string(fieldData))

	echo.Status.Data.Raw = fieldData

	if err := r.Status().Update(ctx, &echo); err != nil {
		log.Error(err, "unable to update echo status")

		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *EchoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testappv1.Echo{}).
		Complete(r)
}
