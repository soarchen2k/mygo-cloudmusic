package netease

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// An eapi request's User-Agent names the same client as its header. It
// once named an iPhone beside the Windows client's os, osver and appver,
// and NetEase refused many songs with -460 for it.
func TestEapiSpeaksAsOneClient(t *testing.T) {
	var ua string
	var cookies []*http.Cookie
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua, cookies = r.UserAgent(), r.Cookies()
		w.Write([]byte(`{"code":200}`))
	}))
	defer srv.Close()
	defer func(d string) { eapiDomain = d }(eapiDomain)
	eapiDomain = srv.URL
	deviceOnce.Do(func() { deviceID = "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123" })

	var res struct{ Code int }
	if err := eapiPost("/api/song/enhance/player/url/v1", map[string]any{"ids": "[1]"}, &res); err != nil {
		t.Fatal(err)
	}
	header := map[string]string{}
	for _, c := range cookies {
		header[c.Name] = c.Value
	}
	if header["os"] != "pc" || !strings.Contains(header["osver"], "Windows") {
		t.Fatalf("header is not the Windows client's: os %q osver %q", header["os"], header["osver"])
	}
	if !strings.Contains(ua, "Windows NT") || !strings.Contains(ua, "NeteaseMusicDesktop/"+header["appver"]) {
		t.Fatalf("User-Agent %q is not the Windows client of appver %q", ua, header["appver"])
	}
}
