// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"github.com/jackc/pgtype"
)

type Bar struct {
	Addr         pgtype.Macaddr
	NullableAddr pgtype.Macaddr
}

type Foo struct {
	PresentIp    pgtype.Inet
	NullableIp   pgtype.Inet
	PresentCidr  pgtype.CIDR
	NullableCidr pgtype.CIDR
}
