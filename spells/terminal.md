# terminal

## scripts

### bru 

```bash
#!/bin/bash

GREEN='\033[0;32m'
PURPLE='\033[0;35m'
NC='\033[0m'

print() {
    printf "${PURPLE}$1${NC}: ${2}\n"
}

mult() {
    value=$(echo "scale=2;$1*$2" | bc -l)
    echo $value
    #return $value
}

if [ $# -eq 0 ]; then
    echo "usage:"
    echo "bru.sh [rate]"
else
    amount=200
    rate=$1
    afipt=0.35
    totalb=$(mult $amount $rate)
    afip=$(mult $totalb 0.35)
    pais=$(mult $totalb 0.3)
    totaln=$(echo "scale=2;$totalb+$afip+$pais" | bc -l)
    
    print "usd" $amount
    print "rate" $rate
    print "total bruto" $totalb
    print "afip" $afip
    print "pais" $pais
    print "total neto" $totaln
fi
```

### cbt

```bash
#!/bin/bash

GREEN='\033[0;32m'
PURPLE='\033[0;35m'
NC='\033[0m'

call() {
    if [ "$3" == "" ]; then
        printf "${PURPLE}$1${NC}\n${2} | jq .\n"
        curl -s $2 | jq .
    else
        printf "${PURPLE}$1${NC}\n${2} | jq ${3}\n"
        curl -s $2 | jq $3
    fi
}

if [ "$1" == "-o" ]; then
    url='[url]'$2'[parameters]'
    jq='{id:.id,date_created:.date_created,last_updated:.last_updated}'
    call order $url $jq
elif [ "$1" == "-m" ]; then
    url='[url]'$2'[parameters]'
    call movement $url
elif [ "$1" == "-i" ]; then
    url='[url]'$2'[parameters]'
    json=$(curl -s $url)
    echo $json | jq '{id:.id,title:.title,date_created:.date_created}'
    echo $json | jq '.sale_terms[] | select(.id == "ID_X") | {id:.id,value:.value_struct.number,unit:.value_struct.unit}'
else
    echo "usage:"
    echo "-o  [id]        get o item"
    echo "-m  [id]        get m item"
    echo "-i  [id]        get i item"
fi
```

### loop

```bash
#!/bin/bash

for id in 1 2 3

do
    curl -s -X GET '[url]'$id'[parameters]' | jq '{id:.id,registration_date:.registration_date,tags:.tags}'

    curl -s -X DELETE '[url]'$id'[parameters]'$id'[more-parameters]' | jq .
done

exit 0
```

## replace underscore by spaces

example: archivo_de_texto.txt >  archivo de texto.txt

```bash
for file in * ; do mv -v "$file" "$(echo $file | sed 's/_/\\ /g')  " ; done

$ echo "$a" | tr '[:upper:]' '[:lower:]'
```

## copy with progress

```bash
rsync -ha --progress [source] [destination]
```

## copy files recursively to another directory

```bash
find . -name '*.3gp' -print -exec mv -t /media/mario/4EFE9676FE9656552/pictures/moto-x-yani {} +
```

## rename multiple files

```bash
find . -name "*.JPEG" -exec rename 's/\.JPEG$/.jpeg/' '{}' \;
```

## rename files, add prefix (in the example, underscore before previous file name)

```bash
for file in *.jpg; do mv "$file" "_$file"; done;
```

## rename multiple files

in this example, prefix with "2016-03_"

```bash
for f in *; do [[ -f "$f" ]] && mv "$f" "2016-03_$f"; done
```

## recursively remove files

```bash
find . -name "*.bak" -type f
```

```bash
find . -name "*.bak" -type f -delete
```

## search and replace text on file

```bash
sed -i 's/localhost:/192.168.1.55:/g' main.js
```

## show all extensions in a directory

```bash
find . -type f | sed -rn 's|.*/[^/]+\.([^/.]+)$|\1|p' | sort -u
```

## search specific file content recursively

```bash
grep -iRl "your-text-to-find" ./
```

Another option

```bash
grep -rnw '/path/to/somewhere/' -e 'pattern'
```

> r: recurse  
> n: line number  
> w: match the whole word  
> alternative we can add -i for case-insensitive (makes search slower)  
> also if we only need the file name we can add -l  

## count word ocurences in file

```bash
grep -o -i mauris example.txt | wc -l
```

## count files by extension

```bash
ls -lR | grep --count \.go$
```

## print a range of lines from file

```bash
sed -n 'n0,n1p' '[path-to-file]'
```

## service log

since the beginning with paging

```bash
journalctl -u [service-name]
```

just current session (boot)

```bash
journalctl -u [service-name] -b
```

just current session (boot) end (no scrolling)
 
```bash
journalctl -u [service-name] -b -e
```

just current session (boot) show [x] lines before and after a specific [word]

