/*
 * Copyright (c) 2014 Michael Wendland
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 *
 * 	Authors:
 * 		Michael Wendland <michael@michiwend.com>
 */

package gomusicbrainz

import "encoding/xml"

type Work struct {
	ID        MBID                `xml:"id,attr"`
	Type      string              `xml:"type,attr"`
	TypeID    MBID                `xml:"type-id,attr"`
	Title     string              `xml:"title"`
	Relations []*RelationAbstract `xml:"relation-list>relation"`
}

func (mbe *Work) lookupResult() interface{} {
	var res struct {
		XMLName xml.Name `xml:"metadata"`
		Ptr     *Work    `xml:"work"`
	}
	res.Ptr = mbe
	return &res
}

func (mbe *Work) apiEndpoint() string {
	return "/work"
}

func (mbe *Work) Id() MBID {
	return mbe.ID
}

// LookupRecording performs an recording lookup request for the given MBID.
func (c *WS2Client) LookupWork(id MBID, inc ...string) (*Work, error) {
	a := &Work{ID: id}
	err := c.Lookup(a, inc...)

	return a, err
}

func (c *WS2Client) SearchWork(searchTerm string, limit, offset int) (*WorkSearchResponse, error) {
	//TODO implement
	return nil, nil
}

type WorkSearchResponse struct {
	//TODO implement
}
