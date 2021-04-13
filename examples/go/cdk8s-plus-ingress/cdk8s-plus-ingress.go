package main

import (
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/awslabs/cdk8s-go/cdk8s"
	"github.com/awslabs/cdk8s-go/cdk8splus17"
	kplus "github.com/awslabs/cdk8s-go/cdk8splus17"
)

type Cdk8SPlusIngressChartProps struct {
	cdk8s.ChartProps
}

func NewCdk8SPlusIngressChart(scope constructs.Construct, id string, props *Cdk8SPlusIngressChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, &id, &cprops)


	ingress := kplus.NewIngressV1Beta1(chart, jsii.String("ingress"), nil);
	ingress.AddRule(jsii.String("/"), echoBackend(chart, jsii.String("root")));
	ingress.AddRule(jsii.String("/foo"), echoBackend(chart, jsii.String("foo")));
	ingress.AddRule(jsii.String("/foo/bar"), echoBackend(chart, jsii.String("foo-bar")));

	ingress.AddHostDefaultBackend(jsii.String("my.host"), echoBackend(chart, jsii.String("my.host/hey")));

	return chart
}

func echoBackend(chart cdk8s.Chart, text *string) cdk8splus17.IngressV1Beta1Backend {
	containers := []*cdk8splus17.ContainerProps{
		{
			Image: jsii.String("hashicorp/http-echo"),
			Args: jsii.Strings("-text", *text),
		},
	}
	deployment := kplus.NewDeployment(chart, text, &kplus.DeploymentProps{
		Containers: &containers,
	})

	service := deployment.Expose(jsii.Number(5678), nil)

	return kplus.IngressV1Beta1Backend_FromService(service, nil);
}

func main() {
	app := cdk8s.NewApp(nil)

	NewCdk8SPlusIngressChart(app, "Cdk8SPlusIngressStack", nil)

	app.Synth()
}
