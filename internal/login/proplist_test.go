package login_test

import (
	"strings"
	"testing"
	"vpn-dns/internal/login"
)

func assertKey(t *testing.T, plist string, k string) {
	t.Helper()
	if !strings.Contains(plist, "<key>"+k+"</key>") {
		t.Errorf("Missing key %s", k)
	}
}

func assertBool(t *testing.T, plist string, v bool) {
	t.Helper()
	var tag string
	if v {
		tag = "<true/>"
	} else {
		tag = "<false/>"
	}
	if !strings.Contains(plist, tag) {
		t.Errorf("Missing bool tag %s", tag)
	}
}

func assertString(t *testing.T, plist string, v string) {
	t.Helper()
	if !strings.Contains(plist, "<string>"+v+"</string>") {
		t.Errorf("Missing string %s", v)
	}
}

func TestPropList(t *testing.T) {
	t.Parallel()
	props := login.NewPropList()
	props.Bool("IsItWorks", true)
	props.String("CanIAddStringValue", "Yes i can")
	props.StringArray("CanIAddMultipleStrings", []string{"first", "second", "third"})
	list := props.Join()
	if !strings.Contains(list, "PropertyList-1.0") {
		t.Errorf("Missing header: %v", list)
	}
	assertKey(t, list, "IsItWorks")
	assertBool(t, list, true)
	assertKey(t, list, "CanIAddStringValue")
	assertString(t, list, "Yes i can")
	assertString(t, list, "first")
	assertString(t, list, "second")
	assertString(t, list, "third")
}
