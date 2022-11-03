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

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	secondexamplev1 "acme.corp/pkg/apis/secondexample/v1"
	secondexamplev1listers "acme.corp/pkg/generated/listers/secondexample/v1"
)

// TestTypeClusterLister can list TestTypes across all workspaces, or scope down to a TestTypeLister for one workspace.
// All objects returned here must be treated as read-only.
type TestTypeClusterLister interface {
	// List lists all TestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*secondexamplev1.TestType, err error)
	// Cluster returns a lister that can list and get TestTypes in one workspace.
	Cluster(cluster logicalcluster.Name) secondexamplev1listers.TestTypeLister
	TestTypeClusterListerExpansion
}

type testTypeClusterLister struct {
	indexer cache.Indexer
}

// NewTestTypeClusterLister returns a new TestTypeClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewTestTypeClusterLister(indexer cache.Indexer) *testTypeClusterLister {
	return &testTypeClusterLister{indexer: indexer}
}

// List lists all TestTypes in the indexer across all workspaces.
func (s *testTypeClusterLister) List(selector labels.Selector) (ret []*secondexamplev1.TestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*secondexamplev1.TestType))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get TestTypes.
func (s *testTypeClusterLister) Cluster(cluster logicalcluster.Name) secondexamplev1listers.TestTypeLister {
	return &testTypeLister{indexer: s.indexer, cluster: cluster}
}

// testTypeLister implements the secondexamplev1listers.TestTypeLister interface.
type testTypeLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all TestTypes in the indexer for a workspace.
func (s *testTypeLister) List(selector labels.Selector) (ret []*secondexamplev1.TestType, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*secondexamplev1.TestType))
	})
	return ret, err
}

// TestTypes returns an object that can list and get TestTypes in one namespace.
func (s *testTypeLister) TestTypes(namespace string) secondexamplev1listers.TestTypeNamespaceLister {
	return &testTypeNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// testTypeNamespaceLister implements the secondexamplev1listers.TestTypeNamespaceLister interface.
type testTypeNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all TestTypes in the indexer for a given workspace and namespace.
func (s *testTypeNamespaceLister) List(selector labels.Selector) (ret []*secondexamplev1.TestType, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*secondexamplev1.TestType))
	})
	return ret, err
}

// Get retrieves the TestType from the indexer for a given workspace, namespace and name.
func (s *testTypeNamespaceLister) Get(name string) (*secondexamplev1.TestType, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(secondexamplev1.Resource("TestType"), name)
	}
	return obj.(*secondexamplev1.TestType), nil
}
