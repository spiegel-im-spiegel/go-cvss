package cvss

import (
	"github.com/spiegel-im-spiegel/go-cvss/v3/metric"
)

//ImportBase creates new metric.Base instance from CVSSv3 vector string.
func ImportBase(vector string) (*metric.Base, error) {
	return metric.NewBase().Decode(vector)
}

//ImportTemporal creates new metric.Temporal instance from CVSSv3 vector string.
func ImportTemporal(vector string) (*metric.Temporal, error) {
	return metric.NewTemporal().Decode(vector)
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */