//go:build e2e

// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	capsulev1beta1 "github.com/clastix/capsule/api/v1beta1"
)

var _ = Describe("enforcing an allowed set of Service external IPs", func() {
	tnt := &capsulev1beta1.Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name: "allowed-external-ip",
		},
		Spec: capsulev1beta1.TenantSpec{
			Owners: capsulev1beta1.OwnerListSpec{
				{
					Name: "google",
					Kind: "User",
				},
			},
			ServiceOptions: &capsulev1beta1.ServiceOptions{
				ExternalServiceIPs: &capsulev1beta1.ExternalServiceIPsSpec{
					Allowed: []capsulev1beta1.AllowedIP{
						"10.20.0.0/16",
						"192.168.1.2/32",
					},
				},
			},
		},
	}

	JustBeforeEach(func() {
		EventuallyCreation(func() error {
			tnt.ResourceVersion = ""
			return k8sClient.Create(context.TODO(), tnt)
		}).Should(Succeed())
	})
	JustAfterEach(func() {
		Expect(k8sClient.Delete(context.TODO(), tnt)).Should(Succeed())
	})

	It("should fail creating an evil service", func() {
		ns := NewNamespace("evil-service")
		NamespaceCreation(ns, tnt.Spec.Owners[0], defaultTimeoutInterval).Should(Succeed())

		svc := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-evil-dns-server",
			},
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Name:       "dns",
						Protocol:   "UDP",
						Port:       53,
						TargetPort: intstr.FromInt(9053),
					},
				},
				Selector: map[string]string{
					"app": "my-evil-dns-server",
				},
				ExternalIPs: []string{
					"8.8.8.8",
					"8.8.4.4",
				},
			},
		}
		EventuallyCreation(func() error {
			cs := ownerClient(tnt.Spec.Owners[0])
			_, err := cs.CoreV1().Services(ns.Name).Create(context.Background(), svc, metav1.CreateOptions{})
			return err
		}).ShouldNot(Succeed())
	})

	It("should allow the first CIDR block", func() {
		ns := NewNamespace("allowed-service-cidr")
		NamespaceCreation(ns, tnt.Spec.Owners[0], defaultTimeoutInterval).Should(Succeed())

		svc := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dns-server",
			},
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Name:       "dns",
						Protocol:   "UDP",
						Port:       53,
						TargetPort: intstr.FromInt(9053),
					},
				},
				Selector: map[string]string{
					"app": "dns-server",
				},
				ExternalIPs: []string{
					"10.20.0.0",
					"10.20.255.255",
				},
			},
		}
		EventuallyCreation(func() error {
			cs := ownerClient(tnt.Spec.Owners[0])
			_, err := cs.CoreV1().Services(ns.Name).Create(context.Background(), svc, metav1.CreateOptions{})
			return err
		}).Should(Succeed())
	})

	It("should allow the /32 CIDR block", func() {
		ns := NewNamespace("allowed-service-strict")
		NamespaceCreation(ns, tnt.Spec.Owners[0], defaultTimeoutInterval).Should(Succeed())

		svc := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dns-server",
			},
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Name:       "dns",
						Protocol:   "UDP",
						Port:       53,
						TargetPort: intstr.FromInt(9053),
					},
				},
				Selector: map[string]string{
					"app": "dns-server",
				},
				ExternalIPs: []string{
					"192.168.1.2",
				},
			},
		}
		EventuallyCreation(func() error {
			cs := ownerClient(tnt.Spec.Owners[0])
			_, err := cs.CoreV1().Services(ns.Name).Create(context.Background(), svc, metav1.CreateOptions{})
			return err
		}).Should(Succeed())
	})
})
