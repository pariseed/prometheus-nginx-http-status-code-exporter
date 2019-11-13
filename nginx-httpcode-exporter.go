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



var code100 = []int{}
var code100count = 0
var code101 = []int{}
var code101count = 0
var code102 = []int{}
var code102count = 0
var code103 = []int{}
var code103count = 0

var code200 = []int{}
var code200count = 0
var code201 = []int{}
var code201count = 0
var code202 = []int{}
var code202count = 0
var code203 = []int{}
var code203count = 0
var code204 = []int{}
var code204count = 0
var code205 = []int{}
var code205count = 0
var code206 = []int{}
var code206count = 0
var code207 = []int{}
var code207count = 0
var code208 = []int{}
var code208count = 0
var code226 = []int{}
var code226count = 0

var code300 = []int{}
var code300count = 0
var code301 = []int{}
var code301count = 0
var code302 = []int{}
var code302count = 0
var code303 = []int{}
var code303count = 0
var code304 = []int{}
var code304count = 0
var code305 = []int{}
var code305count = 0
var code306 = []int{}
var code306count = 0
var code307 = []int{}
var code307count = 0
var code308 = []int{}
var code308count = 0

var code400 = []int{}
var code400count = 0
var code401 = []int{}
var code401count = 0
var code402 = []int{}
var code402count = 0
var code403 = []int{}
var code403count = 0
var code404 = []int{}
var code404count = 0
var code405 = []int{}
var code405count = 0
var code406 = []int{}
var code406count = 0
var code407 = []int{}
var code407count = 0
var code408 = []int{}
var code408count = 0
var code409 = []int{}
var code409count = 0
var code410 = []int{}
var code410count = 0
var code411 = []int{}
var code411count = 0
var code412 = []int{}
var code412count = 0
var code413 = []int{}
var code413count = 0
var code414 = []int{}
var code414count = 0
var code415 = []int{}
var code415count = 0
var code416 = []int{}
var code416count = 0
var code417 = []int{}
var code417count = 0
var code418 = []int{}
var code418count = 0
var code421 = []int{}
var code421count = 0
var code422 = []int{}
var code422count = 0
var code423 = []int{}
var code423count = 0
var code424 = []int{}
var code424count = 0
var code425 = []int{}
var code425count = 0
var code426 = []int{}
var code426count = 0
var code427 = []int{}
var code427count = 0
var code428 = []int{}
var code428count = 0
var code429 = []int{}
var code429count = 0
var code431 = []int{}
var code431count = 0
var code451 = []int{}
var code451count = 0

var code500 = []int{}
var code500count = 0
var code501 = []int{}
var code501count = 0
var code502 = []int{}
var code502count = 0
var code503 = []int{}
var code503count = 0
var code504 = []int{}
var code504count = 0
var code505 = []int{}
var code505count = 0
var code506 = []int{}
var code506count = 0
var code507 = []int{}
var code507count = 0
var code508 = []int{}
var code508count = 0
var code510 = []int{}
var code510count = 0
var code511 = []int{}
var code511count = 0

var code700 = []int{}
var code700count = 0





