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

	"github.com/chenji1990/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_LsdException struct {
	// Primary Style Name
	NameAttr string
	// Latent Style Locking Setting
	LockedAttr *sharedTypes.ST_OnOff
	// Override default sorting order
	UiPriorityAttr *int64
	// Semi hidden text override
	SemiHiddenAttr *sharedTypes.ST_OnOff
	// Unhide when used
	UnhideWhenUsedAttr *sharedTypes.ST_OnOff
	// Latent Style Primary Style Setting
	QFormatAttr *sharedTypes.ST_OnOff
}

func NewCT_LsdException() *CT_LsdException {
	ret := &CT_LsdException{}
	return ret
}

func (m *CT_LsdException) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.LockedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:locked"},
			Value: fmt.Sprintf("%v", *m.LockedAttr)})
	}
	if m.UiPriorityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:uiPriority"},
			Value: fmt.Sprintf("%v", *m.UiPriorityAttr)})
	}
	if m.SemiHiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:semiHidden"},
			Value: fmt.Sprintf("%v", *m.SemiHiddenAttr)})
	}
	if m.UnhideWhenUsedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:unhideWhenUsed"},
			Value: fmt.Sprintf("%v", *m.UnhideWhenUsedAttr)})
	}
	if m.QFormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:qFormat"},
			Value: fmt.Sprintf("%v", *m.QFormatAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LsdException) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "locked" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = &parsed
			continue
		}
		if attr.Name.Local == "uiPriority" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.UiPriorityAttr = &parsed
			continue
		}
		if attr.Name.Local == "semiHidden" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.SemiHiddenAttr = &parsed
			continue
		}
		if attr.Name.Local == "unhideWhenUsed" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.UnhideWhenUsedAttr = &parsed
			continue
		}
		if attr.Name.Local == "qFormat" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.QFormatAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_LsdException: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_LsdException and its children
func (m *CT_LsdException) Validate() error {
	return m.ValidateWithPath("CT_LsdException")
}

// ValidateWithPath validates the CT_LsdException and its children, prefixing error messages with path
func (m *CT_LsdException) ValidateWithPath(path string) error {
	if m.LockedAttr != nil {
		if err := m.LockedAttr.ValidateWithPath(path + "/LockedAttr"); err != nil {
			return err
		}
	}
	if m.SemiHiddenAttr != nil {
		if err := m.SemiHiddenAttr.ValidateWithPath(path + "/SemiHiddenAttr"); err != nil {
			return err
		}
	}
	if m.UnhideWhenUsedAttr != nil {
		if err := m.UnhideWhenUsedAttr.ValidateWithPath(path + "/UnhideWhenUsedAttr"); err != nil {
			return err
		}
	}
	if m.QFormatAttr != nil {
		if err := m.QFormatAttr.ValidateWithPath(path + "/QFormatAttr"); err != nil {
			return err
		}
	}
	return nil
}
