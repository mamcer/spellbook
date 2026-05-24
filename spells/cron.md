# cron

[https://crontab.cronhub.io](https://crontab.cronhub.io)

cron: daemon  
crontab: tool to create cron tasks  

Ensure time is correctly configured

```bash
timedatectl
```

All the cron tasks are in a user file, you can setup the file using crontab

```bash
crontab [file]
```

## crontab file format

```
* * * * * [user] [command]

m h dom mon dow [user] [command]
```

> m: minute [0-59]
> h: hour [0-23]  
> dom: day of month [1-31]  
> mon: month [1-12]  
> dow: day of week [0-6] (sunday is 0) (you can use text like FRI for friday (5))

## reserved words

```
@reboot: every restart
@yearly: onc a year: 0 0 1 1 *
@monthly: one time the first day of the month 0 0 1 * *
@weekly: every week the first minute of the first hour 0 0 * * 0
@daily: daily at 12am  0 0 * * *
@midnight: idem as daily
@hourly: every hour during the first minute: 0 * * * *
```