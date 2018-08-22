// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReportGenerationQueries implements ReportGenerationQueryInterface
type FakeReportGenerationQueries struct {
	Fake *FakeMeteringV1alpha1
	ns   string
}

var reportgenerationqueriesResource = schema.GroupVersionResource{Group: "metering.openshift.io", Version: "v1alpha1", Resource: "reportgenerationqueries"}

var reportgenerationqueriesKind = schema.GroupVersionKind{Group: "metering.openshift.io", Version: "v1alpha1", Kind: "ReportGenerationQuery"}

// Get takes name of the reportGenerationQuery, and returns the corresponding reportGenerationQuery object, and an error if there is any.
func (c *FakeReportGenerationQueries) Get(name string, options v1.GetOptions) (result *v1alpha1.ReportGenerationQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reportgenerationqueriesResource, c.ns, name), &v1alpha1.ReportGenerationQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReportGenerationQuery), err
}

// List takes label and field selectors, and returns the list of ReportGenerationQueries that match those selectors.
func (c *FakeReportGenerationQueries) List(opts v1.ListOptions) (result *v1alpha1.ReportGenerationQueryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reportgenerationqueriesResource, reportgenerationqueriesKind, c.ns, opts), &v1alpha1.ReportGenerationQueryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReportGenerationQueryList{}
	for _, item := range obj.(*v1alpha1.ReportGenerationQueryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reportGenerationQueries.
func (c *FakeReportGenerationQueries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reportgenerationqueriesResource, c.ns, opts))

}

// Create takes the representation of a reportGenerationQuery and creates it.  Returns the server's representation of the reportGenerationQuery, and an error, if there is any.
func (c *FakeReportGenerationQueries) Create(reportGenerationQuery *v1alpha1.ReportGenerationQuery) (result *v1alpha1.ReportGenerationQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reportgenerationqueriesResource, c.ns, reportGenerationQuery), &v1alpha1.ReportGenerationQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReportGenerationQuery), err
}

// Update takes the representation of a reportGenerationQuery and updates it. Returns the server's representation of the reportGenerationQuery, and an error, if there is any.
func (c *FakeReportGenerationQueries) Update(reportGenerationQuery *v1alpha1.ReportGenerationQuery) (result *v1alpha1.ReportGenerationQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reportgenerationqueriesResource, c.ns, reportGenerationQuery), &v1alpha1.ReportGenerationQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReportGenerationQuery), err
}

// Delete takes name of the reportGenerationQuery and deletes it. Returns an error if one occurs.
func (c *FakeReportGenerationQueries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reportgenerationqueriesResource, c.ns, name), &v1alpha1.ReportGenerationQuery{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReportGenerationQueries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reportgenerationqueriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReportGenerationQueryList{})
	return err
}

// Patch applies the patch and returns the patched reportGenerationQuery.
func (c *FakeReportGenerationQueries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReportGenerationQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reportgenerationqueriesResource, c.ns, name, data, subresources...), &v1alpha1.ReportGenerationQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReportGenerationQuery), err
}