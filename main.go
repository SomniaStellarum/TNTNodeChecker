package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	ND NodeData `json:"nodeData"`
}

type NodeData struct {
	A []Audit `json:"audits"`
	N Node    `json:"node"`
}

type Audit struct {
	TNTBal int `json:"tnt_balance_grains"`
}

type Node struct {
	CFail int `json:"consecutive_fails"`
	CPass int `json:"consecutive_passes"`
	Fail  int `json:"fail_count"`
	Pass  int `json:"pass_count"`
}

type ReturnData struct {
	CFail      int    `json:"cfail"`
	CPass      int    `json:"cpass"`
	FailPer    string `json:"failper"`
	TNTBalance int    `json:"tnt"`
	Fail       int    `json:"fail"`
	Pass       int    `json:"pass"`
}

var server = flag.String("host", "http://localhost:80", "Change host")
var tnt = flag.String("tnt", "0x0000000000000000000000000000000000000000", "tnt address")

func main() {
	flag.Parse()
	path := *server + "/stats"
	r, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	r.Header.Add("auth", strings.ToLower(*tnt))
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading body of request: %v", err)
	}
	//fmt.Print(string(b))
	var d Data
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatalf("Error parsing json: %v", err)
	}
	//fmt.Print(d)
	var rData ReturnData
	rData.CFail = d.ND.N.CFail
	rData.CPass = d.ND.N.CPass
	rData.Fail = d.ND.N.Fail
	rData.Pass = d.ND.N.Pass
	p, f := float32(rData.Pass), float32(rData.Fail)
	rData.FailPer = fmt.Sprintf("%.2f", f/(f+p))
	rData.TNTBalance = d.ND.A[0].TNTBal / 100000000
	//fmt.Print(rData)
	v, err := json.Marshal(&rData)
	if err != nil {
		log.Fatalf("Error generating json: %v", err)
	}
	fmt.Print(string(v))
}
