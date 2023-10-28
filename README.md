# Local Backend Service Development on Kubernetes using mirrord (Sample Projects)

## Prerequisites

- Kubernetes local cluster (use: k3d, minikube etc.)
- Skaffold to easily deploy the sample app in local kubernetes, [Installing Skaffold](https://skaffold.dev/docs/install/)
- mirrord as a tool to develop service locally as if it is inside kubernetes cluster, [Introduction](https://mirrord.dev/docs/overview/introduction/)

### Setup

This section we're going to setup our sample cluster the scenario is that we already have a running
service that we're developing.

To make it easy use skaffold so that we don't have to publish any container image
in order to run it on cluster.

Deploy this project to local kubernetes using skaffold

```bash
skaffold run
```

You can observe that if successful a running pod will appear in cluster

### Running Service Locally

In this case pretend that you have uncommitted code changes and want to test it
on the cluster.

Next, we're going to run this service locally using `mirrord`, the target pod
name can be copied from the cluster

```bash
mirrord exec --target pod/my-server-74cd947f55-lpqdq go run .
```

If there's no error, http port `8080` will be accessible and you can hit the
endpoint `localhost:8080/api/get_pods` to test

```bash
curl localhost:8080/api/get_pods | json_pp
```

The endpoint serve by this sample app will return a list of pods inside the cluster,
this shows that the local service is actually running as if it is inside the cluster

API response something like:

```json
{
   "message" : "get pods successful, there are 9 pods",
   "pods" : [
      ...
      {
         "name" : "my-server-74cd947f55-lpqdq",
         "namespace" : "default"
      },
      {
         "name" : "mirrord-agent-zkjdyl9ffi-f8p47",
         "namespace" : "default"
      },
      ...
   ]
}
```

If you look closer to how mirrord works, it actually spawn a kubernetes job 
and run a mirrord agents pod to communicate with the target pod.