# ffmpeg

[https://evanhahn.github.io/ffmpeg-buddy/](https://evanhahn.github.io/ffmpeg-buddy/)

## convert video

```bash
ffmpeg -i input.mp4 output.gif
```

> gif, mkv, avi, mp3, wmv, mp4
> example: ffmpeg -i input.wav  -vn -ar 44100 -ac 2 -b:a 192k output.mp3

## scale %

```bash
ffmpeg -i input.mp4 -vf 'scale=iw*0.3:ih*0.2' output.mp4
```
> 30% width and 20% height in the example

you can also keep aspet ratio

```bash
ffmpeg -i input.mp4 -vf 'scale=iw*0.5:-1' output.mp4
```
> in the example reduce by 50% width and keep height aspect ratio

## rotate 90 counter-clockwise

```bash
ffmpeg -i input.mp4 -vf 'transpose=2' output.mp4
```
## rotate 90 clockwise

```bash
ffmpeg -i input.mp4 -vf 'transpose=1' output.mp4
```

## horizontal flip 

```bash
ffmpeg -i input.mp4 -vf hflip output.mp4
```

## vertical flip 

```bash
ffmpeg -i input.mp4 -vf vflip output.mp4
```

## time crop

```bash
ffmpeg -i input.mp4 -ss '00:01:16' -to '00:02:16' output.mp4
```
> from 1m:16s to 2m:16s on example

## only audio 

```bash
ffmpeg -i input.mp4 -vn output.mp3
```

## only video

```bash
ffmpeg -i input.mp4 -an output.mp4
```

## dvd to mp4

combine two vob files to a mp4

```bash
ffmpeg -i "concat:/media/mamcer/DVDIRECT_DISC_0010002FD43/VIDEO_TS/VIDEO_TS.VOB|/media/mamcer/DVDIRECT_DISC_0010002FD43/VIDEO_TS/VTS_01_1.VOB" -b:v 1500k -r 30 -vcodec h264 -strict -2 -acodec aac -ar 44100 -f mp4 2016-03-31_eco-ciro.mp4
```

## merge avi files

```bash
ffmpeg -i "concat:input1.avi|input2.avi|input3.avi" -c copy output.avi
```

## web friendly

[https://jshakespeare.com/encoding-browser-friendly-video-files-with-ffmpeg/](https://jshakespeare.com/encoding-browser-friendly-video-files-with-ffmpeg/)

mp4

```bash
ffmpeg -i video.wmv -vcodec libx264 -f mp4 -vb 1024k -preset slow new-video.mp4
```
> it didn't work for me on mobile (video.js)

webm 

```bash
ffmpeg -i video.mp4 -f webm -vcodec libvpx-vp9 new-video.webm
```

## script to convert of video files in current directory to mp4 

with 'high' quality (in most cases human can tell the difference and the file size is drastically reduced compared to lossless formats)

```bash
#!/bin/bash

echo "video conversion started..."

for file in *.avi *.mkv *.mov *.wmv *.flv *.AVI *.MKV *.MOV *.WMV *.FLV; do
    if [ -f "$file" ]; then
        # get the filename without the extension
        filename="${file%.*}"

        # convert to mp4
        ffmpeg -i "$file" -c:v libx264 -crf 18 -c:a aac -b:a 192k "${filename}.mp4"
    fi
done

echo "video conversion finished."
```

if you really want lossless, replace the convert video with this line:

```bash
# convert to mp4
ffmpeg -i "$file" -c:v libx264 -crf 0 -c:a aac -b:a 192k "${filename}.mp4"
```
> in a quick test with a small file it increases the output file size by almost 5 times