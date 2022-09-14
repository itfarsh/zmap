## Zabbix URLS for the map

I wrote zmap to open from zabbix map (rdp, winbox, putty, vnc) 

### Instruction:
- Copy zmap to C:\
- Run C:\zmap\zmap.reg
- Add to Zabbix Valid URI schemes: Administration -> General -> Other -> Valid URI schemes
- Add zabbix -> map -> host -> URLS

### Type URLS:
- zmap://winbox@{HOST.CONN}
- zmap://rdp@{HOST.CONN}
- zmap://ssh@{HOST.CONN}
- zmap://vnc@{HOST.CONN}