func getlog() {

	t, _ := tail.TailFile("/var/log/nginx/nginx-access.log", tail.Config{Follow: true})

	for line := range t.Lines {

		words := strings.Split(line.Text, "	")
		code, _ := strconv.Atoi(words[3])

		if (code == 100) {
			code100 = append(code100, code) 

		} else if (code == 101) {
			code101 = append(code101, code) 

		} else if (code == 102) {
			code102 = append(code102, code) 

		} else if (code == 103) {
			code103 = append(code103, code) 

		} else if (code == 200) {
			code200 = append(code200, code) 
		
		} else if (code == 201) {
			code201 = append(code201, code) 
	
		} else if (code == 202) {
			code202 = append(code202, code) 
	
		} else if (code == 203) {
			code203 = append(code203, code) 

		} else if (code == 204) {
			code204 = append(code204, code) 
		
		} else if (code == 205) {
			code205 = append(code205, code) 
	
		} else if (code == 206) {
			code206 = append(code206, code) 
	
		} else if (code == 207) {
			code207 = append(code207, code) 
	
		} else if (code == 208) {
			code208 = append(code208, code) 
	
		} else if (code == 226) {
			code226 = append(code226, code) 
	
		} else if (code == 300) {
			code300 = append(code300, code) 
	
		} else if (code == 301) {
			code301 = append(code301, code) 
	
		} else if (code == 302) {
			code302 = append(code302, code) 

		} else if (code == 303) {
			code303 = append(code303, code) 

		} else if (code == 304) {
			code304 = append(code304, code) 

		} else if (code == 305) {
			code305 = append(code305, code) 

		} else if (code == 306) {
			code306 = append(code306, code) 

		} else if (code == 307) {
			code307 = append(code307, code) 

		} else if (code == 308) {
			code308 = append(code308, code) 

		} else if (code == 400) {
			code400 = append(code400, code) 

		} else if (code == 401) {
			code401 = append(code401, code) 

		} else if (code == 402) {
			code402 = append(code402, code) 

		} else if (code == 403) {
			code403 = append(code403, code) 

		} else if (code == 404) {
			code404 = append(code404, code) 

		} else if (code == 405) {
			code405 = append(code405, code) 

		} else if (code == 406) {
			code406 = append(code406, code) 

		} else if (code == 407) {
			code407 = append(code407, code) 

		} else if (code == 408) {
			code408 = append(code408, code) 

		} else if (code == 409) {
			code409 = append(code409, code) 

		} else if (code == 410) {
			code410 = append(code410, code) 

		} else if (code == 411) {
			code411 = append(code411, code) 

		} else if (code == 412) {
			code412 = append(code412, code) 

		} else if (code == 413) {
			code413 = append(code413, code) 

		} else if (code == 414) {
			code414 = append(code414, code) 

		} else if (code == 415) {
			code415 = append(code415, code) 

		} else if (code == 416) {
			code416 = append(code416, code) 

		} else if (code == 417) {
			code417 = append(code417, code) 

		} else if (code == 418) {
			code418 = append(code418, code) 

		} else if (code == 421) {
			code421 = append(code421, code) 

		} else if (code == 422) {
			code422 = append(code422, code) 

		} else if (code == 423) {
			code423 = append(code423, code) 

		} else if (code == 424) {
			code424 = append(code424, code) 

		} else if (code == 425) {
			code425 = append(code425, code) 

		} else if (code == 426) {
			code426 = append(code426, code) 

		} else if (code == 427) {
			code427 = append(code427, code) 

		} else if (code == 428) {
			code428 = append(code428, code) 

		} else if (code == 429) {
			code429 = append(code429, code) 

		} else if (code == 431) {
			code431 = append(code431, code) 

		} else if (code == 451) {
			code451 = append(code451, code) 

		} else if (code == 500) {
			code500 = append(code500, code) 

		} else if (code == 501) {
			code501 = append(code501, code) 

		} else if (code == 502) {
			code502 = append(code502, code) 

		} else if (code == 503) {
			code503 = append(code503, code) 

		} else if (code == 504) {
			code504 = append(code504, code) 

		} else if (code == 505) {
			code505 = append(code505, code) 

		} else if (code == 506) {
			code506 = append(code506, code) 

		} else if (code == 507) {
			code507 = append(code507, code) 

		} else if (code == 508) {
			code508 = append(code508, code) 

		} else if (code == 510) {
			code510 = append(code510, code) 

		} else if (code == 511) {
			code511 = append(code511, code) 

		} else if (code == 700) {
			code700 = append(code700, code) 
		}

	}

}




var HTTP_100_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_100_CODE",
		Help: "number of nginx http 100 status code"})

var HTTP_101_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_101_CODE",
		Help: "number of nginx http 101 status code"})

var HTTP_102_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_102_CODE",
		Help: "number of nginx http 102 status code"})

var HTTP_103_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_103_CODE",
		Help: "number of nginx http 103 status code"})

var HTTP_200_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_200_CODE",
		Help: "number of nginx http 200 status code"})

var HTTP_201_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_201_CODE",
		Help: "number of nginx http 201 status code"})

var HTTP_202_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_202_CODE",
		Help: "number of nginx http 202 status code"})

var HTTP_203_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_203_CODE",
		Help: "number of nginx http 203 status code"})

var HTTP_204_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_204_CODE",
		Help: "number of nginx http 204 status code"})

var HTTP_205_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_205_CODE",
		Help: "number of nginx http 205 status code"})

var HTTP_206_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_206_CODE",
		Help: "number of nginx http 206 status code"})

var HTTP_207_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_207_CODE",
		Help: "number of nginx http 207 status code"})

