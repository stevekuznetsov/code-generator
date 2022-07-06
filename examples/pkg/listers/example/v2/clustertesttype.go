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

package v2

import (
	apimachinerycache "github.com/kcp-dev/apimachinery/pkg/cache"
	examplev2 "github.com/kcp-dev/code-generator/examples/pkg/apis/example/v2"
	"github.com/kcp-dev/logicalcluster"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterTestTypeLister helps list examplev2.ClusterTestType.
// All objects returned here must be treated as read-only.
type ClusterTestTypeClusterLister interface {
	// List lists all examplev2.ClusterTestType in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*examplev2.ClusterTestType, err error)

	// Cluster returns an object that can list and get examplev2.ClusterTestType from the given logical cluster.
	Cluster(cluster logicalcluster.Name) ClusterTestTypeLister
}

// clusterTestTypeClusterLister implements the ClusterTestTypeClusterLister interface.
type clusterTestTypeClusterLister struct {
	indexer cache.Indexer
}

// NewClusterTestTypeClusterLister returns a new ClusterTestTypeClusterLister.
func NewClusterTestTypeClusterLister(indexer cache.Indexer) ClusterTestTypeClusterLister {
	return &clusterTestTypeClusterLister{indexer: indexer}
}

// List lists all examplev2.ClusterTestType in the indexer.
func (s *clusterTestTypeClusterLister) List(selector labels.Selector) (ret []*examplev2.ClusterTestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*examplev2.ClusterTestType))
	})
	return ret, err
}

// Cluster returns an object that can list and get examplev2.ClusterTestType.
func (s *clusterTestTypeClusterLister) Cluster(cluster logicalcluster.Name) ClusterTestTypeLister {
	return &clusterTestTypeLister{indexer: s.indexer, cluster: cluster}
}

// ClusterTestTypeLister helps list examplev2.ClusterTestType.
// All objects returned here must be treated as read-only.
type ClusterTestTypeLister interface {
	// List lists all examplev2.ClusterTestType in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*examplev2.ClusterTestType, err error)
	// Get retrieves the examplev2.ClusterTestType from the indexer for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*examplev2.ClusterTestType, error)
}

// clusterTestTypeLister implements the ClusterTestTypeLister interface.
type clusterTestTypeLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all examplev2.ClusterTestType in the indexer.
func (s *clusterTestTypeLister) List(selector labels.Selector) (ret []*examplev2.ClusterTestType, err error) {
	selectAll := selector == nil || selector.Empty()

	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), "", "")
	list, err := s.indexer.ByIndex(apimachinerycache.ClusterIndexName, key)
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*examplev2.ClusterTestType)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// Get retrieves the examplev2.ClusterTestType from the indexer for a given name.
func (s clusterTestTypeLister) Get(name string) (*examplev2.ClusterTestType, error) {
	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(examplev2.Resource("clusterTestType"), name)
	}
	return obj.(*examplev2.ClusterTestType), nil
}
