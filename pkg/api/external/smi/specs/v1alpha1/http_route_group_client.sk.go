// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type HTTPRouteGroupWatcher interface {
	// watch namespace-scoped httproutegroups
	Watch(namespace string, opts clients.WatchOpts) (<-chan HTTPRouteGroupList, <-chan error, error)
}

type HTTPRouteGroupClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*HTTPRouteGroup, error)
	Write(resource *HTTPRouteGroup, opts clients.WriteOpts) (*HTTPRouteGroup, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (HTTPRouteGroupList, error)
	HTTPRouteGroupWatcher
}

type hTTPRouteGroupClient struct {
	rc clients.ResourceClient
}

func NewHTTPRouteGroupClient(rcFactory factory.ResourceClientFactory) (HTTPRouteGroupClient, error) {
	return NewHTTPRouteGroupClientWithToken(rcFactory, "")
}

func NewHTTPRouteGroupClientWithToken(rcFactory factory.ResourceClientFactory, token string) (HTTPRouteGroupClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &HTTPRouteGroup{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base HTTPRouteGroup resource client")
	}
	return NewHTTPRouteGroupClientWithBase(rc), nil
}

func NewHTTPRouteGroupClientWithBase(rc clients.ResourceClient) HTTPRouteGroupClient {
	return &hTTPRouteGroupClient{
		rc: rc,
	}
}

func (client *hTTPRouteGroupClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *hTTPRouteGroupClient) Register() error {
	return client.rc.Register()
}

func (client *hTTPRouteGroupClient) Read(namespace, name string, opts clients.ReadOpts) (*HTTPRouteGroup, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*HTTPRouteGroup), nil
}

func (client *hTTPRouteGroupClient) Write(hTTPRouteGroup *HTTPRouteGroup, opts clients.WriteOpts) (*HTTPRouteGroup, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(hTTPRouteGroup, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*HTTPRouteGroup), nil
}

func (client *hTTPRouteGroupClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *hTTPRouteGroupClient) List(namespace string, opts clients.ListOpts) (HTTPRouteGroupList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToHTTPRouteGroup(resourceList), nil
}

func (client *hTTPRouteGroupClient) Watch(namespace string, opts clients.WatchOpts) (<-chan HTTPRouteGroupList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	httproutegroupsChan := make(chan HTTPRouteGroupList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				httproutegroupsChan <- convertToHTTPRouteGroup(resourceList)
			case <-opts.Ctx.Done():
				close(httproutegroupsChan)
				return
			}
		}
	}()
	return httproutegroupsChan, errs, nil
}

func convertToHTTPRouteGroup(resources resources.ResourceList) HTTPRouteGroupList {
	var hTTPRouteGroupList HTTPRouteGroupList
	for _, resource := range resources {
		hTTPRouteGroupList = append(hTTPRouteGroupList, resource.(*HTTPRouteGroup))
	}
	return hTTPRouteGroupList
}
