// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteMeetingInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMeetingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMeetingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMeetingInput"}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMeetingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteMeetingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMeetingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMeetingOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteMeeting = "DeleteMeeting"

// DeleteMeetingRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the specified Amazon Chime SDK meeting. When a meeting is deleted,
// its attendees are also deleted and clients can no longer join it. For more
// information about the Amazon Chime SDK, see Using the Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
//
//    // Example sending a request using DeleteMeetingRequest.
//    req := client.DeleteMeetingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteMeeting
func (c *Client) DeleteMeetingRequest(input *DeleteMeetingInput) DeleteMeetingRequest {
	op := &aws.Operation{
		Name:       opDeleteMeeting,
		HTTPMethod: "DELETE",
		HTTPPath:   "/meetings/{meetingId}",
	}

	if input == nil {
		input = &DeleteMeetingInput{}
	}

	req := c.newRequest(op, input, &DeleteMeetingOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteMeetingRequest{Request: req, Input: input, Copy: c.DeleteMeetingRequest}
}

// DeleteMeetingRequest is the request type for the
// DeleteMeeting API operation.
type DeleteMeetingRequest struct {
	*aws.Request
	Input *DeleteMeetingInput
	Copy  func(*DeleteMeetingInput) DeleteMeetingRequest
}

// Send marshals and sends the DeleteMeeting API request.
func (r DeleteMeetingRequest) Send(ctx context.Context) (*DeleteMeetingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMeetingResponse{
		DeleteMeetingOutput: r.Request.Data.(*DeleteMeetingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMeetingResponse is the response type for the
// DeleteMeeting API operation.
type DeleteMeetingResponse struct {
	*DeleteMeetingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMeeting request.
func (r *DeleteMeetingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
