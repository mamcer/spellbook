# mail

## install 

```bash
sudo apt install mutt -y
```

## configure

[https://github.com/ork/mutt-office365](https://github.com/ork/mutt-office365)

user file

```bash
# User config
set my_realname="Mario Moreno"
set my_username="<user-name>"
set my_domain="<email-domain>"
set my_password="<app_password>"
set my_lang="en_US"
```

## send emails

[https://www.makeuseof.com/install-configure-mutt-with-gmail-on-linux](https://www.makeuseof.com/install-configure-mutt-with-gmail-on-linux/#:~:text=To%20send%20new%20emails%20using,y%20to%20send%20the%20email.)

```bash
echo "Body Message" | mutt -s "Testing Email from mutt" <email-address>
```

```bash
cat body.txt | mutt -s "Another test email" <email-address> -a <file-path>
```