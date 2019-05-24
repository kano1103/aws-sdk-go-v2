// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// UpdateSMBFileShareInput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateSMBFileShareInput
type UpdateSMBFileShareInput struct {
	_ struct{} `type:"structure"`

	// The default storage class for objects put into an Amazon S3 bucket by the
	// file gateway. Possible values are S3_STANDARD, S3_STANDARD_IA, or S3_ONEZONE_IA.
	// If this field is not populated, the default value S3_STANDARD is used. Optional.
	DefaultStorageClass *string `min:"5" type:"string"`

	// The Amazon Resource Name (ARN) of the SMB file share that you want to update.
	//
	// FileShareARN is a required field
	FileShareARN *string `min:"50" type:"string" required:"true"`

	// A value that enables guessing of the MIME type for uploaded objects based
	// on file extensions. Set this value to true to enable MIME type guessing,
	// and otherwise to false. The default value is true.
	GuessMIMETypeEnabled *bool `type:"boolean"`

	// A list of users or groups in the Active Directory that are not allowed to
	// access the file share. A group must be prefixed with the @ character. For
	// example @group1. Can only be set if Authentication is set to ActiveDirectory.
	InvalidUserList []string `type:"list"`

	// True to use Amazon S3 server side encryption with your own AWS KMS key, or
	// false to use a key managed by Amazon S3. Optional.
	KMSEncrypted *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server
	// side encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// A value that sets the access control list permission for objects in the S3
	// bucket that a file gateway puts objects into. The default value is "private".
	ObjectACL ObjectACL `type:"string" enum:"true"`

	// A value that sets the write status of a file share. This value is true if
	// the write status is read-only, and otherwise false.
	ReadOnly *bool `type:"boolean"`

	// A value that sets who pays the cost of the request and the cost associated
	// with data download from the S3 bucket. If this value is set to true, the
	// requester pays the costs. Otherwise the S3 bucket owner pays. However, the
	// S3 bucket owner always pays the cost of storing data.
	//
	// RequesterPays is a configuration for the S3 bucket that backs the file share,
	// so make sure that the configuration on the file share is the same as the
	// S3 bucket configuration.
	RequesterPays *bool `type:"boolean"`

	// Set this value to "true to enable ACL (access control list) on the SMB file
	// share. Set it to "false" to map file and directory permissions to the POSIX
	// permissions.
	SMBACLEnabled *bool `type:"boolean"`

	// A list of users or groups in the Active Directory that are allowed to access
	// the file share. A group must be prefixed with the @ character. For example
	// @group1. Can only be set if Authentication is set to ActiveDirectory.
	ValidUserList []string `type:"list"`
}

// String returns the string representation
func (s UpdateSMBFileShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSMBFileShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSMBFileShareInput"}
	if s.DefaultStorageClass != nil && len(*s.DefaultStorageClass) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultStorageClass", 5))
	}

	if s.FileShareARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileShareARN"))
	}
	if s.FileShareARN != nil && len(*s.FileShareARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareARN", 50))
	}
	if s.KMSKey != nil && len(*s.KMSKey) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("KMSKey", 7))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// UpdateSMBFileShareOutput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateSMBFileShareOutput
type UpdateSMBFileShareOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the updated SMB file share.
	FileShareARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateSMBFileShareOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSMBFileShare = "UpdateSMBFileShare"

// UpdateSMBFileShareRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates a Server Message Block (SMB) file share.
//
// To leave a file share field unchanged, set the corresponding input field
// to null. This operation is only supported for file gateways.
//
// File gateways require AWS Security Token Service (AWS STS) to be activated
// to enable you to create a file share. Make sure that AWS STS is activated
// in the AWS Region you are creating your file gateway in. If AWS STS is not
// activated in this AWS Region, activate it. For information about how to activate
// AWS STS, see Activating and Deactivating AWS STS in an AWS Region (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the AWS Identity and Access Management User Guide.
//
// File gateways don't support creating hard or symbolic links on a file share.
//
//    // Example sending a request using UpdateSMBFileShareRequest.
//    req := client.UpdateSMBFileShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateSMBFileShare
func (c *Client) UpdateSMBFileShareRequest(input *UpdateSMBFileShareInput) UpdateSMBFileShareRequest {
	op := &aws.Operation{
		Name:       opUpdateSMBFileShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSMBFileShareInput{}
	}

	req := c.newRequest(op, input, &UpdateSMBFileShareOutput{})
	return UpdateSMBFileShareRequest{Request: req, Input: input, Copy: c.UpdateSMBFileShareRequest}
}

// UpdateSMBFileShareRequest is the request type for the
// UpdateSMBFileShare API operation.
type UpdateSMBFileShareRequest struct {
	*aws.Request
	Input *UpdateSMBFileShareInput
	Copy  func(*UpdateSMBFileShareInput) UpdateSMBFileShareRequest
}

// Send marshals and sends the UpdateSMBFileShare API request.
func (r UpdateSMBFileShareRequest) Send(ctx context.Context) (*UpdateSMBFileShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSMBFileShareResponse{
		UpdateSMBFileShareOutput: r.Request.Data.(*UpdateSMBFileShareOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSMBFileShareResponse is the response type for the
// UpdateSMBFileShare API operation.
type UpdateSMBFileShareResponse struct {
	*UpdateSMBFileShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSMBFileShare request.
func (r *UpdateSMBFileShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}