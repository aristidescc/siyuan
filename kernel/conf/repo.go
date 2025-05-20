// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package conf

import (
	"path/filepath"

	"github.com/siyuan-note/siyuan/kernel/util"
)

type Repo struct {
	Key []byte `json:"key"` // AES key

	// Sync index timing in milliseconds, if exceeded, notify user about index performance degradation
	// If the data repo indexing time is greater than 12s, prompt user to purge the data repo https://github.com/siyuan-note/siyuan/issues/9613
	// Supports configuring data sync index time-consuming prompts https://github.com/siyuan-note/siyuan/issues/9698
	SyncIndexTiming int64 `json:"syncIndexTiming"`

	// 自动清理数据仓库 Automatic purge for local data repo https://github.com/siyuan-note/siyuan/issues/13091
	IndexRetentionDays    int `json:"indexRetentionDays"`    // Number of days to retain indexes
	RetentionIndexesDaily int `json:"retentionIndexesDaily"` // Number of indexes to retain per day
}

func NewRepo() *Repo {
	return &Repo{
		SyncIndexTiming:       12 * 1000,
		IndexRetentionDays:    180,
		RetentionIndexesDaily: 2,
	}
}

func (*Repo) GetSaveDir() string {
	return filepath.Join(util.WorkspaceDir, "repo")
}
