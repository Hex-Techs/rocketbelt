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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/hex-techs/rocketbelt/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHexRoles implements HexRoleInterface
type FakeHexRoles struct {
	Fake *FakeRocketbeltV1alpha1
}

var hexrolesResource = schema.GroupVersionResource{Group: "rocketbelt.hextech.com", Version: "v1alpha1", Resource: "hexroles"}

var hexrolesKind = schema.GroupVersionKind{Group: "rocketbelt.hextech.com", Version: "v1alpha1", Kind: "HexRole"}

// Get takes name of the hexRole, and returns the corresponding hexRole object, and an error if there is any.
func (c *FakeHexRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HexRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(hexrolesResource, name), &v1alpha1.HexRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HexRole), err
}

// List takes label and field selectors, and returns the list of HexRoles that match those selectors.
func (c *FakeHexRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HexRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(hexrolesResource, hexrolesKind, opts), &v1alpha1.HexRoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HexRoleList{ListMeta: obj.(*v1alpha1.HexRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.HexRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hexRoles.
func (c *FakeHexRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(hexrolesResource, opts))
}

// Create takes the representation of a hexRole and creates it.  Returns the server's representation of the hexRole, and an error, if there is any.
func (c *FakeHexRoles) Create(ctx context.Context, hexRole *v1alpha1.HexRole, opts v1.CreateOptions) (result *v1alpha1.HexRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(hexrolesResource, hexRole), &v1alpha1.HexRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HexRole), err
}

// Update takes the representation of a hexRole and updates it. Returns the server's representation of the hexRole, and an error, if there is any.
func (c *FakeHexRoles) Update(ctx context.Context, hexRole *v1alpha1.HexRole, opts v1.UpdateOptions) (result *v1alpha1.HexRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(hexrolesResource, hexRole), &v1alpha1.HexRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HexRole), err
}

// Delete takes name of the hexRole and deletes it. Returns an error if one occurs.
func (c *FakeHexRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(hexrolesResource, name), &v1alpha1.HexRole{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHexRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(hexrolesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HexRoleList{})
	return err
}

// Patch applies the patch and returns the patched hexRole.
func (c *FakeHexRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HexRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(hexrolesResource, name, pt, data, subresources...), &v1alpha1.HexRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HexRole), err
}
