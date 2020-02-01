// Copyright 2020 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package nstypes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Namespace Types", func() {

	It("namespace textual representations are parsed correctly", func() {
		id, t := IDwithType("net:[1]")
		Expect(t).To(Equal(CLONE_NEWNET))
		Expect(id).To(Equal(NamespaceID(1)))

		id, t = IDwithType("foo:[1]")
		Expect(t).To(Equal(NaNS))
		Expect(id).To(Equal(NoneID))

		id, t = IDwithType("net:[-1]")
		Expect(t).To(Equal(NaNS))
		Expect(id).To(Equal(NoneID))
	})

})