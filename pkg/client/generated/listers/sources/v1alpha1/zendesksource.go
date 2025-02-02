/*
Copyright 2021 TriggerMesh Inc.

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

package v1alpha1

import (
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ZendeskSourceLister helps list ZendeskSources.
// All objects returned here must be treated as read-only.
type ZendeskSourceLister interface {
	// List lists all ZendeskSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ZendeskSource, err error)
	// ZendeskSources returns an object that can list and get ZendeskSources.
	ZendeskSources(namespace string) ZendeskSourceNamespaceLister
	ZendeskSourceListerExpansion
}

// zendeskSourceLister implements the ZendeskSourceLister interface.
type zendeskSourceLister struct {
	indexer cache.Indexer
}

// NewZendeskSourceLister returns a new ZendeskSourceLister.
func NewZendeskSourceLister(indexer cache.Indexer) ZendeskSourceLister {
	return &zendeskSourceLister{indexer: indexer}
}

// List lists all ZendeskSources in the indexer.
func (s *zendeskSourceLister) List(selector labels.Selector) (ret []*v1alpha1.ZendeskSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ZendeskSource))
	})
	return ret, err
}

// ZendeskSources returns an object that can list and get ZendeskSources.
func (s *zendeskSourceLister) ZendeskSources(namespace string) ZendeskSourceNamespaceLister {
	return zendeskSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ZendeskSourceNamespaceLister helps list and get ZendeskSources.
// All objects returned here must be treated as read-only.
type ZendeskSourceNamespaceLister interface {
	// List lists all ZendeskSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ZendeskSource, err error)
	// Get retrieves the ZendeskSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ZendeskSource, error)
	ZendeskSourceNamespaceListerExpansion
}

// zendeskSourceNamespaceLister implements the ZendeskSourceNamespaceLister
// interface.
type zendeskSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ZendeskSources in the indexer for a given namespace.
func (s zendeskSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ZendeskSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ZendeskSource))
	})
	return ret, err
}

// Get retrieves the ZendeskSource from the indexer for a given namespace and name.
func (s zendeskSourceNamespaceLister) Get(name string) (*v1alpha1.ZendeskSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("zendesksource"), name)
	}
	return obj.(*v1alpha1.ZendeskSource), nil
}
