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
)

// In this example, we read a field from the input object and print it to the log.

func Example_aReadField() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(readField)); err != nil {
		os.Exit(1)
	}
}

func readField(rl *fn.ResourceList) (bool, error) {
	for _, obj := range rl.Items {
		if obj.GetAPIVersion() == "apps/v1" && obj.GetKind() == "Deployment" {
			var replicas int
			obj.GetOrDie(&replicas, "spec", "replicas")
			fn.Logf("replicas is %v\n", replicas)
		}
	}
	return true, nil
}