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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenPage(t *testing.T) {
	page := NewPage[int](10, &RequestPage{
		Size: 2,
		Page: 0,
	})
	assert.Equal(t, &Page[int]{
		TotalSize:     10,
		TotalPageSize: 5,
		RequestPage: RequestPage{
			Size: 2,
			Page: 0,
		},
	}, page)

	page = NewPage[int](11, &RequestPage{
		Size: 2,
		Page: 0,
	})
	assert.Equal(t, &Page[int]{
		TotalSize:     11,
		TotalPageSize: 6,
		RequestPage: RequestPage{
			Size: 2,
			Page: 0,
		},
	}, page)

	page = NewPage[int](11, &RequestPage{
		Size: -1,
		Page: 0,
	})
	assert.Equal(t, &Page[int]{
		TotalSize:     11,
		TotalPageSize: 1,
		RequestPage: RequestPage{
			Size: 11,
			Page: 0,
		},
	}, page)

	page = NewPage[int](11, &RequestPage{
		Size: 3,
		Page: 0,
	})
	assert.Equal(t, &Page[int]{
		TotalSize:     11,
		TotalPageSize: 4,
		RequestPage: RequestPage{
			Size: 3,
			Page: 0,
		},
	}, page)
	page = NewPage[int](11, &RequestPage{
		Size: 3,
		Page: 1,
	})

	assert.Equal(t, &Page[int]{
		TotalSize:     11,
		TotalPageSize: 4,
		RequestPage: RequestPage{
			Size: 3,
			Page: 1,
		},
	}, page)

	page = NewPage[int](11, &RequestPage{
		Size: 3,
		Page: 3,
	})

	assert.Equal(t, &Page[int]{
		TotalSize:     11,
		TotalPageSize: 4,
		RequestPage: RequestPage{
			Size: 3,
			Page: 3,
		},
	}, page)

}
