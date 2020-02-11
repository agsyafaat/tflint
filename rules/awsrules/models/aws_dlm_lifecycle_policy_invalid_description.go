// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDlmLifecyclePolicyInvalidDescriptionRule checks the pattern is valid
type AwsDlmLifecyclePolicyInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDlmLifecyclePolicyInvalidDescriptionRule returns new rule with default attributes
func NewAwsDlmLifecyclePolicyInvalidDescriptionRule() *AwsDlmLifecyclePolicyInvalidDescriptionRule {
	return &AwsDlmLifecyclePolicyInvalidDescriptionRule{
		resourceType:  "aws_dlm_lifecycle_policy",
		attributeName: "description",
		max:           500,
		pattern:       regexp.MustCompile(`^[0-9A-Za-z _-]+$`),
	}
}

// Name returns the rule name
func (r *AwsDlmLifecyclePolicyInvalidDescriptionRule) Name() string {
	return "aws_dlm_lifecycle_policy_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDlmLifecyclePolicyInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDlmLifecyclePolicyInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDlmLifecyclePolicyInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDlmLifecyclePolicyInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 500 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[0-9A-Za-z _-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
