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

	BlogNotExist         uint32 = 40000
	BlogUserInfoNotExist uint32 = 40001
	PageViewNotExist     uint32 = 40002

	PostNotExist          uint32 = 50000
	ImageNotExist         uint32 = 50001
	AuthorNotExist        uint32 = 50002
	LabelNotExist         uint32 = 50003
	LocationNotExist      uint32 = 50004
	PostUserInfosNotExist uint32 = 50005

	PageNotExist uint32 = 60000

	CommentNotExist uint32 = 70000
)

var msg = map[uint32]string{
	Database: "database error",
	Service:  "service error",

	UserNotExist:   "user not exist",
	LocaleNotExist: "locale not exist",

	BlogNotExist:         "blog not exist",
	BlogUserInfoNotExist: "blog user info not exist",
	PageViewNotExist:     "page view not exist",

	PostNotExist:          "post not exist",
	ImageNotExist:         "image not exist",
	AuthorNotExist:        "author not exist",
	LabelNotExist:         "label not exist",
	LocationNotExist:      "location not exist",
	PostUserInfosNotExist: "post user info not exist",

	PageNotExist: "page not exist",

	CommentNotExist: "comment not exist",
}

func Wrap(e uint32) error {
	return status.Error(codes.Code(e), msg[e])
}

func Msg(e uint32) string {
	return msg[e]
}
