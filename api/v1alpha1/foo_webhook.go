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

package v1alpha1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var foolog = logf.Log.WithName("foo-resource")

func (r *Foo) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-samplecontroller-k8s-io-v1alpha1-foo,mutating=true,failurePolicy=fail,groups=samplecontroller.k8s.io,resources=foos,verbs=create;update,versions=v1alpha1,name=mfoo.kb.io

var _ webhook.Defaulter = &Foo{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Foo) Default() {
	foolog.Info("default", "name", r.Name)

	if r.Spec.Replicas == nil {
		r.Spec.Replicas = new(int32)
		*r.Spec.Replicas = 1
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-samplecontroller-k8s-io-v1alpha1-foo,mutating=false,failurePolicy=fail,groups=samplecontroller.k8s.io,resources=foos,versions=v1alpha1,name=vfoo.kb.io

var _ webhook.Validator = &Foo{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Foo) ValidateCreate() error {
	foolog.Info("validate create", "name", r.Name)

	return r.validateFoo()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Foo) ValidateUpdate(old runtime.Object) error {
	foolog.Info("validate update", "name", r.Name)

	return r.validateFoo()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Foo) ValidateDelete() error {
	foolog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

// We validate the DeployName of the Foo.
func (r *Foo) validateFoo() error {
	var allErrs field.ErrorList
	if err := r.validateDeploymentName(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "samplecontroller.k8s.io", Kind: "Foo"},
		r.Name, allErrs)
}

// Validating the the length of DeploymentName field.
func (r *Foo) validateDeploymentName() *field.Error {
	// object name must be no more than 253 characters.
	// ref. https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	if len(r.Spec.DeploymentName) > 253 {
		return field.Invalid(field.NewPath("spec").Child("deploymentName"), r.Spec.DeploymentName, "must be no more than 253 characters")
	}
	return nil
}
