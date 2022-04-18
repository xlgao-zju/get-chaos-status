package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	c           client.Client
	someIndexer client.FieldIndexer
)

type ChaosStatus struct {
	Status v1alpha1.ChaosStatus `json:"status"`
}

func ExampleNew() {
	cl, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	// Using an unstructured object.
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "chaos-mesh.org",
		Kind:    "TimeChaos",
		Version: "v1alpha1",
	})

	err = cl.Get(context.Background(), client.ObjectKey{Name: "test-time", Namespace: "xlgao"}, u)
	if err != nil {
		fmt.Printf("failed to list pods in namespace default: %v\n", err)
		os.Exit(1)
	}
	obj, err := u.MarshalJSON()
	status := &ChaosStatus{}
	err = json.Unmarshal(obj, status)
	fmt.Printf("%#v", *status)
}

func main() {
	ExampleNew()
}
