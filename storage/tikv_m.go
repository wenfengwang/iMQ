package storage

import (
"fmt"

"github.com/pingcap/tidb/config"
"github.com/pingcap/tidb/store/tikv"
	"strconv"
	"time"
)

func main() {
	cli, err := tikv.NewRawKVClient([]string{"192.168.179.71:12500"}, config.Security{})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	fmt.Printf("cluster ID: %d\n", cli.ClusterID())

	//key := []byte("Company")
	//val := []byte("PingCAP")

	start := time.Now().UnixNano()
	ch := make(chan interface{}, 10)
	for c := 0; c < 10; c++  {
		go func(ind int) {
			for i := 0 ; i < 1e2  ; i++ {
				err = cli.Put([]byte(strconv.Itoa(i)), value)
				if err != nil {
					panic(err)
				}
			}
			fmt.Printf("%d done.", ind)
			ch <- "done"
		}(c)
	}

	for c := 0; c < 10; c++  {
		<- ch
	}

	fmt.Println("use time:", (time.Now().UnixNano() - start) / 1e6, "ms")

	//// put key into tikv
	//err = cli.Put(key, val)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Successfully put %s:%s to tikv\n", key, val)
	//
	//// get key from tikv
	//val, err = cli.Get(key)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("found val: %s for key: %s\n", val, key)
	//
	//// delete key from tikv
	//err = cli.Delete(key)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("key: %s deleted\n", key)
	//
	//// get key again from tikv
	//val, err = cli.Get(key)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("found val: %s for key: %s\n", val, key)
}

var value = []byte("abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzab" +
	"cdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklm" +
	"nopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwx" +
	"yzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghij" +
	"klmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistu" +
	"vwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdef" +
	"ghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopq" +
	"istuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz")