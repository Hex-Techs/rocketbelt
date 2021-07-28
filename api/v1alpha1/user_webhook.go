/*
Copyright 2021 The.

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
	"crypto/sha1"
	"crypto/subtle"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var userlog = logf.Log.WithName("user-resource")

func (r *User) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:verbs=create;update,path=/mutate-rocketbelt-hextech-com-v1alpha1-user,mutating=true,failurePolicy=fail,sideEffects=None,groups=rocketbelt.hextech.com,resources=users,verbs=create;update,versions=v1alpha1,name=muser.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &User{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *User) Default() {
	userlog.Info("default", "name", r.Name)
	if r.Spec.Password != "" {
		r.Spec.Password = Generate(r.Spec.Password)
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:verbs=create;update,path=/validate-rocketbelt-hextech-com-v1alpha1-user,mutating=false,failurePolicy=fail,sideEffects=None,groups=rocketbelt.hextech.com,resources=users,verbs=create;update,versions=v1alpha1,name=vuser.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &User{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *User) ValidateCreate() error {
	userlog.Info("validate create", "name", r.Name)

	return r.validateUser()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *User) ValidateUpdate(old runtime.Object) error {
	userlog.Info("validate update", "name", r.Name)

	return r.validateUser()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *User) ValidateDelete() error {
	userlog.Info("validate delete", "name", r.Name)

	return nil
}

func (r *User) validateUser() error {
	if r.Spec.DisplayName == "" {
		return errors.New("must given a displayName")
	}
	if r.Spec.Password == "" {
		return errors.New("password can not empty")
	}
	return nil
}

func Generate(p string) string {
	return fmt.Sprintf("%x",
		pbkdf2.Key([]byte(p), []byte("rocketbelt"), 4096, 32, sha1.New))
}

func Validate(p1, p2 string) bool {
	return subtle.ConstantTimeCompare([]byte(p1), []byte(p2)) == 1
}
