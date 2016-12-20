# igmpqd

IGMPQD aka. IGMP Query Daemon is a lightweight utility to periodically generate an IGMPv2 Query message. The message can be formed to contain whatever payload is necessary, within the RFC specifications (https://tools.ietf.org/html/rfc2236#section-2)

If this is useful, great. If you find need to extend or bugfix, even better, please fork and submit merge requests as necessary.

## Usage

Download an appropriate binary for your platform, or build you own.

### Print Version

`./igmpqd version` will produce something along the lines of:

```
$ ./igmpqd version
Version:        0.0.1 (Commit: 4758b59b1966700fbc1665bcd0bed5a35bd9c4a3)
Built:          2016-04-07 16:17:09 +0100 BST
Fingerprint:    gc/darwin/amd64/go1.6
```

### Run the Daemon

Contrary to many other services, we don't fork/background the process. It's anticipated that you would run igmpqd with systemd, and therefore have a unit control the process execution apprpropriately.

To run the daemon, you use the "run" command (`./igmpqd run`). `run` takes several flags, as described below:

```
Usage:
  igmpqd run [flags]

Flags:
      --debug                 Enable debug messages to stderr.
  -d, --dstAddress string     Specified IP address to send the IGMP Query to.
                              (default "224.0.0.1")
  -g, --grpAddress string     Specified IP address to use as the Group Address.
                              Used to query for specific group members.
                              (default "0.0.0.0")
  -I, --interface string      Specified interface to send the IGMP Query.
                              (no default)
  -i, --interval int          The time in seconds to delay between sending IGMP
                              Query messages. (default 30)
  -t, --ttl int               The TTL of the IGMP Query.
                              (default 1)
  -m, --maxResponseTime int   Specifies the maximum allowed time before sending
                              a responding report in units of 1/10 second.
                              (default 100)
  ```
