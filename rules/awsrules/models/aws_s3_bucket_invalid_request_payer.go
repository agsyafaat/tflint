// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsS3BucketInvalidRequestPayerRule checks the pattern is valid
type AwsS3BucketInvalidRequestPayerRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketInvalidRequestPayerRule returns new rule with default attributes
func NewAwsS3BucketInvalidRequestPayerRule() *AwsS3BucketInvalidRequestPayerRule {
	return &AwsS3BucketInvalidRequestPayerRule{
		resourceType:  "aws_s3_bucket",
		attributeName: "request_payer",
		enum: []string{
			"Requester",
			"BucketOwner",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketInvalidRequestPayerRule) Name() string {
	return "aws_s3_bucket_invalid_request_payer"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketInvalidRequestPayerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3BucketInvalidRequestPayerRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketInvalidRequestPayerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketInvalidRequestPayerRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as request_payer`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
