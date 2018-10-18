## Probe Demo

Health check with a /healthz endpoint. If probe failed on the endpoint, kube will kill the container and restart again

##

  * Deployment

    ```
    kubectl apply -f probe.yml
    ```

    ```
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
      initialDelaySeconds: 3
      periodSeconds: 3
    ```
  * Kill a container

    ```
    kubectl get service -o wide
    NAME            TYPE           CLUSTER-IP      EXTERNAL-IP                    PORT(S)        AGE       SELECTOR
prob-external   LoadBalancer   10.100.200.50   10.193.148.199,100.64.224.47   80:30372/TCP   20m       app=probe
    ```

    Set health as false.... will kill one of the containers
    ```
    curl -X PUT --data "{\"health\": false}" http://10.193.148.199/healthz
    ```

  * Watch the pods

    There are a few restarts.....

     ```
    kubectl get pods
    NAME                     READY     STATUS    RESTARTS   AGE
    probe-7f9845f755-9xnb4   1/1       Running   2          13m
    probe-7f9845f755-n85p7   1/1       Running   1          13m
    probe-7f9845f755-rss7l   1/1       Running   2          13m
     ```
