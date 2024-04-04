// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/chenji1990/gooxml"
)

type CT_PageBreak struct {
	// Page Break Count
	CountAttr *uint32
	// Manual Break Count
	ManualBreakCountAttr *uint32
	// Break
	Brk []*CT_Break
}

func NewCT_PageBreak() *CT_PageBreak {
	ret := &CT_PageBreak{}
	return ret
}

func (m *CT_PageBreak) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.ManualBreakCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "manualBreakCount"},
			Value: fmt.Sprintf("%v", *m.ManualBreakCountAttr)})
	}
	e.EncodeToken(start)
	if m.Brk != nil {
		sebrk := xml.StartElement{Name: xml.Name{Local: "ma:brk"}}
		for _, c := range m.Brk {
			e.EncodeElement(c, sebrk)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageBreak) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
		if attr.Name.Local == "manualBreakCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ManualBreakCountAttr = &pt
			continue
		}
	}
lCT_PageBreak:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "brk"}:
				tmp := NewCT_Break()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Brk = append(m.Brk, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_PageBreak %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PageBreak
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PageBreak and its children
func (m *CT_PageBreak) Validate() error {
	return m.ValidateWithPath("CT_PageBreak")
}

// ValidateWithPath validates the CT_PageBreak and its children, prefixing error messages with path
func (m *CT_PageBreak) ValidateWithPath(path string) error {
	for i, v := range m.Brk {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Brk[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
