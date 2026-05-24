# yt-dlp

## install 

[yt-dlp](https://github.com/yt-dlp/yt-dlp/wiki/Installation)

## file format

video

```bash
yt-dlp -F <video_url>
yt-dlp -f 37<video_url>
(where 37 is the index of the desired format)
```

best format

```bash
yt-dlp -f bestvideo[ext=mp4]+bestaudio[ext=m4a]/bestvideo+bestaudio --merge-output-format mp4 <video_url>
```

audio

```bash
yt-dlp -x --audio-format mp3 <video_url>
```