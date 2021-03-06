// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRegexPatternSetInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the set. This ID is returned in the responses to
	// create and list commands. You provide it to operations like update and delete.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// A token used for optimistic locking. AWS WAF returns a token to your get
	// and list requests, to mark the state of the entity at the time of the request.
	// To make changes to the entity associated with the token, you provide the
	// token to operations like update and delete. AWS WAF uses the token to ensure
	// that no changes have been made to the entity since you last retrieved it.
	// If a change has been made, the update fails with a WAFOptimisticLockException.
	// If this happens, perform another get, and use the new token returned by that
	// operation.
	//
	// LockToken is a required field
	LockToken *string `min:"1" type:"string" required:"true"`

	// A friendly name of the set. You cannot change the name after you create the
	// set.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DeleteRegexPatternSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRegexPatternSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRegexPatternSetInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.LockToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("LockToken"))
	}
	if s.LockToken != nil && len(*s.LockToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LockToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRegexPatternSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRegexPatternSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRegexPatternSet = "DeleteRegexPatternSet"

// DeleteRegexPatternSetRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Deletes the specified RegexPatternSet.
//
//    // Example sending a request using DeleteRegexPatternSetRequest.
//    req := client.DeleteRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/DeleteRegexPatternSet
func (c *Client) DeleteRegexPatternSetRequest(input *DeleteRegexPatternSetInput) DeleteRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opDeleteRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &DeleteRegexPatternSetOutput{})
	return DeleteRegexPatternSetRequest{Request: req, Input: input, Copy: c.DeleteRegexPatternSetRequest}
}

// DeleteRegexPatternSetRequest is the request type for the
// DeleteRegexPatternSet API operation.
type DeleteRegexPatternSetRequest struct {
	*aws.Request
	Input *DeleteRegexPatternSetInput
	Copy  func(*DeleteRegexPatternSetInput) DeleteRegexPatternSetRequest
}

// Send marshals and sends the DeleteRegexPatternSet API request.
func (r DeleteRegexPatternSetRequest) Send(ctx context.Context) (*DeleteRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRegexPatternSetResponse{
		DeleteRegexPatternSetOutput: r.Request.Data.(*DeleteRegexPatternSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRegexPatternSetResponse is the response type for the
// DeleteRegexPatternSet API operation.
type DeleteRegexPatternSetResponse struct {
	*DeleteRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRegexPatternSet request.
func (r *DeleteRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
