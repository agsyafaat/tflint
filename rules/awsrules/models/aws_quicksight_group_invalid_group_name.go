// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsQuicksightGroupInvalidGroupNameRule checks the pattern is valid
type AwsQuicksightGroupInvalidGroupNameRule struct {
	resourceType  string
	attributeName string
	min           int
	pattern       *regexp.Regexp
}

// NewAwsQuicksightGroupInvalidGroupNameRule returns new rule with default attributes
func NewAwsQuicksightGroupInvalidGroupNameRule() *AwsQuicksightGroupInvalidGroupNameRule {
	return &AwsQuicksightGroupInvalidGroupNameRule{
		resourceType:  "aws_quicksight_group",
		attributeName: "group_name",
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightGroupInvalidGroupNameRule) Name() string {
	return "aws_quicksight_group_invalid_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightGroupInvalidGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightGroupInvalidGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightGroupInvalidGroupNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightGroupInvalidGroupNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"group_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[\x{0020}-\x{00FF}]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
