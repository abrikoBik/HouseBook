// Package consts contains all non-changing values required by the storage service, such as hostname, filesystem path, etc.
package consts

const (
	Hostname = "0.0.0.0"
	Port     = 3001
	// About 128MB.
	MaxSize     = 1 << 27
	Key         = "file"
	StoragePath = "/var/lib/storage"
)
