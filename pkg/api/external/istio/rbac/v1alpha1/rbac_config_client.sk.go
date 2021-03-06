// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type RbacConfigWatcher interface {
	// watch namespace-scoped Rbacconfigs
	Watch(namespace string, opts clients.WatchOpts) (<-chan RbacConfigList, <-chan error, error)
}

type RbacConfigClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*RbacConfig, error)
	Write(resource *RbacConfig, opts clients.WriteOpts) (*RbacConfig, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (RbacConfigList, error)
	RbacConfigWatcher
}

type rbacConfigClient struct {
	rc clients.ResourceClient
}

func NewRbacConfigClient(rcFactory factory.ResourceClientFactory) (RbacConfigClient, error) {
	return NewRbacConfigClientWithToken(rcFactory, "")
}

func NewRbacConfigClientWithToken(rcFactory factory.ResourceClientFactory, token string) (RbacConfigClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &RbacConfig{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base RbacConfig resource client")
	}
	return NewRbacConfigClientWithBase(rc), nil
}

func NewRbacConfigClientWithBase(rc clients.ResourceClient) RbacConfigClient {
	return &rbacConfigClient{
		rc: rc,
	}
}

func (client *rbacConfigClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *rbacConfigClient) Register() error {
	return client.rc.Register()
}

func (client *rbacConfigClient) Read(namespace, name string, opts clients.ReadOpts) (*RbacConfig, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*RbacConfig), nil
}

func (client *rbacConfigClient) Write(rbacConfig *RbacConfig, opts clients.WriteOpts) (*RbacConfig, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(rbacConfig, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*RbacConfig), nil
}

func (client *rbacConfigClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *rbacConfigClient) List(namespace string, opts clients.ListOpts) (RbacConfigList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToRbacConfig(resourceList), nil
}

func (client *rbacConfigClient) Watch(namespace string, opts clients.WatchOpts) (<-chan RbacConfigList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	rbacconfigsChan := make(chan RbacConfigList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				rbacconfigsChan <- convertToRbacConfig(resourceList)
			case <-opts.Ctx.Done():
				close(rbacconfigsChan)
				return
			}
		}
	}()
	return rbacconfigsChan, errs, nil
}

func convertToRbacConfig(resources resources.ResourceList) RbacConfigList {
	var rbacConfigList RbacConfigList
	for _, resource := range resources {
		rbacConfigList = append(rbacConfigList, resource.(*RbacConfig))
	}
	return rbacConfigList
}
