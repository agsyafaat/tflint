// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsElasticsearchDomainInvalidDomainNameRule checks the pattern is valid
type AwsElasticsearchDomainInvalidDomainNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsElasticsearchDomainInvalidDomainNameRule returns new rule with default attributes
func NewAwsElasticsearchDomainInvalidDomainNameRule() *AwsElasticsearchDomainInvalidDomainNameRule {
	return &AwsElasticsearchDomainInvalidDomainNameRule{
		resourceType:  "aws_elasticsearch_domain",
		attributeName: "domain_name",
		max:           28,
		min:           3,
		pattern:       regexp.MustCompile(`^[a-z][a-z0-9\-]+$`),
	}
}

// Name returns the rule name
func (r *AwsElasticsearchDomainInvalidDomainNameRule) Name() string {
	return "aws_elasticsearch_domain_invalid_domain_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticsearchDomainInvalidDomainNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticsearchDomainInvalidDomainNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticsearchDomainInvalidDomainNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticsearchDomainInvalidDomainNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"domain_name must be 28 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"domain_name must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[a-z][a-z0-9\-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
