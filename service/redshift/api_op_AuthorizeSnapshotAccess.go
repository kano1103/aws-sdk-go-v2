// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AuthorizeSnapshotAccessInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the AWS customer account authorized to restore the specified
	// snapshot.
	//
	// To share a snapshot with AWS support, specify amazon-redshift-support.
	//
	// AccountWithRestoreAccess is a required field
	AccountWithRestoreAccess *string `type:"string" required:"true"`

	// The identifier of the cluster the snapshot was created from. This parameter
	// is required if your IAM user has a policy containing a snapshot resource
	// element that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string `type:"string"`

	// The identifier of the snapshot the account is authorized to restore.
	//
	// SnapshotIdentifier is a required field
	SnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AuthorizeSnapshotAccessInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeSnapshotAccessInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AuthorizeSnapshotAccessInput"}

	if s.AccountWithRestoreAccess == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountWithRestoreAccess"))
	}

	if s.SnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AuthorizeSnapshotAccessOutput struct {
	_ struct{} `type:"structure"`

	// Describes a snapshot.
	Snapshot *Snapshot `type:"structure"`
}

// String returns the string representation
func (s AuthorizeSnapshotAccessOutput) String() string {
	return awsutil.Prettify(s)
}

const opAuthorizeSnapshotAccess = "AuthorizeSnapshotAccess"

// AuthorizeSnapshotAccessRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Authorizes the specified AWS customer account to restore the specified snapshot.
//
// For more information about working with snapshots, go to Amazon Redshift
// Snapshots (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using AuthorizeSnapshotAccessRequest.
//    req := client.AuthorizeSnapshotAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/AuthorizeSnapshotAccess
func (c *Client) AuthorizeSnapshotAccessRequest(input *AuthorizeSnapshotAccessInput) AuthorizeSnapshotAccessRequest {
	op := &aws.Operation{
		Name:       opAuthorizeSnapshotAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeSnapshotAccessInput{}
	}

	req := c.newRequest(op, input, &AuthorizeSnapshotAccessOutput{})
	return AuthorizeSnapshotAccessRequest{Request: req, Input: input, Copy: c.AuthorizeSnapshotAccessRequest}
}

// AuthorizeSnapshotAccessRequest is the request type for the
// AuthorizeSnapshotAccess API operation.
type AuthorizeSnapshotAccessRequest struct {
	*aws.Request
	Input *AuthorizeSnapshotAccessInput
	Copy  func(*AuthorizeSnapshotAccessInput) AuthorizeSnapshotAccessRequest
}

// Send marshals and sends the AuthorizeSnapshotAccess API request.
func (r AuthorizeSnapshotAccessRequest) Send(ctx context.Context) (*AuthorizeSnapshotAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AuthorizeSnapshotAccessResponse{
		AuthorizeSnapshotAccessOutput: r.Request.Data.(*AuthorizeSnapshotAccessOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AuthorizeSnapshotAccessResponse is the response type for the
// AuthorizeSnapshotAccess API operation.
type AuthorizeSnapshotAccessResponse struct {
	*AuthorizeSnapshotAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AuthorizeSnapshotAccess request.
func (r *AuthorizeSnapshotAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
