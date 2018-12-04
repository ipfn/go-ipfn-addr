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
	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
)

// PrefixBinaryCellV1 - Binary cell byte code CID prefix.
var PrefixBinaryCellV1 = cid.Prefix{
	Version:  1,
	Codec:    CodeBinaryCellV1,
	MhType:   0,
	MhLength: -1,
}

// PrefixBinaryCellHashV1 - Binary cell byte code CID prefix.
var PrefixBinaryCellHashV1 = cid.Prefix{
	Version:  1,
	Codec:    CodeBinaryCellV1,
	MhType:   0xb220, // blake2b 256 bits
	MhLength: 32,
}

// PrefixFunctionV1 - Function address V1 CID prefix.
var PrefixFunctionV1 = cid.Prefix{
	Version:  1,
	Codec:    CodeFunctionV1,
	MhType:   0xb220, // blake2b 256 bits
	MhLength: 32,
}

// PrefixIdentityV1 - Identity public key CID prefix.
var PrefixIdentityV1 = cid.Prefix{
	Version:  1,
	Codec:    CodeIdentityV1,
	MhType:   0,
	MhLength: 32,
}
