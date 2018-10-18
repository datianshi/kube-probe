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


     ```
     kubectl describe pod probe-7f9845f755-rss7l

     Normal   Scheduled  15m                default-scheduler                              Successfully assigned default/probe-7f9845f755-rss7l to 51ce836b-d21c-47ec-959b-48c0d5c7d4fe
  Normal   Pulling    10m (x3 over 15m)  kubelet, 51ce836b-d21c-47ec-959b-48c0d5c7d4fe  pulling image "datianshi/kube-probe"
  Warning  Unhealthy  10m (x6 over 11m)  kubelet, 51ce836b-d21c-47ec-959b-48c0d5c7d4fe  Liveness probe failed: HTTP probe failed with statuscode: 500
  Normal   Killing    10m (x2 over 11m)  kubelet, 51ce836b-d21c-47ec-959b-48c0d5c7d4fe  Killing container with id docker://probe:Container failed liveness probe.. Container will be killed and recreated.
  
     ```    
