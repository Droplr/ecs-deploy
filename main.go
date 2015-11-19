package main

import (
	"fmt"
	
	"github.com/travisjeffery/deploy/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	service = kingpin.Flag("service", "Set service name to update.").Required().String()
	image = kingpin.Flag("image", "Set image to update the service to.").Required().String()
	cluster = kingpin.Flag("cluster", "Set cluster name with the service to update.").Default("default").String()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0").Author("Travis Jeffery")
	kingpin.CommandLine.Help = "Update ECS service."
	kingpin.Parse()
	c := client.New()

	arn, err := c.RegisterTaskDefinition(*service, *image)
	if err != nil {
		fmt.Printf("register task definition error: %s\n", err.Error())
		return
	}

	err = c.UpdateService(*cluster, *service, arn)
	if err != nil {
		fmt.Printf("update service error: %s\n", err.Error())
		return
	}
	fmt.Printf("update service success")
}