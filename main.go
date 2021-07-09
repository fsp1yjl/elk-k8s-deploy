package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"github.com/minio/minio/pkg/ioutil"
	"text/template"
)

func main() {
    conf := config{}
    bytes,_ := ioutil.ReadFile("config/config.json")
    json.Unmarshal(bytes,&conf)

    fmt.Println("conf:", conf)


	EsFileRender(conf)
	FilebeatFileRender(conf)
	LogstashFileRender(conf)
	KibanaFileRender(conf)
}

type config struct {
	Es EsYaml `json:"EsYaml"`
	Filebeat FilebeatYaml `json:"FilebeatYaml"`
	Logstash LogstashYaml `json:"LogstashYaml"`
	Kibana KibanaYaml `json:"KibanaYaml"`
}

type pvInfo struct {
	Storage string `json:"Storage"`
	Node  string `json:"Node"`
}

type EsYaml struct {
	PV []pvInfo `json:"Pv"`
	Image string `json:"Image"`
	Ingress string `json:"Ingress"`
	Storage string  `json:"Storage"`
	Namespace string `json:"Namespace"`
}

type FilebeatYaml struct {
	Image string `json:"Image"`
	Namespace string `json:"Namespace"`
}

type LogstashYaml struct {
	Image string `json:"Image"`
	Namespace string `json:"Namespace"`
}

type KibanaYaml struct {
	Image string `json:"Image"`
	Ingress string  `json:"Ingress"`
	Affinity bool `json:"Affinity"`
	AffinityNode string `json:"AffinityNode"`
	Namespace string `json:"Namespace"`
}


func EsFileRender(conf config) {

	/*
	pvarr := []pvInfo {
		pvInfo{
			Storage: "10Gi",
			Node: "001",
		},
		pvInfo {
			Storage: "10Gi",
			Node: "001",
		},
	}

	esYaml := EsYaml{
		PV: pvarr,
		Image: "harbor.cestc.com/paas/elasticsearch:7.6.1-ik",
		Ingress: "paas-es.oc25.wlc.cecloudcs.com",
		Storage: "10Gi",
	}
	*/

	t, err := template.ParseFiles("template/es.yml")
	if err != nil {
		fmt.Println("error parse file", err)
	}

	//buf := bytes.NewBuffer(make([]byte, 0, 10000))
	buf := bytes.NewBuffer(make([]byte,0))





	err = t.Execute(buf, conf.Es)
	if err != nil {
		fmt.Println("excute error:", err)
	}

	//m := make(map[string]string)
	//fmt.Println(buf.String())

	filename := "tmp/es-final.yml"

	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("write es file error:", err)
	}

}


func FilebeatFileRender(conf config) {

/*	filebeatYaml := FilebeatYaml{
		Image: "elastic/filebeat:7.3.1",
	}*/
	filename := "tmp/filebeat-final.yml"
	tempPath := "template/filebeat.yml"

	t, err := template.ParseFiles(tempPath)
	if err != nil {
		fmt.Println("error parse file", err)
	}

	//buf := bytes.NewBuffer(make([]byte, 0, 10000))
	buf := bytes.NewBuffer(make([]byte,0))


	err = t.Execute(buf, conf.Filebeat)
	if err != nil {
		fmt.Println("excute error:", err)
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("write es file error:", err)
	}

}


func LogstashFileRender(conf config) {

/*	logstashYaml := LogstashYaml{
		Image: "elastic/logstash:7.3.1",
	}*/
	filename := "tmp/logstash-final.yml"
	tempPath := "template/logstash.yml"

	t, err := template.ParseFiles(tempPath)
	if err != nil {
		fmt.Println("error parse file", err)
	}

	//buf := bytes.NewBuffer(make([]byte, 0, 10000))
	buf := bytes.NewBuffer(make([]byte,0))

	err = t.Execute(buf, conf.Logstash)
	if err != nil {
		fmt.Println("excute error:", err)
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("write es file error:", err)
	}

}


func KibanaFileRender(conf config) {

/*	kibanaYaml := KibanaYaml{
		Image: "kibana:7.3.1",
		Ingress: "kibana.wlc102.intranet.cecloudcs.com",
		Affinity:  true,
		AffinityNode: "node01",
	}*/
	filename := "tmp/kibana-final.yml"
	tempPath := "template/kibana.yml"

	t, err := template.ParseFiles(tempPath)
	if err != nil {
		fmt.Println("error parse file", err)
	}

	//buf := bytes.NewBuffer(make([]byte, 0, 10000))
	buf := bytes.NewBuffer(make([]byte,0))

	err = t.Execute(buf, conf.Kibana)
	if err != nil {
		fmt.Println("excute error:", err)
	}


	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("write es file error:", err)
	}

}
