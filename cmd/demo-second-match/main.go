// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"open-match.dev/open-match/examples/demo-dropin"
	"open-match.dev/open-match/examples/demo-dropin/components"
	"open-match.dev/open-match/examples/demo-dropin/components/clients"
	"open-match.dev/open-match/examples/demo-dropin/components/director"
	"open-match.dev/open-match/examples/demo-dropin/components/uptime"
)

func main() {
	demo.Run(map[string]func(*components.DemoShared){
		"uptime":   uptime.Run,
		"clients":  clients.Run,
		"director": director.Run,
	})
}
