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
package machine

import (
	"nestos-kubernetes-deployer/pkg/configmanager"
	"nestos-kubernetes-deployer/pkg/configmanager/asset"
	"nestos-kubernetes-deployer/pkg/ignition"
	"path/filepath"

	igntypes "github.com/coreos/ignition/v2/config/v3_2/types"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	ClusterAsset *asset.ClusterAsset
}

func (w *Worker) GenerateFiles() error {
	wtd := ignition.GetTmplData(w.ClusterAsset)

	for i, worker := range w.ClusterAsset.Worker {
		config := &igntypes.Config{}
		wtd.NodeName = worker.Hostname
		generateFile := ignition.Common{
			NodeType:        "worker",
			TmplData:        wtd,
			EnabledServices: ignition.EnabledServices,
			Config:          config,
			UserName:        worker.UserName,
			SSHKey:          worker.SSHKey,
			PassWord:        worker.Password,
		}

		// Generate Ignition data
		if err := generateFile.Generate(); err != nil {
			logrus.Errorf("failed to generate %s ignition file: %v", worker.UserName, err)
			return err
		}

		// Assign the Ignition path to the Worker node
		filePath := filepath.Join(configmanager.GetPersistDir(), w.ClusterAsset.Cluster_ID, "ignition")
		fileName := worker.Hostname + ".ign"
		w.ClusterAsset.Worker[i].Ign_Path = filepath.Join(filePath, fileName)

		ignition.SaveFile(generateFile.Config, filePath, fileName)
		logrus.Infof("Successfully generate %s ignition file", worker.Hostname)
	}

	return nil
}
