# TNTNodeChecker
A small go program to check the health of TNT nodes

## Installation Instructions (Zabbix Agent)

Must have zabbix-agent installed on the node server. Connect to this server and go to the $HOME folder. The following steps must be done for each node server seperately.

Clone the github repository
```
git clone https://github.com/SomniaStellarum/TNTNodeChecker.git $HOME/go/src/github.com/SomniaStellarum/TNTNodeChecker
```
Edit the `/etc/zabbix/zabbix_agentd.conf` file. Recommend
```
sudo vim /etc/zabbix/zabbix_agentd.conf
```
Add the zabbix server to the config file. This is the IP to allow incoming requests from. Good practice to only set `Server` and `HostName` once in configuration file.
```
Server=127.0.0.1,UserSpecifiedIP
```
Set `AllowRoot=1`. Will likely have to uncomment from default.

Add `UserParameter`. This is the script that will be called to check on the health of the TNT Node.
```
UserParameter=healthcheck,HOME/go/src/github.com/SomniaStellarum/TNTNodeChecker/nodeCheck.sh HOME TNT_Address
```
Note: `HOME` and `TNT_Address` should be replaced here with the actual values here. `HOME` is the path the the Home folder and `TNT_Address` is the ethereum address used for the node (holding the TNT balance).

## Zabbix Server

Next we need to configure the zabbix server so that we can collect the data. The easiest way to do this is through a template. Create a template and add an initial item (called `_hc`). This is the main item that will collect the data, so set the collection period (eg 1m). This should be a type `Zabbix Agent`. The key should be `healthcheck`, which will collect the json data from the node checker to be parsed by dependent variables. Recommended history storage is `0` or `3600` (for debugging).

Next, add each of the following dependent variables. There should be a preprocessing step for each (type `JSONPath`)

`BalanceTNT`, key: `healthcheck.tnt`, type: `Text`, PreProc Param: `$.tnt`

`ConsecutiveFail`, key: `healthcheck.cfail`, type: `Numeric (unsigned)`, PreProc Param: `$.cfail`

`ConsecutivePass`, key: `healthcheck.cpass`, type: `Numeric (unsigned)`, PreProc Param: `$.cpass`

`Fail`, key: `healthcheck.fail`, type: `Numeric (unsigned)`, PreProc Param: `$.fail`

`FailPercent`, key: `healthcheck.failper`, type: `Numeric (float)`, PreProc Param: `$.failper`

`Pass`, key: `healthcheck.pass`, type: `Numeric (unsigned)`, PreProc Param: `$.pass`

Finally, add the hosts (ensure the correct IP) and apply the template.
