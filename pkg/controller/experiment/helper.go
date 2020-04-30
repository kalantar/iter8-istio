/*

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
package experiment

import (
	"context"
	"strings"

	"github.com/iter8-tools/iter8-controller/pkg/controller/experiment/cache/abstract"
)

// interState controlls the execution of inter functions in the reconcile logic
type interState struct {
	statusUpdate bool
	refresh      bool
	progress     bool
}

func (r *ReconcileExperiment) initState() {
	r.interState = interState{}
}

func (r *ReconcileExperiment) markStatusUpdate() {
	r.interState.statusUpdate = true
}

func (r *ReconcileExperiment) needStatusUpdate() bool {
	return r.interState.statusUpdate
}

func (r *ReconcileExperiment) markRefresh() {
	r.interState.refresh = true
}

func (r *ReconcileExperiment) needRefresh() bool {
	return r.interState.refresh
}

func (r *ReconcileExperiment) markProgress() {
	r.interState.progress = true
}

func (r *ReconcileExperiment) hasProgress() bool {
	return r.interState.progress
}

func experimentAbstract(ctx context.Context) abstract.Snapshot {
	if ctx.Value(abstract.SnapshotKey) == nil {
		return nil
	}
	return ctx.Value(abstract.SnapshotKey).(abstract.Snapshot)
}

func validUpdateErr(err error) bool {
	if err == nil {
		return true
	}
	benignMsg := "the object has been modified"
	return strings.Contains(err.Error(), benignMsg)
}
