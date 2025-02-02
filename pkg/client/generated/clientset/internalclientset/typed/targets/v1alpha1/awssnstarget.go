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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	scheme "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSSNSTargetsGetter has a method to return a AWSSNSTargetInterface.
// A group's client should implement this interface.
type AWSSNSTargetsGetter interface {
	AWSSNSTargets(namespace string) AWSSNSTargetInterface
}

// AWSSNSTargetInterface has methods to work with AWSSNSTarget resources.
type AWSSNSTargetInterface interface {
	Create(ctx context.Context, aWSSNSTarget *v1alpha1.AWSSNSTarget, opts v1.CreateOptions) (*v1alpha1.AWSSNSTarget, error)
	Update(ctx context.Context, aWSSNSTarget *v1alpha1.AWSSNSTarget, opts v1.UpdateOptions) (*v1alpha1.AWSSNSTarget, error)
	UpdateStatus(ctx context.Context, aWSSNSTarget *v1alpha1.AWSSNSTarget, opts v1.UpdateOptions) (*v1alpha1.AWSSNSTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AWSSNSTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AWSSNSTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSSNSTarget, err error)
	AWSSNSTargetExpansion
}

// aWSSNSTargets implements AWSSNSTargetInterface
type aWSSNSTargets struct {
	client rest.Interface
	ns     string
}

// newAWSSNSTargets returns a AWSSNSTargets
func newAWSSNSTargets(c *TargetsV1alpha1Client, namespace string) *aWSSNSTargets {
	return &aWSSNSTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSSNSTarget, and returns the corresponding aWSSNSTarget object, and an error if there is any.
func (c *aWSSNSTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSSNSTarget, err error) {
	result = &v1alpha1.AWSSNSTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssnstargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSSNSTargets that match those selectors.
func (c *aWSSNSTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSSNSTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AWSSNSTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssnstargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSSNSTargets.
func (c *aWSSNSTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssnstargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aWSSNSTarget and creates it.  Returns the server's representation of the aWSSNSTarget, and an error, if there is any.
func (c *aWSSNSTargets) Create(ctx context.Context, aWSSNSTarget *v1alpha1.AWSSNSTarget, opts v1.CreateOptions) (result *v1alpha1.AWSSNSTarget, err error) {
	result = &v1alpha1.AWSSNSTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssnstargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSSNSTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aWSSNSTarget and updates it. Returns the server's representation of the aWSSNSTarget, and an error, if there is any.
func (c *aWSSNSTargets) Update(ctx context.Context, aWSSNSTarget *v1alpha1.AWSSNSTarget, opts v1.UpdateOptions) (result *v1alpha1.AWSSNSTarget, err error) {
	result = &v1alpha1.AWSSNSTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssnstargets").
		Name(aWSSNSTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSSNSTarget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aWSSNSTargets) UpdateStatus(ctx context.Context, aWSSNSTarget *v1alpha1.AWSSNSTarget, opts v1.UpdateOptions) (result *v1alpha1.AWSSNSTarget, err error) {
	result = &v1alpha1.AWSSNSTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssnstargets").
		Name(aWSSNSTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSSNSTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aWSSNSTarget and deletes it. Returns an error if one occurs.
func (c *aWSSNSTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssnstargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSSNSTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssnstargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aWSSNSTarget.
func (c *aWSSNSTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSSNSTarget, err error) {
	result = &v1alpha1.AWSSNSTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssnstargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