var HTTP_208_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_208_CODE",
		Help: "number of nginx http 208 status code"})

var HTTP_226_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_226_CODE",
		Help: "number of nginx http 226 status code"})

var HTTP_300_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_300_CODE",
		Help: "number of nginx http 300 status code"})

var HTTP_301_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_301_CODE",
		Help: "number of nginx http 301 status code"})

var HTTP_302_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_302_CODE",
		Help: "number of nginx http 302 status code"})

var HTTP_303_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_303_CODE",
		Help: "number of nginx http 303 status code"})

var HTTP_304_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_304_CODE",
		Help: "number of nginx http 304 status code"})

var HTTP_305_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_305_CODE",
		Help: "number of nginx http 305 status code"})

var HTTP_306_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_306_CODE",
		Help: "number of nginx http 306 status code"})

var HTTP_307_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_307_CODE",
		Help: "number of nginx http 307 status code"})

var HTTP_308_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_308_CODE",
		Help: "number of nginx http 308 status code"})

var HTTP_400_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_400_CODE",
		Help: "number of nginx http 400 status code"})

var HTTP_401_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_401_CODE",
		Help: "number of nginx http 401 status code"})

var HTTP_402_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_402_CODE",
		Help: "number of nginx http 402 status code"})

var HTTP_403_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_403_CODE",
		Help: "number of nginx http 403 status code"})

var HTTP_404_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_404_CODE",
		Help: "number of nginx http 404 status code"})

var HTTP_405_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_405_CODE",
		Help: "number of nginx http 405 status code"})

var HTTP_406_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_406_CODE",
		Help: "number of nginx http 406 status code"})

var HTTP_407_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_407_CODE",
		Help: "number of nginx http 407 status code"})

var HTTP_408_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_408_CODE",
		Help: "number of nginx http 408 status code"})

var HTTP_409_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_409_CODE",
		Help: "number of nginx http 409 status code"})

var HTTP_410_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_410_CODE",
		Help: "number of nginx http 410 status code"})

var HTTP_411_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_411_CODE",
		Help: "number of nginx http 411 status code"})

var HTTP_412_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_412_CODE",
		Help: "number of nginx http 412 status code"})

var HTTP_413_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_413_CODE",
		Help: "number of nginx http 413 status code"})

var HTTP_414_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_414_CODE",
		Help: "number of nginx http 414 status code"})

var HTTP_415_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_415_CODE",
		Help: "number of nginx http 415 status code"})

var HTTP_416_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_416_CODE",
		Help: "number of nginx http 416 status code"})

var HTTP_417_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_417_CODE",
		Help: "number of nginx http 417 status code"})

var HTTP_418_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_418_CODE",
		Help: "number of nginx http 418 status code"})

var HTTP_421_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_421_CODE",
		Help: "number of nginx http 421 status code"})

var HTTP_422_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_422_CODE",
		Help: "number of nginx http 422 status code"})

var HTTP_423_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_423_CODE",
		Help: "number of nginx http 423 status code"})

var HTTP_424_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_424_CODE",
		Help: "number of nginx http 424 status code"})

var HTTP_425_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_425_CODE",
		Help: "number of nginx http 425 status code"})

var HTTP_426_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_426_CODE",
		Help: "number of nginx http 426 status code"})

var HTTP_427_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_427_CODE",
		Help: "number of nginx http 427 status code"})

var HTTP_428_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_428_CODE",
		Help: "number of nginx http 428 status code"})

var HTTP_429_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_429_CODE",
		Help: "number of nginx http 429 status code"})

var HTTP_431_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_431_CODE",
		Help: "number of nginx http 431 status code"})

var HTTP_451_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_451_CODE",
		Help: "number of nginx http 451 status code"})

var HTTP_500_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_500_CODE",
		Help: "number of nginx http 500 status code"})

var HTTP_501_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_501_CODE",
		Help: "number of nginx http 501 status code"})

var HTTP_502_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_502_CODE",
		Help: "number of nginx http 502 status code"})

var HTTP_503_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_503_CODE",
		Help: "number of nginx http 503 status code"})

var HTTP_504_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_504_CODE",
		Help: "number of nginx http 504 status code"})

var HTTP_505_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_505_CODE",
		Help: "number of nginx http 505 status code"})

var HTTP_506_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_506_CODE",
		Help: "number of nginx http 506 status code"})

var HTTP_507_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_507_CODE",
		Help: "number of nginx http 507 status code"})

