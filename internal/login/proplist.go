package login

import "strings"

var XMLHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"

var specURL = "http://www.apple.com/DTDs/PropertyList-1.0.dtd"
var TypePropList = "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"" + specURL + "\">"

type PropList struct {
	props []string
}

func (d *PropList) append(value string) {
	d.props = append(d.props, value)
}

func (d *PropList) appendWrapped(tag, value string) {
	d.append("<" + tag + ">" + value + "</" + tag + ">")
}

func (d *PropList) Bool(key string, value bool) {
	d.append("<key>" + key + "</key>")
	if value {
		d.append("<true/>")
	} else {
		d.append("<false/>")
	}
}

func (d *PropList) String(key string, value string) {
	d.appendWrapped("key", key)
	d.appendWrapped("string", value)
}

func (d *PropList) StringArray(key string, values []string) {
	d.appendWrapped("key", key)
	d.append("<array>")
	for _, value := range values {
		d.appendWrapped("string", value)
	}
	d.append("</array>")
}

func (d *PropList) Join() string {
	result := XMLHeader + "\n" + TypePropList + "\n"
	result += "<plist version=\"1.0\">" + "\n"
	result += "<dict>" + "\n"
	result += strings.Join(d.props, "\n") + "\n"
	result += "</dict>" + "\n"
	result += "</plist>" + "\n"
	return result
}

func NewPropList() PropList {
	builder := PropList{
		props: make([]string, 0),
	}
	return builder
}
