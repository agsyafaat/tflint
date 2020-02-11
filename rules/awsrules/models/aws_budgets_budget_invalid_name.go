// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsBudgetsBudgetInvalidNameRule checks the pattern is valid
type AwsBudgetsBudgetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsBudgetsBudgetInvalidNameRule returns new rule with default attributes
func NewAwsBudgetsBudgetInvalidNameRule() *AwsBudgetsBudgetInvalidNameRule {
	return &AwsBudgetsBudgetInvalidNameRule{
		resourceType:  "aws_budgets_budget",
		attributeName: "name",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[^:\\]+$`),
	}
}

// Name returns the rule name
func (r *AwsBudgetsBudgetInvalidNameRule) Name() string {
	return "aws_budgets_budget_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBudgetsBudgetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBudgetsBudgetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBudgetsBudgetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBudgetsBudgetInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 100 characters or less",
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[^:\\]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
