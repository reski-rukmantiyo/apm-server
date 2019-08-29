// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package licenser

// CallbackWatcher defines an addhoc listener for events generated by the manager.
type CallbackWatcher struct {
	New     func(License)
	Stopped func()
}

// OnNewLicense is called when a new license is set in the manager.
func (cb *CallbackWatcher) OnNewLicense(license License) {
	if cb.New == nil {
		return
	}
	cb.New(license)
}

// OnManagerStopped is called when the manager is stopped, watcher are expected to terminates any
// features that depends on a specific license.
func (cb *CallbackWatcher) OnManagerStopped() {
	if cb.Stopped == nil {
		return
	}

	cb.Stopped()
}
