// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package authorizer

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/apigatewayv2/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.Authorizer{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetAuthorizerWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AuthorizerId != nil {
		ko.Status.AuthorizerID = resp.AuthorizerId
	}

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.AuthorizerID == nil || r.ko.Spec.APIID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetAuthorizerInput, error) {
	res := &svcsdk.GetAuthorizerInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.AuthorizerID != nil {
		res.SetAuthorizerId(*r.ko.Status.AuthorizerID)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.GetAuthorizersInput, error) {
	res := &svcsdk.GetAuthorizersInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateAuthorizerWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AuthorizerId != nil {
		ko.Status.AuthorizerID = resp.AuthorizerId
	}

	ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{OwnerAccountID: &rm.awsAccountID}
	ko.Status.Conditions = []*ackv1alpha1.Condition{}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateAuthorizerInput, error) {
	res := &svcsdk.CreateAuthorizerInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.AuthorizerCredentialsARN != nil {
		res.SetAuthorizerCredentialsArn(*r.ko.Spec.AuthorizerCredentialsARN)
	}
	if r.ko.Spec.AuthorizerResultTtlInSeconds != nil {
		res.SetAuthorizerResultTtlInSeconds(*r.ko.Spec.AuthorizerResultTtlInSeconds)
	}
	if r.ko.Spec.AuthorizerType != nil {
		res.SetAuthorizerType(*r.ko.Spec.AuthorizerType)
	}
	if r.ko.Spec.AuthorizerURI != nil {
		res.SetAuthorizerUri(*r.ko.Spec.AuthorizerURI)
	}
	if r.ko.Spec.IDentitySource != nil {
		f5 := []*string{}
		for _, f5iter := range r.ko.Spec.IDentitySource {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		res.SetIdentitySource(f5)
	}
	if r.ko.Spec.IDentityValidationExpression != nil {
		res.SetIdentityValidationExpression(*r.ko.Spec.IDentityValidationExpression)
	}
	if r.ko.Spec.JWTConfiguration != nil {
		f7 := &svcsdk.JWTConfiguration{}
		if r.ko.Spec.JWTConfiguration.Audience != nil {
			f7f0 := []*string{}
			for _, f7f0iter := range r.ko.Spec.JWTConfiguration.Audience {
				var f7f0elem string
				f7f0elem = *f7f0iter
				f7f0 = append(f7f0, &f7f0elem)
			}
			f7.SetAudience(f7f0)
		}
		if r.ko.Spec.JWTConfiguration.Issuer != nil {
			f7.SetIssuer(*r.ko.Spec.JWTConfiguration.Issuer)
		}
		res.SetJwtConfiguration(f7)
	}
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	r *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	input, err := rm.newUpdateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateAuthorizerWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AuthorizerId != nil {
		ko.Status.AuthorizerID = resp.AuthorizerId
	}

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateAuthorizerInput, error) {
	res := &svcsdk.UpdateAuthorizerInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.AuthorizerCredentialsARN != nil {
		res.SetAuthorizerCredentialsArn(*r.ko.Spec.AuthorizerCredentialsARN)
	}
	if r.ko.Status.AuthorizerID != nil {
		res.SetAuthorizerId(*r.ko.Status.AuthorizerID)
	}
	if r.ko.Spec.AuthorizerResultTtlInSeconds != nil {
		res.SetAuthorizerResultTtlInSeconds(*r.ko.Spec.AuthorizerResultTtlInSeconds)
	}
	if r.ko.Spec.AuthorizerType != nil {
		res.SetAuthorizerType(*r.ko.Spec.AuthorizerType)
	}
	if r.ko.Spec.AuthorizerURI != nil {
		res.SetAuthorizerUri(*r.ko.Spec.AuthorizerURI)
	}
	if r.ko.Spec.IDentitySource != nil {
		f6 := []*string{}
		for _, f6iter := range r.ko.Spec.IDentitySource {
			var f6elem string
			f6elem = *f6iter
			f6 = append(f6, &f6elem)
		}
		res.SetIdentitySource(f6)
	}
	if r.ko.Spec.IDentityValidationExpression != nil {
		res.SetIdentityValidationExpression(*r.ko.Spec.IDentityValidationExpression)
	}
	if r.ko.Spec.JWTConfiguration != nil {
		f8 := &svcsdk.JWTConfiguration{}
		if r.ko.Spec.JWTConfiguration.Audience != nil {
			f8f0 := []*string{}
			for _, f8f0iter := range r.ko.Spec.JWTConfiguration.Audience {
				var f8f0elem string
				f8f0elem = *f8f0iter
				f8f0 = append(f8f0, &f8f0elem)
			}
			f8.SetAudience(f8f0)
		}
		if r.ko.Spec.JWTConfiguration.Issuer != nil {
			f8.SetIssuer(*r.ko.Spec.JWTConfiguration.Issuer)
		}
		res.SetJwtConfiguration(f8)
	}
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteAuthorizerWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteAuthorizerInput, error) {
	res := &svcsdk.DeleteAuthorizerInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.AuthorizerID != nil {
		res.SetAuthorizerId(*r.ko.Status.AuthorizerID)
	}

	return res, nil
}
