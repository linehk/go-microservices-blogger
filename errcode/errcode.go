package errcode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	Database uint32 = 20000
	Service  uint32 = 20001

	UserNotExist   uint32 = 30000
	LocaleNotExist uint32 = 30001

	BlogNotExist uint32 = 40000
)

var msg = map[uint32]string{
	Database: "database error",
	Service:  "service error",

	UserNotExist:   "user not exist",
	LocaleNotExist: "locale not exist",

	BlogNotExist: "blog not exist",
}

func Wrap(e uint32) error {
	return status.Error(codes.Code(e), msg[e])
}

func Msg(e uint32) string {
	return msg[e]
}
