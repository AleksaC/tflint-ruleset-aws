// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMacie2ClassificationJobInvalidJobStatusRule checks the pattern is valid
type AwsMacie2ClassificationJobInvalidJobStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsMacie2ClassificationJobInvalidJobStatusRule returns new rule with default attributes
func NewAwsMacie2ClassificationJobInvalidJobStatusRule() *AwsMacie2ClassificationJobInvalidJobStatusRule {
	return &AwsMacie2ClassificationJobInvalidJobStatusRule{
		resourceType:  "aws_macie2_classification_job",
		attributeName: "job_status",
		enum: []string{
			"RUNNING",
			"PAUSED",
			"CANCELLED",
			"COMPLETE",
			"IDLE",
			"USER_PAUSED",
		},
	}
}

// Name returns the rule name
func (r *AwsMacie2ClassificationJobInvalidJobStatusRule) Name() string {
	return "aws_macie2_classification_job_invalid_job_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacie2ClassificationJobInvalidJobStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacie2ClassificationJobInvalidJobStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacie2ClassificationJobInvalidJobStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacie2ClassificationJobInvalidJobStatusRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as job_status`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}