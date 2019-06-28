// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	etcdv1 "github.com/openshift/api/etcd/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterMemberRequests implements ClusterMemberRequestInterface
type FakeClusterMemberRequests struct {
	Fake *FakeEtcdV1
}

var clustermemberrequestsResource = schema.GroupVersionResource{Group: "etcd.openshift.io", Version: "v1", Resource: "clustermemberrequests"}

var clustermemberrequestsKind = schema.GroupVersionKind{Group: "etcd.openshift.io", Version: "v1", Kind: "ClusterMemberRequest"}

// Get takes name of the clusterMemberRequest, and returns the corresponding clusterMemberRequest object, and an error if there is any.
func (c *FakeClusterMemberRequests) Get(name string, options v1.GetOptions) (result *etcdv1.ClusterMemberRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustermemberrequestsResource, name), &etcdv1.ClusterMemberRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*etcdv1.ClusterMemberRequest), err
}

// List takes label and field selectors, and returns the list of ClusterMemberRequests that match those selectors.
func (c *FakeClusterMemberRequests) List(opts v1.ListOptions) (result *etcdv1.ClusterMemberRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustermemberrequestsResource, clustermemberrequestsKind, opts), &etcdv1.ClusterMemberRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &etcdv1.ClusterMemberRequestList{ListMeta: obj.(*etcdv1.ClusterMemberRequestList).ListMeta}
	for _, item := range obj.(*etcdv1.ClusterMemberRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterMemberRequests.
func (c *FakeClusterMemberRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustermemberrequestsResource, opts))
}

// Create takes the representation of a clusterMemberRequest and creates it.  Returns the server's representation of the clusterMemberRequest, and an error, if there is any.
func (c *FakeClusterMemberRequests) Create(clusterMemberRequest *etcdv1.ClusterMemberRequest) (result *etcdv1.ClusterMemberRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustermemberrequestsResource, clusterMemberRequest), &etcdv1.ClusterMemberRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*etcdv1.ClusterMemberRequest), err
}

// Update takes the representation of a clusterMemberRequest and updates it. Returns the server's representation of the clusterMemberRequest, and an error, if there is any.
func (c *FakeClusterMemberRequests) Update(clusterMemberRequest *etcdv1.ClusterMemberRequest) (result *etcdv1.ClusterMemberRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustermemberrequestsResource, clusterMemberRequest), &etcdv1.ClusterMemberRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*etcdv1.ClusterMemberRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterMemberRequests) UpdateStatus(clusterMemberRequest *etcdv1.ClusterMemberRequest) (*etcdv1.ClusterMemberRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustermemberrequestsResource, "status", clusterMemberRequest), &etcdv1.ClusterMemberRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*etcdv1.ClusterMemberRequest), err
}

// Delete takes name of the clusterMemberRequest and deletes it. Returns an error if one occurs.
func (c *FakeClusterMemberRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clustermemberrequestsResource, name), &etcdv1.ClusterMemberRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterMemberRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustermemberrequestsResource, listOptions)

	_, err := c.Fake.Invokes(action, &etcdv1.ClusterMemberRequestList{})
	return err
}

// Patch applies the patch and returns the patched clusterMemberRequest.
func (c *FakeClusterMemberRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *etcdv1.ClusterMemberRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustermemberrequestsResource, name, pt, data, subresources...), &etcdv1.ClusterMemberRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*etcdv1.ClusterMemberRequest), err
}
