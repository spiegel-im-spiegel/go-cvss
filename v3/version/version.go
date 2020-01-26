package version

import "fmt"

//Num is error number for CVSS
type Num int

const (
	Unknown Num = iota //unknown version
	V3_0               //v3.0
	V3_1               //v3.1
)

var verStrings = map[Num]string{
	V3_0: "3.0",
	V3_1: "3.1",
}

//String is Stringer method
func (n Num) String() string {
	if s, ok := verStrings[n]; ok {
		return s
	}
	return fmt.Sprintf("unknown")
}

//Get returns Version number from string
func Get(s string) Num {
	for k, v := range verStrings {
		if s == v {
			return k
		}
	}
	return Unknown
}

/* Copyright 2018-2020 Spiegel
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
