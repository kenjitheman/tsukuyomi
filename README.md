<h3 align="center">tsukuyomi - image, video, audio compressor/re-encoder library</h3>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

- ### third party libraries this library using:

  - `github.com/chai2010/webp`
  
  - `github.com/hajimehoshi/go-mp3`
  
  - `github.com/icza/bitio`
  
  - `github.com/mewkiz/flac`
  
  - `github.com/mewkiz/pkg`

- ### supported audio formats:

  - `MP3`
  - `WAV`
  - `FLAC`
  - `OGG`
  - `AIFF`

- ### supported video formats:

  - `MP4`
  - `AVI`
  - `MKV`
  - `MOV`
  - `WMV`
  - `FLV`
  - `WebM`

- ### supported image formats:

  - `JPEG`
  - `PNG`
  - `GIF`
  - `TIFF`
  - `WebP`

## project structure

```
├── audio
│   ├── aiff.go
│   ├── flac.go
│   ├── mp3.go
│   ├── ogg.go
│   └── wav.go
├── go.mod
├── img
│   ├── gif.go
│   ├── jpeg.go
│   ├── png.go
│   ├── tiff.go
│   └── webp.go
├── LICENSE
├── README.md
└── video
    ├── avi.go
    ├── flv.go
    ├── mkv.go
    ├── mov.go
    ├── mp4.go
    ├── webm.go
    └── wmv.go
```

## installation

```
go get github.com/kenjitheman/tsukuyomi
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
