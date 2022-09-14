## Zabbix URLS for the map

I wrote zmap to open from zabbix map (rdp, winbox, putty, vnc) 

### Instruction:
- Copy zmap to C:\
- Run C:\zmap\zmap.reg
- Add to Zabbix Valid URI schemes: Administration -> General -> Other -> Valid URI schemes
![zabbix_uri.png](https://itfarsh.com/media/posts/images/2022/09/14/8ea31df0fa-zabbix_uri.png) 
- Add zabbix -> map -> host -> URLS
![zabbix_map_add.png](https://itfarsh.com/media/posts/images/2022/09/14/97b47509e6-zabbix_map_add.png) 

### Type URLS:
- zmap://winbox@{HOST.CONN}
- zmap://rdp@{HOST.CONN}
- zmap://ssh@{HOST.CONN}
- zmap://vnc@{HOST.CONN}
