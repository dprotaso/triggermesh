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
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSEventBridgeTargetLister helps list AWSEventBridgeTargets.
// All objects returned here must be treated as read-only.
type AWSEventBridgeTargetLister interface {
	// List lists all AWSEventBridgeTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AWSEventBridgeTarget, err error)
	// AWSEventBridgeTargets returns an object that can list and get AWSEventBridgeTargets.
	AWSEventBridgeTargets(namespace string) AWSEventBridgeTargetNamespaceLister
	AWSEventBridgeTargetListerExpansion
}

// aWSEventBridgeTargetLister implements the AWSEventBridgeTargetLister interface.
type aWSEventBridgeTargetLister struct {
	indexer cache.Indexer
}

// NewAWSEventBridgeTargetLister returns a new AWSEventBridgeTargetLister.
func NewAWSEventBridgeTargetLister(indexer cache.Indexer) AWSEventBridgeTargetLister {
	return &aWSEventBridgeTargetLister{indexer: indexer}
}

// List lists all AWSEventBridgeTargets in the indexer.
func (s *aWSEventBridgeTargetLister) List(selector labels.Selector) (ret []*v1alpha1.AWSEventBridgeTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSEventBridgeTarget))
	})
	return ret, err
}

// AWSEventBridgeTargets returns an object that can list and get AWSEventBridgeTargets.
func (s *aWSEventBridgeTargetLister) AWSEventBridgeTargets(namespace string) AWSEventBridgeTargetNamespaceLister {
	return aWSEventBridgeTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AWSEventBridgeTargetNamespaceLister helps list and get AWSEventBridgeTargets.
// All objects returned here must be treated as read-only.
type AWSEventBridgeTargetNamespaceLister interface {
	// List lists all AWSEventBridgeTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AWSEventBridgeTarget, err error)
	// Get retrieves the AWSEventBridgeTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AWSEventBridgeTarget, error)
	AWSEventBridgeTargetNamespaceListerExpansion
}

// aWSEventBridgeTargetNamespaceLister implements the AWSEventBridgeTargetNamespaceLister
// interface.
type aWSEventBridgeTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AWSEventBridgeTargets in the indexer for a given namespace.
func (s aWSEventBridgeTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AWSEventBridgeTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSEventBridgeTarget))
	})
	return ret, err
}

// Get retrieves the AWSEventBridgeTarget from the indexer for a given namespace and name.
func (s aWSEventBridgeTargetNamespaceLister) Get(name string) (*v1alpha1.AWSEventBridgeTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awseventbridgetarget"), name)
	}
	return obj.(*v1alpha1.AWSEventBridgeTarget), nil
}
