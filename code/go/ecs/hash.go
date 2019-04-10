// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// The hash fields represent different hash algorithms and their values.
type Hash struct {
	// BLAKE2b-256 hash.
	Blake2b256 string `ecs:"blake2b_256"`

	// BLAKE2b-384 hash.
	Blake2b384 string `ecs:"blake2b_384"`

	// BLAKE2b-512 hash.
	Blake2b512 string `ecs:"blake2b_512"`

	// MD5 hash.
	Md5 string `ecs:"md5"`

	// SHA1 hash.
	Sha1 string `ecs:"sha1"`

	// SHA224 hash.
	Sha224 string `ecs:"sha224"`

	// SHA256 hash.
	Sha256 string `ecs:"sha256"`

	// SHA384 hash.
	Sha384 string `ecs:"sha384"`

	// SHA3_224 hash.
	Sha3224 string `ecs:"sha3_224"`

	// SHA3_256 hash.
	Sha3256 string `ecs:"sha3_256"`

	// SHA3_384 hash.
	Sha3384 string `ecs:"sha3_384"`

	// SHA3_512 hash.
	Sha3512 string `ecs:"sha3_512"`

	// SHA512 hash.
	Sha512 string `ecs:"sha512"`

	// SHA512/224 hash.
	Sha512224 string `ecs:"sha512_224"`

	// SHA512/256 hash.
	Sha512256 string `ecs:"sha512_256"`

	// XX64 hash.
	Xxh64 string `ecs:"xxh64"`
}
