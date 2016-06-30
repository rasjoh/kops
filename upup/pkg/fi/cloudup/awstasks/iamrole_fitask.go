// Code generated by ""fitask" -type=IAMRole"; DO NOT EDIT

package awstasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// IAMRole

// JSON marshalling boilerplate
type realIAMRole IAMRole

func (o *IAMRole) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realIAMRole
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = IAMRole(r)
	return nil
}

var _ fi.HasName = &IAMRole{}

func (e *IAMRole) GetName() *string {
	return e.Name
}

func (e *IAMRole) SetName(name string) {
	e.Name = &name
}

func (e *IAMRole) String() string {
	return fi.TaskAsString(e)
}