var HTTP_508_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_508_CODE",
		Help: "number of nginx http 508 status code"})

var HTTP_510_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_510_CODE",
		Help: "number of nginx http 510 status code"})

var HTTP_511_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_511_CODE",
		Help: "number of nginx http 511 status code"})

var HTTP_700_CODE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "HTTP_700_CODE",
		Help: "number of nginx http 700 status code"})






func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(HTTP_100_CODE)
	prometheus.MustRegister(HTTP_101_CODE)
	prometheus.MustRegister(HTTP_102_CODE)
	prometheus.MustRegister(HTTP_103_CODE)
	prometheus.MustRegister(HTTP_200_CODE)
	prometheus.MustRegister(HTTP_201_CODE)
	prometheus.MustRegister(HTTP_202_CODE)
	prometheus.MustRegister(HTTP_203_CODE)
	prometheus.MustRegister(HTTP_204_CODE)
	prometheus.MustRegister(HTTP_205_CODE)
	prometheus.MustRegister(HTTP_206_CODE)
	prometheus.MustRegister(HTTP_207_CODE)
	prometheus.MustRegister(HTTP_208_CODE)
	prometheus.MustRegister(HTTP_226_CODE)
	prometheus.MustRegister(HTTP_300_CODE)
	prometheus.MustRegister(HTTP_301_CODE)
	prometheus.MustRegister(HTTP_302_CODE)
	prometheus.MustRegister(HTTP_303_CODE)
	prometheus.MustRegister(HTTP_304_CODE)
	prometheus.MustRegister(HTTP_305_CODE)
	prometheus.MustRegister(HTTP_306_CODE)
	prometheus.MustRegister(HTTP_307_CODE)
	prometheus.MustRegister(HTTP_308_CODE)
	prometheus.MustRegister(HTTP_400_CODE)
	prometheus.MustRegister(HTTP_401_CODE)
	prometheus.MustRegister(HTTP_402_CODE)
	prometheus.MustRegister(HTTP_403_CODE)
	prometheus.MustRegister(HTTP_404_CODE)
	prometheus.MustRegister(HTTP_405_CODE)
	prometheus.MustRegister(HTTP_406_CODE)
	prometheus.MustRegister(HTTP_407_CODE)
	prometheus.MustRegister(HTTP_408_CODE)
	prometheus.MustRegister(HTTP_409_CODE)
	prometheus.MustRegister(HTTP_410_CODE)
	prometheus.MustRegister(HTTP_411_CODE)
	prometheus.MustRegister(HTTP_412_CODE)
	prometheus.MustRegister(HTTP_413_CODE)
	prometheus.MustRegister(HTTP_414_CODE)
	prometheus.MustRegister(HTTP_415_CODE)
	prometheus.MustRegister(HTTP_416_CODE)
	prometheus.MustRegister(HTTP_417_CODE)
	prometheus.MustRegister(HTTP_418_CODE)
	prometheus.MustRegister(HTTP_421_CODE)
	prometheus.MustRegister(HTTP_422_CODE)
	prometheus.MustRegister(HTTP_423_CODE)
	prometheus.MustRegister(HTTP_424_CODE)
	prometheus.MustRegister(HTTP_425_CODE)
	prometheus.MustRegister(HTTP_426_CODE)
	prometheus.MustRegister(HTTP_427_CODE)
	prometheus.MustRegister(HTTP_428_CODE)
	prometheus.MustRegister(HTTP_429_CODE)
	prometheus.MustRegister(HTTP_431_CODE)
	prometheus.MustRegister(HTTP_451_CODE)
	prometheus.MustRegister(HTTP_500_CODE)
	prometheus.MustRegister(HTTP_501_CODE)
	prometheus.MustRegister(HTTP_502_CODE)
	prometheus.MustRegister(HTTP_503_CODE)
	prometheus.MustRegister(HTTP_504_CODE)
	prometheus.MustRegister(HTTP_505_CODE)
	prometheus.MustRegister(HTTP_506_CODE)
	prometheus.MustRegister(HTTP_507_CODE)
	prometheus.MustRegister(HTTP_508_CODE)
	prometheus.MustRegister(HTTP_510_CODE)
	prometheus.MustRegister(HTTP_511_CODE)
	prometheus.MustRegister(HTTP_700_CODE)


}


