# [go-ipfn-addr][IPv8]: IPFN address aka IPv8™

[![IPFN project][badge-ipfn]][org-ipfn]
[![IPFN Documentation][badge-docs]][docs]
[![See COPYING.txt][badge-copying]][COPYING]
[![GoDoc][badge-godoc]][godoc-ipfn]
[![Travis CI][badge-ci]][ci]
[![Coverage Status][coverage-badge]][coverage-status]

Go implementation of IPFN address aka IPv8™ specification.

IPv8™ - don't take it too seriously it's just a [content ID][cid].

## Standards

* [ed25519][] - Cryptographic Identity Public Keys
* [blake2b][] - Cryptographic Hashing of Contents

Content IDs contain [multihashes][multihash].

Multihash codec ID `0` is used to identify a raw binary contents and everything different identifes a digest [hash code][hash-code].

## License

See [COPYING][COPYING] file for licensing details.

## Project

This source code is part of [IPFN](https://github.com/ipfn) – interplanetary functions project.

[ci]: https://travis-ci.org/ipfn/go-ipfn-addr
[docs]: https://docs.ipfn.io/
[COPYING]: https://github.com/ipfn/go-ipfn-addr/blob/master/COPYING
[badge-ci]: https://travis-ci.org/ipfn/go-ipfn-addr.svg?branch=master
[badge-copying]: https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square
[badge-docs]: https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square
[badge-godoc]: https://godoc.org/github.com/ipfn/go-ipfn-addr/ipv8?status.svg
[badge-ipfn]: https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square
[coverage-badge]: https://coveralls.io/repos/github/ipfn/go-ipfn-addr/badge.svg?branch=master
[coverage-status]: https://coveralls.io/github/ipfn/go-ipfn-addr?branch=master
[org-ipfn]: https://github.com/ipfn
[godoc-ipfn]: https://godoc.org/github.com/ipfn/go-ipfn-addr/ipv8
[cid]:  https://github.com/ipld/cid
[multihash]:  https://github.com/multiformats/multihash
[ipv8]: https://github.com/ipfn/go-ipfn-addr/
[hash-code]: https://github.com/multiformats/multicodec/blob/master/table.csv
[ed25519]: https://ed25519.cr.yp.to/
[blake2b]: https://blake2.net/
