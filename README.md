[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![godoc](https://godoc.org/github.com/KusoKaihatsuSha/appgotapo?status.svg)](https://godoc.org/github.com/KusoKaihatsuSha/appgotapo) [![Go Report Card](https://goreportcard.com/badge/github.com/KusoKaihatsuSha/appgotapo)](https://goreportcard.com/report/github.com/KusoKaihatsuSha/appgotapo)

# App for remoting TAPO Cameras (C200, C210, C220, C310 and other)

> Notes: CLI (NON-GUI) unofficial app for remoting model cameras TP Link Tapo C200/C310

> !!! Caution with update your cameras to new drivers. May be problems with using unofficial apps.

> !!! Don't forget to update the `gotapo` package.

### Accessible flags

`-host string`  - Camera IP or IPs ("192.168.1.111,192.168.1.222")

`-p string`     - Camera password

`-u string`      - Camera login (Set in official app*)

**!!! in various version drivers user login may be equal user login on official app or it may be 'admin'**

`-do stringActions` - You need write string of action list(delimiter `|`) after flag `-do`

### Commands
Short description of functionality:

`test` - test moving by saving position (C200)

`reboot` - reboot device

`next` - move to next saved position

`led_on` - led turn on

`led_off` - led turn off

`night_auto` - use auto mode

`night_on` - use night mode

`night_off` -not use night mode

`alarm_on` - alarm preset 1 on

`alarm_off` - alarm preset 1 off

`alarm2_on` - alarm preset 2 on

`alarm2_off` - alarm preset 2 off

`detect_on` - use visual and sound alarm by setting preset

`detect_off` - off visual and sound alarm

`off` - private mode off

`on` - private mode on

`bigbro_on` - use moving tower on moving objects

`bigbro_off` - off moving eye

`flash_on` - visual alarm on

`flash_off` - visual alarm off

`sound_on` - sound alarm on

`sound_off` - sound alarm off

`text_on` - text on screen mode on

`text_off` - text on screen mode off

`time_on` - time on screen mode on

`time_off` - time on screen mode off

`flip_on` - flip picture 180 degree

`flip_off` - flip normal

`sens:3` -  level sensitivity of finding objects

`text:isTestedText` - set osd text

`refresh_on` - refresh osd

`refresh_off` - not refresh osd

`move:100_10` - move tower on 100 degrees on X and 10 degrees on Y

`sleep:3s` - delay between commands


### Examples

```sh
$ appgotapo -host=IP -u=user -p=pass -do="next|sleep:3s"
$ appgotapo.exe -host=IP -u=user -p=pass -do="next|sleep:3s"
$ appgotapo.exe -host=IP -u=user -p=pass -do="led_on|sleep:3s|led_off|night_on|sleep:3s|night_auto|alarm_on|sleep:30s|alarm_off|sleep:6s|next|sleep:6s|next|sleep:6s|next|off|next|sleep:10s|on"
$ appgotapo.exe -host=IP -u=user -p=pass -do="alarm_off|reboot"
$ appgotapo.exe -host=IP -u=user -p=pass -do="off|next|sleep:10s|on"
$ appgotapo.exe -host=IP -u=user -p=pass -do="text:isTestedText|refresh_off|next"
$ appgotapo.exe -host=IP -u=user -p=pass -do="sens:3|alarm_on"
$ appgotapo.exe -host=IP -u=user -p=pass -do="refresh_on|next"
```

