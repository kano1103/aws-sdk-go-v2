// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstanceSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of the snapshot for which you are requesting information.
	//
	// InstanceSnapshotName is a required field
	InstanceSnapshotName *string `locationName:"instanceSnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceSnapshotInput"}

	if s.InstanceSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInstanceSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the results of your
	// get instance snapshot request.
	InstanceSnapshot *InstanceSnapshot `locationName:"instanceSnapshot" type:"structure"`
}

// String returns the string representation
func (s GetInstanceSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstanceSnapshot = "GetInstanceSnapshot"

// GetInstanceSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about a specific instance snapshot.
//
//    // Example sending a request using GetInstanceSnapshotRequest.
//    req := client.GetInstanceSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstanceSnapshot
func (c *Client) GetInstanceSnapshotRequest(input *GetInstanceSnapshotInput) GetInstanceSnapshotRequest {
	op := &aws.Operation{
		Name:       opGetInstanceSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstanceSnapshotInput{}
	}

	req := c.newRequest(op, input, &GetInstanceSnapshotOutput{})
	return GetInstanceSnapshotRequest{Request: req, Input: input, Copy: c.GetInstanceSnapshotRequest}
}

// GetInstanceSnapshotRequest is the request type for the
// GetInstanceSnapshot API operation.
type GetInstanceSnapshotRequest struct {
	*aws.Request
	Input *GetInstanceSnapshotInput
	Copy  func(*GetInstanceSnapshotInput) GetInstanceSnapshotRequest
}

// Send marshals and sends the GetInstanceSnapshot API request.
func (r GetInstanceSnapshotRequest) Send(ctx context.Context) (*GetInstanceSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstanceSnapshotResponse{
		GetInstanceSnapshotOutput: r.Request.Data.(*GetInstanceSnapshotOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstanceSnapshotResponse is the response type for the
// GetInstanceSnapshot API operation.
type GetInstanceSnapshotResponse struct {
	*GetInstanceSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstanceSnapshot request.
func (r *GetInstanceSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