```bash
journalctl -u [service-name] -b | grep -C [x] [word]
```

## grant SSH Access to User

On the server

```bash
adduser [username]
usermod -aG sudo [username]
```
Test it with

```bash
su - [username]
sudo ls -la /root
```

```bash
mkdir -p ~/.ssh
vim ~/.ssh/authorized_keys
[paste user public ssh key]
chmod -R go= ~/.ssh
```

If you are not logged as the user:

```bash
chown -R [username]:[username] ~/.ssh
```

You can login as the user with:

```bash
su - [username]
```

## fixed IP

```bash
sudo vim /etc/network/interfaces
```

```
auto lo
iface lo inet loopback

auto eno1
iface eno1 inet static
	address 192.168.100.100
	netmask 255.255.255.0
	network 192.168.100.0
	gateway 192.168.100.1
	dns-nameservers 192.168.100.1 8.8.8.8
```

```bash
ifdown eno1; ifup eno1
```

## change Hostname

```bash
sudo vim /etc/hostname
```

## ubuntu Open Port

```bash
sudo apt install ufw
sudo ufw enable
sudo ufw status verbose   
sudo ufw allow 80/tcp
```

## mount usb

```bash
sudo fdisk -l
sudo lsblk
sudo mount /dev/sdb1 /media/kingston
sudo umount /media/kingston
fuser -m /media/kingston
umount -l /media/kingston
```

## mount directory into another

Linux

```bash
sudo mount --bind ~/Desktop/upictures/pictures/ /home/mario/Projects/    upictures/src/UPictures.Web/wwwroot/pictures    
sudo umount /home/mario/Projects/upictures/src/UPictures.Web/wwwroot/    pictures
```

MacOS

```bash
sudo mount localhost:/Users/mario/Desktop/ufiles/output /Users/mario/    Projects/ufiles/src/UFiles.Web/wwwroot/pictures/   
sudo umount /Users/mario/Projects/ufiles/src/UFiles.Web/wwwroot/    pictures/
```

## remove image metadata

```bash
sudo apt install libimage-exiftool-perl
```

Usage example

```bash
exiftool -all= *.jpg
```

## desactivate touchpad

Run

```bash
xinput list
```

To see the list of available devices. In my case the entry was:

```
ImPS/2 Generic Wheel Mouse    id=13
```

Then run the following command to disable it:

```
xinput set-prop 13 "Device Enabled" 0
```

## network

nmap will scan the network for hosts and open ports

```bash
nmap 192.168.1.1-255
```

list network shares

```bash
smbclient -L=192.168.1.146
```

or

```bash
ls /var/lib/samba/usershares
```

mount network shares

```bash
sudo apt install cifs-utils
sudo mount //192.168.1.146/music /media/baal/music/ -o username=[user],password=[password]
```

unmount network share

```bash
sudo umount /media/baal/music
```

## ubuntu server configuration

Rename Server

```bash
hostnamectl set-hostname new-hostname
sudo vim /etc/hosts
```

Static IP

```bash
ip addr show
```

```bash
sudo lshw -class network
```

view the network adapter name, tipically eth0 

```bash
sudo vim /etc/network/interfaces
```

Add:

```
auto eth0
iface eth0 inet static
address 10.0.64.7
netmask 255.255.255.0
gateway 10.0.64.1
```

now you can do sudo ifup eth0 or sudo ifdown eth0

restart, surely there is a way to configure this without restarting

Static Hostnames

```bash
sudo vim /etc/hosts
```

Add a new line like: 10.0.64.7	duriel

## cut, get slices from a file

-d delimiter, -f fields
For example, using ',' as delimiter [-d,] and then selecting the third group [-f3]

```bash
cut -d, -f3 /etc/group 
```

Sort the output in ascending order (as a string -s, as a number -n)

```bash
cut -d, -f3 /etc/group | sort -s
```

Sort the output in descending order

```bash
cut -d, -f3 /etc/group | sort -rs
```

## get your ip address

```bash
ip a
```

## scripting

Write a .sh file, like myscript.sh. 

```bash
#!/bin/bash
chmod +x ./myscript.sh
```

Example of content

```bash
for filename in file1 file2 file3
do
  echo "Important stuff" >> $filename
done
```

## info about computer hardware

```bash
sudo lshw
```

## info about a host

```bash
# host [host]
host google.com
```

## skip first line, get only colum 12 and sum values

```bash
sed 1,1d 2019-01_MXN.csv | cut -d, -f12 | awk '{n += $1}; END{print n}'
```

## scan for ips in current network

```bash
nmap -sn 192.168.100.0/24
```

first you will have to confirm your ip range with another command like: `ip a`

## DVD to iso

```bash
sudo dd if= /dev/sr0 of= ~/Desktop/2016-03-31_eco-ciro.iso status=progress
```

