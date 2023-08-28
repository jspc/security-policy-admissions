package main

import (
	"errors"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestValdidatedeployment(t *testing.T) {
	for _, test := range []struct {
		name        string
		deployment  *appsv1.Deployment
		expectError error
	}{
		{"empty deployment", emptyDeployment, nil},
		{"deployment with automountSAToken set to true errors", automountDeployment, automountServiceTokenErr},
		{"deployment with allowPrivEsc set to true errors", allowPrivDeployment, allowPrivilegeEscalationErr},
		{"deployment with privileged set true false errors", privilegedDeployment, privilegedErr},
		{"deployment with runAsNonRoot set to false errors", runAsNonRootDeployment, runAsNonRootErr},
		{"deployment with roFilesystem set to true errors", readonlyFSDeployment, readOnlyFilesystemErr},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := validateDeployment(test.deployment)
			if err != nil && test.expectError == nil {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && test.expectError != nil {
				t.Error("expected error")
			}

			if test.expectError != nil && !errors.Is(err, test.expectError) {
				t.Errorf("expected error %q, received %q", test.expectError, err)
			}
		})
	}
}

var (
	fv = false
	tv = true

	emptyDeployment = &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					AutomountServiceAccountToken: &fv,
					Containers:                   []corev1.Container{},
				},
			},
		},
	}

	automountDeployment = &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					AutomountServiceAccountToken: &tv,
					Containers:                   []corev1.Container{},
				},
			},
		},
	}

	allowPrivDeployment = &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					AutomountServiceAccountToken: &fv,
					Containers: []corev1.Container{
						{
							SecurityContext: &corev1.SecurityContext{
								AllowPrivilegeEscalation: &tv,
								Privileged:               &fv,
								RunAsNonRoot:             &tv,
								ReadOnlyRootFilesystem:   &tv,
							},
						},
					},
				},
			},
		},
	}

	privilegedDeployment = &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					AutomountServiceAccountToken: &fv,
					Containers: []corev1.Container{
						{
							SecurityContext: &corev1.SecurityContext{
								AllowPrivilegeEscalation: &fv,
								Privileged:               &tv,
								RunAsNonRoot:             &tv,
								ReadOnlyRootFilesystem:   &tv,
							},
						},
					},
				},
			},
		},
	}

	runAsNonRootDeployment = &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					AutomountServiceAccountToken: &fv,
					Containers: []corev1.Container{
						{
							SecurityContext: &corev1.SecurityContext{
								AllowPrivilegeEscalation: &fv,
								Privileged:               &fv,
								RunAsNonRoot:             &fv,
								ReadOnlyRootFilesystem:   &tv,
							},
						},
					},
				},
			},
		},
	}

	readonlyFSDeployment = &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					AutomountServiceAccountToken: &fv,
					Containers: []corev1.Container{
						{
							SecurityContext: &corev1.SecurityContext{
								AllowPrivilegeEscalation: &fv,
								Privileged:               &fv,
								RunAsNonRoot:             &tv,
								ReadOnlyRootFilesystem:   &fv,
							},
						},
					},
				},
			},
		},
	}
)
