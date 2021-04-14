// This is a generated file, please do not edit by hand
package filetype

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// FileType represents the type of a file.

type FileType int32

// FileTypePtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func FileTypePtr(in FileType) *FileType { return &in }

// FileTypeFromPtr provides a way to convert back from an enum pointer to a
// guaranteed FileType value of some kind (0 is default)
func FileTypeFromPtr(in *FileType) FileType {
	if in == nil {
		return 0
	}
	return *in
}

// FileType constant values
const (
	PICTURE    = FileType(pb.FileType_PICTURE)
	MOVIE      = FileType(pb.FileType_MOVIE)
	DOCUMENT   = FileType(pb.FileType_DOCUMENT)
	EXECUTABLE = FileType(pb.FileType_EXECUTABLE)
	MUSIC      = FileType(pb.FileType_MUSIC)
	GENERIC    = FileType(pb.FileType_GENERIC)
)
