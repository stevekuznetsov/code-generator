//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1beta1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	examplev1beta1 "acme.corp/pkg/apis/example/v1beta1"
	examplev1beta1client "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1"
)

// ClusterTestTypesClusterGetter has a method to return a ClusterTestTypeClusterInterface.
// A group's cluster client should implement this interface.
type ClusterTestTypesClusterGetter interface {
	ClusterTestTypes() ClusterTestTypeClusterInterface
}

// ClusterTestTypeClusterInterface can operate on ClusterTestTypes across all clusters,
// or scope down to one cluster and return a examplev1beta1client.ClusterTestTypeInterface.
type ClusterTestTypeClusterInterface interface {
	Cluster(logicalcluster.Name) examplev1beta1client.ClusterTestTypeInterface
	List(ctx context.Context, opts metav1.ListOptions) (*examplev1beta1.ClusterTestTypeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type clusterTestTypesClusterInterface struct {
	clientCache kcpclient.Cache[*examplev1beta1client.ExampleV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterTestTypesClusterInterface) Cluster(name logicalcluster.Name) examplev1beta1client.ClusterTestTypeInterface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(name).ClusterTestTypes()
}

// List returns the entire collection of all ClusterTestTypes across all clusters.
func (c *clusterTestTypesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*examplev1beta1.ClusterTestTypeList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterTestTypes().List(ctx, opts)
}

// Watch begins to watch all ClusterTestTypes across all clusters.
func (c *clusterTestTypesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterTestTypes().Watch(ctx, opts)
}
