// Copyright 2017 The etcd Authors
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

//go:build cov
// +build cov

package ctlv2

import (
	"os"
	"strings"

	"github.com/urfave/cli"
)

func runCtlV2(app *cli.App) error {
	return app.Run(strings.Split(os.Getenv("ETCDCTL_ARGS"), "\xe7\xcd"))
}
