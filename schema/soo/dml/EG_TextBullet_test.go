
// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/chenji1990/gooxml/schema/soo/dml"
)

func TestEG_TextBulletConstructor(t *testing.T) {
	v := dml.NewEG_TextBullet()
	if v == nil {
		t.Errorf("dml.NewEG_TextBullet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextBullet should validate: %s", err)
	}
}

func TestEG_TextBulletMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextBullet()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextBullet()
	xml.Unmarshal(buf, v2)
}
