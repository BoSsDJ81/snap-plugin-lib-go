/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import "fmt"

// floatRule defines a type to contain Float specific rule data.
type floatRule struct {
	Key string

	Default    float64
	HasDefault bool

	Required bool

	Minimum float64
	HasMin  bool

	Maximum float64
	HasMax  bool
}

type floatRuleOpt func(*floatRule)

// SetDefaultFloat Allows easy setting of the Default value for an floatRule.
// Usage:
//		NewFloatRule(key, req, config.SetDefaultFloat(default))
func SetDefaultFloat(in float64) floatRuleOpt {
	return func(i *floatRule) {
		i.Default = in
		i.HasDefault = true
	}
}

// SetMaxFloat Allows easy setting of the Max value for an floatRule.
// Usage:
//		NewFloatRule(key, req, config.SetMaxFloat(max))
func SetMaxFloat(max float64) floatRuleOpt {
	return func(i *floatRule) {
		i.Maximum = max
		i.HasMax = true
	}
}

// SetMinFloat Allows easy setting of the Min value for an floatRule.
// Usage:
//		NewFloatRule(key, req, config.SetMinFloat(min))
func SetMinFloat(min float64) floatRuleOpt {
	return func(i *floatRule) {
		i.Minimum = min
		i.HasMin = true
	}
}

// NewFloatRule returns a new floatRule with the specified args.
// The required arguments are key(string), req(bool)
// and optionally:
//		config.SetDefaultFloat(float64),
//		config.SetMinFloat(float64),
//		config.SetMaxFloat(float64),
func NewFloatRule(key string, req bool, opts ...floatRuleOpt) (floatRule, error) {
	if key == "" {
		return floatRule{}, fmt.Errorf(errEmptyKey)
	}
	rule := floatRule{
		Key:      key,
		Required: req,
	}

	for _, opt := range opts {
		opt(&rule)
	}

	return rule, nil
}
