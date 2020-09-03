# SubSaaS
A tool for enumerating if an organization uses certain SaaS products via information from subdomains.

SubSaaS uses a number of techniques for determing if an organization uses a certain SaaS products. Some products, such as Splunk, simply do not resolve. Others, like Adobe Creative Cloud, reveal information on if an organization is federated via the password login field.

If you have an organization you'd like added and have an idea for determining if it's valid, add an issue.

I won't add support for wordlist generation or advanced false positive detection. If you notice numerous false positives, feel free to make an issue and I'll see about adding more verbose output (like I did with email addresses and Slack).

## Installation

```
git clone git@github.com:daniel-infosec/subsaas.git
cd subsaas/
go build
```

## Execution

You can use the provided wordlist as an example or run with just one organization

```
./subsaas -org praetorian
Splunk
[]
Slack
[{praetorian @praetorian.com}]
Zoom
[praetorian]
Atlassian
[praetorian]
Okta
[praetorian]
Box
[praetorian]
Adobe Creative Cloud
[]
```

```
./subsaas -orglist complist.txt 
Splunk
[]
Slack
[{praetorian @praetorian.com} {microsoft @microsoft.com} {amazon @amazon.com} {riot @coderoad.com} {apple nil} {ge @ge.com, @bitstew.com, @nurego.com, or @mediamonks.com} {boeing nil} {walmart nil} {mckesson nil} {ford @razorfish.com} {costco @costco.com} {intel @intel.com} {fedex @drivingpurple.com} {aetna nil}]
Zoom
[praetorian apple walmart intel aetna]
Atlassian
[praetorian microsoft amazon riot apple walmart ford costco]
Okta
[praetorian microsoft amazon apple ge boeing walmart mckesson ford costco intel fedex aetna]
Box
[praetorian microsoft amazon riot apple ge boeing walmart mckesson intel fedex aetna]
Adobe Creative Cloud
[amazon apple ge boeing walmart ford intel fedex]
```

## Current Supported SaaS Checks

* Splunkcloud
* Slack (low fidelity)
* Zoom
* Atlassian
* Okta
* Box
* Adobe Creative Cloud
* Servicenow
* Snowflake Computing
* Workday
* Pagerduty

## Docker

Alternatively, if you have Docker installed and don't want to bother installing Go or getting the binary allowlisted, you only have to do the following. Note that if you make any changes or update, you'll have to rebuild the docker file manually.

### Windows

`.\subsaas.ps1 -org praetorian`

### Linux

`./subsaas.sh -org praetorian`

# golang

This is my first attempt at pivoting from Python to Go. If you notice any bad habits or bad code, please feel free to call me out (gently if you can).
