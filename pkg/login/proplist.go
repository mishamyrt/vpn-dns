// Package login provides tools for working with autorun on macOS.
package login

import "strings"

const xmlHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
const specURL = "http://www.apple.com/DTDs/PropertyList-1.0.dtd"
const typePropList = "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"" + specURL + "\">"

// PropList represents Apple property list file.
type PropList struct {
	props []string
}

// Bool adds logical tag to PropList.
func (d *PropList) Bool(key string, value bool) {
	d.appendRaw("<key>" + key + "</key>")
	if value {
		d.appendRaw("<true/>")
	} else {
		d.appendRaw("<false/>")
	}
}

// String adds literal tag to PropList.
func (d *PropList) String(key string, value string) {
	d.append("key", key)
	d.append("string", value)
}

// StringArray adds multiple literal tags to PropList.
func (d *PropList) StringArray(key string, values []string) {
	d.append("key", key)
	d.appendRaw("<array>")
	for _, value := range values {
		d.append("string", value)
	}
	d.appendRaw("</array>")
}

// Join list to string.
func (d *PropList) Join() string {
	result := xmlHeader + "\n" + typePropList + "\n"
	result += "<plist version=\"1.0\">" + "\n"
	result += "<dict>" + "\n"
	result += strings.Join(d.props, "\n") + "\n"
	result += "</dict>" + "\n"
	result += "</plist>" + "\n"
	return result
}

func (d *PropList) appendRaw(value string) {
	d.props = append(d.props, value)
}

func (d *PropList) append(tag, value string) {
	d.appendRaw("<" + tag + ">" + value + "</" + tag + ">")
}

// NewPropList creates new property list.
func NewPropList() PropList {
	builder := PropList{
		props: make([]string, 0),
	}
	return builder
}
