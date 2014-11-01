// iptoloc_test.go
package ipdb

import (
	"strings"
	"testing"
)

func TestIPToLocation(t *testing.T) {
	var result string
	var err error
	var validIPdata = map[string]string{
		"0.0.0.0": "保留地址	保留地址",
		"127.0.0.1": "本机地址	本机地址",
		"171.34.113.168": "中国	江西	南昌",
		"223.27.38.232": "中国	台湾",
		"112.64.137.205": "中国	上海	上海"}

	var invalidIpdata = map[string]string{
		"123.456.789.0":   "",
		"a.b.c.d":         "",
		"123":             "",
		"256.256.256.256": ""}

	for ip, location := range validIPdata {
		result, err = IPToLocation(ip)
		if err != nil ||
			location != strings.TrimSpace(result) {
			t.Log("Testing Failed on ip: " + ip)
			t.Log(err)
			t.Fail()
		}
	}

	for ip, _ := range invalidIpdata {
		result, err = IPToLocation(ip)
		if err == nil {
			t.Log("Testing Failed on ip: " + ip)
			t.Log(err)
			t.Fail()
		}
	}
}
