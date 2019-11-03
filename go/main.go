package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/api/compute/v1"
)

const (
	instanceName = "vm-by-go"
	zone         = "asia-east1-b"
)

func main() {
	fmt.Println("Create VM by Go library")

	projectID := os.Args[1]
	filePath := os.Args[2]

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	instance := &compute.Instance{}
	err = json.Unmarshal(b, instance)
	if err != nil {
		panic(err)
	}
	instance.Name = instanceName

	ctx := context.Background()
	service, err := compute.NewService(ctx)
	if err != nil {
		panic(err)
	}
	op, err := service.Instances.Insert(projectID, zone, instance).Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("operation: %v\n", op.Id)
}
