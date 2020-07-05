# Ebiten

[![Build Status](https://travis-ci.org/hajimehoshi/ebiten.svg?branch=master)](https://travis-ci.org/hajimehoshi/ebiten)
[![Go Report Card](https://goreportcard.com/badge/github.com/hajimehoshi/ebiten)](https://goreportcard.com/report/github.com/hajimehoshi/ebiten)

**一个简单的2D游戏库**

Ebiten是一个基于Go的开源游戏库。Ebiten的简单API允许您快速轻松地开发跨平台的2D游戏。

* [官方主页 (ebiten.org)](https://ebiten.org)
* [API 手册](https://pkg.go.dev/github.com/hajimehoshi/ebiten)
* [Cheat Sheet](https://ebiten.org/documents/cheatsheet.html)

![Overview](https://ebiten.org/images/overview1.11.png)

## Platforms

* [Windows](https://ebiten.org/documents/install.html?os=windows) (No Cgo!)
* [macOS](https://ebiten.org/documents/install.html?os=darwin)
* [Linux](https://ebiten.org/documents/install.html?os=linux)
* [FreeBSD](https://ebiten.org/documents/install.html?os=freebsd)
* [Android](https://ebiten.org/documents/mobile.html)
* [iOS](https://ebiten.org/documents/mobile.html)
* Web browsers (Chrome, Firefox, Safari and Edge)
  * [WebAssembly](https://ebiten.org/documents/webassembly.html)
  * [GopherJS](https://ebiten.org/documents/gopherjs.html)

Note: Gamepad and keyboard are not available on Android/iOS.

For installation on desktops, see [the installation instruction](https://ebiten.org/documents/install.html).

## Features

* 2D Graphics (Geometry/Color matrix transformation, Various composition modes, Offscreen rendering, Fullscreen, Text rendering, Automatic batches, Automatic texture atlas)
* Input (Mouse, Keyboard, Gamepads, Touches)
* Audio (Ogg/Vorbis, MP3, WAV, PCM)

## Packages

* [ebiten](https://pkg.go.dev/github.com/hajimehoshi/ebiten)
  * [audio](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio)
    * [mp3](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio/mp3)
    * [vorbis](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio/vorbis)
    * [wav](https://pkg.go.dev/github.com/hajimehoshi/ebiten/audio/wav)
  * [ebitenutil](https://pkg.go.dev/github.com/hajimehoshi/ebiten/ebitenutil)
  * [inpututil](https://pkg.go.dev/github.com/hajimehoshi/ebiten/inpututil)
  * [mobile](https://pkg.go.dev/github.com/hajimehoshi/ebiten/mobile)
  * [text](https://pkg.go.dev/github.com/hajimehoshi/ebiten/text)

## Community

### Slack

`#ebiten` channel in [Gophers Slack](https://blog.gopheracademy.com/gophers-slack-community/)

## License

Ebiten is licensed under Apache license version 2.0. See LICENSE file.
