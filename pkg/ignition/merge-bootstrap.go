/*
Copyright 2023 KylinSoft  Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package ignition

import (
	"fmt"
	"net/url"

	ignutil "github.com/coreos/ignition/v2/config/util"
	igntypes "github.com/coreos/ignition/v2/config/v3_2/types"
)

func GenerateMergeIgnition(bootstrapIgnitionHost string, role string) *igntypes.Config {
	ign := igntypes.Config{
		Ignition: igntypes.Ignition{
			Version: igntypes.MaxVersion.String(),
			Config: igntypes.IgnitionConfig{
				Merge: []igntypes.Resource{{
					Source: ignutil.StrToPtr(func() *url.URL {
						return &url.URL{
							Scheme: "http",
							Host:   bootstrapIgnitionHost,
							Path:   fmt.Sprintf("%s", role),
						}
					}().String()),
				}},
			},
		},
	}
	return &ign
}