func commit100() {

	diff := math.Abs(float64(code100count)-float64(len(code100)))
	
	


	HTTP_100_CODE.Set(diff)

	code100count = len(code100)

	code100 = code100[:0]

	time.Sleep(1 * time.Second)

}

func commit101() {

	diff := math.Abs(float64(code101count)-float64(len(code101)))
	
	


	HTTP_101_CODE.Set(diff)

	code101count = len(code101)

	code101 = code101[:0]

	time.Sleep(1 * time.Second)

}

func commit102() {

	diff := math.Abs(float64(code102count)-float64(len(code102)))
	
	


	HTTP_102_CODE.Set(diff)

	code102count = len(code102)

	code102 = code102[:0]

	time.Sleep(1 * time.Second)

}

func commit103() {

	diff := math.Abs(float64(code103count)-float64(len(code103)))
	
	


	HTTP_103_CODE.Set(diff)

	code103count = len(code103)

	code103 = code103[:0]

	time.Sleep(1 * time.Second)

}

func commit200() {

	diff := math.Abs(float64(code200count)-float64(len(code200)))
	
	


	HTTP_200_CODE.Set(diff)

	code200count = len(code200)

	code200 = code200[:0]

	time.Sleep(1 * time.Second)

}

func commit201() {

	diff := math.Abs(float64(code201count)-float64(len(code201)))
	
	


	HTTP_201_CODE.Set(diff)

	code201count = len(code201)

	code201 = code201[:0]

	time.Sleep(1 * time.Second)

}

func commit202() {

	diff := math.Abs(float64(code202count)-float64(len(code202)))
	
	


	HTTP_202_CODE.Set(diff)

	code202count = len(code202)

	code202 = code202[:0]

	time.Sleep(1 * time.Second)

}

func commit203() {

	diff := math.Abs(float64(code203count)-float64(len(code203)))
	
	


	HTTP_203_CODE.Set(diff)

	code203count = len(code203)

	code203 = code203[:0]

	time.Sleep(1 * time.Second)

}

func commit204() {

	diff := math.Abs(float64(code204count)-float64(len(code204)))
	
	


	HTTP_204_CODE.Set(diff)

	code204count = len(code204)

	code204 = code204[:0]

	time.Sleep(1 * time.Second)

}

func commit205() {

	diff := math.Abs(float64(code205count)-float64(len(code205)))
	
	


	HTTP_205_CODE.Set(diff)

	code205count = len(code205)

	code205 = code205[:0]

	time.Sleep(1 * time.Second)

}

func commit206() {

	diff := math.Abs(float64(code206count)-float64(len(code206)))
	
	


	HTTP_206_CODE.Set(diff)

	code206count = len(code206)

	code206 = code206[:0]

	time.Sleep(1 * time.Second)

}

func commit207() {

	diff := math.Abs(float64(code207count)-float64(len(code207)))
	
	


	HTTP_207_CODE.Set(diff)

	code207count = len(code207)

	code207 = code207[:0]

	time.Sleep(1 * time.Second)

}

func commit208() {

	diff := math.Abs(float64(code208count)-float64(len(code208)))
	
	


	HTTP_208_CODE.Set(diff)

	code208count = len(code208)

	code208 = code208[:0]

	time.Sleep(1 * time.Second)

}

func commit226() {

	diff := math.Abs(float64(code226count)-float64(len(code226)))
	
	


	HTTP_226_CODE.Set(diff)

	code226count = len(code226)

	code226 = code226[:0]

	time.Sleep(1 * time.Second)

}

func commit300() {

	diff := math.Abs(float64(code300count)-float64(len(code300)))
	
	


	HTTP_300_CODE.Set(diff)

	code300count = len(code300)

	code300 = code300[:0]

	time.Sleep(1 * time.Second)

}

func commit301() {

	diff := math.Abs(float64(code301count)-float64(len(code301)))
	
	


	HTTP_301_CODE.Set(diff)

	code301count = len(code301)

	code301 = code301[:0]

	time.Sleep(1 * time.Second)

}

func commit302() {

	diff := math.Abs(float64(code302count)-float64(len(code302)))
	
	


	HTTP_302_CODE.Set(diff)

	code302count = len(code302)

	code302 = code302[:0]

	time.Sleep(1 * time.Second)

}

func commit303() {

	diff := math.Abs(float64(code303count)-float64(len(code303)))
	
	


	HTTP_303_CODE.Set(diff)

	code303count = len(code303)

	code303 = code303[:0]

	time.Sleep(1 * time.Second)

}

