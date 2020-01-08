// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetAccountSettingsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type GetAccountSettingsOutput struct {
	_ struct{} `type:"structure"`

	// Limits that are related to concurrency and code storage.
	AccountLimit *AccountLimit `type:"structure"`

	// The number of functions and amount of storage in use.
	AccountUsage *AccountUsage `type:"structure"`
}

// String returns the string representation
func (s GetAccountSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAccountSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccountLimit != nil {
		v := s.AccountLimit

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "AccountLimit", v, metadata)
	}
	if s.AccountUsage != nil {
		v := s.AccountUsage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "AccountUsage", v, metadata)
	}
	return nil
}

const opGetAccountSettings = "GetAccountSettings"

// GetAccountSettingsRequest returns a request value for making API operation for
// AWS Lambda.
//
// Retrieves details about your account's limits (https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
// and usage in an AWS Region.
//
//    // Example sending a request using GetAccountSettingsRequest.
//    req := client.GetAccountSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetAccountSettings
func (c *Client) GetAccountSettingsRequest(input *GetAccountSettingsInput) GetAccountSettingsRequest {
	op := &aws.Operation{
		Name:       opGetAccountSettings,
		HTTPMethod: "GET",
		HTTPPath:   "/2016-08-19/account-settings/",
	}

	if input == nil {
		input = &GetAccountSettingsInput{}
	}

	req := c.newRequest(op, input, &GetAccountSettingsOutput{})
	return GetAccountSettingsRequest{Request: req, Input: input, Copy: c.GetAccountSettingsRequest}
}

// GetAccountSettingsRequest is the request type for the
// GetAccountSettings API operation.
type GetAccountSettingsRequest struct {
	*aws.Request
	Input *GetAccountSettingsInput
	Copy  func(*GetAccountSettingsInput) GetAccountSettingsRequest
}

// Send marshals and sends the GetAccountSettings API request.
func (r GetAccountSettingsRequest) Send(ctx context.Context) (*GetAccountSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountSettingsResponse{
		GetAccountSettingsOutput: r.Request.Data.(*GetAccountSettingsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountSettingsResponse is the response type for the
// GetAccountSettings API operation.
type GetAccountSettingsResponse struct {
	*GetAccountSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccountSettings request.
func (r *GetAccountSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
