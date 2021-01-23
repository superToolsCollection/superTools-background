package util

import "testing"

/**
* @Author: super
* @Date: 2021-01-23 20:55
* @Description:
**/

func TestGeneratePassword(t *testing.T) {
	bytes, err := GeneratePassword("admin")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(bytes))
}
