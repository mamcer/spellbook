# prometheus

Prometheus is a pull based architecture, it actively scrapes or 'pull' metrics from services. 

**prometheus server**: central component that scrapes and stores time series data  
**targets**: services or applications you want to monitor  
**exporters**: agent-like components that run alongside a target and expose its metrics for example a *node exporter* would expose system metrics (cpu, memory) from a host. For custom application metrics you 'instrument' your code with a prometheus client library.  
**promQL** prometheus query language.  
**grafana** a popular visualization tool to create custom dashboards.  

the prometheus library exposes a /metrics http endpoint on your service. When a request hits the endpoint the library returns the current state of the metrics in a specific text-based format.

metric types: the library lets you define different types of metrics:
- counters: things that only increase, like requests
- gauges: for values that can go up and down, like cpu usage
- histograms: for measuring distributions of values, like request durations

query and visualize

the build in UI can be accessed at http://localhost:9090/graph and let youy run promql queries and see the resulta in a table or a basic graph. For example a query like http_requests_total would show the cumulative number of http requests.
while prometheus ui is good for basic queries, grafana is the industry standard for professional dashboards. 

## node exporter

install

```bash
# optionally
sudo mkdir /opt/node-exporter
sudo chmod mario /opt/node-exporter
cd node-exporter

wget https://github.com/prometheus/node_exporter/releases/download/v*/node_exporter-*.*-amd64.tar.gz
tar xvfz node_exporter-*.*-amd64.tar.gz
cd node_exporter-*.*-amd64
./node_exporter
```

you can configure it as a service file: [https://www.crybit.com/install-and-configure-node-exporter/](https://www.crybit.com/install-and-configure-node-exporter/)

test it i working: 

```bash
curl http://localhost:9100/metrics
```

## prometheus

[https://prometheus.io/docs/prometheus/latest/getting_started/](https://prometheus.io/docs/prometheus/latest/getting_started/)

install

```bash
# optionally
sudo mkdir /opt/prometheus
sudo chmod mario /opt/prometheus
cd prometheus

wget https://github.com/prometheus/prometheus/releases/download/v*/prometheus-*.*-amd64.tar.gz
tar xvf prometheus-*.*-amd64.tar.gz
cd prometheus-*.*
```

configure

```bash
# configure the previously installed node_exporter
vim prometheus.yml
```

```yaml
global:
  scrape_interval: 15s

scrape_configs:
- job_name: node
  static_configs:
  - targets: ['localhost:9100']
```
> the job name is the descriptive name of the node: jenkins, db, web, etc

run

```bash
./prometheus --config.file=./prometheus.yml
```