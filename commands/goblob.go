// Copyright 2017-Present Pivotal Software, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http:#www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

type GoblobCommand struct {
	Version func() `command:"version" description:"Print version information and exit"`

	Migrate MigrateCommand `command:"migrate" description:"Migrate blobs from one blobstore to another"`
	MigrateToAzure MigrateToAzureBlobCommand `command:"migrate2azure" description:"Migrate blobs from NFS blobstore to Azure blobstore"`
}

var Goblob GoblobCommand
