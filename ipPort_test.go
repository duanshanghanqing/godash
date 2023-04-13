package godash

import "testing"

// go test -v -run Test_GetFreePort ipPort_test.go ipPort.go
func Test_GetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(port)
}

// go test -v -run Test_GetIPv4 ipPort_test.go ipPort.go
func Test_GetIPv4(t *testing.T) {
	ip, err := GetIPv4()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(ip.String())
}
