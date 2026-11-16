#!/bin/sh
# Puts the libraries of the tray icon into resources/linux-amd64/lib, for
# `mygo build` to ship next to the Linux app. Systems without them, as
# SteamOS, whose system cannot be changed, get no tray icon otherwise: the
# app then starts again with them (see main_linux.go). They come from
# Ubuntu's packages, so this runs on Ubuntu or Debian, as the release build
# does; their glibc is the oldest they load with.
set -eu

cd "$(dirname -- "$0")/.."
out=resources/linux-amd64/lib
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

(cd "$work" && apt-get download libayatana-appindicator3-1:amd64 libayatana-indicator3-7:amd64 libayatana-ido3-0.4-0:amd64)
for deb in "$work"/*.deb; do
	dpkg-deb -x "$deb" "$work/root"
done

rm -rf "$out"
mkdir -p "$out"
# Under the names the system loads them by, as real files: links would not
# survive the archives.
for name in libayatana-appindicator3.so.1 libayatana-indicator3.so.7 libayatana-ido3-0.4.so.0; do
	cp -L "$work/root/usr/lib/x86_64-linux-gnu/$name" "$out/$name"
done
ls -l "$out"
