# Manual Installation

If you want to install metering without OLM, using what's currently in master, first start by cloning the repo.
Next, decide which namespace you want to install Metering into, and set the `METERING_NAMESPACE` environment variable to the namespace you want to use.
By default, if it's unset, it will use the `metering` namespace.

## Install

Depending on your Kubernetes platform (regular Kubernetes, or Openshift)

For a standard Kubernetes cluster:

```
$ export METERING_NAMESPACE=metering-$USER
$ ./hack/install.sh
```

If your using Openshift, use openshift-install.sh:

```
$ export METERING_NAMESPACE=metering-$USER
$ ./hack/openshift-install.sh
```

## Uninstall

To uninstall the process is the same, pick the right uninstall script for your platform, and run it.

For a standard Kubernetes cluster:

```
$ export METERING_NAMESPACE=metering-$USER
$ ./hack/uninstall.sh
```

If your using Openshift, use openshift-uninstall.sh:

```
$ export METERING_NAMESPACE=metering-$USER
$ ./hack/openshift-uninstall.sh
```

## Customize installation


If you wish to customize the installation, such as to modify configuration
options, change the image tag or repository, then you can use a custom
`metering` resource. To start, copy the default metering resource to a
separate file that we can modify:

```
$ cp manifests/metering-config/default.yaml metering-custom.yaml
```

For developers, the most common change is modifying the image tag, config, and resource limits.
Take a look at the [common configuration docs](common-configuration.md) to get an
idea of what you can modify that relates to configuration and resource limits, and
`manifests/metering-config/latest-versions.yaml` to see how to change the
image tag of each component.

```
$ export METERING_NAMESPACE=metering-$USER
$ export METERING_CR_FILE=metering-custom.yaml
```

Then run the installation script for your platform:

- `./hack/install.sh`
- `./hack/openshift-install.sh`

For more details on configuration options, most are documented in the [configuring metering document][configuring-metering].

## Run reporting operator locally

It's also possible to run the operator locally.
To simplify this, we've got a few `Makefile` targets to handle the building and running of the operator.

First, we still need to run Presto, Hive, and HDFS in the cluster, and also set reporting-operator replicas to 0 so that our local operator can obtain the leader election lease when we start it.

To do this, update your `metering-custom.yaml` to set `spec.reporting-operator.replicas` to `0` like so:

```
spec:
  reporting-operator:
    replicas: 0
```

Next, run the install script for your platform (see above).

After running the install script, figure out where your Prometheus pod is running.
By default the `run-reporting-operator-local` Makefile target assumes that the pod is in the `openshift-monitoring` namespace and can be queried using the label selector `app=prometheus`.

If you're Prometheus is located somewhere, else, you can override the defaults using the environment variables `METERING_PROMETHEUS_NAMESPACE` and `METERING_PROMTHEUS_LABEL_SELECTOR` to the namespace your Prometheus pod is in, and the label selector for querying Prometheus.

Ex (these are the defaults):
```
export METERING_PROMETHEUS_NAMESPACE=openshift-monitoring
export METERING_PROMTHEUS_LABEL_SELECTOR=app=prometheus
```

Finally, use the following command to build & run the operator:

```
make run-reporting-operator-local
```

The above command builds the operator for your local OS (by default it only builds for Linux), uses kubectl port-forward to make Prometheus, Presto, and Hive available locally for your operator to communicate with, and then starts the operator with configuration set to use these local port-forwards.
Lastly, the operator automatically uses your `$KUBECONFIG` to connect and authenticate to your cluster and perform Kubernetes API calls.

## Run metering operator locally

The metering operator is the top-level operator which deploys other components using helm charts.
It's possible to also run this locally so you can iterate on charts and test them with the metering-operator before they're built and pushed to Quay for CI.

To run it locally you need to have the following:

- A connection to a docker daemon.
- Your `$KUBECONFIG` environment variable must be set and accessible to your Docker daemon.
- Your `$METERING_NAMESPACE` environment variable must be set, and unless `$LOCAL_METERING_OPERATOR_RUN_INSTALL` to `true`, the namespace must already exist.

This will just build and run the metering-operator docker image, which will watch for `Metering` resources in the namespace specified by `$METERING_NAMESPACE`, using your `$KUBECONFIG` to communicate with the API server.

```
make run-metering-operator-local
```

If you want to also create the namespace using the install scripts you can do that manually and set `$SKIP_METERING_OPERATOR_DEPLOYMENT` to `true`:

```
export SKIP_METERING_OPERATOR_DEPLOYMENT=true
./hack/openshift-install.sh
```

Or you can set `$LOCAL_METERING_OPERATOR_RUN_INSTALL` to `true` and run the same make command as above:

```
export LOCAL_METERING_OPERATOR_RUN_INSTALL=true
make run-metering-operator-local
```

[configuring-metering]: metering-config.md
