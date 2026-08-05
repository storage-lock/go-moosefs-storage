package moosefs_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// MoosefsStorage moosefs对象存储实现 (S3兼容)
type MoosefsStorage struct {
	*s3_storage.S3Storage
}

func NewMoosefsStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*MoosefsStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &MoosefsStorage{S3Storage: s3Storage}, nil
}

const StorageName = "moosefs-storage"

func (x *MoosefsStorage) GetName() string {
	return StorageName
}

func (x *MoosefsStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}
