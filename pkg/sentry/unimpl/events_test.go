// Copyright 2023 The gVisor Authors.
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

// Package events_test verifies that kernel.Kernel implements interface unimpl.Events.
package events_test

import (
	"testing"

	"gvisor.dev/gvisor/pkg/sentry/kernel"
	"gvisor.dev/gvisor/pkg/sentry/unimpl"
)

// TestInterfaceMatch verifies that kernel.Kernel implements interface unimpl.Events.
func TestInterfaceMatch(t *testing.T) {
	var _ = (unimpl.Events)((*kernel.Kernel)(nil))
}