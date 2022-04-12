// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package example

import (
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	yaml2 "sigs.k8s.io/kustomize/kyaml/yaml"
)

// In this example, we convert the functionConfig as strong typed object and then
// read a field from the functionConfig object.

func Example_bReadFunctionConfig() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(readFunctionConfig)); err != nil {
		os.Exit(1)
	}
}

func readFunctionConfig(rl *fn.ResourceList) (bool, error) {
	var sr SetReplicas
	rl.FunctionConfig.AsOrDie(&sr)
	fn.Logf("desired replicas is %v\n", sr.DesiredReplicas)
	return true, nil
}

// SetReplicas is the type definition of the functionConfig
type SetReplicas struct {
	yaml2.ResourceIdentifier `json:",inline" yaml:",inline"`
	DesiredReplicas          int `json:"desiredReplicas,omitempty" yaml:"desiredReplicas,omitempty"`
}