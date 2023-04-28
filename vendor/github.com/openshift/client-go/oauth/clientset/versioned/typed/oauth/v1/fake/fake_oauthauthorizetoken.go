// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	oauthv1 "github.com/openshift/api/oauth/v1"
	applyconfigurationsoauthv1 "github.com/openshift/client-go/oauth/applyconfigurations/oauth/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOAuthAuthorizeTokens implements OAuthAuthorizeTokenInterface
type FakeOAuthAuthorizeTokens struct {
	Fake *FakeOauthV1
}

var oauthauthorizetokensResource = schema.GroupVersionResource{Group: "oauth.openshift.io", Version: "v1", Resource: "oauthauthorizetokens"}

var oauthauthorizetokensKind = schema.GroupVersionKind{Group: "oauth.openshift.io", Version: "v1", Kind: "OAuthAuthorizeToken"}

// Get takes name of the oAuthAuthorizeToken, and returns the corresponding oAuthAuthorizeToken object, and an error if there is any.
func (c *FakeOAuthAuthorizeTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *oauthv1.OAuthAuthorizeToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(oauthauthorizetokensResource, name), &oauthv1.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAuthorizeToken), err
}

// List takes label and field selectors, and returns the list of OAuthAuthorizeTokens that match those selectors.
func (c *FakeOAuthAuthorizeTokens) List(ctx context.Context, opts v1.ListOptions) (result *oauthv1.OAuthAuthorizeTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(oauthauthorizetokensResource, oauthauthorizetokensKind, opts), &oauthv1.OAuthAuthorizeTokenList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &oauthv1.OAuthAuthorizeTokenList{ListMeta: obj.(*oauthv1.OAuthAuthorizeTokenList).ListMeta}
	for _, item := range obj.(*oauthv1.OAuthAuthorizeTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oAuthAuthorizeTokens.
func (c *FakeOAuthAuthorizeTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(oauthauthorizetokensResource, opts))
}

// Create takes the representation of a oAuthAuthorizeToken and creates it.  Returns the server's representation of the oAuthAuthorizeToken, and an error, if there is any.
func (c *FakeOAuthAuthorizeTokens) Create(ctx context.Context, oAuthAuthorizeToken *oauthv1.OAuthAuthorizeToken, opts v1.CreateOptions) (result *oauthv1.OAuthAuthorizeToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(oauthauthorizetokensResource, oAuthAuthorizeToken), &oauthv1.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAuthorizeToken), err
}

// Update takes the representation of a oAuthAuthorizeToken and updates it. Returns the server's representation of the oAuthAuthorizeToken, and an error, if there is any.
func (c *FakeOAuthAuthorizeTokens) Update(ctx context.Context, oAuthAuthorizeToken *oauthv1.OAuthAuthorizeToken, opts v1.UpdateOptions) (result *oauthv1.OAuthAuthorizeToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(oauthauthorizetokensResource, oAuthAuthorizeToken), &oauthv1.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAuthorizeToken), err
}

// Delete takes name of the oAuthAuthorizeToken and deletes it. Returns an error if one occurs.
func (c *FakeOAuthAuthorizeTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(oauthauthorizetokensResource, name, opts), &oauthv1.OAuthAuthorizeToken{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOAuthAuthorizeTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(oauthauthorizetokensResource, listOpts)

	_, err := c.Fake.Invokes(action, &oauthv1.OAuthAuthorizeTokenList{})
	return err
}

// Patch applies the patch and returns the patched oAuthAuthorizeToken.
func (c *FakeOAuthAuthorizeTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *oauthv1.OAuthAuthorizeToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(oauthauthorizetokensResource, name, pt, data, subresources...), &oauthv1.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAuthorizeToken), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied oAuthAuthorizeToken.
func (c *FakeOAuthAuthorizeTokens) Apply(ctx context.Context, oAuthAuthorizeToken *applyconfigurationsoauthv1.OAuthAuthorizeTokenApplyConfiguration, opts v1.ApplyOptions) (result *oauthv1.OAuthAuthorizeToken, err error) {
	if oAuthAuthorizeToken == nil {
		return nil, fmt.Errorf("oAuthAuthorizeToken provided to Apply must not be nil")
	}
	data, err := json.Marshal(oAuthAuthorizeToken)
	if err != nil {
		return nil, err
	}
	name := oAuthAuthorizeToken.Name
	if name == nil {
		return nil, fmt.Errorf("oAuthAuthorizeToken.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(oauthauthorizetokensResource, *name, types.ApplyPatchType, data), &oauthv1.OAuthAuthorizeToken{})
	if obj == nil {
		return nil, err
	}
	return obj.(*oauthv1.OAuthAuthorizeToken), err
}
