# ubuntu server install 

## version

```bash
cat /etc/os-release
```

## release

```bash
cat /etc/lsb-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=22.04
DISTRIB_CODENAME=jammy
DISTRIB_DESCRIPTION="Ubuntu 22.04.3 LTS"
```

## apt update-upgrade

```bash
sudo apt update
sudo apt upgrade
```

## install sudo (debian)

```bash
su
apt install sudo
sudo adduser [username] sudo
exit
```

## kernel version

```bash
uname --all
Linux nuc 5.15.0-91-generic #101-Ubuntu SMP Tue Nov 14 13:30:08 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux
```

## timezone

```bash
timedatectl
```
```
Local time: Sat 2024-01-20 13:17:54 UTC
Universal time: Sat 2024-01-20 13:17:54 UTC
RTC time: Sat 2024-01-20 13:17:54
Time zone: Etc/UTC (UTC, +0000)
System clock synchronized: yes
NTP service: active
RTC in local TZ: no
```
```bash
sudo timedatectl set-timezone America/Buenos_Aires
```

```
Local time: Sat 2024-01-20 10:19:03 -03
Universal time: Sat 2024-01-20 13:19:03 UTC
RTC time: Sat 2024-01-20 13:19:03
Time zone: America/Buenos_Aires (-03, -0300)
System clock synchronized: yes
NTP service: active
RTC in local TZ: no
```

## ssh server

```bash
sudo apt install openssh-server
sudo systemctl enable ssh
sudo systemctl start ssh    
ip a
```

connect to a server without need to write your credential every time

```bash
ssh-copy-id username@server_address
```

## setup wifi 

Debian 12 example

```bash
ip a
```
> to get the name of the network interface 

```bash
vi /etc/network/interfaces

auto wlan0
iface wlan0 inet dhcp
    wpa-ssid JCMM
    wpa-psk [password]
```
> auto make it auto connect at start

example oculus (intel stick, debian 12.5)

```bash
# This file describes the network interfaces available on your system
# and how to activate them. For more information, see interfaces(5).
source /etc/network/interfaces.d/*
# The loopback network interface
#auto lo
#iface lo inet loopback
# The primary network interface
#allow-hotplug wlan0
auto wlan0
iface wlan0 inet static
  wpa-ssid [network-ssid]
      wpa-psk  [password]
      address 192.168.100.101
      netmask 255.255.255.0
      gateway 192.168.100.1
      dns-nameservers 8.8.8.8 8.8.4.4
```

> logout to take changes

##  static ip

```bash
sudo ip a
```

> look for interface name

```bash
sudo vim /etc/netplan/01-netcfg.yaml
```

```bash
network:
  version: 2
  renderer: networkd
  ethernets:
      enp2s0:
        dhcp4: no
        addresses:
          - 192.168.100.100/24
        nameservers:
          addresses: [8.8.8.8, 8.8.4.4]
        routes:
            - to: default
            via: 192.168.10.1
```

```bash
sudo netplan apply
```

Debian 12 example over Wifi

```bash
vi /etc/network/interfaces
```

```bash
auto wlan0
iface wlan0 inet static
    wpa-ssid JCMM
    wpa-psk [password]
    address 192.168.100.101
    netmask 255.255.255.0
    gateway 192.168.100.1
    dns-nameservers 8.8.8.8 8.8.4.4
```

## disk space

```bash
df
df -h 
df -h /
df -h /dev/sdb1
```

## torrents (old)

