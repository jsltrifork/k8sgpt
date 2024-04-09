package kubernetes

import (
	openapi_v2 "github.com/google/gnostic/openapiv2"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	ctrl "sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	MetricsClient versioned.Interface
	Client        kubernetes.Interface
	CtrlClient    ctrl.Client
	Config        *rest.Config
	ServerVersion *version.Info
}

type K8sApiReference struct {
	ApiVersion    schema.GroupVersion
	Kind          string
	OpenapiSchema *openapi_v2.Document
}
