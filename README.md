# prometheus-loader
WIP
this tool will simulate the Grafana workloads for promethues DS
based on kubernetes-mixin assets.

# How it works
the following tool will use the kubernetes-mixin dashboard files, but it can be
used by any other Grafana dashboard file.
the app will scan the file and create a set of dashboards with dashboard name
and dashboard queries and than it will simulate the queries workload as long it runs

Note: this tool will fit only to backend performance simulation.

# How to
currently anything is adjustable, when using the mixin defaults just start the
tool.

defaults:
```python prometheus-loader.py -i 20 -p 60```

custom
```python prometheus-loader.py -f ./example.txt -i 20 -p 60 -t 20```

# Log
the tool will log the query perforamnce for each dashboard in the following format
2018-08-13 09:09:09.01 - duration: 0.5 - [nodes] - concurrency:18 - query: sum (rate (container_cpu_usage_seconds_total{id="/",instance=~"^.*"}[2m]))
