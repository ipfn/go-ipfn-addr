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
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ipfn/go-base32i/base32i"
)

func TestIdentityV1(t *testing.T) {
	pk, _ := hex.DecodeString("4652486ebc271520d844e5bdda9ac243c05dcbe7bc9b93807073a32177a6f73d")
	c, err := PrefixIdentityV1.Sum(pk)
	assert.NoError(t, err)
	assert.Equal(t, "z4k889HeEVsMVFxdtVB3SE4uzAf2z6LD7aoLf1mxVqdFokG1cXv", c.String())
	assert.Equal(t, "08ksz0p0befysm4uyu2jpkzyuk7a4xkzbd09mjl8hjre80qsww3jzaax7u7sc", base32i.CheckEncodeToString(c.Bytes()))
}

func TestCellHTTPGet(t *testing.T) {
	body, _ := base32i.DecodeString("psx0yzn8rahkwmq99e3k7mb")
	cid, _ := PrefixBinaryCellV1.Sum(body)
	assert.Equal(t, "z2ceLLU67MxkSCRZ7nnfso5jcAA8", cid.String())
	assert.Equal(t, "08lcb0s0pcx0c0s2vahk7emvv5hxxmmrpc", base32i.CheckEncodeToString(cid.Bytes()))
}

func TestCellHTTPGet_hash(t *testing.T) {
	body, _ := base32i.DecodeString("psx0yzn8rahkwmq99e3k7mb")
	cid, _ := PrefixBinaryCellHashV1.Sum(body)
	assert.Equal(t, "z6ceLLNmpHqWREov94EFz2Hwh7AQ2dhAGG8VyhThbaWy81AAz5ZPWyY", cid.String())
	assert.Equal(t, "08et7zr0uspz0ua52chtrtc7km7ah5z89wxhyta86pnn5kc0ml7ef3jr897rfkvnp0", base32i.CheckEncodeToString(cid.Bytes()))
}

func TestPrefixBinaryCellHashV1(t *testing.T) {
	cid, _ := PrefixBinaryCellHashV1.Sum([]byte("test"))
	assert.Equal(t, "z6ceLLNmpHqYVd12b8xhDeBqEhu4LDPnEBkyvRVokJSgVEuGtSbwp5F", cid.String())
	// assert.Equal(t, "z6fncY87LynykjazXaHpRoYb7tvFzzvkXyhhU2y6oNuugn9VqiJ84X3", cid.String())	// 0x1007f7
}
