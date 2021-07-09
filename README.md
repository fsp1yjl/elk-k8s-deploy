# elk-k8s-deploy

本项目启动后会基于config/config.json的配置，基于配置模版，
分别生成es, filebeat,logstash,kibana的k8s 部署yaml文件放在tmp目录下 



## 配置说明 

```code
{
  "EsYaml": {
    "Namespace": "logging",  // es部署使用的namespace
    "PV": [   // es部署要使用的localpv 信息，根据es的pod个数去配置
      {"Storage":  "10Gi", "Node":"001"},// storage代表pv容量， node代表选择的节点
      {"Storage":  "10Gi", "Node":"001"}
    ],
    "Image": "harbor.cestc.com/paas/elasticsearch:7.6.1-ik",  // es使用的镜像
    "Ingress": "paas-es.oc25.wlc.cecloudcs.com",   // es的ingress 域名
    "Storage": "10Gi"      //代表es节点的数据容量，小于等于pv的容量即可
  },
  "FilebeatYaml": {
    "Image": "elastic/filebeat:7.3.1",  // filebeat部署使用的镜像
    "Namespace": "logging"   // filebeat部署使用的namespace
  },
  "LogstashYaml": {
    "Namespace": "logging",  // logstash部署使用的namespace   
    "Image": "elastic/logstash:7.3.1" // logstash 部署使用的镜像
  },
  "KibanaYaml": {
    "Namespace": "logging",  //kibana 部署使用的namespace   
    "Image": "kibana:7.3.1",  // kibana 部署使用的镜像
    "Ingress": "kibana.wlc102.intranet.cecloudcs.com",  //kibana的ingress域名地址
    "Affinity": "true",   //
    "AffinityNode": "kibana-node1"
  }
}

```

## 操作前准备

确认elfk环境是否已经创建：

查看logging namespace下全部资源：
```
kubectl get all -n logging 
```

查看es statefulset 是否存在：
```
kubectl get sts es7-cluster -n logging
```

查看logstash  deployment部署情况：
```code 
kubectl get deployment paas-logstash -n logging
```

查看 filebeat daemonset的部署情况：
```code
kubectl get daemonset filebeat -n logging
```

查看kibana的部署情况：
```code 
 kubectl get deployment kibana  -n logging
```

## 生成elfk 部署yaml  

配置好config/config.json 这个配置文件，执行如下：
go run main.go 

生成的yaml会放在tmp目录下，scp tmp到对应集群节点上，
kubectl apply -f tmp 可以一次性部署elfk 
也可以分别执行各自的yaml分别部署es, filebeat, logstash, kibana:
kubectl apply -f  es-final.yml
kubectl apply -f  filebeat-final.yml
kubectl apply -f  kibana-final.yml
kubectl apply -f  logstash-final.yml

