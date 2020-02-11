// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCognitoUserPoolInvalidEmailVerificationSubjectRule checks the pattern is valid
type AwsCognitoUserPoolInvalidEmailVerificationSubjectRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolInvalidEmailVerificationSubjectRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidEmailVerificationSubjectRule() *AwsCognitoUserPoolInvalidEmailVerificationSubjectRule {
	return &AwsCognitoUserPoolInvalidEmailVerificationSubjectRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "email_verification_subject",
		max:           140,
		min:           1,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{S}\p{N}\p{P}\s]+$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidEmailVerificationSubjectRule) Name() string {
	return "aws_cognito_user_pool_invalid_email_verification_subject"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidEmailVerificationSubjectRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolInvalidEmailVerificationSubjectRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidEmailVerificationSubjectRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidEmailVerificationSubjectRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"email_verification_subject must be 140 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"email_verification_subject must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[\p{L}\p{M}\p{S}\p{N}\p{P}\s]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
