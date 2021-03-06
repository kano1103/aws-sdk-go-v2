// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateResolverRuleInput struct {
	_ struct{} `type:"structure"`

	// A unique string that identifies the request and that allows failed requests
	// to be retried without the risk of executing the operation twice. CreatorRequestId
	// can be any unique string, for example, a date/time stamp.
	//
	// CreatorRequestId is a required field
	CreatorRequestId *string `min:"1" type:"string" required:"true"`

	// DNS queries for this domain name are forwarded to the IP addresses that you
	// specify in TargetIps. If a query matches multiple resolver rules (example.com
	// and www.example.com), outbound DNS queries are routed using the resolver
	// rule that contains the most specific domain name (www.example.com).
	//
	// DomainName is a required field
	DomainName *string `min:"1" type:"string" required:"true"`

	// A friendly name that lets you easily find a rule in the Resolver dashboard
	// in the Route 53 console.
	Name *string `type:"string"`

	// The ID of the outbound resolver endpoint that you want to use to route DNS
	// queries to the IP addresses that you specify in TargetIps.
	ResolverEndpointId *string `min:"1" type:"string"`

	// Specify FORWARD. Other resolver rule types aren't supported.
	//
	// RuleType is a required field
	RuleType RuleTypeOption `type:"string" required:"true" enum:"true"`

	// A list of the tag keys and values that you want to associate with the endpoint.
	Tags []Tag `type:"list"`

	// The IPs that you want Resolver to forward DNS queries to. You can specify
	// only IPv4 addresses. Separate IP addresses with a comma.
	TargetIps []TargetAddress `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateResolverRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResolverRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResolverRuleInput"}

	if s.CreatorRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreatorRequestId"))
	}
	if s.CreatorRequestId != nil && len(*s.CreatorRequestId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CreatorRequestId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 1))
	}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}
	if len(s.RuleType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RuleType"))
	}
	if s.TargetIps != nil && len(s.TargetIps) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetIps", 1))
	}
	if s.TargetIps != nil {
		for i, v := range s.TargetIps {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetIps", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateResolverRuleOutput struct {
	_ struct{} `type:"structure"`

	// Information about the CreateResolverRule request, including the status of
	// the request.
	ResolverRule *ResolverRule `type:"structure"`
}

// String returns the string representation
func (s CreateResolverRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateResolverRule = "CreateResolverRule"

// CreateResolverRuleRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// For DNS queries that originate in your VPCs, specifies which resolver endpoint
// the queries pass through, one domain name that you want to forward to your
// network, and the IP addresses of the DNS resolvers in your network.
//
//    // Example sending a request using CreateResolverRuleRequest.
//    req := client.CreateResolverRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/CreateResolverRule
func (c *Client) CreateResolverRuleRequest(input *CreateResolverRuleInput) CreateResolverRuleRequest {
	op := &aws.Operation{
		Name:       opCreateResolverRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateResolverRuleInput{}
	}

	req := c.newRequest(op, input, &CreateResolverRuleOutput{})
	return CreateResolverRuleRequest{Request: req, Input: input, Copy: c.CreateResolverRuleRequest}
}

// CreateResolverRuleRequest is the request type for the
// CreateResolverRule API operation.
type CreateResolverRuleRequest struct {
	*aws.Request
	Input *CreateResolverRuleInput
	Copy  func(*CreateResolverRuleInput) CreateResolverRuleRequest
}

// Send marshals and sends the CreateResolverRule API request.
func (r CreateResolverRuleRequest) Send(ctx context.Context) (*CreateResolverRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateResolverRuleResponse{
		CreateResolverRuleOutput: r.Request.Data.(*CreateResolverRuleOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateResolverRuleResponse is the response type for the
// CreateResolverRule API operation.
type CreateResolverRuleResponse struct {
	*CreateResolverRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateResolverRule request.
func (r *CreateResolverRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
