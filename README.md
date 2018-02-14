### Example implementation of OpenTracing using InfluxDB
This is a small, simple implementation of an OpenTracing library
and storing traces into InfluxDB.

This stores each span of a trace as a unique series within InfluxDB.
This needs to be InfluxDB 1.4 or greater with TSI.

#### Schema
Each span of a trace is recorded as a unique series.

Here is an example line protocol representation of a trace with id 1;
each line represents a span:

```
traces,trace_id=1,parent_id=1,id=1,name=tier1 duration_ns=120357159i 1518581993555323961
traces,trace_id=1,parent_id=1,id=2,name=tier2 duration_ns=115352090i 1518581993558045808
traces,trace_id=1,parent_id=2,id=3,name=tier3 duration_ns=82872204i 1518581993560018841
```

Notice that the `parent_id` is the `id` of the parent span. In the case of a first span in the
trace, the `parent_id` is the same as the `id` of the span.

The `duration_ns` is the duration of the span in nanoseconds

#### Example
The example is tracing a 3-tier service.

To run the examples:

```sh
influx -execute "create database influxdays with duration 2d shard duration 1h"
go run tier3/main.go &
go run tier2/main.go &
go run tier1/main.go
```

The `tier1` program sends three requests to `tier2`.  `tier2` waits 20ms and
sends a request to `tier3`.  Finally, `tier3` waits between 1ms and 1s to respond.

The data is recorded into InfluxDB under the `influxdays` database in the `traces`
measurement.

#### Queries
To find all traces generated from `tier1` in the last 3 hours:

```
SHOW TAG VALUES FROM "traces" with key="trace_id" WHERE "name" = 'tier1' and time > now() - 3h`
```

To find the longest trace duration generated from `tier1` in the last 20 minutes:

```
SELECT max("duration_ns"),trace_id,"name" FROM "traces" WHERE "name" = 'tier1' AND time > now() - 20m GROUP BY time(5m) fill(none)
```

To find all spans from a single trace with id:

```
SELECT * FROM "traces" WHERE "trace_id" = 'ba1rhq8ifaqd6nsgvvq0'
```

Count the number of spans generated by each tier in the last 15 minutes:

```
SELECT count("duration_ns") FROM traces WHERE time > now() - 15m GROUP BY "name", time(15m) fill(none)
```

See all the dependencies between services in the last hour:

```
SHOW SERIES FROM traces where time > now() - 1h
```
