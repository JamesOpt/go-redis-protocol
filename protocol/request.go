package protocol

import (
	"strconv"
	"strings"
)

func GetRequest(args []string) []byte {
	// start *+len(args)
	req := []string{
		"*" + strconv.Itoa(len(args)),
	}

	// second $ len(arg)
	// third arg
	for _, arg := range args {
		req = append(req, "$" + strconv.Itoa(len(arg)), arg)
	}

	// last join byte with CR LF
	str := strings.Join(req, "\r\n")
	return []byte(str + "\r\n")
}