// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type AbortDocumentVersionUploadInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// The ID of the version.
	//
	// VersionId is a required field
	VersionId *string `location:"uri" locationName:"VersionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AbortDocumentVersionUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AbortDocumentVersionUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AbortDocumentVersionUploadInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}
	if s.VersionId != nil && len(*s.VersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortDocumentVersionUploadInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentId != nil {
		v := *s.DocumentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DocumentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AbortDocumentVersionUploadOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AbortDocumentVersionUploadOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortDocumentVersionUploadOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAbortDocumentVersionUpload = "AbortDocumentVersionUpload"

// AbortDocumentVersionUploadRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Aborts the upload of the specified document version that was previously initiated
// by InitiateDocumentVersionUpload. The client should make this call only when
// it no longer intends to upload the document version, or fails to do so.
//
//    // Example sending a request using AbortDocumentVersionUploadRequest.
//    req := client.AbortDocumentVersionUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/AbortDocumentVersionUpload
func (c *Client) AbortDocumentVersionUploadRequest(input *AbortDocumentVersionUploadInput) AbortDocumentVersionUploadRequest {
	op := &aws.Operation{
		Name:       opAbortDocumentVersionUpload,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/documents/{DocumentId}/versions/{VersionId}",
	}

	if input == nil {
		input = &AbortDocumentVersionUploadInput{}
	}

	req := c.newRequest(op, input, &AbortDocumentVersionUploadOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AbortDocumentVersionUploadRequest{Request: req, Input: input, Copy: c.AbortDocumentVersionUploadRequest}
}

// AbortDocumentVersionUploadRequest is the request type for the
// AbortDocumentVersionUpload API operation.
type AbortDocumentVersionUploadRequest struct {
	*aws.Request
	Input *AbortDocumentVersionUploadInput
	Copy  func(*AbortDocumentVersionUploadInput) AbortDocumentVersionUploadRequest
}

// Send marshals and sends the AbortDocumentVersionUpload API request.
func (r AbortDocumentVersionUploadRequest) Send(ctx context.Context) (*AbortDocumentVersionUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AbortDocumentVersionUploadResponse{
		AbortDocumentVersionUploadOutput: r.Request.Data.(*AbortDocumentVersionUploadOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AbortDocumentVersionUploadResponse is the response type for the
// AbortDocumentVersionUpload API operation.
type AbortDocumentVersionUploadResponse struct {
	*AbortDocumentVersionUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AbortDocumentVersionUpload request.
func (r *AbortDocumentVersionUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
