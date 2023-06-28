package rdbstmt

import "strings"

var replacers = map[string]Replacer{
	// Clusters

	// Connection

	// Geo
	"geoadd":  trimAt6,
	"geohash": trimAt6,
	"geopos":  trimAt6,

	// Hashes
	"hmset":  replaceHMSet,
	"hset":   replace3,
	"hsetnx": replace3,

	// HyperLogLog
	"pfadd": trimAt3,

	// Keys
	"restore": replace3,

	// Lists
	"linsert": replaceFrom3,
	"lpush":   replaceFrom2,
	"lpushx":  replace2,
	"lrem":    replace3,
	"lset":    replace3,
	"rpushx":  replace2,

	// Pub/Sub
	"publish": replace2,

	// Scripting
	"eval":    trimAt4,
	"evalsha": trimAt4,

	// Server
	"auth": replaceFrom1,

	// Sets
	"sadd":      replaceFrom2,
	"sismember": replace2,
	"smove":     replace3,
	"srem":      replaceFrom2,

	// Sorted Sets
	"zadd":      replaceZAdd,
	"zincrby":   replace3,
	"zrank":     replace2,
	"zrem":      replaceFrom2,
	"zreverank": replace2,
	"zscore":    replace2,

	// Streams
	"xack":       trimAt3,
	"xadd":       trimAt2,
	"xclaim":     trimAt5,
	"xdel":       trimAt2,
	"xread":      trimStreams,
	"xreadgroup": trimStreams,

	// strings
	"append":   replace2,
	"mset":     replaceMSet,
	"msetnx":   replaceMSet,
	"psetex":   replace2,
	"set":      replace2,
	"setex":    replace3,
	"setnx":    replace2,
	"setrange": replace3,

	// Transactions
}

func getReplacer(name string) Replacer {
	name = strings.ToLower(name)
	if r, ok := replacers[name]; ok {
		return r
	}
	return replaceNoop
}

var replaceChars = "?"

func SetReplaceChars(str string) {
	replaceChars = str
}

func GetReplaceChars(str string) string {
	return replaceChars
}

func replaceNoop(*BaseStatement) {
}

var (
	replace2 = replaceAt(2, 2)
	replace3 = replaceAt(3, 3)

	replaceFrom1 = replaceAt(1, -1)
	replaceFrom2 = replaceAt(2, -1)
	replaceFrom3 = replaceAt(3, -1)
)

func replaceAt(start, end int) Replacer {
	return func(b *BaseStatement) {
		if len(b.args) < start {
			return
		}
		if end == -1 || end > len(b.args)-1 {
			end = len(b.args) - 1
		}

		for i := start; i <= end; i++ {
			b.args[i] = replaceChars
		}
	}
}

var (
	appendString = "..."

	trimAt2 = trimAt(2, appendString)
	trimAt3 = trimAt(3, appendString)
	trimAt4 = trimAt(4, appendString)
	trimAt5 = trimAt(5, appendString)
	trimAt6 = trimAt(6, appendString)
)

func trimAt(n int, append string) Replacer {
	return func(b *BaseStatement) {
		if len(b.args) <= n {
			return
		}
		m := n
		if append != "" {
			m++
		}
		args := make([]interface{}, m)
		copy(args, b.args[0:n])
		if m > n {
			args[m-1] = append
		}
		b.args = args
	}
}

func replaceMSet(b *BaseStatement) {
	for i := 2; i < len(b.args); i++ {
		if !isOdd(i) {
			b.args[i] = replaceChars
		}
	}
}

func replaceHMSet(b *BaseStatement) {
	for i := 3; i < len(b.args); i++ {
		if isOdd(i) {
			b.args[i] = replaceChars
		}
	}
}

func replaceZAdd(b *BaseStatement) {
	for i := 3; i < len(b.args); i++ {
		if isOdd(i) {
			b.args[i] = replaceChars
		}
	}
}

func trimStreams(b *BaseStatement) {
	const cmp = "streams"
	index := stringIndex(b.args, cmp)
	f := trimAt6
	if index != -1 {
		f = trimAt(index+1, appendString)
	}
	f(b)
}

func stringCompare(v interface{}, str string) bool {
	if s, ok := v.(string); ok {
		if strings.ToLower(s) == str {
			return true
		}
	}
	return false
}

func stringIndex(args []interface{}, str string) int {
	for i := 0; i < len(args); i++ {
		if stringCompare(args[i], str) {
			return i
		}
	}
	return -1
}

func isOdd(i int) bool {
	return i%2 == 1
}
