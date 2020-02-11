// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule checks the pattern is valid
type AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule returns new rule with default attributes
func NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule() *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule {
	return &AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule{
		resourceType:  "aws_ec2_capacity_reservation",
		attributeName: "instance_match_criteria",
		enum: []string{
			"open",
			"targeted",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Name() string {
	return "aws_ec2_capacity_reservation_invalid_instance_match_criteria"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as instance_match_criteria`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
