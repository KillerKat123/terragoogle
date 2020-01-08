// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SetInstanceProtectionInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// One or more instance IDs.
	//
	// InstanceIds is a required field
	InstanceIds []string `type:"list" required:"true"`

	// Indicates whether the instance is protected from termination by Amazon EC2
	// Auto Scaling when scaling in.
	//
	// ProtectedFromScaleIn is a required field
	ProtectedFromScaleIn *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s SetInstanceProtectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetInstanceProtectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetInstanceProtectionInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.InstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceIds"))
	}

	if s.ProtectedFromScaleIn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProtectedFromScaleIn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetInstanceProtectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetInstanceProtectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetInstanceProtection = "SetInstanceProtection"

// SetInstanceProtectionRequest returns a request value for making API operation for
// Auto Scaling.
//
// Updates the instance protection settings of the specified instances.
//
// For more information about preventing instances that are part of an Auto
// Scaling group from terminating on scale in, see Instance Protection (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using SetInstanceProtectionRequest.
//    req := client.SetInstanceProtectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/SetInstanceProtection
func (c *Client) SetInstanceProtectionRequest(input *SetInstanceProtectionInput) SetInstanceProtectionRequest {
	op := &aws.Operation{
		Name:       opSetInstanceProtection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetInstanceProtectionInput{}
	}

	req := c.newRequest(op, input, &SetInstanceProtectionOutput{})
	return SetInstanceProtectionRequest{Request: req, Input: input, Copy: c.SetInstanceProtectionRequest}
}

// SetInstanceProtectionRequest is the request type for the
// SetInstanceProtection API operation.
type SetInstanceProtectionRequest struct {
	*aws.Request
	Input *SetInstanceProtectionInput
	Copy  func(*SetInstanceProtectionInput) SetInstanceProtectionRequest
}

// Send marshals and sends the SetInstanceProtection API request.
func (r SetInstanceProtectionRequest) Send(ctx context.Context) (*SetInstanceProtectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetInstanceProtectionResponse{
		SetInstanceProtectionOutput: r.Request.Data.(*SetInstanceProtectionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetInstanceProtectionResponse is the response type for the
// SetInstanceProtection API operation.
type SetInstanceProtectionResponse struct {
	*SetInstanceProtectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetInstanceProtection request.
func (r *SetInstanceProtectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
