// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type RemoveRoleFromDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster to disassociate the IAM role from.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the feature for the DB cluster that the IAM role is to be disassociated
	// from. For the list of supported feature names, see DBEngineVersion.
	FeatureName *string `type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role to disassociate from the Aurora
	// DB cluster, for example arn:aws:iam::123456789012:role/AuroraAccessRole.
	//
	// RoleArn is a required field
	RoleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveRoleFromDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveRoleFromDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveRoleFromDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveRoleFromDBClusterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveRoleFromDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveRoleFromDBCluster = "RemoveRoleFromDBCluster"

// RemoveRoleFromDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Disassociates an AWS Identity and Access Management (IAM) role from an Amazon
// Aurora DB cluster. For more information, see Authorizing Amazon Aurora MySQL
// to Access Other AWS Services on Your Behalf (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using RemoveRoleFromDBClusterRequest.
//    req := client.RemoveRoleFromDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RemoveRoleFromDBCluster
func (c *Client) RemoveRoleFromDBClusterRequest(input *RemoveRoleFromDBClusterInput) RemoveRoleFromDBClusterRequest {
	op := &aws.Operation{
		Name:       opRemoveRoleFromDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveRoleFromDBClusterInput{}
	}

	req := c.newRequest(op, input, &RemoveRoleFromDBClusterOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveRoleFromDBClusterRequest{Request: req, Input: input, Copy: c.RemoveRoleFromDBClusterRequest}
}

// RemoveRoleFromDBClusterRequest is the request type for the
// RemoveRoleFromDBCluster API operation.
type RemoveRoleFromDBClusterRequest struct {
	*aws.Request
	Input *RemoveRoleFromDBClusterInput
	Copy  func(*RemoveRoleFromDBClusterInput) RemoveRoleFromDBClusterRequest
}

// Send marshals and sends the RemoveRoleFromDBCluster API request.
func (r RemoveRoleFromDBClusterRequest) Send(ctx context.Context) (*RemoveRoleFromDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveRoleFromDBClusterResponse{
		RemoveRoleFromDBClusterOutput: r.Request.Data.(*RemoveRoleFromDBClusterOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveRoleFromDBClusterResponse is the response type for the
// RemoveRoleFromDBCluster API operation.
type RemoveRoleFromDBClusterResponse struct {
	*RemoveRoleFromDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveRoleFromDBCluster request.
func (r *RemoveRoleFromDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
