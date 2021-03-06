// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/xmudrii/etcdproxy-controller/pkg/apis/etcd/v1alpha1"
	scheme "github.com/xmudrii/etcdproxy-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EtcdStoragesGetter has a method to return a EtcdStorageInterface.
// A group's client should implement this interface.
type EtcdStoragesGetter interface {
	EtcdStorages() EtcdStorageInterface
}

// EtcdStorageInterface has methods to work with EtcdStorage resources.
type EtcdStorageInterface interface {
	Create(*v1alpha1.EtcdStorage) (*v1alpha1.EtcdStorage, error)
	Update(*v1alpha1.EtcdStorage) (*v1alpha1.EtcdStorage, error)
	UpdateStatus(*v1alpha1.EtcdStorage) (*v1alpha1.EtcdStorage, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EtcdStorage, error)
	List(opts v1.ListOptions) (*v1alpha1.EtcdStorageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EtcdStorage, err error)
	EtcdStorageExpansion
}

// etcdStorages implements EtcdStorageInterface
type etcdStorages struct {
	client rest.Interface
}

// newEtcdStorages returns a EtcdStorages
func newEtcdStorages(c *EtcdV1alpha1Client) *etcdStorages {
	return &etcdStorages{
		client: c.RESTClient(),
	}
}

// Get takes name of the etcdStorage, and returns the corresponding etcdStorage object, and an error if there is any.
func (c *etcdStorages) Get(name string, options v1.GetOptions) (result *v1alpha1.EtcdStorage, err error) {
	result = &v1alpha1.EtcdStorage{}
	err = c.client.Get().
		Resource("etcdstorages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EtcdStorages that match those selectors.
func (c *etcdStorages) List(opts v1.ListOptions) (result *v1alpha1.EtcdStorageList, err error) {
	result = &v1alpha1.EtcdStorageList{}
	err = c.client.Get().
		Resource("etcdstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested etcdStorages.
func (c *etcdStorages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("etcdstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a etcdStorage and creates it.  Returns the server's representation of the etcdStorage, and an error, if there is any.
func (c *etcdStorages) Create(etcdStorage *v1alpha1.EtcdStorage) (result *v1alpha1.EtcdStorage, err error) {
	result = &v1alpha1.EtcdStorage{}
	err = c.client.Post().
		Resource("etcdstorages").
		Body(etcdStorage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a etcdStorage and updates it. Returns the server's representation of the etcdStorage, and an error, if there is any.
func (c *etcdStorages) Update(etcdStorage *v1alpha1.EtcdStorage) (result *v1alpha1.EtcdStorage, err error) {
	result = &v1alpha1.EtcdStorage{}
	err = c.client.Put().
		Resource("etcdstorages").
		Name(etcdStorage.Name).
		Body(etcdStorage).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *etcdStorages) UpdateStatus(etcdStorage *v1alpha1.EtcdStorage) (result *v1alpha1.EtcdStorage, err error) {
	result = &v1alpha1.EtcdStorage{}
	err = c.client.Put().
		Resource("etcdstorages").
		Name(etcdStorage.Name).
		SubResource("status").
		Body(etcdStorage).
		Do().
		Into(result)
	return
}

// Delete takes name of the etcdStorage and deletes it. Returns an error if one occurs.
func (c *etcdStorages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("etcdstorages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *etcdStorages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("etcdstorages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched etcdStorage.
func (c *etcdStorages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EtcdStorage, err error) {
	result = &v1alpha1.EtcdStorage{}
	err = c.client.Patch(pt).
		Resource("etcdstorages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
