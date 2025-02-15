/*
Copyright 2022 The KCP Authors.

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

package informergen

import (
	"io"
	"text/template"

	"k8s.io/code-generator/cmd/client-gen/types"

	"github.com/kcp-dev/code-generator/pkg/parser"
)

type Generic struct {
	// Groups are the groups in this informer factory.
	Groups []types.GroupVersionInfo

	// GroupVersionKinds are all the kinds we need to support,indexed by group and version.
	GroupVersionKinds map[types.Group]map[types.Version][]parser.Kind

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string
}

func (g *Generic) WriteContent(w io.Writer) error {
	templ, err := template.New("generic").Funcs(templateFuncs).Parse(genericInformer)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"groups":            g.Groups,
		"groupVersionKinds": g.GroupVersionKinds,
		"apiPackagePath":    g.APIPackagePath,
	}
	return templ.Execute(w, m)
}

var genericInformer = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"

{{range .groups}}	{{.PackageAlias}} "{{$.apiPackagePath}}/{{.Group.PackageName}}/{{.Version.PackageName}}"
{{end -}}
)

type GenericClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kcpcache.GenericClusterLister
}

type genericClusterInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericClusterInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericClusterLister.
func (f *genericClusterInformer) Lister() kcpcache.GenericClusterLister {
	return kcpcache.NewGenericClusterLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error) {
	switch resource {
{{range $group := .groups}}	// Group={{.Group.NonEmpty}}, Version={{.Version}}
{{range $kind := index (index $.groupVersionKinds .Group) .Version}}	case {{$group.PackageAlias}}.SchemeGroupVersion.WithResource("{{$kind.Plural|toLower}}"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.{{$group.GroupGoName}}().{{$group.Version}}().{{$kind.Plural}}().Informer()}, nil
{{end -}}
{{end -}}
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
`
