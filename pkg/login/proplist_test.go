package login_test

import (
	"strings"
	"testing"
	"vpn-dns/pkg/login"
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

func assertHeader(t *testing.T, plist string) {
	t.Helper()
	if !strings.Contains(plist, "PropertyList-1.0") {
		t.Errorf("Missing header: %v", plist)
	}
}

func TestPropList(t *testing.T) {
	t.Parallel()

	boolKey := "IsItWorks"
	boolValue := true

	stringKey := "CanIAddStringValue"
	stringValue := "Yes i can"

	stringArrayKey := "CanIAddMultipleStrings"
	stringArrayValue := []string{"first", "second", "third"}

	props := login.NewPropList()
	props.Bool(boolKey, true)
	props.String(stringKey, stringValue)
	props.StringArray(stringArrayKey, stringArrayValue)
	list := props.Join()

	assertHeader(t, list)

	assertKey(t, list, boolKey)
	assertBool(t, list, boolValue)

	assertKey(t, list, stringKey)
	assertString(t, list, stringValue)

	assertString(t, list, stringArrayValue[0])
	assertString(t, list, stringArrayValue[1])
	assertString(t, list, stringArrayValue[2])
}
