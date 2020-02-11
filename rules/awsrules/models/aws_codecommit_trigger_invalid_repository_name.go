// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCodecommitTriggerInvalidRepositoryNameRule checks the pattern is valid
type AwsCodecommitTriggerInvalidRepositoryNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodecommitTriggerInvalidRepositoryNameRule returns new rule with default attributes
func NewAwsCodecommitTriggerInvalidRepositoryNameRule() *AwsCodecommitTriggerInvalidRepositoryNameRule {
	return &AwsCodecommitTriggerInvalidRepositoryNameRule{
		resourceType:  "aws_codecommit_trigger",
		attributeName: "repository_name",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w\.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsCodecommitTriggerInvalidRepositoryNameRule) Name() string {
	return "aws_codecommit_trigger_invalid_repository_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodecommitTriggerInvalidRepositoryNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodecommitTriggerInvalidRepositoryNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodecommitTriggerInvalidRepositoryNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodecommitTriggerInvalidRepositoryNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"repository_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"repository_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[\w\.-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
