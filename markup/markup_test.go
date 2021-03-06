// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package markup

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/gohugoio/hugo/markup/converter"

	qt "github.com/frankban/quicktest"
)

func TestConverterRegistry(t *testing.T) {
	c := qt.New(t)

	r, err := NewConverterProvider(converter.ProviderConfig{Cfg: viper.New()})

	c.Assert(err, qt.IsNil)

	c.Assert(r.Get("foo"), qt.IsNil)
	c.Assert(r.Get("markdown"), qt.Not(qt.IsNil))
	c.Assert(r.Get("mmark"), qt.Not(qt.IsNil))
	c.Assert(r.Get("asciidoc"), qt.Not(qt.IsNil))
	c.Assert(r.Get("rst"), qt.Not(qt.IsNil))
	c.Assert(r.Get("pandoc"), qt.Not(qt.IsNil))
	c.Assert(r.Get("org"), qt.Not(qt.IsNil))

}
