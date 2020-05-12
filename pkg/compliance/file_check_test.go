// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.
package compliance

import (
	"testing"
	"time"

	"github.com/DataDog/datadog-agent/pkg/collector/check"
	"github.com/stretchr/testify/assert"
)

func TestFileCheck(t *testing.T) {
	reporter := &MockReporter{}

	const (
		framework    = "cis-docker"
		version      = "1.2.0"
		ruleID       = "rule"
		resourceID   = "host"
		resourceType = "docker"
	)

	fc := &fileCheck{
		baseCheck: baseCheck{
			id:        check.ID("check-1"),
			interval:  time.Minute,
			framework: framework,
			version:   version,

			ruleID:       ruleID,
			resourceType: resourceType,
			resourceID:   resourceID,
			reporter:     reporter,
		},
		File: &File{
			Path: "./testdata/644.dat",
			Report: Report{
				{
					Attribute: "permissions",
				},
			},
		},
	}
	reporter.On(
		"Report",
		&RuleEvent{
			RuleID:       ruleID,
			Framework:    framework,
			Version:      version,
			ResourceType: resourceType,
			ResourceID:   resourceID,
			Data: KV{
				"permissions": "644",
			},
		},
	).Once()

	err := fc.Run()
	assert.NoError(t, err)
}
