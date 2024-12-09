package util

import (
	"fmt"
	"testing"
)

func TestAvatarUrl(t *testing.T) {
	baseUrl := "http://58.87.88.132:3000/"
	t.Run("TestAvatarUrl", func(t *testing.T) {
		url := GetRandomAvatar("Jack", baseUrl)
		fmt.Println(url)
		t.Log(url)
	})
}
