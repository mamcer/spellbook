# sorceress

## performance

list enabled services

```bash
systemctl list-unit-files --type=service --state=enabled
```

printer

```bash
sudo systemctl disable --now cups.service
sudo systemctl disable --now cups-browsed.service
```

modem

```bash
sudo systemctl disable --now ModemManager.service
```

reduce memory pressure

Debian 13 no longer reads `/etc/sysctl.conf` at boot; use a drop-in under `/etc/sysctl.d/`.

```bash
sudo vim /etc/sysctl.d/99-sorceress.conf
```

add

```conf
# high swappiness: swap goes to zram (compressed RAM), not to disk
vm.swappiness=100
vm.vfs_cache_pressure=50
```

apply

```bash
sudo sysctl --system
```

clean system

```bash
sudo apt autoremove --purge
sudo apt clean
```

## golang

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOBIN=$HOME/.local/bin
export PATH=$PATH:$GOBIN
```

## caja

```bash
sudo apt install caja-open-terminal
caja -q
```

## video player

```bash
sudo apt install mpv
```

```bash
mkdir -p ~/.config/mpv
vim ~/.config/mpv/mpv.conf
```

```
hwdec=auto
framedrop=yes
keep-open=yes
```

### cheat sheet

```
# cycle audio tracks
9/0 volume down/up
m mute
j cycle subtitle tracks
J subtitles ON/OFF
v toggle subtitle visibility
← / → seek 5 seconds
↓ / ↑ seek 1 minute
shift + ← / → seek 1 second
shift ↓ / ↑ seek 5 minutes
space play/pause
. frame by frame forward
, frame by frame backward
[/] decrease increase playback speed
backspace reset speed to normal
f toggle full screen
0 show progress bar
```


```bash
yt-dlp -f "bestvideo[height<=1080]+bestaudio/best[height<=1080]" -o - "URL_DEL_VIDEO" | mpv -
```

## bluetooth

```bash
sudo apt install bluez bluetooth rfkill blueman
sudo systemctl enable --now bluetooth
systemctl status bluetooth
```

## ncspot

```
https://github.com/hrkfdn/ncspot/releases
```

```bash
tar xzf ncspot-v1.3.3-linux-x86_64.tar.gz
sudo mv ncspot /usr/local/bin/
sudo chmod +x /usr/local/bin/ncspot
mkdir -p ~/.config/ncspot
vim ~/.config/ncspot/config.toml
```

```
bitrate = 192
audio_cache = true
cache_size = 2048
backend = "pulseaudio"
gapless = true
shuffle = false
repeat = "off"
notify = false
volnorm = false
```

### cheat sheet

```
o open playlist
backspace go back
q add to queue
```


## zram-tools

```bash
sudo apt update && sudo apt install zram-tools
```

```bash
sudo vim /etc/default/zramswap
```

Uncomment and adjust lines:

```
ALGO=lz4
PERCENT=50
PRIORITY=100
```

```bash
sudo systemctl restart zramswap
```

verify (ALGORITHM should show `lz4`)

```bash
zramctl
```

## minimize logs

```bash
sudo vim /etc/systemd/journald.conf
```

uncomment line `#SystemMaxUse=` and set value to `SystemMaxUse=50M`

```bash
sudo systemctl restart systemd-journald
```

## tmux

```bash
sudo apt update
sudo apt install tmux
```

```bash
tmux                 # new session
ctrl+b, d            # detach session
tmux ls              # list sessions
tmux attach          # reconnect last session
tmux attach -t 0     # reconnect specific session
tmux kill-session -t 0
ctrl+b c             # new window
ctrl+b |             # vertical split
ctrl+b -             # horizontal split
```

config

```bash
vim ~/.tmux.conf
```

```
set -g mouse on
setw -g mode-keys vi
set -g history-limit 10000

unbind '"'
unbind %

bind | split-window -h
bind - split-window -v

bind h select-pane -L
bind j select-pane -D
bind k select-pane -U
bind l select-pane -R

bind r source-file ~/.tmux.conf \; display "Reloaded!"
```

## alacritty

```bash
sudo apt install alacritty
```

```bash
sudo apt install fonts-jetbrains-mono
```

```bash
mkdir -p ~/.config/alacritty
vim ~/.config/alacritty/alacritty.toml
```

