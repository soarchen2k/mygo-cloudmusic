package app

import (
	"testing"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

// Choosing a size for 界面 in the settings scales the text of the settings
// themselves and no other part, and 恢复默认 brings it back.
func TestChooseFontScale(t *testing.T) {
	defer func(f func(mygo.ThemeSource)) { setThemeSource = f }(setThemeSource)
	setThemeSource = func(mygo.ThemeSource) {}
	a := newTestApp()
	a.settings = defaultSettings()
	tt := ui.NewTester(a.settingsView, 560, 3000)
	tt.Frame()
	usual, ok := tt.Find("实验性功能")
	if !ok {
		t.Fatal("no group 实验性功能")
	}
	if tt.HasText("恢复默认") {
		t.Fatal("恢复默认 shows with every part at its usual size")
	}

	if err := tt.Click("界面字体大小"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("150%"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if got := a.settings.fontScale(fontUI); got != 1.5 {
		t.Fatalf("界面 at %v after choosing 150%%", got)
	}
	if got := a.settings.fontScale(fontLyrics); got != 1 {
		t.Fatalf("歌词 at %v after choosing a size for 界面", got)
	}
	if r, _ := tt.Find("实验性功能"); r.H < usual.H*1.3 {
		t.Fatalf("the settings' text %v high at 150%%, %v as usual", r.H, usual.H)
	}

	if err := tt.Click("恢复默认"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if a.settings.FontScale != nil {
		t.Fatalf("sizes %v kept after 恢复默认", a.settings.FontScale)
	}
	if r, _ := tt.Find("实验性功能"); r.H != usual.H {
		t.Fatalf("the settings' text %v high after 恢复默认, %v as usual", r.H, usual.H)
	}
}
