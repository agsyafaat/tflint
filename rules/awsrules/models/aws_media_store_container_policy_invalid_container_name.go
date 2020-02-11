// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsMediaStoreContainerPolicyInvalidContainerNameRule checks the pattern is valid
type AwsMediaStoreContainerPolicyInvalidContainerNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsMediaStoreContainerPolicyInvalidContainerNameRule returns new rule with default attributes
func NewAwsMediaStoreContainerPolicyInvalidContainerNameRule() *AwsMediaStoreContainerPolicyInvalidContainerNameRule {
	return &AwsMediaStoreContainerPolicyInvalidContainerNameRule{
		resourceType:  "aws_media_store_container_policy",
		attributeName: "container_name",
		max:           255,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w-]+$`),
	}
}

// Name returns the rule name
func (r *AwsMediaStoreContainerPolicyInvalidContainerNameRule) Name() string {
	return "aws_media_store_container_policy_invalid_container_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMediaStoreContainerPolicyInvalidContainerNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMediaStoreContainerPolicyInvalidContainerNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMediaStoreContainerPolicyInvalidContainerNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMediaStoreContainerPolicyInvalidContainerNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"container_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"container_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[\w-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
