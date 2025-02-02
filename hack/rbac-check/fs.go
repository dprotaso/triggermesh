/*
Copyright 2021 TriggerMesh Inc.

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

package main

import (
	"io/fs"
	"os"
)

// osFS is a fs.FS implementation that uses functions provided by the 'os' package.
type osFS struct{}

var _ fs.ReadDirFS = (*osFS)(nil)

// Open implements fs.ReadDirFS.
func (*osFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

// ReadDirFS implements fs.ReadDirFS.
func (*osFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}
