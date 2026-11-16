package app

// fontPart is a part of the app whose text the listener may make larger or
// smaller, as a whole: settings.FontScale keeps how much, by its id.
type fontPart string

const (
	// fontUI is the pages, the sidebar, the lists, dialogs and the
	// settings.
	fontUI fontPart = "ui"
	// fontPlayer is the player bar at the foot of the window.
	fontPlayer fontPart = "player"
	// fontNowPlaying is Now Playing but its lyrics: the song, the queue,
	// the comments and the wall.
	fontNowPlaying fontPart = "nowPlaying"
	// fontLyrics is the lyrics, scrolling or on the stage, the hot
	// comments in their quiet parts and the captions of episodes.
	fontLyrics fontPart = "lyrics"
	// fontNotch is the notch player, on macOS.
	fontNotch fontPart = "notch"
)

// fontParts are the parts the settings offer, in their order.
var fontParts = []struct {
	part        fontPart
	label, desc string
}{
	{fontUI, "界面", "侧边栏、各个页面、列表、对话框与设置"},
	{fontPlayer, "播放栏", "窗口底部的播放栏"},
	{fontNowPlaying, "正在播放", "歌曲信息、播放队列、评论与封面墙"},
	{fontLyrics, "歌词", "滚动歌词、舞台歌词、间奏热评与播客字幕"},
	{fontNotch, "刘海播放器", "屏幕顶部的迷你播放器"},
}

// fontScales are how much larger or smaller than usual a part's text may
// be.
var fontScales = []struct {
	scale float64
	label string
}{
	{0.85, "85%"},
	{0.9, "90%"},
	{1, "100%（默认）"},
	{1.1, "110%"},
	{1.2, "120%"},
	{1.35, "135%"},
	{1.5, "150%"},
}

// fontScale is how many times larger than usual the text of part is: 1
// unless the listener chose otherwise, and never so far either way that it
// cannot be read or laid out.
func (s *settings) fontScale(part fontPart) float32 {
	v, ok := s.FontScale[string(part)]
	if !ok || v <= 0 {
		return 1
	}
	return float32(min(max(v, 0.5), 2))
}

// setFontScale makes the text of part scale times its usual size; at 1 the
// part is forgotten, so that the settings keep only what was changed.
func (s *settings) setFontScale(part fontPart, scale float64) {
	if scale == 1 {
		delete(s.FontScale, string(part))
		if len(s.FontScale) == 0 {
			s.FontScale = nil
		}
		return
	}
	if s.FontScale == nil {
		s.FontScale = map[string]float64{}
	}
	s.FontScale[string(part)] = scale
}

// fs is the size of text of the given usual size in part, as the listener
// scaled it.
func (a *app) fs(part fontPart, size float32) float32 {
	return size * a.settings.fontScale(part)
}
