# OpenFaaS - Serverless Functions Made Simple

## Requirements

 - A Docker EE cluster with admin credentials (see [here]() for how to install one)
 - A configured [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) (follow [this guide](https://docs.docker.com/ee/ucp/user-access/kubectl/) to configure it on your cluster)
 - The [faas-cli](https://github.com/openfaas/faas-cli)

### Golang (HTTP) Example

The Golang HTTP template is a runtime for openfaas geared towards speeding up function invocation latency.

```
mkdir simple-auth && cd simple-auth
faas-cli template pull https://github.com/openfaas-incubator/golang-http-template
faas-cli new --lang go simple-auth
```

Now edit the `simple-auth.yml` file as follows:

```
provider:
  name: faas
  gateway: http://$UCP_HOST:33000

functions:
  hello-world:
    lang: golang-http
    handler: ./simple-auth
    image: $DOCKER_HUB_USERNAME/simple-auth
```

where `$DOCKER_EE_IP` is your docker ee ip and `$DOCKER_HUB_USERNAME` is your Docker Hub username.

Finally, here is how to build, ship and deploy the function:

```
gateway=http://$UCP_HOST:33000
faas-cli login -u admin --password $password --gateway $gateway
faas-cli build -f simple-auth.yml
faas-cli push -f simple-auth.yml
faas-cli deploy -f simple-auth.yml --gateway $gateway
```

Finally, invoke your function:

```
faas-cli invoke simple-auth --gateway $gateway
```
