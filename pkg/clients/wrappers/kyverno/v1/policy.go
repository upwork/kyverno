package v1

import (
	"context"

	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	v1 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v1"
	"github.com/kyverno/kyverno/pkg/clients/wrappers/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

type policies struct {
	inner             v1.PolicyInterface
	clientQueryMetric utils.ClientQueryMetric
	ns                string
}

func wrapPolicies(c v1.PolicyInterface, m utils.ClientQueryMetric, namespace string) v1.PolicyInterface {
	return &policies{
		inner:             c,
		clientQueryMetric: m,
		ns:                namespace,
	}
}

func (c *policies) Create(ctx context.Context, o *kyvernov1.Policy, opts metav1.CreateOptions) (*kyvernov1.Policy, error) {
	return utils.Create(ctx, c.clientQueryMetric, "Policy", c.ns, o, opts, c.inner.Create)
}

func (c *policies) Update(ctx context.Context, o *kyvernov1.Policy, opts metav1.UpdateOptions) (*kyvernov1.Policy, error) {
	return utils.Update(ctx, c.clientQueryMetric, "Policy", c.ns, o, opts, c.inner.Update)
}

func (c *policies) UpdateStatus(ctx context.Context, o *kyvernov1.Policy, opts metav1.UpdateOptions) (*kyvernov1.Policy, error) {
	return utils.UpdateStatus(ctx, c.clientQueryMetric, "Policy", c.ns, o, opts, c.inner.UpdateStatus)
}

func (c *policies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return utils.Delete(ctx, c.clientQueryMetric, "Policy", c.ns, name, opts, c.inner.Delete)
}

func (c *policies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return utils.DeleteCollection(ctx, c.clientQueryMetric, "Policy", c.ns, opts, listOpts, c.inner.DeleteCollection)
}

func (c *policies) Get(ctx context.Context, name string, opts metav1.GetOptions) (*kyvernov1.Policy, error) {
	return utils.Get(ctx, c.clientQueryMetric, "Policy", c.ns, name, opts, c.inner.Get)
}

func (c *policies) List(ctx context.Context, opts metav1.ListOptions) (*kyvernov1.PolicyList, error) {
	return utils.List(ctx, c.clientQueryMetric, "Policy", c.ns, opts, c.inner.List)
}

func (c *policies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return utils.Watch(ctx, c.clientQueryMetric, "Policy", c.ns, opts, c.inner.Watch)
}

func (c *policies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kyvernov1.Policy, error) {
	return utils.Patch(ctx, c.clientQueryMetric, "Policy", c.ns, name, pt, data, opts, c.inner.Patch, subresources...)
}