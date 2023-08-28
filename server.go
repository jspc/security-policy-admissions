package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
	admissionv1 "k8s.io/api/admission/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	automountServiceTokenErr    = fmt.Errorf("automounting of Sevice Account Tokens must be false")
	allowPrivilegeEscalationErr = fmt.Errorf("allowPrivilegeEscalation must be false")
	privilegedErr               = fmt.Errorf("privileged must be false")
	runAsNonRootErr             = fmt.Errorf("runAsNonRoot must be true")
	readOnlyFilesystemErr       = fmt.Errorf("readOnlyRootFilesystem must be true")
)

func response(allow bool, msg string) admissionv1.AdmissionReview {
	return admissionv1.AdmissionReview{
		Response: &admissionv1.AdmissionResponse{
			Allowed: allow,
			Result: &metav1.Status{
				Message: msg,
			},
		},
	}
}

func Validate(ctx *fasthttp.RequestCtx) {
	request := admissionv1.AdmissionReview{}

	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString("unknown body type")

		return
	}

	if request.Request.Namespace == "kube-system" ||
		request.Request.RequestKind.Kind != "Deployment" ||
		request.Request.RequestKind.Group != "apps" {
		json.NewEncoder(ctx).Encode(response(true, "skipping this resource"))

		return
	}

	d := new(appsv1.Deployment)
	err = json.Unmarshal(request.Request.Object.Raw, d)
	if err != nil {
		json.NewEncoder(ctx).Encode(response(false, "could not parse request, erroring from caution"))

		return
	}

	err = validateDeployment(d)
	if err != nil {
		json.NewEncoder(ctx).Encode(response(false, err.Error()))

		return
	}

	json.NewEncoder(ctx).Encode(response(true, "ok"))
}

func validateDeployment(d *appsv1.Deployment) error {
	if *d.Spec.Template.Spec.AutomountServiceAccountToken != false {
		return automountServiceTokenErr
	}

	for _, c := range d.Spec.Template.Spec.Containers {
		for _, r := range []struct {
			v, expect bool
			err       error
		}{
			{*c.SecurityContext.AllowPrivilegeEscalation, false, allowPrivilegeEscalationErr},
			{*c.SecurityContext.Privileged, false, privilegedErr},
			{*c.SecurityContext.RunAsNonRoot, true, runAsNonRootErr},
			{*c.SecurityContext.ReadOnlyRootFilesystem, true, readOnlyFilesystemErr},
		} {
			if r.v != r.expect {
				return r.err
			}

		}
	}

	return nil
}
