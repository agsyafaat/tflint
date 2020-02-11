// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIAMSamlProviderInvalidNameRule checks the pattern is valid
type AwsIAMSamlProviderInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMSamlProviderInvalidNameRule returns new rule with default attributes
func NewAwsIAMSamlProviderInvalidNameRule() *AwsIAMSamlProviderInvalidNameRule {
	return &AwsIAMSamlProviderInvalidNameRule{
		resourceType:  "aws_iam_saml_provider",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w._-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMSamlProviderInvalidNameRule) Name() string {
	return "aws_iam_saml_provider_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMSamlProviderInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMSamlProviderInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMSamlProviderInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMSamlProviderInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[\w._-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