func commit304() {

	diff := math.Abs(float64(code304count)-float64(len(code304)))
	
	


	HTTP_304_CODE.Set(diff)

	code304count = len(code304)

	code304 = code304[:0]

	time.Sleep(1 * time.Second)

}

func commit305() {

	diff := math.Abs(float64(code305count)-float64(len(code305)))
	
	


	HTTP_305_CODE.Set(diff)

	code305count = len(code305)

	code305 = code305[:0]

	time.Sleep(1 * time.Second)

}

func commit306() {

	diff := math.Abs(float64(code306count)-float64(len(code306)))
	
	


	HTTP_306_CODE.Set(diff)

	code306count = len(code306)

	code306 = code306[:0]

	time.Sleep(1 * time.Second)

}

func commit307() {

	diff := math.Abs(float64(code307count)-float64(len(code307)))
	
	


	HTTP_307_CODE.Set(diff)

	code307count = len(code307)

	code307 = code307[:0]

	time.Sleep(1 * time.Second)

}

func commit308() {

	diff := math.Abs(float64(code308count)-float64(len(code308)))
	
	


	HTTP_308_CODE.Set(diff)

	code308count = len(code308)

	code308 = code308[:0]

	time.Sleep(1 * time.Second)

}

func commit400() {

	diff := math.Abs(float64(code400count)-float64(len(code400)))
	
	


	HTTP_400_CODE.Set(diff)

	code400count = len(code400)

	code400 = code400[:0]

	time.Sleep(1 * time.Second)

}

func commit401() {

	diff := math.Abs(float64(code401count)-float64(len(code401)))
	
	


	HTTP_401_CODE.Set(diff)

	code401count = len(code401)

	code401 = code401[:0]

	time.Sleep(1 * time.Second)

}

func commit402() {

	diff := math.Abs(float64(code402count)-float64(len(code402)))
	
	


	HTTP_402_CODE.Set(diff)

	code402count = len(code402)

	code402 = code402[:0]

	time.Sleep(1 * time.Second)

}

func commit403() {

	diff := math.Abs(float64(code403count)-float64(len(code403)))
	
	


	HTTP_403_CODE.Set(diff)

	code403count = len(code403)

	code403 = code403[:0]

	time.Sleep(1 * time.Second)

}

func commit404() {

	diff := math.Abs(float64(code404count)-float64(len(code404)))
	
	


	HTTP_404_CODE.Set(diff)

	code404count = len(code404)

	code404 = code404[:0]

	time.Sleep(1 * time.Second)

}

func commit405() {

	diff := math.Abs(float64(code405count)-float64(len(code405)))
	
	


	HTTP_405_CODE.Set(diff)

	code405count = len(code405)

	code405 = code405[:0]

	time.Sleep(1 * time.Second)

}

func commit406() {

	diff := math.Abs(float64(code406count)-float64(len(code406)))
	
	


	HTTP_406_CODE.Set(diff)

	code406count = len(code406)

	code406 = code406[:0]

	time.Sleep(1 * time.Second)

}

func commit407() {

	diff := math.Abs(float64(code407count)-float64(len(code407)))
	
	


	HTTP_407_CODE.Set(diff)

	code407count = len(code407)

	code407 = code407[:0]

	time.Sleep(1 * time.Second)

}

func commit408() {

	diff := math.Abs(float64(code408count)-float64(len(code408)))
	
	


	HTTP_408_CODE.Set(diff)

	code408count = len(code408)

	code408 = code408[:0]

	time.Sleep(1 * time.Second)

}

func commit409() {

	diff := math.Abs(float64(code409count)-float64(len(code409)))
	
	


	HTTP_409_CODE.Set(diff)

	code409count = len(code409)

	code409 = code409[:0]

	time.Sleep(1 * time.Second)

}

func commit410() {

	diff := math.Abs(float64(code410count)-float64(len(code410)))
	
	


	HTTP_410_CODE.Set(diff)

	code410count = len(code410)

	code410 = code410[:0]

	time.Sleep(1 * time.Second)

}

func commit411() {

	diff := math.Abs(float64(code411count)-float64(len(code411)))
	
	


	HTTP_411_CODE.Set(diff)

	code411count = len(code411)

	code411 = code411[:0]

	time.Sleep(1 * time.Second)

}

