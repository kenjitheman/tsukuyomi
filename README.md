<h3 align="center">tsukuyomi - image, video, audio compressor/re-encoder without quality loss</h3>

###

<img align="right" height="220" src="https://media.tenor.com/3wBClQGtDkgAAAAC/hououin-kyouma.gif"  />

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

### supported audio formats:

- `MP3`
- `WAV`
- `FLAC`
- `OGG`
- `AIFF`

### supported video formats:

- `MP4`
- `AVI`
- `MKV`
- `MOV`
- `WMV`
- `FLV`
- `WebM`

### supported image formats:

- `JPEG`
- `PNG`
- `GIF`
- `TIFF`
- `WebP`

## project structure

```
.
├── cmd
│   └── main.go
├── core
│   ├── audio
│   │   ├── aiff.go
│   │   ├── flac.go
│   │   ├── mp3.go
│   │   ├── ogg.go
│   │   └── wav.go
│   ├── img
│   │   ├── gif.go
│   │   ├── jpeg.go
│   │   ├── png.go
│   │   ├── tiff.go
│   │   └── webp.go
│   └── video
│       ├── avi.go
│       ├── flv.go
│       ├── mkv.go
│       ├── mov.go
│       ├── mp4.go
│       ├── webm.go
│       └── wmv.go
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## installation

- use go get ->

```
go get github.com/kenjitheman/tsukuyomi
```

- or use git clone to clone the repo ->

```
git clone https://github.com/kenjitheman/tsukuyomi
```

## usage

- you can run it using: go run main.go

```
cd cmd
go run main.go
```

- or using docker:

```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## contributing

- pull requests are welcome. For major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
