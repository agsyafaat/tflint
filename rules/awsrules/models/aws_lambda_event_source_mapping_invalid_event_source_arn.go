// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLambdaEventSourceMappingInvalidEventSourceArnRule checks the pattern is valid
type AwsLambdaEventSourceMappingInvalidEventSourceArnRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLambdaEventSourceMappingInvalidEventSourceArnRule returns new rule with default attributes
func NewAwsLambdaEventSourceMappingInvalidEventSourceArnRule() *AwsLambdaEventSourceMappingInvalidEventSourceArnRule {
	return &AwsLambdaEventSourceMappingInvalidEventSourceArnRule{
		resourceType:  "aws_lambda_event_source_mapping",
		attributeName: "event_source_arn",
		pattern:       regexp.MustCompile(`^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:([a-z]{2}(-gov)?-[a-z]+-\d{1})?:(\d{12})?:(.*)$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaEventSourceMappingInvalidEventSourceArnRule) Name() string {
	return "aws_lambda_event_source_mapping_invalid_event_source_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaEventSourceMappingInvalidEventSourceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaEventSourceMappingInvalidEventSourceArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaEventSourceMappingInvalidEventSourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaEventSourceMappingInvalidEventSourceArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:([a-z]{2}(-gov)?-[a-z]+-\d{1})?:(\d{12})?:(.*)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
