package transport

import (
	"github.com/beclab/lldap-client/pkg/config"
	"k8s.io/klog/v2"
	"net/http"
)

func New(cfg *config.Config) (http.RoundTripper, error) {
	var rt http.RoundTripper
	rt = http.DefaultTransport
	if cfg.Transport != nil {
		rt = cfg.Transport
	}
	klog.Infof("nw ......")
	return HTTPWrappersForConfig(cfg, rt)
}
