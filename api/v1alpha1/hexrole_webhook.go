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
	"fmt"

	"github.com/gookit/goutil/arrutil"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var hexrolelog = logf.Log.WithName("hexrole-resource")

func (r *HexRole) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-rocketbelt-hextech-com-v1alpha1-hexrole,mutating=false,failurePolicy=fail,sideEffects=None,groups=rocketbelt.hextech.com,resources=hexroles,verbs=create;update,versions=v1alpha1,name=vhexrole.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &HexRole{}

var verbs = []string{string(Get), string(Create), string(Delete), string(Update), string(List)}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *HexRole) ValidateCreate() error {
	hexrolelog.Info("validate create", "name", r.Name)

	return r.validateHexRole()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *HexRole) ValidateUpdate(old runtime.Object) error {
	hexrolelog.Info("validate update", "name", r.Name)

	return r.validateHexRole()
}

func (r *HexRole) validateHexRole() error {
	for _, rule := range r.Spec.Rules {
		for _, v := range rule.Verbs {
			if !arrutil.HasValue(verbs, string(v)) {
				return fmt.Errorf("%s is unknow verbs", v)
			}
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *HexRole) ValidateDelete() error {
	hexrolelog.Info("validate delete", "name", r.Name)

	return nil
}
