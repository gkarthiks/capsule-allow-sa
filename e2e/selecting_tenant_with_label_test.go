//go:build e2e

// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	capsulev1beta1 "github.com/clastix/capsule/api/v1beta1"
)

var _ = Describe("creating a Namespace with Tenant selector when user owns multiple tenants", func() {
	t1 := &capsulev1beta1.Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tenant-one",
		},
		Spec: capsulev1beta1.TenantSpec{
			Owners: capsulev1beta1.OwnerListSpec{
				{
					Name: "john",
					Kind: "User",
				},
			},
		},
	}
	t2 := &capsulev1beta1.Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tenant-two",
		},
		Spec: capsulev1beta1.TenantSpec{
			Owners: capsulev1beta1.OwnerListSpec{
				{
					Name: "john",
					Kind: "User",
				},
			},
		},
	}

	JustBeforeEach(func() {
		EventuallyCreation(func() error {
			return k8sClient.Create(context.TODO(), t1)
		}).Should(Succeed())
		EventuallyCreation(func() error {
			return k8sClient.Create(context.TODO(), t2)
		}).Should(Succeed())
	})
	JustAfterEach(func() {
		Expect(k8sClient.Delete(context.TODO(), t1)).Should(Succeed())
		Expect(k8sClient.Delete(context.TODO(), t2)).Should(Succeed())
	})

	It("should be assigned to the selected Tenant", func() {
		ns := NewNamespace("tenant-2-ns")
		By("assigning to the Namespace the Capsule Tenant label", func() {
			l, err := capsulev1beta1.GetTypeLabel(&capsulev1beta1.Tenant{})
			Expect(err).ToNot(HaveOccurred())
			ns.Labels = map[string]string{
				l: t2.Name,
			}
		})
		NamespaceCreation(ns, t2.Spec.Owners[0], defaultTimeoutInterval).Should(Succeed())
		TenantNamespaceList(t2, defaultTimeoutInterval).Should(ContainElement(ns.GetName()))
	})
})
