package clientgen

import (
	"io"
	"strings"
	"text/template"

	"k8s.io/code-generator/cmd/client-gen/types"

	"github.com/kcp-dev/code-generator/pkg/util"
)

type ClientSet struct {
	// Name is the name of the clientset, e.g. "kubernetes"
	Name string

	// Groups are the groups in this client-set.
	Groups []types.GroupVersionInfo

	// PackagePath is the package under which this client-set will be exposed.
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	PackagePath string

	// SingleClusterClientPackagePath is the root directory under which single-cluster-aware clients exist.
	// e.g. "k8s.io/client-go/kubernetes"
	SingleClusterClientPackagePath string
}

func (c *ClientSet) WriteContent(w io.Writer) error {
	templ, err := template.New("clientset").Funcs(template.FuncMap{
		"upperFirst": util.UpperFirst,
		"lowerFirst": util.LowerFirst,
		"toLower":    strings.ToLower,
	}).Parse(clientset)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"name":                           c.Name,
		"packagePath":                    c.PackagePath,
		"groups":                         c.Groups,
		"singleClusterClientPackagePath": c.SingleClusterClientPackagePath,
	}
	return templ.Execute(w, m)
}

var clientset = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package {{.name}}

import (
	"fmt"
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	client "{{.singleClusterClientPackagePath}}"
	
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"

{{range .groups}}	{{.PackageAlias}} "{{$.packagePath}}/typed/{{.Group.PackageName}}/{{.Version.PackageName}}"
{{end -}}
)

type ClusterInterface interface {
	Cluster(logicalcluster.Path) client.Interface
	Discovery() discovery.DiscoveryInterface
{{range .groups}}	{{.GroupGoName}}{{.Version}}() {{.PackageAlias}}.{{.GroupGoName}}{{.Version}}ClusterInterface
{{end -}}
}

// ClusterClientset contains the clients for groups.
type ClusterClientset struct {
	*discovery.DiscoveryClient
	clientCache kcpclient.Cache[*client.Clientset]
{{range .groups}}	{{.LowerCaseGroupGoName}}{{.Version}} *{{.PackageAlias}}.{{.GroupGoName}}{{.Version}}ClusterClient
{{end -}}
}

// Discovery retrieves the DiscoveryClient
func (c *ClusterClientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

{{range .groups}}
// {{.GroupGoName}}{{.Version}} retrieves the {{.GroupGoName}}{{.Version}}ClusterClient.  
func (c *ClusterClientset) {{.GroupGoName}}{{.Version}}() {{.PackageAlias}}.{{.GroupGoName}}{{.Version}}ClusterInterface {
	return c.{{.LowerCaseGroupGoName}}{{.Version}}
}
{{end -}}

// Cluster scopes this clientset to one cluster.
func (c *ClusterClientset) Cluster(clusterPath logicalcluster.Path) client.Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

// NewForConfig creates a new ClusterClientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable, 
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ClusterClientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new ClusterClientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*ClusterClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	cache := kcpclient.NewCache(c, httpClient, &kcpclient.Constructor[*client.Clientset]{
		NewForConfigAndClient: client.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}

	var cs ClusterClientset
	cs.clientCache = cache
	var err error
{{range .groups}}    cs.{{.LowerCaseGroupGoName}}{{.Version}}, err = {{.PackageAlias}}.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
{{end}}
	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new ClusterClientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterClientset {
	cs, err := NewForConfig(c)
	if err!=nil {
		panic(err)
	}
	return cs
}
`
