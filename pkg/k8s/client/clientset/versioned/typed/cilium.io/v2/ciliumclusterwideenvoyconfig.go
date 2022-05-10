// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	"context"
	"time"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumClusterwideEnvoyConfigsGetter has a method to return a CiliumClusterwideEnvoyConfigInterface.
// A group's client should implement this interface.
type CiliumClusterwideEnvoyConfigsGetter interface {
	CiliumClusterwideEnvoyConfigs() CiliumClusterwideEnvoyConfigInterface
}

// CiliumClusterwideEnvoyConfigInterface has methods to work with CiliumClusterwideEnvoyConfig resources.
type CiliumClusterwideEnvoyConfigInterface interface {
	Create(ctx context.Context, ciliumClusterwideEnvoyConfig *v2.CiliumClusterwideEnvoyConfig, opts v1.CreateOptions) (*v2.CiliumClusterwideEnvoyConfig, error)
	Update(ctx context.Context, ciliumClusterwideEnvoyConfig *v2.CiliumClusterwideEnvoyConfig, opts v1.UpdateOptions) (*v2.CiliumClusterwideEnvoyConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.CiliumClusterwideEnvoyConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.CiliumClusterwideEnvoyConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumClusterwideEnvoyConfig, err error)
	CiliumClusterwideEnvoyConfigExpansion
}

// ciliumClusterwideEnvoyConfigs implements CiliumClusterwideEnvoyConfigInterface
type ciliumClusterwideEnvoyConfigs struct {
	client rest.Interface
}

// newCiliumClusterwideEnvoyConfigs returns a CiliumClusterwideEnvoyConfigs
func newCiliumClusterwideEnvoyConfigs(c *CiliumV2Client) *ciliumClusterwideEnvoyConfigs {
	return &ciliumClusterwideEnvoyConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumClusterwideEnvoyConfig, and returns the corresponding ciliumClusterwideEnvoyConfig object, and an error if there is any.
func (c *ciliumClusterwideEnvoyConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.CiliumClusterwideEnvoyConfig, err error) {
	result = &v2.CiliumClusterwideEnvoyConfig{}
	err = c.client.Get().
		Resource("ciliumclusterwideenvoyconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumClusterwideEnvoyConfigs that match those selectors.
func (c *ciliumClusterwideEnvoyConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v2.CiliumClusterwideEnvoyConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.CiliumClusterwideEnvoyConfigList{}
	err = c.client.Get().
		Resource("ciliumclusterwideenvoyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumClusterwideEnvoyConfigs.
func (c *ciliumClusterwideEnvoyConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumclusterwideenvoyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ciliumClusterwideEnvoyConfig and creates it.  Returns the server's representation of the ciliumClusterwideEnvoyConfig, and an error, if there is any.
func (c *ciliumClusterwideEnvoyConfigs) Create(ctx context.Context, ciliumClusterwideEnvoyConfig *v2.CiliumClusterwideEnvoyConfig, opts v1.CreateOptions) (result *v2.CiliumClusterwideEnvoyConfig, err error) {
	result = &v2.CiliumClusterwideEnvoyConfig{}
	err = c.client.Post().
		Resource("ciliumclusterwideenvoyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumClusterwideEnvoyConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ciliumClusterwideEnvoyConfig and updates it. Returns the server's representation of the ciliumClusterwideEnvoyConfig, and an error, if there is any.
func (c *ciliumClusterwideEnvoyConfigs) Update(ctx context.Context, ciliumClusterwideEnvoyConfig *v2.CiliumClusterwideEnvoyConfig, opts v1.UpdateOptions) (result *v2.CiliumClusterwideEnvoyConfig, err error) {
	result = &v2.CiliumClusterwideEnvoyConfig{}
	err = c.client.Put().
		Resource("ciliumclusterwideenvoyconfigs").
		Name(ciliumClusterwideEnvoyConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ciliumClusterwideEnvoyConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ciliumClusterwideEnvoyConfig and deletes it. Returns an error if one occurs.
func (c *ciliumClusterwideEnvoyConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumclusterwideenvoyconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumClusterwideEnvoyConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumclusterwideenvoyconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ciliumClusterwideEnvoyConfig.
func (c *ciliumClusterwideEnvoyConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.CiliumClusterwideEnvoyConfig, err error) {
	result = &v2.CiliumClusterwideEnvoyConfig{}
	err = c.client.Patch(pt).
		Resource("ciliumclusterwideenvoyconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}