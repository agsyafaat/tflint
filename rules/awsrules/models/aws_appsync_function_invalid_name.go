// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAppsyncFunctionInvalidNameRule checks the pattern is valid
type AwsAppsyncFunctionInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAppsyncFunctionInvalidNameRule returns new rule with default attributes
func NewAwsAppsyncFunctionInvalidNameRule() *AwsAppsyncFunctionInvalidNameRule {
	return &AwsAppsyncFunctionInvalidNameRule{
		resourceType:  "aws_appsync_function",
		attributeName: "name",
		max:           65536,
		min:           1,
		pattern:       regexp.MustCompile(`^[_A-Za-z][_0-9A-Za-z]*$`),
	}
}

// Name returns the rule name
func (r *AwsAppsyncFunctionInvalidNameRule) Name() string {
	return "aws_appsync_function_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncFunctionInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncFunctionInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncFunctionInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncFunctionInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 65536 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[_A-Za-z][_0-9A-Za-z]*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
