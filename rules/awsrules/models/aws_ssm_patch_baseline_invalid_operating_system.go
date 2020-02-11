// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmPatchBaselineInvalidOperatingSystemRule checks the pattern is valid
type AwsSsmPatchBaselineInvalidOperatingSystemRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmPatchBaselineInvalidOperatingSystemRule returns new rule with default attributes
func NewAwsSsmPatchBaselineInvalidOperatingSystemRule() *AwsSsmPatchBaselineInvalidOperatingSystemRule {
	return &AwsSsmPatchBaselineInvalidOperatingSystemRule{
		resourceType:  "aws_ssm_patch_baseline",
		attributeName: "operating_system",
		enum: []string{
			"WINDOWS",
			"AMAZON_LINUX",
			"AMAZON_LINUX_2",
			"UBUNTU",
			"REDHAT_ENTERPRISE_LINUX",
			"SUSE",
			"CENTOS",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmPatchBaselineInvalidOperatingSystemRule) Name() string {
	return "aws_ssm_patch_baseline_invalid_operating_system"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmPatchBaselineInvalidOperatingSystemRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmPatchBaselineInvalidOperatingSystemRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmPatchBaselineInvalidOperatingSystemRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmPatchBaselineInvalidOperatingSystemRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as operating_system`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
