package vmware

import (
	"github.com/kanisterio/kanister/pkg/blockstorage"
	"github.com/vmware/govmomi/vim25/types"
)

func convertFromObjectToVolume(vso *types.VStorageObject) *blockstorage.Volume {
	return &blockstorage.Volume{
		Type:         blockstorage.TypeFCD,
		ID:           vso.Config.Id.Id,
		CreationTime: blockstorage.TimeStamp(vso.Config.CreateTime),
		Size:         vso.Config.CapacityInMB / 1024,
		Az:           "",
		Iops:         0,
		Encrypted:    false,
	}
}

func convertFromObjectToSnapshot(vso *types.VStorageObjectSnapshotInfoVStorageObjectSnapshot) *blockstorage.Snapshot {
	return &blockstorage.Snapshot{
		Type:         blockstorage.TypeFCD,
		CreationTime: blockstorage.TimeStamp(vso.CreateTime),
		ID:           vso.Id.Id,
		Size:         0,
		Region:       "",
		Encrypted:    false,
	}
}

// ID wraps ID string with vim25.ID struct.
func ID(id string) types.ID {
	return types.ID{
		Id: id,
	}
}