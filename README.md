## ping-probe-register

ping-probe-register基于Consul，实现ping监控探针的自动注册和注销功能。

在各个VPC中部署ping-probe-register探针之前，请保证网络可达：

- ping-probe需要访问consul，注册探针，10.67.47.50:8500
- Consul需要访问各个ping-probe探针的9347端口，做Service的健康检查
- Prometheus采集器需要访问各个ping-probe探针的9346端口

探针安全组开放：
- 9346 开放10.107.0.0/16
- 9347 开放10.67.0.0/16
- icmp 开放 10.0.0.0/8 172.16.0.0/12 172.32.0.0/11

请使用deploy/manifest.yaml文件进行部署

![image.png](https://i.loli.net/2021/03/31/SEBTlU4WmJ9qPGk.png)

注意：以上框选的字段请按照所在的集群信息填写。


