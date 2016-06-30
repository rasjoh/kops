// Code generated by ""fitask" -type=LoadBalancer"; DO NOT EDIT

package awstasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// LoadBalancer

// JSON marshalling boilerplate
type realLoadBalancer LoadBalancer

func (o *LoadBalancer) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realLoadBalancer
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = LoadBalancer(r)
	return nil
}

var _ fi.HasName = &LoadBalancer{}

func (e *LoadBalancer) GetName() *string {
	return e.Name
}

func (e *LoadBalancer) SetName(name string) {
	e.Name = &name
}

func (e *LoadBalancer) String() string {
	return fi.TaskAsString(e)
}
