// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsFlowLogInvalidTrafficTypeRule checks the pattern is valid
type AwsFlowLogInvalidTrafficTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsFlowLogInvalidTrafficTypeRule returns new rule with default attributes
func NewAwsFlowLogInvalidTrafficTypeRule() *AwsFlowLogInvalidTrafficTypeRule {
	return &AwsFlowLogInvalidTrafficTypeRule{
		resourceType:  "aws_flow_log",
		attributeName: "traffic_type",
		enum: []string{
			"ACCEPT",
			"REJECT",
			"ALL",
		},
	}
}

// Name returns the rule name
func (r *AwsFlowLogInvalidTrafficTypeRule) Name() string {
	return "aws_flow_log_invalid_traffic_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFlowLogInvalidTrafficTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFlowLogInvalidTrafficTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFlowLogInvalidTrafficTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFlowLogInvalidTrafficTypeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as traffic_type`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
