# Using Metering

This document demonstrating using Metering to create reports and fetch their results.
If you want a guide on how to extend Metering with custom queries and reports, read the [writing custom queries guide][writing-custom-queries]

*Note: Wait at least 15 minutes after installing Metering before generating reports to enable Metering to reach a consistent state and fetch some data from Prometheus*

## Writing a report

First, read the [Reports][report-md] guide for a list of available options.

Select a `ReportGenerationQuery` and a `reportingStart` and `reportingEnd`.
Use 'kubectl' to query the Metering namespace for a list of available  `ReportGenerationQueries`:

```
kubectl get reportgenerationqueries -n $METERING_NAMESPACE
```

Each ReportGenerationQuery is designed to report on a specific resource, usually a `pod`, `namespace` or `node`, and on a specific metric, like `cpu` or `memory`, on a specific resource. Some reports correlate several of these metrics in a single report. See the [Reports][report-md] guide for more information on the returns provided by each report query.

## Creating a report

A report can be created for Metering to run using `kubectl`.
The report should be created in the same namespace as Metering is installed.

First, create an example report. Save the following into a file called `report.yaml` (times are UTC):

```
apiVersion: metering.openshift.io/v1alpha1
kind: Report
metadata:
  name: namespace-cpu-request
spec:
  reportingStart: '2018-01-01T00:00:00Z'
  reportingEnd: '2018-12-30T23:59:59Z'
  generationQuery: "namespace-cpu-request"
  runImmediately: true
```

Once the report YAML is written, use `kubectl` to create the report:

```
$ kubectl -n $METERING_NAMESPACE create -f report.yaml
```

Existing reports can be viewed in Kubernetes with the following command:

```
$ kubectl -n $METERING_NAMESPACE get reports
```

A report's status can be inspected by viewing the object with the `-o json`
flag:

```
$ kubectl -n $METERING_NAMESPACE get report namespace-cpu-request -o json
```

## Viewing reports

Once a report's status has changed to `Finished`, the report is ready to be
downloaded. The Metering Pod exposes an HTTP API for this.

First, use `kubectl` to set up a proxy for accessing Kubernetes services:

```
$ kubectl proxy
```

The URL used to fetch a report changes based on the report's name and format.
The `format` parameter may be either `csv` or `json`. The URL scheme is:

```
/api/v1/reports/get?name=[Report Name]&format=[Format]
```

Using `kubectl proxy` requires that the URL be accessed through a prefix that
points to the Kubernetes service. (See the upstream documentation on
[Manually constructing apiserver proxy URLs][accessing-services] for more details.) The following example assumes Metering is deployed in the `metering` namespace.

```
http://127.0.0.1:8001/api/v1/namespaces/metering/services/http:reporting-operator:http/proxy/api/v1/reports/get?name=[Report Name]&format=[Format]
```

If your using Openshift, you'll need to change to the following, which uses the `openshift-metering` namespace, and which uses HTTPS by default:

```
http://127.0.0.1:8001/api/v1/namespaces/metering/services/https:reporting-operator:http/proxy/api/v1/reports/get?name=[Report Name]&format=[Format]
```

For example, the results of a report with the name `namespace-cpu-request` report can be fetched in
CSV, with the following command:

```
$ curl "http://127.0.0.1:8001/api/v1/namespaces/metering/services/http:reporting-operator:http/proxy/api/v1/reports/get?name=namespace-cpu-request&format=csv"
```


[accessing-services]: https://kubernetes.io/docs/tasks/administer-cluster/access-cluster-services/#manually-constructing-apiserver-proxy-urls
[report-md]: report.md
[writing-custom-queries]: writing-custom-queries.md
