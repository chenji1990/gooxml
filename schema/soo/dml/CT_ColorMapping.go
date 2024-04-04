// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/chenji1990/gooxml"
)

type CT_ColorMapping struct {
	Bg1Attr      ST_ColorSchemeIndex
	Tx1Attr      ST_ColorSchemeIndex
	Bg2Attr      ST_ColorSchemeIndex
	Tx2Attr      ST_ColorSchemeIndex
	Accent1Attr  ST_ColorSchemeIndex
	Accent2Attr  ST_ColorSchemeIndex
	Accent3Attr  ST_ColorSchemeIndex
	Accent4Attr  ST_ColorSchemeIndex
	Accent5Attr  ST_ColorSchemeIndex
	Accent6Attr  ST_ColorSchemeIndex
	HlinkAttr    ST_ColorSchemeIndex
	FolHlinkAttr ST_ColorSchemeIndex
	ExtLst       *CT_OfficeArtExtensionList
}

func NewCT_ColorMapping() *CT_ColorMapping {
	ret := &CT_ColorMapping{}
	ret.Bg1Attr = ST_ColorSchemeIndex(1)
	ret.Tx1Attr = ST_ColorSchemeIndex(1)
	ret.Bg2Attr = ST_ColorSchemeIndex(1)
	ret.Tx2Attr = ST_ColorSchemeIndex(1)
	ret.Accent1Attr = ST_ColorSchemeIndex(1)
	ret.Accent2Attr = ST_ColorSchemeIndex(1)
	ret.Accent3Attr = ST_ColorSchemeIndex(1)
	ret.Accent4Attr = ST_ColorSchemeIndex(1)
	ret.Accent5Attr = ST_ColorSchemeIndex(1)
	ret.Accent6Attr = ST_ColorSchemeIndex(1)
	ret.HlinkAttr = ST_ColorSchemeIndex(1)
	ret.FolHlinkAttr = ST_ColorSchemeIndex(1)
	return ret
}

func (m *CT_ColorMapping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.Bg1Attr.MarshalXMLAttr(xml.Name{Local: "bg1"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Tx1Attr.MarshalXMLAttr(xml.Name{Local: "tx1"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Bg2Attr.MarshalXMLAttr(xml.Name{Local: "bg2"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Tx2Attr.MarshalXMLAttr(xml.Name{Local: "tx2"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Accent1Attr.MarshalXMLAttr(xml.Name{Local: "accent1"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Accent2Attr.MarshalXMLAttr(xml.Name{Local: "accent2"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Accent3Attr.MarshalXMLAttr(xml.Name{Local: "accent3"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Accent4Attr.MarshalXMLAttr(xml.Name{Local: "accent4"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Accent5Attr.MarshalXMLAttr(xml.Name{Local: "accent5"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.Accent6Attr.MarshalXMLAttr(xml.Name{Local: "accent6"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.HlinkAttr.MarshalXMLAttr(xml.Name{Local: "hlink"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.FolHlinkAttr.MarshalXMLAttr(xml.Name{Local: "folHlink"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorMapping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Bg1Attr = ST_ColorSchemeIndex(1)
	m.Tx1Attr = ST_ColorSchemeIndex(1)
	m.Bg2Attr = ST_ColorSchemeIndex(1)
	m.Tx2Attr = ST_ColorSchemeIndex(1)
	m.Accent1Attr = ST_ColorSchemeIndex(1)
	m.Accent2Attr = ST_ColorSchemeIndex(1)
	m.Accent3Attr = ST_ColorSchemeIndex(1)
	m.Accent4Attr = ST_ColorSchemeIndex(1)
	m.Accent5Attr = ST_ColorSchemeIndex(1)
	m.Accent6Attr = ST_ColorSchemeIndex(1)
	m.HlinkAttr = ST_ColorSchemeIndex(1)
	m.FolHlinkAttr = ST_ColorSchemeIndex(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "accent3" {
			m.Accent3Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "tx1" {
			m.Tx1Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "bg2" {
			m.Bg2Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "tx2" {
			m.Tx2Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent1" {
			m.Accent1Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent2" {
			m.Accent2Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "bg1" {
			m.Bg1Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent4" {
			m.Accent4Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent5" {
			m.Accent5Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "accent6" {
			m.Accent6Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "hlink" {
			m.HlinkAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "folHlink" {
			m.FolHlinkAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_ColorMapping:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ColorMapping %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorMapping
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorMapping and its children
func (m *CT_ColorMapping) Validate() error {
	return m.ValidateWithPath("CT_ColorMapping")
}

// ValidateWithPath validates the CT_ColorMapping and its children, prefixing error messages with path
func (m *CT_ColorMapping) ValidateWithPath(path string) error {
	if m.Bg1Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Bg1Attr is a mandatory field", path)
	}
	if err := m.Bg1Attr.ValidateWithPath(path + "/Bg1Attr"); err != nil {
		return err
	}
	if m.Tx1Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Tx1Attr is a mandatory field", path)
	}
	if err := m.Tx1Attr.ValidateWithPath(path + "/Tx1Attr"); err != nil {
		return err
	}
	if m.Bg2Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Bg2Attr is a mandatory field", path)
	}
	if err := m.Bg2Attr.ValidateWithPath(path + "/Bg2Attr"); err != nil {
		return err
	}
	if m.Tx2Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Tx2Attr is a mandatory field", path)
	}
	if err := m.Tx2Attr.ValidateWithPath(path + "/Tx2Attr"); err != nil {
		return err
	}
	if m.Accent1Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Accent1Attr is a mandatory field", path)
	}
	if err := m.Accent1Attr.ValidateWithPath(path + "/Accent1Attr"); err != nil {
		return err
	}
	if m.Accent2Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Accent2Attr is a mandatory field", path)
	}
	if err := m.Accent2Attr.ValidateWithPath(path + "/Accent2Attr"); err != nil {
		return err
	}
	if m.Accent3Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Accent3Attr is a mandatory field", path)
	}
	if err := m.Accent3Attr.ValidateWithPath(path + "/Accent3Attr"); err != nil {
		return err
	}
	if m.Accent4Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Accent4Attr is a mandatory field", path)
	}
	if err := m.Accent4Attr.ValidateWithPath(path + "/Accent4Attr"); err != nil {
		return err
	}
	if m.Accent5Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Accent5Attr is a mandatory field", path)
	}
	if err := m.Accent5Attr.ValidateWithPath(path + "/Accent5Attr"); err != nil {
		return err
	}
	if m.Accent6Attr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/Accent6Attr is a mandatory field", path)
	}
	if err := m.Accent6Attr.ValidateWithPath(path + "/Accent6Attr"); err != nil {
		return err
	}
	if m.HlinkAttr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/HlinkAttr is a mandatory field", path)
	}
	if err := m.HlinkAttr.ValidateWithPath(path + "/HlinkAttr"); err != nil {
		return err
	}
	if m.FolHlinkAttr == ST_ColorSchemeIndexUnset {
		return fmt.Errorf("%s/FolHlinkAttr is a mandatory field", path)
	}
	if err := m.FolHlinkAttr.ValidateWithPath(path + "/FolHlinkAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