func commit412() {

	diff := math.Abs(float64(code412count)-float64(len(code412)))
	
	


	HTTP_412_CODE.Set(diff)

	code412count = len(code412)

	code412 = code412[:0]

	time.Sleep(1 * time.Second)

}

func commit413() {

	diff := math.Abs(float64(code413count)-float64(len(code413)))
	
	


	HTTP_413_CODE.Set(diff)

	code413count = len(code413)

	code413 = code413[:0]

	time.Sleep(1 * time.Second)

}

func commit414() {

	diff := math.Abs(float64(code414count)-float64(len(code414)))
	
	


	HTTP_414_CODE.Set(diff)

	code414count = len(code414)

	code414 = code414[:0]

	time.Sleep(1 * time.Second)

}

func commit415() {

	diff := math.Abs(float64(code415count)-float64(len(code415)))
	
	


	HTTP_415_CODE.Set(diff)

	code415count = len(code415)

	code415 = code415[:0]

	time.Sleep(1 * time.Second)

}

func commit416() {

	diff := math.Abs(float64(code416count)-float64(len(code416)))
	
	


	HTTP_416_CODE.Set(diff)

	code416count = len(code416)

	code416 = code416[:0]

	time.Sleep(1 * time.Second)

}

func commit417() {

	diff := math.Abs(float64(code417count)-float64(len(code417)))
	
	


	HTTP_417_CODE.Set(diff)

	code417count = len(code417)

	code417 = code417[:0]

	time.Sleep(1 * time.Second)

}

func commit418() {

	diff := math.Abs(float64(code418count)-float64(len(code418)))
	
	


	HTTP_418_CODE.Set(diff)

	code418count = len(code418)

	code418 = code418[:0]

	time.Sleep(1 * time.Second)

}

func commit421() {

	diff := math.Abs(float64(code421count)-float64(len(code421)))
	
	


	HTTP_421_CODE.Set(diff)

	code421count = len(code421)

	code421 = code421[:0]

	time.Sleep(1 * time.Second)

}

func commit422() {

	diff := math.Abs(float64(code422count)-float64(len(code422)))
	
	


	HTTP_422_CODE.Set(diff)

	code422count = len(code422)

	code422 = code422[:0]

	time.Sleep(1 * time.Second)

}

func commit423() {

	diff := math.Abs(float64(code423count)-float64(len(code423)))
	
	


	HTTP_423_CODE.Set(diff)

	code423count = len(code423)

	code423 = code423[:0]

	time.Sleep(1 * time.Second)

}

func commit424() {

	diff := math.Abs(float64(code424count)-float64(len(code424)))
	
	


	HTTP_424_CODE.Set(diff)

	code424count = len(code424)

	code424 = code424[:0]

	time.Sleep(1 * time.Second)

}

func commit425() {

	diff := math.Abs(float64(code425count)-float64(len(code425)))
	
	


	HTTP_425_CODE.Set(diff)

	code425count = len(code425)

	code425 = code425[:0]

	time.Sleep(1 * time.Second)

}

func commit426() {

	diff := math.Abs(float64(code426count)-float64(len(code426)))
	
	


	HTTP_426_CODE.Set(diff)

	code426count = len(code426)

	code426 = code426[:0]

	time.Sleep(1 * time.Second)

}

func commit427() {

	diff := math.Abs(float64(code427count)-float64(len(code427)))
	
	


	HTTP_427_CODE.Set(diff)

	code427count = len(code427)

	code427 = code427[:0]

	time.Sleep(1 * time.Second)

}

func commit428() {

	diff := math.Abs(float64(code428count)-float64(len(code428)))
	
	


	HTTP_428_CODE.Set(diff)

	code428count = len(code428)

	code428 = code428[:0]

	time.Sleep(1 * time.Second)

}

func commit429() {

	diff := math.Abs(float64(code429count)-float64(len(code429)))
	
	


	HTTP_429_CODE.Set(diff)

	code429count = len(code429)

	code429 = code429[:0]

	time.Sleep(1 * time.Second)

}

func commit431() {

	diff := math.Abs(float64(code431count)-float64(len(code431)))
	
	


	HTTP_431_CODE.Set(diff)

	code431count = len(code431)

	code431 = code431[:0]

	time.Sleep(1 * time.Second)

}

func commit451() {

	diff := math.Abs(float64(code451count)-float64(len(code451)))
	
	


	HTTP_451_CODE.Set(diff)

	code451count = len(code451)

	code451 = code451[:0]

	time.Sleep(1 * time.Second)

}

