// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListInfrastructureConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	Filters []Filter `locationName:"filters" min:"1" type:"list"`

	// The maximum items to return in a request.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token to specify where to start paginating. This is the NextToken from
	// a previously truncated response.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListInfrastructureConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInfrastructureConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListInfrastructureConfigurationsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListInfrastructureConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListInfrastructureConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// The list of infrastructure configurations.
	InfrastructureConfigurationSummaryList []InfrastructureConfigurationSummary `locationName:"infrastructureConfigurationSummaryList" type:"list"`

	// The next token used for paginated responses. When this is not empty, there
	// are additional elements that the service has not included in this request.
	// Use this token with the next request to retrieve additional objects.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s ListInfrastructureConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListInfrastructureConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InfrastructureConfigurationSummaryList != nil {
		v := s.InfrastructureConfigurationSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "infrastructureConfigurationSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListInfrastructureConfigurations = "ListInfrastructureConfigurations"

// ListInfrastructureConfigurationsRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Returns a list of infrastructure configurations.
//
//    // Example sending a request using ListInfrastructureConfigurationsRequest.
//    req := client.ListInfrastructureConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/ListInfrastructureConfigurations
func (c *Client) ListInfrastructureConfigurationsRequest(input *ListInfrastructureConfigurationsInput) ListInfrastructureConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListInfrastructureConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/ListInfrastructureConfigurations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListInfrastructureConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListInfrastructureConfigurationsOutput{})
	return ListInfrastructureConfigurationsRequest{Request: req, Input: input, Copy: c.ListInfrastructureConfigurationsRequest}
}

// ListInfrastructureConfigurationsRequest is the request type for the
// ListInfrastructureConfigurations API operation.
type ListInfrastructureConfigurationsRequest struct {
	*aws.Request
	Input *ListInfrastructureConfigurationsInput
	Copy  func(*ListInfrastructureConfigurationsInput) ListInfrastructureConfigurationsRequest
}

// Send marshals and sends the ListInfrastructureConfigurations API request.
func (r ListInfrastructureConfigurationsRequest) Send(ctx context.Context) (*ListInfrastructureConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInfrastructureConfigurationsResponse{
		ListInfrastructureConfigurationsOutput: r.Request.Data.(*ListInfrastructureConfigurationsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListInfrastructureConfigurationsRequestPaginator returns a paginator for ListInfrastructureConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListInfrastructureConfigurationsRequest(input)
//   p := imagebuilder.NewListInfrastructureConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListInfrastructureConfigurationsPaginator(req ListInfrastructureConfigurationsRequest) ListInfrastructureConfigurationsPaginator {
	return ListInfrastructureConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListInfrastructureConfigurationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListInfrastructureConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListInfrastructureConfigurationsPaginator struct {
	aws.Pager
}

func (p *ListInfrastructureConfigurationsPaginator) CurrentPage() *ListInfrastructureConfigurationsOutput {
	return p.Pager.CurrentPage().(*ListInfrastructureConfigurationsOutput)
}

// ListInfrastructureConfigurationsResponse is the response type for the
// ListInfrastructureConfigurations API operation.
type ListInfrastructureConfigurationsResponse struct {
	*ListInfrastructureConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInfrastructureConfigurations request.
func (r *ListInfrastructureConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
