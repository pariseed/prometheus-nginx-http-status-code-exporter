package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/hpcloud/tail"
	"math"
	"strconv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var s = []string{}

var code200 = []int{}
var code200count = 0


func getlog() {

	t, _ := tail.TailFile("/var/log/nginx/nginx-access.log", tail.Config{Follow: true})
	for line := range t.Lines {
	    	s = append(s, line.Text)
	}
}

func parsecode() {
	for elem := range s {

		words := strings.Split(s[elem], "	")
		code, _ := strconv.Atoi(words[3])

		if (code == 200) {

			code200 = append(code200, code) 
	

		}

	}
}




var HTTP_200_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_200_CODE",
		Help: "number of nginx http 200 status code"})


func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(HTTP_200_CODE)
}


func commit200() {

	diff := math.Abs(float64(code200count)-float64(len(code200)))
	//fmt.Printf("%.1f\n", x)
	fmt.Printf("%d\n", int(diff))

//	if (diff != 0) {

	HTTP_200_CODE.Set(diff)
//		HTTP_200_CODE.Set(diff)
//	}// else {
//		HTTP_200_CODE.Set(float64(len(code200)))
//	}

	code200count = len(code200)

	code200 = code200[:0]
//	s = s[:0]
	time.Sleep(1 * time.Second)

}


func exec() {
	go getlog()

	for {
		parsecode()
		go commit200()

		time.Sleep(10 * time.Second)
	}

}


func main() {

	go exec()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8005", nil))
  
}
