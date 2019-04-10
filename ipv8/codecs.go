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
	// CodeIdentityV1 - Content ID of IPFN Identity Version 1.
	// It is a standard used in IPFS and libp2p libraries.
	// Multihash digest contents is a ed25519 public key.
	CodeIdentityV1 = 0xed

	// CodeBinaryCellV1 - Content ID of IPFN Binary Cell Version 1.
	CodeBinaryCellV1 = 33407

	// CodeBinaryCellHashV1 - Content ID of IPFN Function Version 1.
	CodeBinaryCellHashV1 = 0x25ff2
)

func init() {
	codecs.Register(map[string]uint64{
		"ed25519-pub":       CodeIdentityV1,
		"ipfn-cell-v1":      CodeBinaryCellV1,
		"ipfn-cell-hash-v1": CodeBinaryCellHashV1,
	})
}
