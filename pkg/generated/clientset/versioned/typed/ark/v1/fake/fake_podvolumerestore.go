/*
Copyright the Heptio Ark contributors.

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
	arkv1 "github.com/heptio/velero/pkg/apis/ark/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodVolumeRestores implements PodVolumeRestoreInterface
type FakePodVolumeRestores struct {
	Fake *FakeArkV1
	ns   string
}

var podvolumerestoresResource = schema.GroupVersionResource{Group: "ark.heptio.com", Version: "v1", Resource: "podvolumerestores"}

var podvolumerestoresKind = schema.GroupVersionKind{Group: "ark.heptio.com", Version: "v1", Kind: "PodVolumeRestore"}

// Get takes name of the podVolumeRestore, and returns the corresponding podVolumeRestore object, and an error if there is any.
func (c *FakePodVolumeRestores) Get(name string, options v1.GetOptions) (result *arkv1.PodVolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podvolumerestoresResource, c.ns, name), &arkv1.PodVolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.PodVolumeRestore), err
}

// List takes label and field selectors, and returns the list of PodVolumeRestores that match those selectors.
func (c *FakePodVolumeRestores) List(opts v1.ListOptions) (result *arkv1.PodVolumeRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podvolumerestoresResource, podvolumerestoresKind, c.ns, opts), &arkv1.PodVolumeRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &arkv1.PodVolumeRestoreList{ListMeta: obj.(*arkv1.PodVolumeRestoreList).ListMeta}
	for _, item := range obj.(*arkv1.PodVolumeRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podVolumeRestores.
func (c *FakePodVolumeRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podvolumerestoresResource, c.ns, opts))

}

// Create takes the representation of a podVolumeRestore and creates it.  Returns the server's representation of the podVolumeRestore, and an error, if there is any.
func (c *FakePodVolumeRestores) Create(podVolumeRestore *arkv1.PodVolumeRestore) (result *arkv1.PodVolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podvolumerestoresResource, c.ns, podVolumeRestore), &arkv1.PodVolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.PodVolumeRestore), err
}

// Update takes the representation of a podVolumeRestore and updates it. Returns the server's representation of the podVolumeRestore, and an error, if there is any.
func (c *FakePodVolumeRestores) Update(podVolumeRestore *arkv1.PodVolumeRestore) (result *arkv1.PodVolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podvolumerestoresResource, c.ns, podVolumeRestore), &arkv1.PodVolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.PodVolumeRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodVolumeRestores) UpdateStatus(podVolumeRestore *arkv1.PodVolumeRestore) (*arkv1.PodVolumeRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podvolumerestoresResource, "status", c.ns, podVolumeRestore), &arkv1.PodVolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.PodVolumeRestore), err
}

// Delete takes name of the podVolumeRestore and deletes it. Returns an error if one occurs.
func (c *FakePodVolumeRestores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podvolumerestoresResource, c.ns, name), &arkv1.PodVolumeRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodVolumeRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podvolumerestoresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &arkv1.PodVolumeRestoreList{})
	return err
}

// Patch applies the patch and returns the patched podVolumeRestore.
func (c *FakePodVolumeRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *arkv1.PodVolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podvolumerestoresResource, c.ns, name, data, subresources...), &arkv1.PodVolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*arkv1.PodVolumeRestore), err
}