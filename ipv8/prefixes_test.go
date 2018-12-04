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

// Package ipv8 implements helpers for generating public-key CID.
package ipv8

import (
	"encoding/hex"
	"testing"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	mh "gx/ipfs/QmerPMzPk1mJVowm8KgmoknWa4yCYvvugMPsgWmDNUvDLW/go-multihash"

	"github.com/stretchr/testify/assert"

	"github.com/ipfn/go-base32i/base32i"
)

func TestIdentityV1(t *testing.T) {
	pk, _ := hex.DecodeString("4652486ebc271520d844e5bdda9ac243c05dcbe7bc9b93807073a32177a6f73d")
	c, err := PrefixIdentityV1.Sum(pk)
	assert.NoError(t, err)
	assert.Equal(t, "zFNScXjsiATFzu45Xo9vYezYEcF8n11ytkyCa4rZkzX3QuZ4Yjfr", c.String())
	assert.Equal(t, "0xkvz0b0ypq9yjqwhsn32bxcbnjmmk56cfpu0hwtu77fhyu0wpe6xbth5mmn6qc", base32i.CheckEncodeToString(c.Bytes()))

	// Less important...
	pcid, err := cid.Decode(c.String())
	assert.NoError(t, err)
	assert.Equal(t, PrefixIdentityV1.Codec, pcid.Type())
	assert.Equal(t, PrefixIdentityV1.Version, pcid.Version())
	assert.Equal(t, c.String(), pcid.String())
	h, _ := mh.Decode(pcid.Hash())
	assert.Equal(t, PrefixIdentityV1.MhType, h.Code)
	assert.Equal(t, PrefixIdentityV1.MhLength, h.Length)
	assert.Equal(t, pk, h.Digest)
}

func TestIdentityHashV1(t *testing.T) {
	pk, _ := hex.DecodeString("4652486ebc271520d844e5bdda9ac243c05dcbe7bc9b93807073a32177a6f73d")
	c, err := PrefixIdentityHashV1.Sum(pk)
	assert.NoError(t, err)
	assert.Equal(t, "z5pwkYgEhukNP1aHsUe6PRExPPuBN8eHmLp2QYJVnYRmrAfFNJx7GBE", c.String())
	assert.Equal(t, "0xkvz0r0uspz0t2pfexlcs9f27e0vv8vu7xvhmdqquhb59pretbxq77v8j99sm9ryu", base32i.CheckEncodeToString(c.Bytes()))

	// Less important...
	pcid, err := cid.Decode(c.String())
	assert.NoError(t, err)
	assert.Equal(t, PrefixIdentityHashV1.Codec, pcid.Type())
	assert.Equal(t, PrefixIdentityHashV1.Version, pcid.Version())
	assert.Equal(t, c.String(), pcid.String())
	h, _ := mh.Decode(pcid.Hash())
	assert.Equal(t, PrefixIdentityHashV1.MhType, h.Code)
	assert.Equal(t, PrefixIdentityHashV1.MhLength, h.Length)
	assert.NotEqual(t, pk, h.Digest)
}

func TestCellHTTPGet(t *testing.T) {
	body, _ := base32i.DecodeString("psx0yzn8rahkwmq99e3k7mb")
	cid, _ := PrefixBinaryCellV1.Sum(body)

	assert.Equal(t, "z2ZkmggiKq2ABF8pHnereqW6jrPE", cid.String())
	assert.Equal(t, "08hmc800pcx0c0s2vahk7emvv5hxxmmr", base32i.EncodeToString(cid.Bytes()))
	assert.Equal(t, "08hmc800pcx0c0s2vahk7emvv5hxxmmrps", base32i.CheckEncodeToString(cid.Bytes()))
}

func TestCellHTTPGet_hash(t *testing.T) {
	body, _ := base32i.DecodeString("psx0yzn8rahkwmq99e3k7mb")
	cid, _ := PrefixBinaryCellHashV1.Sum(body)

	assert.Equal(t, "z6agHvA6jso95xevjS2GniZ8dG3ERuTAyafdKZKL1JbS5PXT9XcwXJa", cid.String())
	// assert.Equal(t, "zHceLLKf7yq1dJCt1ZjEh5yj4aStE3BiLyhnKDHfVYpBxYQLGUS8", cid.String())
}

func TestPrefixFunctionV1(t *testing.T) {
	cid, _ := PrefixFunctionV1.Sum([]byte("test"))
	assert.Equal(t, "zFunckvxnXmr6CvoSrsmApsvH7Em8mJ1fntbARqQn8Ea8ZgwZ5tA", cid.String())
	assert.Equal(t, "z6fncY87LynykjazXaHpRoYb7tvFzzvkXyhhU2y6oNuugn9VqiJ84X3", cid.String()) // blake2b
	assert.Equal(t, "0x7wz0b5yzdvuzqwnwkyj8avtsw303k2z8tndwf29vhtmyls0htmwy03ps9xw", base32i.EncodeToString(cid.Bytes()))
}

func TestPrefixFunctionV2(t *testing.T) {
	cid, _ := PrefixFunctionV2.Sum([]byte("test"))
	assert.Equal(t, "zFuncrid3XKFUVoXbvHE5Mkwf1WBnFSJiyX7MNPYPossbRHRqQd4", cid.String())
	assert.Equal(t, "0x7wz0s5yzdvuzqwnwkyj8avtsw303k2z8tndwf29vhtmyls0htmwy03ps9xw", base32i.EncodeToString(cid.Bytes()))
}

func TestPfx(t *testing.T) {
	pfx := cid.Prefix{
		Version:  1,
		Codec:    0xb0bc,
		MhType:   0,
		MhLength: -1,
	}
	cid, _ := pfx.Sum([]byte("test"))
	assert.Equal(t, "z6fSR3G2u1EFhd", cid.String())
}
