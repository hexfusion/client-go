// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/client-go/etcd/clientset/versioned/typed/etcd/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEtcdV1 struct {
	*testing.Fake
}

func (c *FakeEtcdV1) ClusterMemberRequests() v1.ClusterMemberRequestInterface {
	return &FakeClusterMemberRequests{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEtcdV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
