
## SELinux Module for Forwarding SELinux Audit Logs

This is based entirely on the work posted here:
http://www.reddit.com/r/sysadmin/comments/2fn013/forwarding_audit_logs_via_rsyslog_with_selinux/

To install 

1. Install audit.log into /etc/rsyslog.d/
2. Run make rsyslog.pp
3. Run semodule -i rsyslog.pp

