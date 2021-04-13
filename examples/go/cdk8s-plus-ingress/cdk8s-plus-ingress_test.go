package main

import (
	"testing"

	"github.com/awslabs/cdk8s-go/cdk8s"
	"github.com/stretchr/testify/assert"
)

func TestCdk8SPlusIngressChart(t *testing.T) {
	// GIVEN
	app := cdk8s.NewApp(nil)

	// WHEN
	chart := NewCdk8SPlusIngressChart(app, "MyChart", nil)

	// THEN
	manifest := *chart.ToJson()
	ingress, ok := manifest[0].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "Ingress", ingress["kind"])
}