```bash
sudo chown mario:mario ~/Desktop/2016-03-31_eco-ciro.iso
```

## bin to iso

```bash
bchunk backup7.bin backup7.cue backup_7.iso
```

## ssh tunnel

```bash
ssh -D 8080 [user]@[host-ip]
```

Firefox

```
Manual proxy configuration
SOCKS Host: 127.0.0.1
Port: 8080
SOCKS v5
```

## curl 

[https://catonmat.net/cookbooks/curl](https://catonmat.net/cookbooks/curl)

## safely Remove Flash Drive

```bash
fdisk -l
```

look for device path like for example `/dev/sdb1`

```bash
sudo eject /dev/sdb1
```

## disable sleep, do nothing on lid close

```bash
# sudo su
echo 'HandleLidSwitch=ignore' | tee --append /etc/systemd/logind.conf
echo 'HandleLidSwitchDocked=ignore' | tee --append /etc/systemd/logind.conf
service systemd-logind restart
```

## battery level

```bash
upower -i /org/freedesktop/UPower/devices/battery_BAT0
```

basic details

```bash
upower -i /org/freedesktop/UPower/devices/battery_BAT0 | grep -E "state|to\ full|percentage"
```

## sShare Folder

```bash
net usershare add videos /home/mario/Videos/movies "" everyone:R guest_ok=y
```

## tar exclude folders

```bash
tar -czf deploy.tar.gz * --exclude="c:\\\windows\\\temp"
```

## kernel info

```bash
uname -a for all information regarding the kernel version,
```

```bash
uname -r for the exact kernel version
```

```bash
lsb_release -afor all information related to the Ubuntu version,
```

```bash
lsb_release -r for the exact version
```

```bash
sudo fdisk -l for partition information with all details.
```

## network

nmap will scan the network for hosts and open ports

```bash
nmap 192.168.1.1-255
```

list network shares

```bash
smbclient -L=192.168.1.146
```
mount network shares

```bash
sudo mount //192.168.1.146/music /media/baal/music/ -o username=[user],password=[password]
```

unmount network share

```bash
sudo umount /media/baal/music
```

## long running script trough ssh

```bash
nohup ./apis-startup.sh > apis.out 2> apis.err < /dev/null &
```

## open images

```bash
sudo apt install feh
feh -F ./Pictures/563139.jpg
```

## compare with diff

[https://www.howtogeek.com/410532/how-to-compare-two-text-files-in-the-linux-terminal/](https://www.howtogeek.com/410532/how-to-compare-two-text-files-in-the-linux-terminal/)

## weather

```bash
curl 'https://wttr.in/:help'
```

[https://github.com/chubin/wttr.in](https://github.com/chubin/wttr.in)

## curl ip publica

```bash
curl ifconfig.me
```

## ps con parent processes

```bash
ps fax
```

## jq to csv

```bash
jq -r '(.[0] | keys_unsorted) as $keys | $keys, map([.[ $keys[] ]])[] | @csv'
```

## search

## search specific file content recursively

```bash
grep -rn 'pattern' .
```

```bash
grep -iRl "your-text-to-find" ./
```

copy&paste:

```bash
grep -iRl "" .
```

Another option

```bash
grep -rnw '/path/to/somewhere/' -e 'pattern'
```

copy & paste:

```bash
grep -rnw . -e ''
```

> r: recurse
> n: line number
> w: match the whole word
> alternative we can add -i for case-insensitive (makes search slower)
> also if we only need the file name we can add -l

## print specific range of lines from a file

```bash
sed -n '60,65p' './kb/terminal.md'
```

copy&paste:

```bash
sed -n ',p' ''
```

## clear cached memory

```bash
sudo -i
sync; echo 1 > /proc/sys/vm/drop_caches
```

> clear cached memory: The page cache contains any memory mappings to blocks on disk. That could be buffered I/O, memory mapped files, paged areas of executables

## see logged users

```bash
w
```

```bash
pkill -9 -t pts/[x]
```

```bash
pkill -f ms-dotnettools.csdevkit
```

## man pages tldr

[https://tldr.sh/](https://tldr.sh/)

[https://furbo.org/2014/09/03/the-terminal/](https://furbo.org/2014/09/03/the-terminal/)

## luks encryption

```bash
lsblk
# assume luks encrypted volume is sda
```

mount

```bash
sudo cryptsetup open /dev/sda homunculus --type luks
sudo mount /dev/mapper/homunculus /media/homunculus/
```

umount 

```bash
sudo umount /media/homunculus
sudo cryptsetup close homunculus
```

## lm sensors

cpu temperature and different metrics

```bash
sudo apt install lm-sensors
sensors
```

or    

```bash
watch sensors
```

intel stick

```
Every 2.0s: sensors                                     oculus: Wed Apr  3 15:48:54 2024
soc_dts1-virtual-0
Adapter: Virtual device
temp1:        +54.0 C
coretemp-isa-0000
Adapter: ISA adapter
Core 0:       +53.0 C  (high = +90.0 C, crit = +90.0 C)
Core 1:       +53.0 C  (high = +90.0 C, crit = +90.0 C)
Core 2:       +53.0 C  (high = +90.0 C, crit = +90.0 C)
Core 3:       +53.0 C  (high = +90.0 C, crit = +90.0 C)
axp288_fuel_gauge-isa-0000
Adapter: ISA adapter
in0:           4.49 V
curr1:         0.00 A
```

dell xps

```
dell_smm-isa-0000
Adapter: ISA adapter
Processor Fan: 4720 RPM  (min =    0 RPM, max = 7000 RPM)
Video Fan:     4413 RPM  (min =    0 RPM, max = 7000 RPM)
CPU:            +51.0°C  
SODIMM:         +42.0°C  
Other:          +44.0°C  
Ambient:        +33.0°C  
Other:          +41.0°C  
Ambient:        +44.0°C  
Ambient:        +41.0°C  

ucsi_source_psy_USBC000:003-isa-0000
Adapter: ISA adapter
in0:           5.00 V  (min =  +5.00 V, max =  +5.00 V)
curr1:         0.00 A  (max =  +1.50 A)

ucsi_source_psy_USBC000:001-isa-0000
Adapter: ISA adapter
in0:           5.00 V  (min =  +5.00 V, max =  +5.00 V)
curr1:         0.00 A  (max =  +0.00 A)

BAT0-acpi-0
Adapter: ACPI interface
in0:           6.99 V  
curr1:         2.08 A  

ucsi_source_psy_USBC000:002-isa-0000
Adapter: ISA adapter
in0:           5.00 V  (min =  +5.00 V, max =  +5.00 V)
curr1:         0.00 A  (max =  +0.00 A)

ath10k_hwmon-pci-0200
Adapter: PCI adapter
temp1:        +41.0°C  

pch_skylake-virtual-0
Adapter: Virtual device
temp1:        +46.0°C  
nvme-pci-6e00
Adapter: PCI adapter
Composite:    +43.9°C  (low  = -20.1°C, high = +77.8°C)
                       (crit = +81.8°C)
Sensor 1:     +43.9°C  (low  = -273.1°C, high = +65261.8°C)
acpitz-acpi-0
Adapter: ACPI interface
temp1:        +25.0°C  (crit = +107.0°C)
```

##  flash drive

```bash
lsblk
# identify flash drive, for example sdb
```

format

```bash
sudo mkfs.ext4 /dev/sdb
```

merge partitions

```bash
sudo fdisk /dev/sdb
# list curent partitions
p
# delete partition type 'd'
# create partition type 'n'
# write changes
w
```

show file system details

```bash
lsblk -f
```

## wifi

```bash
// restart wifi
sudo service NetworkManager restart
```

## show battery percentage

```bash
upower -i /org/freedesktop/UPower/devices/battery_BAT0 | grep -E "state|to\ full|percentage"
```

## chmod

owner/group/world

```bash
chmod 777 // rwx for everyone
chmod 755 // rw for owner, rx for group world
```

## ps 

```bash
ps 
ps aux
```

```bash
kill [pid]
killall [name] 
```

## access devices with a vpn

[https://tailscale.com/](https://tailscale.com/)

[https://blog.6nok.org/tailscale-is-pretty-useful/](https://blog.6nok.org/tailscale-is-pretty-useful/)

## calendar txt

[https://terokarvinen.com/2021/calendar-txt/](https://terokarvinen.com/2021/calendar-txt/)

## mp4 files, create thumbnails

```bash
#!/bin/bash

# Directory to process (default: current directory)
DIR="${1:-.}"

# Loop through all .mp4 files in the directory
for video in "$DIR"/*.mp4; do
  # Skip if no .mp4 files are found
  [ -e "$video" ] || continue

  # Get the base filename without extension
  base="${video%.*}"
  # Set output thumbnail filename
  thumbnail="${base}.png"

  # Create thumbnail at 10 seconds (change -ss value if needed)
  ffmpeg -y -ss 00:00:10 -i "$video" -vframes 1 "$thumbnail"
  echo "created thumbnail: $thumbnail"
done
```

## todo project

```bash
x 0 YYYY-MM-DD YYYY-MM-DD "description +project @context" due:YYYY-MM-DD-hh-mm rem:3 
```

## disable gnome remote desktop

```bash
systemctl --user disable --now gnome-remote-desktop.service
```

## check if a port is open

```bash
ss -tuln | grep :3306

# kill it
lsof -ti :5000 | xargs kill 

# if you don't want to wait
lsof -ti :5000 | xargs kill -9
```