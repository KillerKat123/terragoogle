// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AcceptTransitGatewayVpcAttachmentInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptTransitGatewayVpcAttachmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptTransitGatewayVpcAttachmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptTransitGatewayVpcAttachmentInput"}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AcceptTransitGatewayVpcAttachmentOutput struct {
	_ struct{} `type:"structure"`

	// The VPC attachment.
	TransitGatewayVpcAttachment *TransitGatewayVpcAttachment `locationName:"transitGatewayVpcAttachment" type:"structure"`
}

// String returns the string representation
func (s AcceptTransitGatewayVpcAttachmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptTransitGatewayVpcAttachment = "AcceptTransitGatewayVpcAttachment"

// AcceptTransitGatewayVpcAttachmentRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Accepts a request to attach a VPC to a transit gateway.
//
// The VPC attachment must be in the pendingAcceptance state. Use DescribeTransitGatewayVpcAttachments
// to view your pending VPC attachment requests. Use RejectTransitGatewayVpcAttachment
// to reject a VPC attachment request.
//
//    // Example sending a request using AcceptTransitGatewayVpcAttachmentRequest.
//    req := client.AcceptTransitGatewayVpcAttachmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AcceptTransitGatewayVpcAttachment
func (c *Client) AcceptTransitGatewayVpcAttachmentRequest(input *AcceptTransitGatewayVpcAttachmentInput) AcceptTransitGatewayVpcAttachmentRequest {
	op := &aws.Operation{
		Name:       opAcceptTransitGatewayVpcAttachment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptTransitGatewayVpcAttachmentInput{}
	}

	req := c.newRequest(op, input, &AcceptTransitGatewayVpcAttachmentOutput{})
	return AcceptTransitGatewayVpcAttachmentRequest{Request: req, Input: input, Copy: c.AcceptTransitGatewayVpcAttachmentRequest}
}

// AcceptTransitGatewayVpcAttachmentRequest is the request type for the
// AcceptTransitGatewayVpcAttachment API operation.
type AcceptTransitGatewayVpcAttachmentRequest struct {
	*aws.Request
	Input *AcceptTransitGatewayVpcAttachmentInput
	Copy  func(*AcceptTransitGatewayVpcAttachmentInput) AcceptTransitGatewayVpcAttachmentRequest
}

// Send marshals and sends the AcceptTransitGatewayVpcAttachment API request.
func (r AcceptTransitGatewayVpcAttachmentRequest) Send(ctx context.Context) (*AcceptTransitGatewayVpcAttachmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptTransitGatewayVpcAttachmentResponse{
		AcceptTransitGatewayVpcAttachmentOutput: r.Request.Data.(*AcceptTransitGatewayVpcAttachmentOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptTransitGatewayVpcAttachmentResponse is the response type for the
// AcceptTransitGatewayVpcAttachment API operation.
type AcceptTransitGatewayVpcAttachmentResponse struct {
	*AcceptTransitGatewayVpcAttachmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptTransitGatewayVpcAttachment request.
func (r *AcceptTransitGatewayVpcAttachmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
