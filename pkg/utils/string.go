/*
  Copyright 2021 The Rama Authors.

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

package utils

import (
	"github.com/gogf/gf/container/gset"
	jsoniter "github.com/json-iterator/go"
)

func PickFirstNonEmptyString(ss ...string) string {
	for _, s := range ss {
		if len(s) > 0 {
			return s
		}
	}
	return ""
}

func ToJSONString(i interface{}) string {
	s, _ := jsoniter.MarshalToString(i)
	return s
}

// DifferentSetFromStringSlice returns if two string slice contains different item
// True for two different string slice
func DifferentSetFromStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aSet := gset.NewStrSetFrom(a)
	bSet := gset.NewStrSetFrom(b)
	return !aSet.Equal(bSet)
}
