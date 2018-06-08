# Deploying and Benchmarking Open Source Serverless Frameworks

Since the advent of [AWS Lambda](https://aws.amazon.com/lambda/) in 2014, the function-as-a-service programming paradigm has gained a lot of traction in the cloud community. At first, only large cloud providers such as AWS, GKE or Azure provided such services with a pay-per-invocation model, but since then interest has increased for developers and entreprises to build their own solutions on an open source model.

The rise of container orchestrators such as Kubernetes, Docker Swarm or DC/OS has made this process even easier, resulting in a number of competing frameworks in this space. These frameworks vary a lot in feature set, but all have the feature we come know to expect from a FaaS offering in the style of AWS Lambda:

 - The ability to create and modify language-specific scripts with a simple interface in either "inline" (within browser) or packaged form.
 - The ability to invoke functions through an HTTP API.

Finally, all of these frameworks can easily be deploying using a Kubernetes distribution.

In this project, we provide common ways to both **deploy** and **benchmark** these open source serverless frameworks. The scope of the benchmarking is still in progress.

### Deploying Kubernetes Clusters

In order to make benchmarking reproduceable and verifiable, we provide instructions on deploying clusters on AWS in the [deploy-cluster](https://github.com/docker/serverlessbenchmark/tree/master/deploy-cluster) folder. Clusters are configured to have:

 - 1 kubernetes manager using an M4.Xlarge instance
 - 2 kubernetes workers using a T2.large instance
 - Each machine runs the ubuntu-16.04 operating system

The instructions may vary according to kubernetes distribution. The only instructions currently available are for Docker EE.

### Deploying Serverless Frameworks

In the [deploy-framework folder](https://github.com/docker/serverlessbenchmark/tree/master/deploy-framework), we provide instructions to deploy each framework for a specific kubernetes distribution. The only instructions available are for Docker EE.

### Deploying Functions

TBD

### Benchmarking Serverless Frameworks

TBD

### Open Source Serverless Frameworks

Here is the list of the supported open source serverless frameworks.

 - [OpenFaaS](https://github.com/openfaas)
 - [Nuclio](https://github.com/nuclio/nuclio)
 - [Gestalt](http://www.galacticfog.com/product/)
 - [Riff](https://github.com/projectriff/riff)
 - [Fn](https://github.com/fnproject/fn)
 - [OpenWhisk](https://openwhisk.apache.org/)

### Kubernetes Distributions

For now, the only vendor that is suppported is Docker Entreprise Edition.

### Notes

 - The benchmarking component of this repository has not been released yet.
