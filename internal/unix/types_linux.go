// +build linux

package unix

import (
	linux "golang.org/x/sys/unix"
)

const (
	AF_INET                   = linux.AF_INET
	AF_INET6                  = linux.AF_INET6
	NFNETLINK_V0              = linux.NFNETLINK_V0
	NFNL_SUBSYS_CTNETLINK     = linux.NFNL_SUBSYS_CTNETLINK
	NFNL_SUBSYS_CTNETLINK_EXP = linux.NFNL_SUBSYS_CTNETLINK_EXP
	NETLINK_NETFILTER         = linux.NETLINK_NETFILTER
)
