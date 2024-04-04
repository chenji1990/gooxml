// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/chenji1990/gooxml"
	"github.com/chenji1990/gooxml/schema/soo/dml"
)

type WdCT_WordprocessingShape struct {
	NormalEastAsianFlowAttr *bool
	CNvPr                   *dml.CT_NonVisualDrawingProps
	Choice                  *WdCT_WordprocessingShapeChoice
	SpPr                    *dml.CT_ShapeProperties
	Style                   *dml.CT_ShapeStyle
	ExtLst                  *dml.CT_OfficeArtExtensionList
	WChoice                 *WdCT_WordprocessingShapeChoice1
	BodyPr                  *dml.CT_TextBodyProperties
}

func NewWdCT_WordprocessingShape() *WdCT_WordprocessingShape {
	ret := &WdCT_WordprocessingShape{}
	ret.Choice = NewWdCT_WordprocessingShapeChoice()
	ret.SpPr = dml.NewCT_ShapeProperties()
	ret.BodyPr = dml.NewCT_TextBodyProperties()
	return ret
}

func (m *WdCT_WordprocessingShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NormalEastAsianFlowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "normalEastAsianFlow"},
			Value: fmt.Sprintf("%d", b2i(*m.NormalEastAsianFlowAttr))})
	}
	e.EncodeToken(start)
	if m.CNvPr != nil {
		secNvPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvPr"}}
		e.EncodeElement(m.CNvPr, secNvPr)
	}
	m.Choice.MarshalXML(e, xml.StartElement{})
	sespPr := xml.StartElement{Name: xml.Name{Local: "wp:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "wp:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	if m.WChoice != nil {
		m.WChoice.MarshalXML(e, xml.StartElement{})
	}
	sebodyPr := xml.StartElement{Name: xml.Name{Local: "wp:bodyPr"}}
	e.EncodeElement(m.BodyPr, sebodyPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WordprocessingShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewWdCT_WordprocessingShapeChoice()
	m.SpPr = dml.NewCT_ShapeProperties()
	m.BodyPr = dml.NewCT_TextBodyProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "normalEastAsianFlow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NormalEastAsianFlowAttr = &parsed
			continue
		}
	}
lWdCT_WordprocessingShape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvPr"}:
				m.CNvPr = dml.NewCT_NonVisualDrawingProps()
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvSpPr"}:
				m.Choice = NewWdCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvCnPr"}:
				m.Choice = NewWdCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvCnPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "style"}:
				m.Style = dml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "txbx"}:
				m.WChoice = NewWdCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.WChoice.Txbx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "linkedTxbx"}:
				m.WChoice = NewWdCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.WChoice.LinkedTxbx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "bodyPr"}:
				if err := d.DecodeElement(m.BodyPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_WordprocessingShape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingShape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingShape and its children
func (m *WdCT_WordprocessingShape) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingShape")
}

// ValidateWithPath validates the WdCT_WordprocessingShape and its children, prefixing error messages with path
func (m *WdCT_WordprocessingShape) ValidateWithPath(path string) error {
	if m.CNvPr != nil {
		if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
			return err
		}
	}
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	if m.WChoice != nil {
		if err := m.WChoice.ValidateWithPath(path + "/WChoice"); err != nil {
			return err
		}
	}
	if err := m.BodyPr.ValidateWithPath(path + "/BodyPr"); err != nil {
		return err
	}
	return nil
}
