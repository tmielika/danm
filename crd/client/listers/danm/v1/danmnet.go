/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/nokia/danm/crd/apis/danm/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DanmNetLister helps list DanmNets.
type DanmNetLister interface {
	// List lists all DanmNets in the indexer.
	List(selector labels.Selector) (ret []*v1.DanmNet, err error)
	// DanmNets returns an object that can list and get DanmNets.
	DanmNets(namespace string) DanmNetNamespaceLister
	DanmNetListerExpansion
}

// danmNetLister implements the DanmNetLister interface.
type danmNetLister struct {
	indexer cache.Indexer
}

// NewDanmNetLister returns a new DanmNetLister.
func NewDanmNetLister(indexer cache.Indexer) DanmNetLister {
	return &danmNetLister{indexer: indexer}
}

// List lists all DanmNets in the indexer.
func (s *danmNetLister) List(selector labels.Selector) (ret []*v1.DanmNet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DanmNet))
	})
	return ret, err
}

// DanmNets returns an object that can list and get DanmNets.
func (s *danmNetLister) DanmNets(namespace string) DanmNetNamespaceLister {
	return danmNetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DanmNetNamespaceLister helps list and get DanmNets.
type DanmNetNamespaceLister interface {
	// List lists all DanmNets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.DanmNet, err error)
	// Get retrieves the DanmNet from the indexer for a given namespace and name.
	Get(name string) (*v1.DanmNet, error)
	DanmNetNamespaceListerExpansion
}

// danmNetNamespaceLister implements the DanmNetNamespaceLister
// interface.
type danmNetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DanmNets in the indexer for a given namespace.
func (s danmNetNamespaceLister) List(selector labels.Selector) (ret []*v1.DanmNet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DanmNet))
	})
	return ret, err
}

// Get retrieves the DanmNet from the indexer for a given namespace and name.
func (s danmNetNamespaceLister) Get(name string) (*v1.DanmNet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("danmnet"), name)
	}
	return obj.(*v1.DanmNet), nil
}
