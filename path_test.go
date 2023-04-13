package godash

import "testing"

// go test -v -run Test_GetRootPath path_test.go path.go
func Test_GetRootPath(t *testing.T) {
	port, err := GetRootPath("godash")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(port)
}
