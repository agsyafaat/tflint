// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAPIGatewayGatewayResponseInvalidResponseTypeRule checks the pattern is valid
type AwsAPIGatewayGatewayResponseInvalidResponseTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayGatewayResponseInvalidResponseTypeRule returns new rule with default attributes
func NewAwsAPIGatewayGatewayResponseInvalidResponseTypeRule() *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule {
	return &AwsAPIGatewayGatewayResponseInvalidResponseTypeRule{
		resourceType:  "aws_api_gateway_gateway_response",
		attributeName: "response_type",
		enum: []string{
			"DEFAULT_4XX",
			"DEFAULT_5XX",
			"RESOURCE_NOT_FOUND",
			"UNAUTHORIZED",
			"INVALID_API_KEY",
			"ACCESS_DENIED",
			"AUTHORIZER_FAILURE",
			"AUTHORIZER_CONFIGURATION_ERROR",
			"INVALID_SIGNATURE",
			"EXPIRED_TOKEN",
			"MISSING_AUTHENTICATION_TOKEN",
			"INTEGRATION_FAILURE",
			"INTEGRATION_TIMEOUT",
			"API_CONFIGURATION_ERROR",
			"UNSUPPORTED_MEDIA_TYPE",
			"BAD_REQUEST_PARAMETERS",
			"BAD_REQUEST_BODY",
			"REQUEST_TOO_LARGE",
			"THROTTLED",
			"QUOTA_EXCEEDED",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Name() string {
	return "aws_api_gateway_gateway_response_invalid_response_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as response_type`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
