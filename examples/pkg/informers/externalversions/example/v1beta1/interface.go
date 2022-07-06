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
	"github.com/kcp-dev/code-generator/examples/pkg/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterTestTypes returns a ClusterTestTypeInformer.
	ClusterTestTypes() ClusterTestTypeInformer
	// TestTypes returns a TestTypeInformer.
	TestTypes() TestTypeInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterTestTypes returns a ClusterTestTypeInformer.
func (v *version) ClusterTestTypes() ClusterTestTypeInformer {
	return &clusterTestTypeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// TestTypes returns a TestTypeInformer.
func (v *version) TestTypes() TestTypeInformer {
	return &testTypeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
