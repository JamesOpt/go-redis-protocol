package protocol

import (
	//"fmt"
	"strings"
)

const(
	STATUS_REPLY = '+'
	ERROR_REPLY = '-'
	INTEGER_REPLY = ':'
	BULK_REPLY = '$'
	MULTI_BULK_REPLY = '*'

	OK = "OK"
)

func GetReply(reply []byte) (interface{}, error) {
	switch reply[0] {
	case STATUS_REPLY:
		return doStatusReplay(reply)
	case ERROR_REPLY:
		return string(reply), nil
	case INTEGER_REPLY:
		return doIntegerReply(reply)
	case BULK_REPLY:
		return doBulkRely(reply)
	case MULTI_BULK_REPLY:
		return doMultiBulkReply(reply)
	}

	return "",nil
}

func doStatusReplay(reply []byte) (string, error) {
	if len(reply) == 5 && 'O' == reply[1] && 'K' == reply[2] {
		return OK, nil
	}

	return string(reply), nil
}

func doIntegerReply(reply []byte) (string, error) {
	result := strings.Split(string(reply[1:]), "\r")[0]

	return result, nil
}

/**
$6\r\nfoobar\r\n"
 */
func doBulkRely(reply []byte) (string, error) {
	result := strings.Split(string(reply[1:]), "\r\n")[0]

	return result, nil
}

/**
客户端： LRANGE mylist 0 3
服务器： *4
服务器： $3
服务器： foo
服务器： $3
服务器： bar
服务器： $5
服务器： Hello
服务器： $5
服务器： World
 */
func doMultiBulkReply(reply []byte) (interface{}, error) {
	result := strings.Split(string(reply[1:]), "\r\n")

	len := len(result)

	r := []string{}
	for i := 1; i < len - 1; i++ {
		if i%2 == 0 {
			r = append(r, result[i])
		}
	}

	return strings.Join(r, "\r\n"), nil
}