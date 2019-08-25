/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package rootfs

import (
	"github.com/talos-systems/talos/internal/app/machined/internal/phase"
	"github.com/talos-systems/talos/internal/app/machined/internal/platform"
	"github.com/talos-systems/talos/internal/app/machined/internal/runtime"
	"github.com/talos-systems/talos/pkg/userdata"
	"golang.org/x/sys/unix"
)

// MountShared represents the MountShared task.
type MountShared struct{}

// NewMountSharedTask initializes and returns an MountShared task.
func NewMountSharedTask() phase.Task {
	return &MountShared{}
}

// RuntimeFunc returns the runtime function.
func (task *MountShared) RuntimeFunc(mode runtime.Mode) phase.RuntimeFunc {
	switch mode {
	case runtime.Container:
		return task.container
	default:
		return nil
	}
}

func (task *MountShared) container(platform platform.Platform, data *userdata.UserData) (err error) {
	targets := []string{"/", "/var/lib/kubelet", "/etc/cni"}
	for _, t := range targets {
		if err = unix.Mount("", t, "", unix.MS_SHARED, ""); err != nil {
			return err
		}
	}

	return nil
}
