//
// Copyright (c) 2014 The tm Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tm

import (
	"fmt"
	"time"
)

type TimeDuration struct {
	Duration int64 `json:"duration"`
	Counter  int64 `json:"count"`
}

func (d *TimeDuration) Add(delta time.Duration) {
	d.Duration += delta.Nanoseconds()
	d.Counter++
}

func (d *TimeDuration) MeanTimeUsecs() float64 {
	if d.Counter == 0 {
		return 0.0
	}
	return (float64(d.Duration) / float64(d.Counter)) / 1000.0
}

func (d *TimeDuration) Csv() string {
	return fmt.Sprintf("%v,"+
		"%v,",
		d.Duration,
		d.Counter)
}

func (d *TimeDuration) DeltaMeanTimeUsecs(prev *TimeDuration) float64 {
	delta := TimeDuration{}
	delta.Duration = d.Duration - prev.Duration
	delta.Counter = d.Counter - prev.Counter
	return delta.MeanTimeUsecs()
}

func (d *TimeDuration) Copy() *TimeDuration {
	tdcopy := &TimeDuration{}
	*tdcopy = *d
	return tdcopy
}

func (d *TimeDuration) String() string {
	return fmt.Sprintf("Duration = %v\n"+
		"Counter = %v\n",
		d.Duration,
		d.Counter)
}
