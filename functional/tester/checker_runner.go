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

type runnerChecker struct {
	ctype              rpcpb.Checker
	etcdClientEndpoint string
	errc               chan error
}

func newRunnerChecker(ep string, errc chan error) Checker {
	return &runnerChecker{
		ctype:              rpcpb.Checker_RUNNER,
		etcdClientEndpoint: ep,
		errc:               errc,
	}
}

func (rc *runnerChecker) Type() rpcpb.Checker {
	return rc.ctype
}

func (rc *runnerChecker) EtcdClientEndpoints() []string {
	return []string{rc.etcdClientEndpoint}
}

func (rc *runnerChecker) Check() error {
	select {
	case err := <-rc.errc:
		return err
	default:
		return nil
	}
}
