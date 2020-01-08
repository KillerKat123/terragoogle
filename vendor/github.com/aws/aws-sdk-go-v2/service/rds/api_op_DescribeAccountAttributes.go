// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAccountAttributesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeAccountAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Data returned by the DescribeAccountAttributes action.
type DescribeAccountAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A list of AccountQuota objects. Within this list, each quota has a name,
	// a count of usage toward the quota maximum, and a maximum value for the quota.
	AccountQuotas []AccountQuota `locationNameList:"AccountQuota" type:"list"`
}

// String returns the string representation
func (s DescribeAccountAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountAttributes = "DescribeAccountAttributes"

// DescribeAccountAttributesRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Lists all of the attributes for a customer account. The attributes include
// Amazon RDS quotas for the account, such as the number of DB instances allowed.
// The description for a quota includes the quota name, current usage toward
// that quota, and the quota's maximum value.
//
// This command doesn't take any parameters.
//
//    // Example sending a request using DescribeAccountAttributesRequest.
//    req := client.DescribeAccountAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeAccountAttributes
func (c *Client) DescribeAccountAttributesRequest(input *DescribeAccountAttributesInput) DescribeAccountAttributesRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountAttributesInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountAttributesOutput{})
	return DescribeAccountAttributesRequest{Request: req, Input: input, Copy: c.DescribeAccountAttributesRequest}
}

// DescribeAccountAttributesRequest is the request type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesRequest struct {
	*aws.Request
	Input *DescribeAccountAttributesInput
	Copy  func(*DescribeAccountAttributesInput) DescribeAccountAttributesRequest
}

// Send marshals and sends the DescribeAccountAttributes API request.
func (r DescribeAccountAttributesRequest) Send(ctx context.Context) (*DescribeAccountAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountAttributesResponse{
		DescribeAccountAttributesOutput: r.Request.Data.(*DescribeAccountAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountAttributesResponse is the response type for the
// DescribeAccountAttributes API operation.
type DescribeAccountAttributesResponse struct {
	*DescribeAccountAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountAttributes request.
func (r *DescribeAccountAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