[https://cli-ck.io/transmission-cli-user-guide/](https://cli-ck.io/transmission-cli-user-guide/)

## close the lid do-nothing

```bash
echo 'HandleLidSwitch=ignore' | sudo tee --append /etc/systemd/logind.conf
echo 'HandleLidSwitchDocked=ignore' | sudo tee --append /etc/systemd/logind.conf
service systemd-logind restart
```

## battery power?

```bash
upower -i /org/freedesktop/UPower/devices/battery_BAT0
```

## mount 

```bash
lsblk
sudo mount /dev/sdb1 /media/rose/
udisksctl unmount -b /dev/sdb1
udisksctl power-off -b /dev/sdb1
```

## share directory

```bash
sudo apt install -y samba
sudo smbpasswd -a mario
```

```bash
sudo vim /etc/samba/smb.conf		
```

```bash
[music]
path=/home/mario/music
available=yes
valid users=mario
read only=yes
browsable=yes
public=yes
writable=no    
```

```bash
sudo service smbd restart
```

```bash
smb://[user]@[ip-address]/music
```

## golang

```bash
wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz
vim ~/.profile
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$Home/go:$Home/src
source ~/.profile
go version
go version go1.17.6 linux/amd64
```

## veracrypt

```bash
wget https://launchpad.net/veracrypt/trunk/1.26.7/+download/veracrypt-console-1.26.7-Ubuntu-22.04-amd64.deb
sudo dpkg -i veracrypt-console-1.26.7-Ubuntu-22.04-amd64.deb
```

> veracrypt-console depends on libfuse2; however:
> Package libfuse2 is not installed.
> veracrypt-console depends on pcscd; however:
> Package pcscd is not installed.
> sudo apt --fix-broken install
> run again dpkg -i ...

```bash
veracrypt --help
```

```bash
veracrypt /dev/sdb1 /media/darkforce/
veracrypt -d 
```

create file container

```bash
veracrypt -t --create
```

You'll then be prompted for:

```
Volume type: Normal or Hidden.
Volume path: The file path for your new container (e.g., /home/user/my_encrypted_data.hc).
Volume size: The size of the container (e.g., 1G for 1 Gigabyte, 500M for 500 Megabytes).
Encryption Algorithm: (e.g., AES, Serpent, Twofish, etc.)
Hash Algorithm: (e.g., SHA-512, Whirlpool, etc.)
Filesystem: (e.g., Ext4, FAT, NTFS, ExFAT, None).
Password: Your chosen password for the volume.
PIM (Personal Iterations Multiplier): Usually leave this as 0 unless you have a specific reason to change it.
Keyfiles: If you want to use keyfiles in addition to or instead of a password.
Random characters: You'll be asked to type random characters or use a random source to generate cryptographic strength.
```

or

```bash
sudo veracrypt -t --create /path/to/your/volume.vc \
  --volume-type=normal \
  --size=1G \
  --encryption=AES \
  --hash=SHA-512 \
  --filesystem=ext4 \
  -p "YourStrongPasswordHere" \
  --pim=0 \
  --quick
```

## android 

```bash
    sudo apt install -y jmtpfs
    mkdir -p ~/android/xiaomi
    jmtpfs ~/android/xiaomi
```

## docker

[https://docs.docker.com/engine/install/ubuntu/](https://docs.docker.com/engine/install/ubuntu/)

```bash
# after installation
sudo groupadd docker
sudo usermod -aG docker $USER
```

## automatic session

```bash
sudo -e /etc/lightdm/lightdm.conf
```

```bash
[Seat:*]
autologin-session=xubuntu
autologin-user=mario
autologin-user-timeout=0
```

## jobs

example

```bash
vim notes.txt &
[1] 17965
jobs
[1]+  Stopped                 vim notes.txt
fg %1
```

basically, with '&' at the end of a command and you will create a job in 'stopped' state. With 'jobs' you can see the status and send a job to foreground with 'fg %n' where n is the job number or 'bg %n' to change his status to 'running' (in background)

Also you can create a job by pressing ctrl-z when running an application or command

you can kill a stopped or running in background job with 'kill %n'

## create booteable flash drive

```bash
lsblk
// make sure it is not mounted
sudo dd bs=4M if=/home/mario/Downloads/ubuntu-20.04.3-live-server-amd64.iso of=/dev/sdb1 status=progress oflag=sync
```

## create iso file from cd/dvd

```bash
lsblk
// locate your cd/dvd drive, example sr0
sudo dd if=/dev/sr0 of=backup-01.iso bs=2048 status=progress
```

## rename server

```bash
hostnamectl set-hostname new-hostname
sudo vim /etc/hosts
```

## static hostnames

```bash
sudo vim /etc/hosts
```

Add a new line like: 10.0.64.7	duriel