// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLightsailInstanceInvalidBlueprintIDRule checks the pattern is valid
type AwsLightsailInstanceInvalidBlueprintIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLightsailInstanceInvalidBlueprintIDRule returns new rule with default attributes
func NewAwsLightsailInstanceInvalidBlueprintIDRule() *AwsLightsailInstanceInvalidBlueprintIDRule {
	return &AwsLightsailInstanceInvalidBlueprintIDRule{
		resourceType:  "aws_lightsail_instance",
		attributeName: "blueprint_id",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsLightsailInstanceInvalidBlueprintIDRule) Name() string {
	return "aws_lightsail_instance_invalid_blueprint_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLightsailInstanceInvalidBlueprintIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLightsailInstanceInvalidBlueprintIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLightsailInstanceInvalidBlueprintIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLightsailInstanceInvalidBlueprintIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^.*\S.*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
