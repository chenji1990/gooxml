// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcCT_IdMap struct {
	DataAttr *string
	ExtAttr  ST_Ext
}

func NewOfcCT_IdMap() *OfcCT_IdMap {
	ret := &OfcCT_IdMap{}
	return ret
}

func (m *OfcCT_IdMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "data"},
			Value: fmt.Sprintf("%v", *m.DataAttr)})
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_IdMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "data" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DataAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_IdMap: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_IdMap and its children
func (m *OfcCT_IdMap) Validate() error {
	return m.ValidateWithPath("OfcCT_IdMap")
}

// ValidateWithPath validates the OfcCT_IdMap and its children, prefixing error messages with path
func (m *OfcCT_IdMap) ValidateWithPath(path string) error {
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
