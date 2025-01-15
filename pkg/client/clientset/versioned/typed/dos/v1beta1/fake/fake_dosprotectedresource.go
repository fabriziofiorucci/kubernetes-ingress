// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/nginx/kubernetes-ingress/pkg/apis/dos/v1beta1"
	dosv1beta1 "github.com/nginx/kubernetes-ingress/pkg/client/clientset/versioned/typed/dos/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeDosProtectedResources implements DosProtectedResourceInterface
type fakeDosProtectedResources struct {
	*gentype.FakeClientWithList[*v1beta1.DosProtectedResource, *v1beta1.DosProtectedResourceList]
	Fake *FakeAppprotectdosV1beta1
}

func newFakeDosProtectedResources(fake *FakeAppprotectdosV1beta1, namespace string) dosv1beta1.DosProtectedResourceInterface {
	return &fakeDosProtectedResources{
		gentype.NewFakeClientWithList[*v1beta1.DosProtectedResource, *v1beta1.DosProtectedResourceList](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("dosprotectedresources"),
			v1beta1.SchemeGroupVersion.WithKind("DosProtectedResource"),
			func() *v1beta1.DosProtectedResource { return &v1beta1.DosProtectedResource{} },
			func() *v1beta1.DosProtectedResourceList { return &v1beta1.DosProtectedResourceList{} },
			func(dst, src *v1beta1.DosProtectedResourceList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.DosProtectedResourceList) []*v1beta1.DosProtectedResource {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.DosProtectedResourceList, items []*v1beta1.DosProtectedResource) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
