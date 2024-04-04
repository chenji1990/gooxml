// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"

	"github.com/chenji1990/gooxml"
)

type CT_EqArr struct {
	EqArrPr *CT_EqArrPr
	E       []*CT_OMathArg
}

func NewCT_EqArr() *CT_EqArr {
	ret := &CT_EqArr{}
	return ret
}

func (m *CT_EqArr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.EqArrPr != nil {
		seeqArrPr := xml.StartElement{Name: xml.Name{Local: "m:eqArrPr"}}
		e.EncodeElement(m.EqArrPr, seeqArrPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	for _, c := range m.E {
		e.EncodeElement(c, see)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EqArr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EqArr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "eqArrPr"}:
				m.EqArrPr = NewCT_EqArrPr()
				if err := d.DecodeElement(m.EqArrPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"}:
				tmp := NewCT_OMathArg()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_EqArr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EqArr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EqArr and its children
func (m *CT_EqArr) Validate() error {
	return m.ValidateWithPath("CT_EqArr")
}

// ValidateWithPath validates the CT_EqArr and its children, prefixing error messages with path
func (m *CT_EqArr) ValidateWithPath(path string) error {
	if m.EqArrPr != nil {
		if err := m.EqArrPr.ValidateWithPath(path + "/EqArrPr"); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
