//    Copyright 2022 unbounded
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package page

import "math"

type (
	RequestPage struct {
		Size int64 `json:"size"` // 当前页条数
		Page int64 `json:"page"` // 当前页码
	}

	Page[T any] struct {
		TotalSize     int64 `json:"totalSize"`     // 总条数
		TotalPageSize int64 `json:"totalPageSize"` // 总页数
		RequestPage
		Data []T `json:"data"` // 数据
	}
)

// NewPage returns a Page.
func NewPage[T any](totalSize int64, requestPage *RequestPage) (p *Page[T]) {
	p = &Page[T]{TotalSize: totalSize, RequestPage: *requestPage}
	if requestPage.Size == -1 {
		p.Size = totalSize
		p.TotalPageSize = 1
		return
	}

	p.TotalPageSize = int64(math.Ceil(float64(totalSize) / float64(requestPage.Size)))
	return
}
