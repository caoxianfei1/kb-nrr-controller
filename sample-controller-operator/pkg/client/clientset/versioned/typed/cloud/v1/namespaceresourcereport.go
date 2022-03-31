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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	v1 "nrr-controller/pkg/apis/cloud/v1"
	scheme "nrr-controller/pkg/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NamespaceResourceReportsGetter has a method to return a NamespaceResourceReportInterface.
// A group's client should implement this interface.
type NamespaceResourceReportsGetter interface {
	NamespaceResourceReports(namespace string) NamespaceResourceReportInterface
}

// NamespaceResourceReportInterface has methods to work with NamespaceResourceReport resources.
type NamespaceResourceReportInterface interface {
	Create(ctx context.Context, namespaceResourceReport *v1.NamespaceResourceReport, opts metav1.CreateOptions) (*v1.NamespaceResourceReport, error)
	Update(ctx context.Context, namespaceResourceReport *v1.NamespaceResourceReport, opts metav1.UpdateOptions) (*v1.NamespaceResourceReport, error)
	UpdateStatus(ctx context.Context, namespaceResourceReport *v1.NamespaceResourceReport, opts metav1.UpdateOptions) (*v1.NamespaceResourceReport, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NamespaceResourceReport, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NamespaceResourceReportList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NamespaceResourceReport, err error)
	NamespaceResourceReportExpansion
}

// namespaceResourceReports implements NamespaceResourceReportInterface
type namespaceResourceReports struct {
	client rest.Interface
	ns     string
}

// newNamespaceResourceReports returns a NamespaceResourceReports
func newNamespaceResourceReports(c *CloudV1Client, namespace string) *namespaceResourceReports {
	return &namespaceResourceReports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the namespaceResourceReport, and returns the corresponding namespaceResourceReport object, and an error if there is any.
func (c *namespaceResourceReports) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NamespaceResourceReport, err error) {
	result = &v1.NamespaceResourceReport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NamespaceResourceReports that match those selectors.
func (c *namespaceResourceReports) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NamespaceResourceReportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NamespaceResourceReportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested namespaceResourceReports.
func (c *namespaceResourceReports) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a namespaceResourceReport and creates it.  Returns the server's representation of the namespaceResourceReport, and an error, if there is any.
func (c *namespaceResourceReports) Create(ctx context.Context, namespaceResourceReport *v1.NamespaceResourceReport, opts metav1.CreateOptions) (result *v1.NamespaceResourceReport, err error) {
	result = &v1.NamespaceResourceReport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespaceResourceReport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a namespaceResourceReport and updates it. Returns the server's representation of the namespaceResourceReport, and an error, if there is any.
func (c *namespaceResourceReports) Update(ctx context.Context, namespaceResourceReport *v1.NamespaceResourceReport, opts metav1.UpdateOptions) (result *v1.NamespaceResourceReport, err error) {
	result = &v1.NamespaceResourceReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		Name(namespaceResourceReport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespaceResourceReport).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *namespaceResourceReports) UpdateStatus(ctx context.Context, namespaceResourceReport *v1.NamespaceResourceReport, opts metav1.UpdateOptions) (result *v1.NamespaceResourceReport, err error) {
	result = &v1.NamespaceResourceReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		Name(namespaceResourceReport.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(namespaceResourceReport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the namespaceResourceReport and deletes it. Returns an error if one occurs.
func (c *namespaceResourceReports) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *namespaceResourceReports) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched namespaceResourceReport.
func (c *namespaceResourceReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NamespaceResourceReport, err error) {
	result = &v1.NamespaceResourceReport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("namespaceresourcereports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
