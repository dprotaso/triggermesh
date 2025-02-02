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

// AWSCognitoUserPoolSourceLister helps list AWSCognitoUserPoolSources.
// All objects returned here must be treated as read-only.
type AWSCognitoUserPoolSourceLister interface {
	// List lists all AWSCognitoUserPoolSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AWSCognitoUserPoolSource, err error)
	// AWSCognitoUserPoolSources returns an object that can list and get AWSCognitoUserPoolSources.
	AWSCognitoUserPoolSources(namespace string) AWSCognitoUserPoolSourceNamespaceLister
	AWSCognitoUserPoolSourceListerExpansion
}

// aWSCognitoUserPoolSourceLister implements the AWSCognitoUserPoolSourceLister interface.
type aWSCognitoUserPoolSourceLister struct {
	indexer cache.Indexer
}

// NewAWSCognitoUserPoolSourceLister returns a new AWSCognitoUserPoolSourceLister.
func NewAWSCognitoUserPoolSourceLister(indexer cache.Indexer) AWSCognitoUserPoolSourceLister {
	return &aWSCognitoUserPoolSourceLister{indexer: indexer}
}

// List lists all AWSCognitoUserPoolSources in the indexer.
func (s *aWSCognitoUserPoolSourceLister) List(selector labels.Selector) (ret []*v1alpha1.AWSCognitoUserPoolSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSCognitoUserPoolSource))
	})
	return ret, err
}

// AWSCognitoUserPoolSources returns an object that can list and get AWSCognitoUserPoolSources.
func (s *aWSCognitoUserPoolSourceLister) AWSCognitoUserPoolSources(namespace string) AWSCognitoUserPoolSourceNamespaceLister {
	return aWSCognitoUserPoolSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AWSCognitoUserPoolSourceNamespaceLister helps list and get AWSCognitoUserPoolSources.
// All objects returned here must be treated as read-only.
type AWSCognitoUserPoolSourceNamespaceLister interface {
	// List lists all AWSCognitoUserPoolSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AWSCognitoUserPoolSource, err error)
	// Get retrieves the AWSCognitoUserPoolSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AWSCognitoUserPoolSource, error)
	AWSCognitoUserPoolSourceNamespaceListerExpansion
}

// aWSCognitoUserPoolSourceNamespaceLister implements the AWSCognitoUserPoolSourceNamespaceLister
// interface.
type aWSCognitoUserPoolSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AWSCognitoUserPoolSources in the indexer for a given namespace.
func (s aWSCognitoUserPoolSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AWSCognitoUserPoolSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSCognitoUserPoolSource))
	})
	return ret, err
}

// Get retrieves the AWSCognitoUserPoolSource from the indexer for a given namespace and name.
func (s aWSCognitoUserPoolSourceNamespaceLister) Get(name string) (*v1alpha1.AWSCognitoUserPoolSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awscognitouserpoolsource"), name)
	}
	return obj.(*v1alpha1.AWSCognitoUserPoolSource), nil
}