func commit500() {

	diff := math.Abs(float64(code500count)-float64(len(code500)))
	
	


	HTTP_500_CODE.Set(diff)

	code500count = len(code500)

	code500 = code500[:0]

	time.Sleep(1 * time.Second)

}

func commit501() {

	diff := math.Abs(float64(code501count)-float64(len(code501)))
	
	


	HTTP_501_CODE.Set(diff)

	code501count = len(code501)

	code501 = code501[:0]

	time.Sleep(1 * time.Second)

}

func commit502() {

	diff := math.Abs(float64(code502count)-float64(len(code502)))
	
	


	HTTP_502_CODE.Set(diff)

	code502count = len(code502)

	code502 = code502[:0]

	time.Sleep(1 * time.Second)

}

func commit503() {

	diff := math.Abs(float64(code503count)-float64(len(code503)))
	
	


	HTTP_503_CODE.Set(diff)

	code503count = len(code503)

	code503 = code503[:0]

	time.Sleep(1 * time.Second)

}

func commit504() {

	diff := math.Abs(float64(code504count)-float64(len(code504)))
	
	


	HTTP_504_CODE.Set(diff)

	code504count = len(code504)

	code504 = code504[:0]

	time.Sleep(1 * time.Second)

}

func commit505() {

	diff := math.Abs(float64(code505count)-float64(len(code505)))
	
	


	HTTP_505_CODE.Set(diff)

	code505count = len(code505)

	code505 = code505[:0]

	time.Sleep(1 * time.Second)

}

func commit506() {

	diff := math.Abs(float64(code506count)-float64(len(code506)))
	
	


	HTTP_506_CODE.Set(diff)

	code506count = len(code506)

	code506 = code506[:0]

	time.Sleep(1 * time.Second)

}

func commit507() {

	diff := math.Abs(float64(code507count)-float64(len(code507)))
	
	


	HTTP_507_CODE.Set(diff)

	code507count = len(code507)

	code507 = code507[:0]

	time.Sleep(1 * time.Second)

}

func commit508() {

	diff := math.Abs(float64(code508count)-float64(len(code508)))
	
	


	HTTP_508_CODE.Set(diff)

	code508count = len(code508)

	code508 = code508[:0]

	time.Sleep(1 * time.Second)

}

func commit510() {

	diff := math.Abs(float64(code510count)-float64(len(code510)))
	
	


	HTTP_510_CODE.Set(diff)

	code510count = len(code510)

	code510 = code510[:0]

	time.Sleep(1 * time.Second)

}

func commit511() {

	diff := math.Abs(float64(code511count)-float64(len(code511)))
	
	


	HTTP_511_CODE.Set(diff)

	code511count = len(code511)

	code511 = code511[:0]

	time.Sleep(1 * time.Second)

}

func commit700() {

	diff := math.Abs(float64(code700count)-float64(len(code700)))
	
	


	HTTP_700_CODE.Set(diff)

	code700count = len(code700)

	code700 = code700[:0]

	time.Sleep(1 * time.Second)

}



func exec() {
	go getlog()

	for {
		go commit100()
		go commit101()
		go commit102()
		go commit103()
		go commit200()
		go commit201()
		go commit202()
		go commit203()
		go commit204()
		go commit205()
		go commit206()
		go commit207()
		go commit208()
		go commit226()
		go commit300()
		go commit301()
		go commit302()
		go commit303()
		go commit304()
		go commit305()
		go commit306()
		go commit307()
		go commit308()
		go commit400()
		go commit401()
		go commit402()
		go commit403()
		go commit404()
		go commit405()
		go commit406()
		go commit407()
		go commit408()
		go commit409()
		go commit410()
		go commit411()
		go commit412()
		go commit413()
		go commit414()
		go commit415()
		go commit416()
		go commit417()
		go commit418()
		go commit421()
		go commit422()
		go commit423()
		go commit424()
		go commit425()
		go commit426()
		go commit427()
		go commit428()
		go commit429()
		go commit431()
		go commit451()
		go commit500()
		go commit501()
		go commit502()
		go commit503()
		go commit504()
		go commit505()
		go commit506()
		go commit507()
		go commit508()
		go commit510()
		go commit511()
		go commit700()
		time.Sleep(5 * time.Second)
	}

}


func main() {
	go exec()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8005", nil))

}



