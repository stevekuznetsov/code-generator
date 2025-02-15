package clientgen

import (
	"io"
	"strings"
	"text/template"

	"k8s.io/code-generator/cmd/client-gen/types"

	"github.com/kcp-dev/code-generator/pkg/parser"
	"github.com/kcp-dev/code-generator/pkg/util"
)

type Group struct {
	// Group is the group in this client.
	Group types.GroupVersionInfo

	// Kinds are the kinds in the group.
	Kinds []parser.Kind

	// SingleClusterClientPackagePath is the root directory under which single-cluster-aware clients exist.
	// e.g. "k8s.io/client-go/kubernetes"
	SingleClusterClientPackagePath string
}

func (g *Group) WriteContent(w io.Writer) error {
	templ, err := template.New("group").Funcs(template.FuncMap{
		"upperFirst": util.UpperFirst,
		"lowerFirst": util.LowerFirst,
		"toLower":    strings.ToLower,
	}).Parse(group)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"group":                          g.Group,
		"kinds":                          g.Kinds,
		"singleClusterClientPackagePath": g.SingleClusterClientPackagePath,
	}
	return templ.Execute(w, m)
}

var group = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package {{.group.Version.PackageName}}

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v2"

	"k8s.io/client-go/rest"

	{{.group.PackageAlias}} "{{.singleClusterClientPackagePath}}/typed/{{.group.Group.PackageName}}/{{.group.Version.PackageName}}"
)

type {{.group.GroupGoName}}{{.group.Version}}ClusterInterface interface {
	{{.group.GroupGoName}}{{.group.Version}}ClusterScoper
{{range .kinds}}	{{.Plural}}ClusterGetter
{{end -}}
}

type {{.group.GroupGoName}}{{.group.Version}}ClusterScoper interface {
	Cluster(logicalcluster.Name) {{.group.PackageAlias}}.{{.group.GroupGoName}}{{.group.Version}}Interface
}

type {{.group.GroupGoName}}{{.group.Version}}ClusterClient struct {
	clientCache kcpclient.Cache[*{{.group.PackageAlias}}.{{.group.GroupGoName}}{{.group.Version}}Client]
}

func (c *{{.group.GroupGoName}}{{.group.Version}}ClusterClient) Cluster(name logicalcluster.Name) {{.group.PackageAlias}}.{{.group.GroupGoName}}{{.group.Version}}Interface {
	if name == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(name)
}

{{ range .kinds}}
func (c *{{$.group.GroupGoName}}{{$.group.Version}}ClusterClient) {{.Plural}}() {{.String}}ClusterInterface {
	return &{{.Plural | lowerFirst}}ClusterInterface{clientCache: c.clientCache}
}
{{end -}}

// NewForConfig creates a new {{.group.GroupGoName}}{{.group.Version}}ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*{{.group.GroupGoName}}{{.group.Version}}ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new {{.group.GroupGoName}}{{.group.Version}}ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*{{.group.GroupGoName}}{{.group.Version}}ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*{{.group.PackageAlias}}.{{.group.GroupGoName}}{{.group.Version}}Client]{
		NewForConfigAndClient: {{.group.PackageAlias}}.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.New("root")); err != nil {
		return nil, err
	}
	return &{{.group.GroupGoName}}{{.group.Version}}ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new {{.group.GroupGoName}}{{.group.Version}}ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *{{.group.GroupGoName}}{{.group.Version}}ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
`
