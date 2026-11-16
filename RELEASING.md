# Releasing

GitHub holds only the source. There are no GitHub releases, no release builds in Actions and no `v*` tags on the remote: a new version is built and installed on your own machine.

1. Commit pending work, bring main up to date with origin, and run `go test ./internal/...`. Others push to main too, so a version cut from a stale main leaves out their work:

   ```bash
   git pull --rebase origin main
   ```

2. Bump `version` in `mygo.json`, and add a `## 0.1.27` section at the top of `CHANGELOG.md` saying, in Chinese, what changed for people using the app since the last version. Commit both in a commit titled just the version, e.g. `0.1.27`, and push main. `mygo build` fails without the changelog section.

3. Build. On Linux, fetch the tray libraries first, or the app ships without them and has no tray icon where the system lacks them (SteamOS):

   ```bash
   sh scripts/tray-libs.sh
   go tool mygo build -platform linux/amd64
   ```

4. Install the archive it made, into `~/.local/mygo-cloudmusic.app` with its entry in the applications menu:

   ```bash
   sh build/linux-amd64/install.sh build/linux-amd64/mygo-cloudmusic-0.1.27-linux-amd64.tar.gz
   ```

   Quit the running app first, then open it from the menu and check 设置 → 更新 shows the new version.

The app's updater, the install command on the website and the website's download page all look for the latest GitHub release, of which there is none: they find nothing to update to or download.

## The speech recognition libraries

Podcast captions run sherpa-onnx, whose libraries the app downloads with the model from the `sherpa-onnx-<version>` release. With no releases on the remote that download has nowhere to come from. `go run ./scripts/sherpa` still packs them into `build/sherpa` and writes `internal/speech/natives.go` when sherpa-onnx's version changes in `scripts/sherpa`.
