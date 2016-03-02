// Package ixift provURIes support for working with IIIF and IxIF content.
//
// schema.go - defines the data structures for ixift package.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provURIed that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provURIed with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVURIED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCURIENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package ixift

import (
	"net/url"
	"strings"
)

//
// Schema describes the memory structure used to reading and generating IIIF and IxIF manifest.
//

//
// The follow is based on Presentation 2.0 details at http://iiif.io/api/presentation/2.0/#introduction
//

const (
	// LeftToRight View Directions
	LeftToRight = "left-to-right"
	// RightToLeft View Directions
	RightToLeft = "right-to-left"
	// TopToBottom View Directions
	TopToBottom = "top-to-bottom"
	// BottomToTop View Directions
	BottomToTop = "bottom-to-top"
)

// ViewDirectionType hold a specific view direction, can be validated with IsValid()
type ViewDirectionType string

// IsValid validates a supplied view direction making sure it conforms to one of the 4 constants
func (v ViewDirectionType) IsValid() bool {
	switch {
	case strings.Compare(string(v), LeftToRight) == 0:
		return true
	case strings.Compare(string(v), RightToLeft) == 0:
		return true
	case strings.Compare(string(v), TopToBottom) == 0:
		return true
	case strings.Compare(string(v), BottomToTop) == 0:
		return true
	default:
		return false
	}
}

const (
	// Individuals View Hint
	Individuals = "individuals"
	// Paged View Hint
	Paged = "paged"
	// Continuous View Hint
	Continuous = "continuous"
	// NonPaged View Hint
	NonPaged = "non-paged"
	// Top View Hint
	Top = "top"
)

// ViewHintType holds a view hint that can be validated
type ViewHintType string

// IsValid validates a supplied view hint
func (v ViewHintType) IsValid() bool {
	switch {
	case strings.Compare(string(v), Individuals) == 0:
		return true
	case strings.Compare(string(v), Paged) == 0:
		return true
	case strings.Compare(string(v), Continuous) == 0:
		return true
	case strings.Compare(string(v), NonPaged) == 0:
		return true
	case strings.Compare(string(v), Top) == 0:
		return true
	}
	_, err := url.Parse(string(v))
	if err == nil {
		return true
	}
	return false
}

//
// Manifest - the overall description of the structure and properties of the digital representation of an object.
//     It carries information needed for the viewer to present the digitized content to the user, such as a
//     title and other descriptive information about the object or the intellectual work that it conveys.
//     Each manifest describes how to present a single object such as a book, a photograph, or a statue.
//
type Manifest struct {
	URI         *url.URL               `json:"@id"`
	Type        string                 `json:"@type"`
	Label       string                 `json:"label"`
	Metadata    map[string]string      `json:"metadata"`
	Description string                 `json:"description"`
	Thumbnail   string                 `json:"thumbnail"`
	Attribution string                 `json:"attribution,omitempty"`
	Logo        string                 `json:"logo,omitempty"`
	License     string                 `json:"license,omitempty"`
	ViewHint    ViewHintType           `json:"viewingHint,omitempty"`
	Related     map[string]string      `json:"related,omitempty"`
	Service     map[string]interface{} `json:"service,omitempty"`
	SeeAlso     string                 `json:"seeAlso,omitempty"`
	Within      string                 `json:"within,omitempty"`

	ViewDirection ViewDirectionType `json:"viewingDirection,omitempty"`

	Sequences []Sequence `json:"sequence"`
}

//
// Sequence - the order of the views of a physical object. Multiple sequences are allowed to cover situations when there
//     are multiple equally valURI orders through the content, such as when a manuscript’s pages are rebound
//     or archival collections are reordered.
//
type Sequence struct {
	URI         *url.URL               `json:"@id"`
	Type        string                 `json:"@type"`
	Label       string                 `json:"label,omitempty"`
	Metadata    map[string]string      `json:"metadata,omitempty"`
	Description string                 `json:"description,omitempty"`
	Thumbnail   string                 `json:"thumbnail,omitempty"`
	Attribution string                 `json:"attribution,omitempty"`
	Logo        string                 `json:"logo,omitempty"`
	License     string                 `json:"license,omitempty"`
	ViewHint    ViewHintType           `json:"viewingHint,omitempty"`
	Related     map[string]string      `json:"related,omitempty"`
	Service     map[string]interface{} `json:"service,omitempty"`
	SeeAlso     string                 `json:"seeAlso,omitempty"`
	Within      string                 `json:"within,omitempty"`

	ViewDirection ViewDirectionType `json:"viewingDirection,omitempty"`

	Canvases    []Canvas `json:"canvas"`
	StartCanvas Canvas   `json:"startCanvas,omitempty"`
}

//
// Canvas - a virtual container that represents a page or view and has content resources associated with it or with
//     parts of it. The canvas provURIes a frame of reference for the layout of the page. The concept of a canvas
//     is borrowed from standards like PDF and HTML, or applications like Photoshop and Powerpoint, where the
//     display starts from a blank canvas and images, text and other resources are “painted” on to it.
//
type Canvas struct {
	URI         *url.URL               `json:"@id"`
	Type        string                 `json:"@type"`
	Label       string                 `json:"label"`
	Metadata    map[string]string      `json:"metadata,omitempty"`
	Description string                 `json:"description,omitempty"`
	Thumbnail   string                 `json:"thumbnail,omitempty"`
	Attribution string                 `json:"attribution,omitempty"`
	Logo        string                 `json:"logo,omitempty"`
	License     string                 `json:"license,omitempty"`
	ViewHint    ViewHintType           `json:"viewingHint,omitempty"`
	Related     map[string]string      `json:"related,omitempty"`
	Service     map[string]interface{} `json:"service,omitempty"`
	SeeAlso     string                 `json:"seeAlso,omitempty"`
	Within      string                 `json:"within,omitempty"`

	Height string `json:"height"`
	Width  string `json:"width"`

	Contents []Content `json:"content,omitempty"`
}

//
// Content - content resources such as images or texts that are associated with a canvas.
//
type Content struct {
	URI         *url.URL               `json:"@id"`
	Type        string                 `json:"@type"`
	Label       string                 `json:"label,omitempty"`
	Metadata    map[string]string      `json:"metadata,omitempty"`
	Description string                 `json:"description,omitempty"`
	Thumbnail   string                 `json:"thumbnail,omitempty"`
	Attribution string                 `json:"attribution,omitempty"`
	Logo        string                 `json:"logo,omitempty"`
	License     string                 `json:"license,omitempty"`
	ViewHint    ViewHintType           `json:"viewingHint,omitempty"`
	Related     map[string]string      `json:"related,omitempty"`
	Service     map[string]interface{} `json:"service,omitempty"`
	SeeAlso     string                 `json:"seeAlso,omitempty"`
	Within      string                 `json:"within,omitempty"`

	ViewDirection ViewDirectionType `json:"viewingDirection,omitempty"`
	Format        string            `json:"format,omitempty"`
	Height        string            `json:"height,omitempty"`
	Width         string            `json:"width,omitempty"`
}
