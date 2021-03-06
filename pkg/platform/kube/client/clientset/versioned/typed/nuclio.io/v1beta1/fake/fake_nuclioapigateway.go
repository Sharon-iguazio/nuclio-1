/*
Copyright The Kubernetes Authors.

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
	v1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/apis/nuclio.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNuclioAPIGateways implements NuclioAPIGatewayInterface
type FakeNuclioAPIGateways struct {
	Fake *FakeNuclioV1beta1
	ns   string
}

var nuclioapigatewaysResource = schema.GroupVersionResource{Group: "nuclio.io", Version: "v1beta1", Resource: "nuclioapigateways"}

var nuclioapigatewaysKind = schema.GroupVersionKind{Group: "nuclio.io", Version: "v1beta1", Kind: "NuclioAPIGateway"}

// Get takes name of the nuclioAPIGateway, and returns the corresponding nuclioAPIGateway object, and an error if there is any.
func (c *FakeNuclioAPIGateways) Get(name string, options v1.GetOptions) (result *v1beta1.NuclioAPIGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nuclioapigatewaysResource, c.ns, name), &v1beta1.NuclioAPIGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NuclioAPIGateway), err
}

// List takes label and field selectors, and returns the list of NuclioAPIGateways that match those selectors.
func (c *FakeNuclioAPIGateways) List(opts v1.ListOptions) (result *v1beta1.NuclioAPIGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nuclioapigatewaysResource, nuclioapigatewaysKind, c.ns, opts), &v1beta1.NuclioAPIGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NuclioAPIGatewayList{ListMeta: obj.(*v1beta1.NuclioAPIGatewayList).ListMeta}
	for _, item := range obj.(*v1beta1.NuclioAPIGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nuclioAPIGateways.
func (c *FakeNuclioAPIGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nuclioapigatewaysResource, c.ns, opts))

}

// Create takes the representation of a nuclioAPIGateway and creates it.  Returns the server's representation of the nuclioAPIGateway, and an error, if there is any.
func (c *FakeNuclioAPIGateways) Create(nuclioAPIGateway *v1beta1.NuclioAPIGateway) (result *v1beta1.NuclioAPIGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nuclioapigatewaysResource, c.ns, nuclioAPIGateway), &v1beta1.NuclioAPIGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NuclioAPIGateway), err
}

// Update takes the representation of a nuclioAPIGateway and updates it. Returns the server's representation of the nuclioAPIGateway, and an error, if there is any.
func (c *FakeNuclioAPIGateways) Update(nuclioAPIGateway *v1beta1.NuclioAPIGateway) (result *v1beta1.NuclioAPIGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nuclioapigatewaysResource, c.ns, nuclioAPIGateway), &v1beta1.NuclioAPIGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NuclioAPIGateway), err
}

// Delete takes name of the nuclioAPIGateway and deletes it. Returns an error if one occurs.
func (c *FakeNuclioAPIGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nuclioapigatewaysResource, c.ns, name), &v1beta1.NuclioAPIGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNuclioAPIGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nuclioapigatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.NuclioAPIGatewayList{})
	return err
}

// Patch applies the patch and returns the patched nuclioAPIGateway.
func (c *FakeNuclioAPIGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.NuclioAPIGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nuclioapigatewaysResource, c.ns, name, pt, data, subresources...), &v1beta1.NuclioAPIGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NuclioAPIGateway), err
}
