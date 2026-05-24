# new linux

## before all

```bash
sudo apt update && sudo apt upgrade -y
```

## firefox

[https://markosaric.com/firefox/](https://markosaric.com/firefox/)

extensions:
- disconnect
- ublock origins

## all the stuff

```bash
sudo apt install git -y
sudo apt install vim -y
sudo apt install htop -y
sudo apt install fastfetch -y
sudo apt install ffmpeg -y
sudo apt install mplayer -y
sudo apt install vlc -y
sudo apt install cmus -y
sudo apt install mycli -y
sudo apt install transmission -y
```

> fastfetch: sudo add-apt-repository ppa:zhangsongcui3371/fastfetch

script

```bash
printf '#!/bin/bash\n\nsudo apt install git -y\nsudo apt install vim -y\nsudo apt install htop -y\nsudo apt install neofetch -y\nsudo apt install ffmpeg -y\nsudo apt install mplayer -y\nsudo apt install vlc -y\nsudo apt install cmus -y\nsudo apt install mycli -y\nsudo apt install transmission -y' > install.sh
```
```bash
sudo chmod +x install.sh
```

font scaling factor 1.15

```bash
sudo apt install gnome-tweaks
```

## github

Generate new ssh key

[https://docs.github.com/en/authentication/connecting-to-github-with-ssh](https://docs.github.com/en/authentication/connecting-to-github-with-ssh)

steps:

```bash
ssh-keygen -t ed25519 -C "<email-address>"
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519
cat ~/.ssh/id_ed25519.pub
git clone git@github.com:mamcer/spellbook.git
```

## optional

- libreoffice
- [docker](https://docs.docker.com/engine/install/ubuntu/)  
- [golang](https://go.dev/dl/)  
- [netcore](https://learn.microsoft.com/en-us/dotnet/core/install/linux-ubuntu-install)  

## programs

- video editing: kdenlive (https://kdenlive.org/), shotcut (https://www.shotcut.org/), openshot (https://www.openshot.org/)
- image editing: gimp  
- mp3 tags: easytag
- audio editing: audacity

## ubuntu after installation

remove snap

```bash
sudo systemctl stop snapd
sudo systemctl disable snapd
sudo apt purge --auto-remove snapd -y
sudo rm -rf ~/snap /snap /var/snap /var/lib/snapd

echo "Package: snapd
Pin: release a=*
Pin-Priority: -10" | sudo tee /etc/apt/preferences.d/no-snap.pref
```

if firefox disappears

```bash
sudo add-apt-repository ppa:mozillateam/ppa -y
sudo apt update
sudo apt install firefox -y
```

remove telemetry

```bash
sudo apt purge apport whoopsie popularity-contest ubuntu-report -y
sudo systemctl disable motd-news.timer
sudo systemctl mask motd-news.timer
sudo systemctl disable motd-news.service
sudo systemctl mask motd-news.service
```

install basic dev tools

```bash
sudo apt autoremove --purge -y
sudo apt install build-essential curl git wget vim gnome-tweaks software-properties-common -y
```

vs code

```bash
wget -qO- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/microsoft.gpg > /dev/null
echo "deb [arch=amd64,arm64,armhf] https://packages.microsoft.com/repos/code stable main" | sudo tee /etc/apt/sources.list.d/vscode.list
sudo apt update
sudo apt install code -y
```

optional

```bash
sudo apt purge thunderbird libreoffice* rhythmbox -y
sudo apt autoremove -y
```

## debian 13 mate

remove home folder from desktop

```bash
gsettings set org.mate.caja.desktop home-icon-visible false
```