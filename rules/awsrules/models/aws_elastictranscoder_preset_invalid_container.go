// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsElastictranscoderPresetInvalidContainerRule checks the pattern is valid
type AwsElastictranscoderPresetInvalidContainerRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsElastictranscoderPresetInvalidContainerRule returns new rule with default attributes
func NewAwsElastictranscoderPresetInvalidContainerRule() *AwsElastictranscoderPresetInvalidContainerRule {
	return &AwsElastictranscoderPresetInvalidContainerRule{
		resourceType:  "aws_elastictranscoder_preset",
		attributeName: "container",
		pattern:       regexp.MustCompile(`^(^mp4$)|(^ts$)|(^webm$)|(^mp3$)|(^flac$)|(^oga$)|(^ogg$)|(^fmp4$)|(^mpg$)|(^flv$)|(^gif$)|(^mxf$)|(^wav$)|(^mp2$)$`),
	}
}

// Name returns the rule name
func (r *AwsElastictranscoderPresetInvalidContainerRule) Name() string {
	return "aws_elastictranscoder_preset_invalid_container"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastictranscoderPresetInvalidContainerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastictranscoderPresetInvalidContainerRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastictranscoderPresetInvalidContainerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastictranscoderPresetInvalidContainerRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^(^mp4$)|(^ts$)|(^webm$)|(^mp3$)|(^flac$)|(^oga$)|(^ogg$)|(^fmp4$)|(^mpg$)|(^flv$)|(^gif$)|(^mxf$)|(^wav$)|(^mp2$)$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
