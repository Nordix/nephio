/*
 Copyright 2025 The Nephio Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 You may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"

	fnr "github.com/nephio-project/nephio/krm-functions/gen-configmap-fn/fn"
)

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(fnr.Process)); err != nil {
		os.Exit(1)
	}
}
