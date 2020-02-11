// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule checks the pattern is valid
type AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule returns new rule with default attributes
func NewAwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule() *AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule {
	return &AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule{
		resourceType:  "aws_cloudwatch_log_destination_policy",
		attributeName: "destination_name",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^[^:*]*$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule) Name() string {
	return "aws_cloudwatch_log_destination_policy_invalid_destination_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"destination_name must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"destination_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[^:*]*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
