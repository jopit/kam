/*
Copyright Red Hat, Inc.

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

	v1alpha2 "github.com/operator-framework/api/pkg/operators/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperatorGroups implements OperatorGroupInterface
type FakeOperatorGroups struct {
	Fake *FakeOperatorsV1alpha2
	ns   string
}

var operatorgroupsResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "v1alpha2", Resource: "operatorgroups"}

var operatorgroupsKind = schema.GroupVersionKind{Group: "operators.coreos.com", Version: "v1alpha2", Kind: "OperatorGroup"}

// Get takes name of the operatorGroup, and returns the corresponding operatorGroup object, and an error if there is any.
func (c *FakeOperatorGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operatorgroupsResource, c.ns, name), &v1alpha2.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.OperatorGroup), err
}

// List takes label and field selectors, and returns the list of OperatorGroups that match those selectors.
func (c *FakeOperatorGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.OperatorGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(operatorgroupsResource, operatorgroupsKind, c.ns, opts), &v1alpha2.OperatorGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.OperatorGroupList{ListMeta: obj.(*v1alpha2.OperatorGroupList).ListMeta}
	for _, item := range obj.(*v1alpha2.OperatorGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operatorGroups.
func (c *FakeOperatorGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(operatorgroupsResource, c.ns, opts))

}

// Create takes the representation of a operatorGroup and creates it.  Returns the server's representation of the operatorGroup, and an error, if there is any.
func (c *FakeOperatorGroups) Create(ctx context.Context, operatorGroup *v1alpha2.OperatorGroup, opts v1.CreateOptions) (result *v1alpha2.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(operatorgroupsResource, c.ns, operatorGroup), &v1alpha2.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.OperatorGroup), err
}

// Update takes the representation of a operatorGroup and updates it. Returns the server's representation of the operatorGroup, and an error, if there is any.
func (c *FakeOperatorGroups) Update(ctx context.Context, operatorGroup *v1alpha2.OperatorGroup, opts v1.UpdateOptions) (result *v1alpha2.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(operatorgroupsResource, c.ns, operatorGroup), &v1alpha2.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.OperatorGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperatorGroups) UpdateStatus(ctx context.Context, operatorGroup *v1alpha2.OperatorGroup, opts v1.UpdateOptions) (*v1alpha2.OperatorGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(operatorgroupsResource, "status", c.ns, operatorGroup), &v1alpha2.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.OperatorGroup), err
}

// Delete takes name of the operatorGroup and deletes it. Returns an error if one occurs.
func (c *FakeOperatorGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(operatorgroupsResource, c.ns, name, opts), &v1alpha2.OperatorGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperatorGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(operatorgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.OperatorGroupList{})
	return err
}

// Patch applies the patch and returns the patched operatorGroup.
func (c *FakeOperatorGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.OperatorGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(operatorgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha2.OperatorGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.OperatorGroup), err
}
