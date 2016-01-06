package main

import (
	"os"
	"flag"
	"fmt"
	"crypto/md5"
	"encoding/binary"
)

var verbose bool

// Hash msgid and get offset with 4 octet value
// Example offset=0, MD5(msgid)=f1980b66e0a43405f3199d92 [ 774695cf(16) ] = 2001114575(32)
func hash(msgid string, offset int) uint32 {
	var digest []byte
	{
		h := md5.New()
		h.Write([]byte(msgid))
		digest = h.Sum(nil)
		// 16bytes
		if verbose {
			fmt.Printf("MD5(%s)=%x\n", msgid, digest)
		}
	}

	begin := 16-offset-4
	ret := binary.BigEndian.Uint32(digest[begin:])
	if verbose {
		fmt.Printf("Offset=%x (hex)\n", ret)
	}
	return ret
}

// Determine if msgid can be found based on feed match
func match(msgid string, feed string) (bool, error) {
	var (
		from uint32
		to uint32
		base uint32
		offset int
	)
	n, e := fmt.Sscanf(feed, "%d-%d/%d:%d", &from, &to, &base, &offset)
	if e != nil {
		return false, e
	}
	if n != 4 {
		return false, fmt.Errorf("Failed parsing feed-str. n=%d", n)
	}

	pos := hash(msgid, offset)
	modulo := pos % base +1
	if verbose {
		fmt.Printf("Modulo (%d %% %d +1) = %d\n", pos, base, modulo)
	}
	if modulo >= from && modulo <= to {
		return true, nil
	}
	return false, nil
}

func main() {
	var (
		msgid string
		hashfeed string
	)
	flag.BoolVar(&verbose, "v", true, "Verbose")
	flag.StringVar(&msgid, "msgid", "<>", "Msgid")
	flag.StringVar(&hashfeed, "hashfeed", "1-120/360:8", "Diablo Hashfeed from dserver.hosts/dnewsfeeds")
	flag.Parse()

	ok, e := match(msgid, hashfeed)
	if e != nil {
		fmt.Fprintf(os.Stderr, e.Error())
	}
	if ok {
		if verbose {
			fmt.Printf("Match\n")
		}
		os.Exit(0)
	}
	if verbose {
		fmt.Printf("Mismatch\n")
	}
	os.Exit(1)
}