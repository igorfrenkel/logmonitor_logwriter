package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const TimestampFormat = "02/January/2006:15:04:05 -0700"

var tpls = []string{
	"127.0.0.1 - mary [%v] \"POST /api/user HTTP/1.0\" 200 120\n",
	"127.0.0.1 - mary [%v] \"POST /shmapapi/user HTTP/1.0\" 200 120\n",
	"127.0.0.1 - mary [%v] \"POST /shmoo/user HTTP/1.0\" 200 120\n",
	"127.0.0.1 - mary [%v] \"POST /flu/user HTTP/1.0\" 200 111\n",
	"127.0.0.1 - mary [%v] \"POST /api14/user HTTP/1.0\" 201 111\n",
	"127.0.0.1 - mary [%v] \"POST /api15/user HTTP/1.0\" 201 111\n",
	"127.0.0.1 - mary [%v] \"POST /api16/user HTTP/1.0\" 201 111\n",
	"127.0.0.1 - mary [%v] \"POST /api17/user HTTP/1.0\" 400 111\n",
	"127.0.0.1 - mary [%v] \"POST /api18/user HTTP/1.0\" 400 120000\n",
	"127.0.0.1 - mary [%v] \"POST /api/user HTTP/1.0\" 400 1020000\n",
	"127.0.0.1 - mary [%v] \"POST /apiv1/user HTTP/1.0\" 501 1001002\n",
	"127.0.0.1 - mary [%v] \"POST /apiv2/user HTTP/1.0\" 501 1202002\n",
	"127.0.0.1 - mary [%v] \"POST /apiv1/user HTTP/1.0\" 401 123333\n",
	"127.0.0.1 - mary [%v] \"POST /apiv2/user HTTP/1.0\" 401 1244444\n",
	"127.0.0.1 - mary [%v] \"POST /apiv2/user HTTP/1.0\" 401 1244\n",
	"127.0.0.1 - mary [%v] \"POST /apiv3/user HTTP/1.0\" 404 14442\n",
	"127.0.0.1 - mary [%v] \"POST /apiv3/user HTTP/1.0\" 404 14442\n",
	"127.0.0.1 - mary [%v] \"POST /api1/user HTTP/1.0\" 404 14442\n",
	"127.0.0.1 - mary [%v] \"POST /api2/user HTTP/1.0\" 200 12\n",
	"127.0.0.1 - mary [%v] \"POST /api3/user HTTP/1.0\" 200 1442\n",
	"127.0.0.1 - mary [%v] \"POST /api4/user HTTP/1.0\" 200 12\n",
	"127.0.0.1 - mary [%v] \"POST /api5/user HTTP/1.0\" 200 1442\n",
	"127.0.0.1 - mary [%v] \"POST /api6/user HTTP/1.0\" 200 12\n",
	"127.0.0.1 - mary [%v] \"POST /api7/user HTTP/1.0\" 200 1442\n",
	"127.0.0.1 - mary [%v] \"POST /api8/user HTTP/1.0\" 200 12\n",
	"127.0.0.1 - mary [%v] \"POST /api9/user HTTP/1.0\" 201 1244\n",
	"127.0.0.1 - mary [%v] \"POST /api10/user HTTP/1.0\" 201 142\n",
	"127.0.0.1 - mary [%v] \"POST /api11/user HTTP/1.0\" 201 1244\n",
	"127.0.0.1 - mary [%v] \"POST /api12/user HTTP/1.0\" 201 1442\n",
	"127.0.0.1 - mary [%v] \"POST /api13/user HTTP/1.0\" 201 1244\n",
}

func getLogLine() string {
	tpl := tpls[rand.Intn(len(tpls))]
	now := time.Now()
	return fmt.Sprintf(tpl, now.Format(TimestampFormat))
}

func main() {
	f, err := os.OpenFile("/tmp/access.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for {
		if _, err = f.WriteString(getLogLine()); err != nil {
			panic(err)
		}
		// <-time.After(time.Second * 1)
		<-time.After(time.Millisecond * 1)
	}
}
