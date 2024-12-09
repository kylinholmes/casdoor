package util

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/logs"
)

// / return a random avatar url or empty string if failed
func GetRandomAvatar(seed string, baseUrl string) string {
	// <img
	//   src="https://api.dicebear.com/9.x/bottts-neutral/svg?seed=Jack&scale=130&radius=10&backgroundColor=00897b,00acc1,ffd5dc&mouth=smile02,bite,diagram"
	//   alt="avatar" />
	if baseUrl[len(baseUrl)-1] == '/' {
		baseUrl = baseUrl[:len(baseUrl)-1]
	}
	ver := "9.x"
	style := "bottts-neutral"
	urltype := "svg"
	suffix := "&scale=130&radius=10&backgroundColor=00897b,00acc1,ffd5dc&mouth=smile02,bite,diagram"
	avt := fmt.Sprintf("%s/%s/%s/%s?seed=%s%s", baseUrl, ver, style, urltype, seed, suffix)

	resp, err := http.Get(avt)
	if err != nil || resp.StatusCode != 200 {
		logs.Error("GetRandomAvatar: %s", err)
		return ""
	}
	return avt
}
