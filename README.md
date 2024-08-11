# go-example-shell-operator

An example of a hook in Golang for [shell-operator](https://github.com/flant/shell-operator).

This example is a rewrite of [`monitor-pods`](https://github.com/flant/shell-operator/tree/main/examples/101-monitor-pods) Bash example. Unlike Bash, Golang must first be compiled in order to act as a hook.

You can use the following steps to run the example:

1. Start a local cluster, e.g. with [`kind`](https://kind.sigs.k8s.io/docs/user/quick-start/).

2. Build a Docker image containing the Golang hook and load it for Kubernetes nodes:
    ```bash
    docker build -t shell-operator:monitor-pods . && kind load docker-image shell-operator:monitor-pods
    ```

3. Deploy the newly build image as a pod, together with other Kubernetes resources:
    ```bash
    kubectl apply -f example.yaml
    ```

4. In order to trigger the hook, try to deploy some other pod or deployment, for example:
    ```bash
    kubectl apply -f https://k8s.io/examples/application/deployment.yaml
    ```

5. See the logs of shell-operator pod.


To clean up:
```bash
kubectl delete  deployments.apps nginx-deployment
kubectl delete -f example.yaml
```