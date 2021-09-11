// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualServerRouteLister helps list VirtualServerRoutes.
// All objects returned here must be treated as read-only.
type VirtualServerRouteLister interface {
	// List lists all VirtualServerRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualServerRoute, err error)
	// VirtualServerRoutes returns an object that can list and get VirtualServerRoutes.
	VirtualServerRoutes(namespace string) VirtualServerRouteNamespaceLister
	VirtualServerRouteListerExpansion
}

// virtualServerRouteLister implements the VirtualServerRouteLister interface.
type virtualServerRouteLister struct {
	indexer cache.Indexer
}

// NewVirtualServerRouteLister returns a new VirtualServerRouteLister.
func NewVirtualServerRouteLister(indexer cache.Indexer) VirtualServerRouteLister {
	return &virtualServerRouteLister{indexer: indexer}
}

// List lists all VirtualServerRoutes in the indexer.
func (s *virtualServerRouteLister) List(selector labels.Selector) (ret []*v1.VirtualServerRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualServerRoute))
	})
	return ret, err
}

// VirtualServerRoutes returns an object that can list and get VirtualServerRoutes.
func (s *virtualServerRouteLister) VirtualServerRoutes(namespace string) VirtualServerRouteNamespaceLister {
	return virtualServerRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualServerRouteNamespaceLister helps list and get VirtualServerRoutes.
// All objects returned here must be treated as read-only.
type VirtualServerRouteNamespaceLister interface {
	// List lists all VirtualServerRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualServerRoute, err error)
	// Get retrieves the VirtualServerRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.VirtualServerRoute, error)
	VirtualServerRouteNamespaceListerExpansion
}

// virtualServerRouteNamespaceLister implements the VirtualServerRouteNamespaceLister
// interface.
type virtualServerRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualServerRoutes in the indexer for a given namespace.
func (s virtualServerRouteNamespaceLister) List(selector labels.Selector) (ret []*v1.VirtualServerRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualServerRoute))
	})
	return ret, err
}

// Get retrieves the VirtualServerRoute from the indexer for a given namespace and name.
func (s virtualServerRouteNamespaceLister) Get(name string) (*v1.VirtualServerRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("virtualserverroute"), name)
	}
	return obj.(*v1.VirtualServerRoute), nil
}