# Nuclio

**In this file, we detail how to deploy OpenFaaS for the serverless benchmark on Docker EE.**

## Requirements

 - A Docker EE cluster with admin credentials (see [here]() for how to install one)
 - A configured [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) (follow [this guide](https://docs.docker.com/ee/ucp/user-access/kubectl/) to configure it on your cluster)
 - The [nuclio cli tool `nuctl`](https://github.com/nuclio/nuclio/releases) 

## Hello World Example

```
nuctl deploy helloworld -n nuclio -v -p https://raw.githubusercontent.com/nuclio/nuclio/master/hack/examples/golang/helloworld/helloworld.go --registry $username
nuctl invoke -n nuclio helloworld
```


