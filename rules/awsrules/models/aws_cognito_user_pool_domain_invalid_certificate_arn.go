// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCognitoUserPoolDomainInvalidCertificateArnRule checks the pattern is valid
type AwsCognitoUserPoolDomainInvalidCertificateArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolDomainInvalidCertificateArnRule returns new rule with default attributes
func NewAwsCognitoUserPoolDomainInvalidCertificateArnRule() *AwsCognitoUserPoolDomainInvalidCertificateArnRule {
	return &AwsCognitoUserPoolDomainInvalidCertificateArnRule{
		resourceType:  "aws_cognito_user_pool_domain",
		attributeName: "certificate_arn",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolDomainInvalidCertificateArnRule) Name() string {
	return "aws_cognito_user_pool_domain_invalid_certificate_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolDomainInvalidCertificateArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolDomainInvalidCertificateArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolDomainInvalidCertificateArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolDomainInvalidCertificateArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"certificate_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"certificate_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
