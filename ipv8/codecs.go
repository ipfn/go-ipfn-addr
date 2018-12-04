// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
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

package ipv8

import (
	"github.com/ipfn/go-ipfn-codecs/codecs"
)

const (
	// BinaryCellV1 - Content ID of Binary Cell Version 1.
	BinaryCellV1 = 0x25ff2
	// FunctionV1 - Content ID of Public Kt Version 1.
	FunctionV1 = 0x70bc
	// IdentityV1 - Content ID of IPFN Identity Version 1.
	// It can contain a multihash cryptographic digest or ed25519 public key contents itself.
	IdentityV1 = 0x60ac
)

func init() {
	codecs.Register(map[string]uint64{
		"ipfn-cell-v1":     BinaryCellV1,
		"ipfn-func-v1":     FunctionV1,
		"ipfn-identity-v1": IdentityV1,
	})
}

// Protocol reservation:
// const (
// 	// FunctionV2 - Content ID of Public Kt Version 1. (name = "ipfn-func-v2", id = 45244)
// 	FunctionV2 = 0xb0bc
// )
