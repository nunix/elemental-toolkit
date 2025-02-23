/*
Copyright © 2022 - 2023 SUSE LLC

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

package v1

import (
	"github.com/rancher/elemental-toolkit/pkg/constants"
)

func (i InstallSpec) GetGrubLabels() map[string]string {
	grubEnv := map[string]string{
		"state_label":    i.Partitions.State.FilesystemLabel,
		"active_label":   i.Active.Label,
		"passive_label":  i.Passive.Label,
		"recovery_label": i.Partitions.Recovery.FilesystemLabel,
		"system_label":   i.Recovery.Label,
		"oem_label":      i.Partitions.OEM.FilesystemLabel,
	}

	if i.Partitions.Persistent != nil {
		grubEnv["persistent_label"] = i.Partitions.Persistent.FilesystemLabel
	}

	return grubEnv
}

func (u UpgradeSpec) GetGrubLabels() map[string]string {
	grubVars := map[string]string{
		"state_label":    u.Partitions.State.FilesystemLabel,
		"active_label":   u.Active.Label,
		"passive_label":  u.Passive.Label,
		"recovery_label": u.Partitions.Recovery.FilesystemLabel,
		"system_label":   u.Recovery.Label,
		"oem_label":      u.Partitions.OEM.FilesystemLabel,
	}

	if u.Partitions.Persistent != nil {
		grubVars["persistent_label"] = u.Partitions.Persistent.FilesystemLabel
	}

	return grubVars
}

func (r ResetSpec) GetGrubLabels() map[string]string {
	grubVars := map[string]string{
		"state_label":    r.Partitions.State.FilesystemLabel,
		"active_label":   r.Active.Label,
		"passive_label":  r.Passive.Label,
		"recovery_label": r.Partitions.Recovery.FilesystemLabel,
		"oem_label":      r.Partitions.OEM.FilesystemLabel,
	}

	if r.State != nil {
		if recoveryPart, ok := r.State.Partitions[constants.RecoveryPartName]; ok {
			grubVars["recovery_label"] = recoveryPart.FSLabel
			if recoveryImg, ok := recoveryPart.Images[constants.RecoveryImgName]; ok {
				grubVars["system_label"] = recoveryImg.Label
			}
		}
	}

	if r.Partitions.Persistent != nil {
		grubVars["persistent_label"] = r.Partitions.Persistent.FilesystemLabel
	}

	return grubVars
}
