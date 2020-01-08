// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to describe one or more applications.
type DescribeApplicationsInput struct {
	_ struct{} `type:"structure"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// only include those with the specified names.
	ApplicationNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Result message containing a list of application descriptions.
type DescribeApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// This parameter contains a list of ApplicationDescription.
	Applications []ApplicationDescription `type:"list"`
}

// String returns the string representation
func (s DescribeApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeApplications = "DescribeApplications"

// DescribeApplicationsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Returns the descriptions of existing applications.
//
//    // Example sending a request using DescribeApplicationsRequest.
//    req := client.DescribeApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeApplications
func (c *Client) DescribeApplicationsRequest(input *DescribeApplicationsInput) DescribeApplicationsRequest {
	op := &aws.Operation{
		Name:       opDescribeApplications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeApplicationsInput{}
	}

	req := c.newRequest(op, input, &DescribeApplicationsOutput{})
	return DescribeApplicationsRequest{Request: req, Input: input, Copy: c.DescribeApplicationsRequest}
}

// DescribeApplicationsRequest is the request type for the
// DescribeApplications API operation.
type DescribeApplicationsRequest struct {
	*aws.Request
	Input *DescribeApplicationsInput
	Copy  func(*DescribeApplicationsInput) DescribeApplicationsRequest
}

// Send marshals and sends the DescribeApplications API request.
func (r DescribeApplicationsRequest) Send(ctx context.Context) (*DescribeApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeApplicationsResponse{
		DescribeApplicationsOutput: r.Request.Data.(*DescribeApplicationsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeApplicationsResponse is the response type for the
// DescribeApplications API operation.
type DescribeApplicationsResponse struct {
	*DescribeApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeApplications request.
func (r *DescribeApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
