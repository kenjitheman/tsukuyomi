## Image, video, audio compressor/re-encoder library/toolkit/another abstraction

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

- ### Third party libraries this library using:

  - `github.com/chai2010/webp`
  
  - `github.com/hajimehoshi/go-mp3`
  
  - `github.com/icza/bitio`
  
  - `github.com/mewkiz/flac`
  
  - `github.com/mewkiz/pkg`

- ### Supported audio formats:

  - `MP3`
  - `WAV`
  - `FLAC`
  - `OGG`
  - `AIFF`

- ### Supported video formats:

  - `MP4`
  - `AVI`
  - `MKV`
  - `MOV`
  - `WMV`
  - `FLV`
  - `WebM`

- ### Supported image formats:

  - `JPEG`
  - `PNG`
  - `GIF`
  - `TIFF`
  - `WebP`

## Project structure

```go
tsukuyomi
│
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

## Installation

```sh
go get github.com/kenjitheman/tsukuyomi
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## License

- [MIT](https://choosealicense.com/licenses/mit/)
