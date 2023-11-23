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

package configmanager

import (
	"errors"
	"nestos-kubernetes-deployer/pkg/configmanager/asset"
	"nestos-kubernetes-deployer/pkg/configmanager/globalconfig"

	"github.com/spf13/cobra"
)

// Set global data
var GlobalConfig *globalconfig.GlobalAsset
var ClusterConfig = map[string]*asset.ClusterAsset{}

func Initial(cmd *cobra.Command) error {
	// Init global
	globalConfig, err := globalconfig.InitGlobalConfig(cmd)
	if err != nil {
		return err
	}

	// Init cluster
	clusterAsset, err := asset.InitClusterAsset(globalConfig, cmd)
	if err != nil {
		return err
	}
	ClusterConfig[clusterAsset.ClusterID] = clusterAsset

	return nil
}

func GetGlobalConfig() (*globalconfig.GlobalAsset, error) {
	return GlobalConfig, nil
}

func GetClusterConfig(clusterID string) (*asset.ClusterAsset, error) {
	clusterConfig, ok := ClusterConfig[clusterID]
	if !ok {
		return nil, errors.New("ClusterID not found")
	}

	return clusterConfig, nil
}

func Persist() error {
	// Persist global
	globalConfig, err := GetGlobalConfig()
	if err != nil {
		return err
	}
	if err := globalConfig.Persist(); err != nil {
		return err
	}

	// Persist cluster
	for _, clusterConfig := range ClusterConfig {
		if err := clusterConfig.Persist(); err != nil {
			return err
		}
	}

	return nil
}

func Delete() error {
	return nil
}