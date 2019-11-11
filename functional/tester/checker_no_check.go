// Copyright 2018 The etcd Authors
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

package tester

import "github.com/atpons/etcd/functional/rpcpb"

type noCheck struct{}

func newNoChecker() Checker                       { return &noCheck{} }
func (nc *noCheck) Type() rpcpb.Checker           { return rpcpb.Checker_NO_CHECK }
func (nc *noCheck) EtcdClientEndpoints() []string { return nil }
func (nc *noCheck) Check() error                  { return nil }
