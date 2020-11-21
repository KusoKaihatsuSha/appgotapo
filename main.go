// main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KusoKaihatsuSha/gotapo"
)

var p = fmt.Println

func main() {
	var host, user, password, commands string

	flag.StringVar(&host, "host", "192.168.1.111", "write cam ip\n")
	flag.StringVar(&user, "u", "camname", "write username\n")
	flag.StringVar(&password, "p", "campass", "write password\n")
	flag.StringVar(&commands, "do", "refresh_on|next", "Example: \n1) led_on|sleep:3s|led_off|night_on|sleep:3s|night_auto|alarm_on|sleep:30s|alarm_off|sleep:6s|next|sleep:6s|next|sleep:6s|next|off|next|sleep:10s|on\n2) alarm_off|reboot\n3) off|next|sleep:10s|on\n4) text:isTestedText|refresh_off|next:5s\n5) sens:3|alarm_on\n")
	//flag.StringVar(&commands, "do", "next", "write password")
	flag.Parse()

	if flag.NFlag() < 3 {
		fmt.Printf("usage: appgotapo -host=IP -u=user -p=pass -do=\"next|sleep:3s\"\nusage: appgotapo.exe -host=IP -u=user -p=pass -do=\"next|sleep:3s\"\n Info:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	conn := gotapo.Connect(host, user, password)
	commandsList := strings.Split(commands, "|")
	for _, v := range commandsList {
		switch {
		case v == "":
		case v == "test":
			conn.MoveTest()
		case v == "reboot":
			conn.Reboot()
		case v == "next":
			conn.NextPreset()
		case v == "led_on":
			conn.On(conn.Elements.Indicator)
		case v == "led_off":
			conn.Off(conn.Elements.Indicator)
		case v == "led_on":
			conn.On(conn.Elements.Indicator)
		case v == "led_off":
			conn.Off(conn.Elements.Indicator)
		case v == "night_auto":
			conn.On(conn.Elements.NightModeAuto)
		case v == "night_on":
			conn.On(conn.Elements.NightMode)
		case v == "night_off":
			conn.Off(conn.Elements.NightMode)
		case v == "alarm_on":
			conn.On(conn.Elements.AlarmMode)
			conn.On(conn.Settings.DetectEnableSound)
			conn.On(conn.Elements.DetectMode)
		case v == "alarm_off":
			conn.Off(conn.Elements.AlarmMode)
			conn.Off(conn.Settings.DetectEnableSound)
			conn.Off(conn.Elements.DetectMode)
		case v == "alarm2_on":
			conn.On(conn.Elements.AlarmMode)
			conn.On(conn.Settings.DetectEnableFlash)
			conn.On(conn.Settings.DetectEnableSound)
			conn.On(conn.Elements.DetectMode)
		case v == "alarm2_off":
			conn.Off(conn.Elements.AlarmMode)
			conn.Off(conn.Settings.DetectEnableFlash)
			conn.Off(conn.Settings.DetectEnableSound)
			conn.Off(conn.Elements.DetectMode)
		case v == "detect_on":
			conn.On(conn.Elements.DetectMode)
		case v == "detect_off":
			conn.Off(conn.Elements.DetectMode)
		case v == "off":
			conn.On(conn.Elements.PrivacyMode)
		case v == "on":
			conn.Off(conn.Elements.PrivacyMode)
		case v == "led_on":
			conn.On(conn.Elements.Indicator)
		case v == "led_off":
			conn.Off(conn.Elements.Indicator)
		case v == "bigbro_on":
			conn.On(conn.Elements.AutotrackingMode)
		case v == "bigbro_off":
			conn.Off(conn.Elements.AutotrackingMode)
		case v == "flash_on":
			conn.On(conn.Settings.DetectEnableFlash)
		case v == "flash_off":
			conn.Off(conn.Settings.DetectEnableFlash)
		case v == "sound_on":
			conn.On(conn.Settings.DetectEnableSound)
		case v == "sound_off":
			conn.Off(conn.Settings.DetectEnableSound)
		case v == "text_on":
			conn.On(conn.Settings.VisibleOsdText)
		case v == "text_off":
			conn.Off(conn.Settings.VisibleOsdText)
		case v == "time_on":
			conn.On(conn.Settings.VisibleOsdTime)
		case v == "time_off":
			conn.Off(conn.Settings.VisibleOsdTime)
		case v[0:4] == "sens":
			sens, _ := strconv.Atoi(v[5:6])
			conn.Settings.DetectSensitivity = sens
		case v[0:4] == "text":
			conn.Settings.OsdText = v[5:len(v)]
			conn.On(conn.Settings.VisibleOsdText)
		case v == "refresh_on":
			conn.On(conn.Settings.PresetChangeOsd)
		case v == "refresh_off":
			conn.Off(conn.Settings.PresetChangeOsd)
		case v[0:5] == "sleep":
			dur_, _ := time.ParseDuration(v[6:len(v)])
			time.Sleep(dur_)
		default:
			p("do nothing")
		}
	}
}
