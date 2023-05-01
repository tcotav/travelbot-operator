# travelbot-operator

The Travelbot Operator is a Kubernetes operator that manages the deployment of the Travelbot application inside a Kubernetes cluster.

## Overview

The Travelbot application is a simple demo program that generates a map of 3-5 stops as x,y coordinates and then travels along the path to the end. At the end, it turns around and follows the path back. The Travelbot Operator manages the state of the Travelbot deployment inside the cluster, including creating and updating deployments, managing replicas, and monitoring the health of the application.

## Operator Levels

The Travelbot Operator starts as a level I operator and will progress along the levels of the operator capacity as new features and capabilities are added. The following levels are planned:

* Level I: Basic Install 
* Level II: Seamless Upgrades
* Level III: Full Lifecycle - app lifecycle, storage lifecycle (backups, failure recovery)
* Level IV: Deep Insights - metrics, alerts, logs processing, workload analysis 
* Level V: Auto Pilot - horizontal/vertical scaling, auto config tuning, abnormality detection, scheduling tuning 

### Level I and Level II Features 

To hit level I, we've added parameters to our custom resource to control the deployment version and number of replicas (in addition to some internal variables):

    - replicas - number of replicas current deployment should have running
    - deployImage - container image for the pods 

Change `deployImage`, deployment rolls out new pods.


## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/travelbot-operator:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/travelbot-operator:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

