### play with HDMI audio

```bash
mplayer -ao alsa:device=hdmi [video-file]
```

### play through ssh

```bash
export DISPLAY=:0
mplayer -fs -vo xv -ao alsa:device=hdmi [video-file]
```

```bash
mplayer -fs -lavdopts threads=4 [video-file]
```

you can add a function to your `.bashrc` file 

```
mplay() { mplayer -fs -lavdopts threads=4 "$@" ;}
```

and then call it:

```bash
mplay "[video-file-name]"
```

start at a specific position

```bash
mplayer -ss [time-in-seconds]
```

alternative you can set an end position, for example from second 80 for 3 seconds

```bash
player -ss 80 -endpos 3 
```

## play directory 

```bash
mplayer -fs -lavdopts threads=4 [directory-path]/*
```

## reduce subtitle size

```bash
vim ~/.mplayer/config
```

add the following line 

```
subfont-text-scale=2
```

## keyboard control 

```
<- -> seek backward forward 10 seconds  
up and down seek backward forward 1 minute  
pgup and pgdown seek backward forward 10 minutes  
[ and ] decrease/increase current playback speed by 10%  
{ and } halve/double current playback speed  
backspace reset playback speed to normal  
< and > go backward/forward in the playlist  
enter go forward in the playlist  
p/space pause  
q quit  
+ and - adjust audio delay  
/ and * decrease/increase volume  
m mute  
f toogle fullscreen  
v toogle subtitle visibility  
j cycle trough available subtitles  
o osd information
. when paused will move frame by frame
```