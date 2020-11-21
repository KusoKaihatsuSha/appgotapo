# appgotapo
CLI app for manage TP Link Tapo C200

usage: appgotapo -host=IP -u=user -p=pass -do="next|sleep:3s"

usage: appgotapo.exe -host=IP -u=user -p=pass -do="next|sleep:3s"

Info:
  -do string
  
        Example:
        
        1) led_on|sleep:3s|led_off|night_on|sleep:3s|night_auto|alarm_on|sleep:30s|alarm_off|sleep:6s|next|sleep:6s|next|sleep:6s|next|off|next|sleep:10s|on
        
        2) alarm_off|reboot
        
        3) off|next|sleep:10s|on
        
        4) text:isTestedText|refresh_off|next:5s
        
        5) sens:3|alarm_on
        
        (default "refresh_on|next")
         
  -host string
  
        write cam ip
        
        (default "192.168.1.111")
         
  -p string
  
        write password
        
        (default "campass")
         
  -u string
  
        write username
        
        (default "camname")