```
[window]
# inner padding so text doesn't touch the edges
padding = { x = 10, y = 10 }
startup_mode = "Fullscreen"
decorations = "None"
opacity = 0.95

[font]
# JetBrains Mono, very readable for code
normal = { family = "JetBrains Mono", style = "Regular" }
bold = { family = "JetBrains Mono", style = "Bold" }
size = 13.0 # adjust to taste

[colors.primary]
background = "#1e1f29"
foreground = "#f8f8f2"

[colors.normal]
black   = "#000000"
red     = "#ff5555"
green   = "#50fa7b"
yellow  = "#f1fa8c"
blue    = "#bd93f9"
magenta = "#ff79c6"
cyan    = "#8be9fd"
white   = "#f8f8f2"

[colors.bright]
black   = "#44475a"
red     = "#ff6e6e"
green   = "#69ff94"
yellow  = "#ffffa5"
blue    = "#d6acff"
magenta = "#ff92df"
cyan    = "#a4ffff"
white   = "#ffffff"

[cursor]
# blinking block cursor
style = { shape = "Block", blinking = "On" }

[terminal.shell]
# force bash
program = "/bin/bash"
```

## maria-db

```bash
sudo apt install mariadb-server
sudo systemctl enable --now mariadb
sudo vim /etc/mysql/mariadb.conf.d/50-server.cnf
```

```
innodb_buffer_pool_size = 256M
innodb_log_file_size = 64M
max_connections = 20
thread_cache_size = 8
tmp_table_size = 32M
max_heap_table_size = 32M
```

```bash
sudo mysql_secure_installation
```

```bash
sudo mariadb
CREATE USER 'mario'@'localhost' IDENTIFIED VIA unix_socket;
GRANT ALL PRIVILEGES ON *.* TO 'mario'@'localhost';
FLUSH PRIVILEGES;
```

```bash
mysql -u mario
```

## vscode

user settings

```json
{
  // =========================
  // MINIMAL UI
  // =========================
  "window.zoomLevel": 1.7,
  "workbench.colorTheme": "Shades of Purple",
  "workbench.activityBar.location": "hidden",
  "window.commandCenter": false,
  "workbench.layoutControl.enabled": false,
  "breadcrumbs.enabled": false,
  "editor.minimap.enabled": false,
  "editor.codeLens": false,
  "editor.glyphMargin": false,
  "editor.stickyScroll.enabled": false,
  "editor.renderLineHighlight": "none",
  "editor.occurrencesHighlight": "off",
  "editor.selectionHighlight": false,
  "editor.overviewRulerBorder": false,
  "editor.hideCursorInOverviewRuler": true,
  "editor.folding": false,
  "editor.guides.indentation": false,
  "editor.guides.highlightActiveIndentation": false,
  "editor.matchBrackets": "never",
  "workbench.tips.enabled": false,
  "extensions.ignoreRecommendations": true,
  "workbench.startupEditor": "none",

  // =========================
  // PERFORMANCE
  // =========================

  "editor.semanticHighlighting.enabled": false,
  "editor.lightbulb.enabled": "off",
  "editor.inlineSuggest.enabled": false,
  "editor.quickSuggestions": {
    "comments": "off",
    "strings": "off",
    "other": "off"
  },
  "editor.parameterHints.enabled": false,
  "editor.hover.enabled": "off",
  "editor.suggestOnTriggerCharacters": false,
  "editor.acceptSuggestionOnCommitCharacter": false,
  "editor.wordBasedSuggestions": "off",

  // =========================
  // FILE WATCHERS
  // =========================

  "files.watcherExclude": {
    "**/.git/**": true,
    "**/node_modules/**": true,
    "**/dist/**": true,
    "**/build/**": true,
    "**/.next/**": true,
    "**/coverage/**": true
  },

  "search.exclude": {
    "**/node_modules": true,
    "**/dist": true,
    "**/build": true,
    "**/.next": true,
    "**/coverage": true
  },

  // =========================
  // GIT OFF
  // =========================

  "git.enabled": false,
  "git.autorefresh": false,
  "git.decorations.enabled": false,

  // =========================
  // TELEMETRY
  // =========================

  "telemetry.telemetryLevel": "off",

  // =========================
  // TERMINAL
  // =========================

  "terminal.integrated.gpuAcceleration": "on",
  "terminal.integrated.smoothScrolling": false,

  // =========================
  // EXPLORER
  // =========================

  "explorer.decorations.badges": false,
  "explorer.decorations.colors": false,

  // =========================
  // SCROLL
  // =========================

  "editor.smoothScrolling": false,
  "workbench.list.smoothScrolling": false,

  // =========================
  // AUTO SAVE
  // =========================

  "files.autoSave": "off",

  "editor.largeFileOptimizations": true,
  "files.simpleDialog.enable": true,
  "workbench.enableExperiments": false,
  "chat.titleBar.signIn.enabled": false,
  "workbench.browser.showInTitleBar": false,
  "workbench.editor.enablePreview": false,
  "update.showReleaseNotes": false
}
```