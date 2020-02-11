// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAPIGatewayIntegrationInvalidConnectionTypeRule checks the pattern is valid
type AwsAPIGatewayIntegrationInvalidConnectionTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayIntegrationInvalidConnectionTypeRule returns new rule with default attributes
func NewAwsAPIGatewayIntegrationInvalidConnectionTypeRule() *AwsAPIGatewayIntegrationInvalidConnectionTypeRule {
	return &AwsAPIGatewayIntegrationInvalidConnectionTypeRule{
		resourceType:  "aws_api_gateway_integration",
		attributeName: "connection_type",
		enum: []string{
			"INTERNET",
			"VPC_LINK",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Name() string {
	return "aws_api_gateway_integration_invalid_connection_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayIntegrationInvalidConnectionTypeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as connection_type`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
