## ping-probe-register

ping-probe-register基于Consul，实现ping监控探针的自动注册和注销功能。

在各个VPC中部署ping-probe-register探针之前，请保证网络可达：

- ping-probe-register需要访问consul 8500端口，注册探针
- Consul需要访问各个ping-probe探针的9347端口，做Service的健康检查
- Prometheus采集器需要访问各个ping-probe探针的9346端口
- 各VPC之间需要开放icmp访问
