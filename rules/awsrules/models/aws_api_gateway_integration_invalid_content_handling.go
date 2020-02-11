// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAPIGatewayIntegrationInvalidContentHandlingRule checks the pattern is valid
type AwsAPIGatewayIntegrationInvalidContentHandlingRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayIntegrationInvalidContentHandlingRule returns new rule with default attributes
func NewAwsAPIGatewayIntegrationInvalidContentHandlingRule() *AwsAPIGatewayIntegrationInvalidContentHandlingRule {
	return &AwsAPIGatewayIntegrationInvalidContentHandlingRule{
		resourceType:  "aws_api_gateway_integration",
		attributeName: "content_handling",
		enum: []string{
			"CONVERT_TO_BINARY",
			"CONVERT_TO_TEXT",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayIntegrationInvalidContentHandlingRule) Name() string {
	return "aws_api_gateway_integration_invalid_content_handling"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayIntegrationInvalidContentHandlingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayIntegrationInvalidContentHandlingRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayIntegrationInvalidContentHandlingRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayIntegrationInvalidContentHandlingRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as content_handling`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
