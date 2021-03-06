// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the DataSource at creation.
	//
	// DataSourceId is a required field
	DataSourceId *string `min:"1" type:"string" required:"true"`

	// Specifies whether the GetDataSource operation should return DataSourceSchema.
	//
	// If true, DataSourceSchema is returned.
	//
	// If false, DataSourceSchema is not returned.
	Verbose *bool `type:"boolean"`
}

// String returns the string representation
func (s GetDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDataSourceInput"}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}
	if s.DataSourceId != nil && len(*s.DataSourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataSourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetDataSource operation and describes a DataSource.
type GetDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The parameter is true if statistics need to be generated from the observation
	// data.
	ComputeStatistics *bool `type:"boolean"`

	// The approximate CPU time in milliseconds that Amazon Machine Learning spent
	// processing the DataSource, normalized and scaled on computation resources.
	// ComputeTime is only available if the DataSource is in the COMPLETED state
	// and the ComputeStatistics is set to true.
	ComputeTime *int64 `type:"long"`

	// The time that the DataSource was created. The time is expressed in epoch
	// time.
	CreatedAt *time.Time `type:"timestamp"`

	// The AWS user account from which the DataSource was created. The account type
	// can be either an AWS root account or an AWS Identity and Access Management
	// (IAM) user account.
	CreatedByIamUser *string `type:"string"`

	// The location of the data file or directory in Amazon Simple Storage Service
	// (Amazon S3).
	DataLocationS3 *string `type:"string"`

	// A JSON string that represents the splitting and rearrangement requirement
	// used when this DataSource was created.
	DataRearrangement *string `type:"string"`

	// The total size of observations in the data files.
	DataSizeInBytes *int64 `type:"long"`

	// The ID assigned to the DataSource at creation. This value should be identical
	// to the value of the DataSourceId in the request.
	DataSourceId *string `min:"1" type:"string"`

	// The schema used by all of the data files of this DataSource.
	//  Note
	// This parameter is provided as part of the verbose format.
	DataSourceSchema *string `type:"string"`

	// The epoch time when Amazon Machine Learning marked the DataSource as COMPLETED
	// or FAILED. FinishedAt is only available when the DataSource is in the COMPLETED
	// or FAILED state.
	FinishedAt *time.Time `type:"timestamp"`

	// The time of the most recent edit to the DataSource. The time is expressed
	// in epoch time.
	LastUpdatedAt *time.Time `type:"timestamp"`

	// A link to the file containing logs of CreateDataSourceFrom* operations.
	LogUri *string `type:"string"`

	// The user-supplied description of the most recent details about creating the
	// DataSource.
	Message *string `type:"string"`

	// A user-supplied name or description of the DataSource.
	Name *string `type:"string"`

	// The number of data files referenced by the DataSource.
	NumberOfFiles *int64 `type:"long"`

	// The datasource details that are specific to Amazon RDS.
	RDSMetadata *RDSMetadata `type:"structure"`

	// Describes the DataSource details specific to Amazon Redshift.
	RedshiftMetadata *RedshiftMetadata `type:"structure"`

	// The Amazon Resource Name (ARN) of an AWS IAM Role (http://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html#roles-about-termsandconcepts),
	// such as the following: arn:aws:iam::account:role/rolename.
	RoleARN *string `min:"1" type:"string"`

	// The epoch time when Amazon Machine Learning marked the DataSource as INPROGRESS.
	// StartedAt isn't available if the DataSource is in the PENDING state.
	StartedAt *time.Time `type:"timestamp"`

	// The current status of the DataSource. This element can have one of the following
	// values:
	//
	//    * PENDING - Amazon ML submitted a request to create a DataSource.
	//
	//    * INPROGRESS - The creation process is underway.
	//
	//    * FAILED - The request to create a DataSource did not run to completion.
	//    It is not usable.
	//
	//    * COMPLETED - The creation process completed successfully.
	//
	//    * DELETED - The DataSource is marked as deleted. It is not usable.
	Status EntityStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDataSource = "GetDataSource"

// GetDataSourceRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Returns a DataSource that includes metadata and data file information, as
// well as the current status of the DataSource.
//
// GetDataSource provides results in normal or verbose format. The verbose format
// adds the schema description and the list of files pointed to by the DataSource
// to the normal format.
//
//    // Example sending a request using GetDataSourceRequest.
//    req := client.GetDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetDataSourceRequest(input *GetDataSourceInput) GetDataSourceRequest {
	op := &aws.Operation{
		Name:       opGetDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDataSourceInput{}
	}

	req := c.newRequest(op, input, &GetDataSourceOutput{})
	return GetDataSourceRequest{Request: req, Input: input, Copy: c.GetDataSourceRequest}
}

// GetDataSourceRequest is the request type for the
// GetDataSource API operation.
type GetDataSourceRequest struct {
	*aws.Request
	Input *GetDataSourceInput
	Copy  func(*GetDataSourceInput) GetDataSourceRequest
}

// Send marshals and sends the GetDataSource API request.
func (r GetDataSourceRequest) Send(ctx context.Context) (*GetDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDataSourceResponse{
		GetDataSourceOutput: r.Request.Data.(*GetDataSourceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDataSourceResponse is the response type for the
// GetDataSource API operation.
type GetDataSourceResponse struct {
	*GetDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDataSource request.
func (r *GetDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
