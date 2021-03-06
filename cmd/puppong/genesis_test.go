// Copyright 2018 The go-orange Authors
// This file is part of go-orange.
//
// go-orange is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-orange is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-orange. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ong2020/go-orange/core"
)

// Tests the go-orange to Along chainspec conversion for the Stureby testnet.
func TestAlongSturebyConverter(t *testing.T) {
	blob, err := ioutil.ReadFile("testdata/stureby_gong.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	var genesis core.Genesis
	if err := json.Unmarshal(blob, &genesis); err != nil {
		t.Fatalf("failed parsing genesis: %v", err)
	}
	spec, err := newAlongGenesisSpec("stureby", &genesis)
	if err != nil {
		t.Fatalf("failed creating chainspec: %v", err)
	}

	expBlob, err := ioutil.ReadFile("testdata/stureby_along.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	expspec := &alongGenesisSpec{}
	if err := json.Unmarshal(expBlob, expspec); err != nil {
		t.Fatalf("failed parsing genesis: %v", err)
	}
	if !reflect.DeepEqual(expspec, spec) {
		t.Errorf("chainspec mismatch")
		c := spew.ConfigState{
			DisablePointerAddresses: true,
			SortKeys:                true,
		}
		exp := strings.Split(c.Sdump(expspec), "\n")
		got := strings.Split(c.Sdump(spec), "\n")
		for i := 0; i < len(exp) && i < len(got); i++ {
			if exp[i] != got[i] {
				t.Logf("got: %v\nexp: %v\n", exp[i], got[i])
			}
		}
	}
}

// Tests the go-orange to Parity chainspec conversion for the Stureby testnet.
func TestParitySturebyConverter(t *testing.T) {
	blob, err := ioutil.ReadFile("testdata/stureby_gong.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	var genesis core.Genesis
	if err := json.Unmarshal(blob, &genesis); err != nil {
		t.Fatalf("failed parsing genesis: %v", err)
	}
	spec, err := newParityChainSpec("stureby", &genesis, []string{})
	if err != nil {
		t.Fatalf("failed creating chainspec: %v", err)
	}
	enc, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		t.Fatalf("failed encoding chainspec: %v", err)
	}
	expBlob, err := ioutil.ReadFile("testdata/stureby_parity.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}
	if !bytes.Equal(expBlob, enc) {
		t.Fatalf("chainspec mismatch")
	}
}
