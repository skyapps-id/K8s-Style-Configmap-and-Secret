# Style Configmap and Secret in Kubernetes

🚢[Docker Image](https://hub.docker.com/_/mysql)

### Running On Premise 
- Ubuntu 20.04 LTS
- Storage Class Rook Ceph Block

### Specification My Cluster
- Master node vCPU 2 Memory 8Gb
- Worker node1 vCPU 2 Memory 4Gb
- Worker node2 vCPU 2 Memory 4Gb
- Worker node3 vCPU 2 Memory 4Gb

### Installation
1. Clone this project.
    ```sh
    $ git clone https://github.com/skyapps-id/K8s-Style-Configmap-and-Secret.git 
    ```

2. Environment by configMapKeyRef and secretKeyRef.
    ```sh
    $ cd  K8s-Style-Configmap-and-Secret/

    $ kubectl apply -f configmap-and-secret-style1.yaml
    configmap/style1-configmap created
    secret/style1-secret created
    ```
    Testing environment on project Go
    ```sh
    $ kubectl apply -f deployment-style1.yaml
    deployment.apps/golang-test created

    $ kubectl get pods
    NAME                           READY   STATUS    RESTARTS   AGE
    golang-test-59c66dccf9-q6dq7   1/1     Running   0          62s
    golang-test-59c66dccf9-st2l8   1/1     Running   0          62s
    golang-test-59c66dccf9-zkswh   1/1     Running   0          62s

    $ kubectl logs golang-test-798fbf86cf-nf48p
    Database Host :  localhost_style1
    Database Name :  fullstack_api_style1
    Database User :  username_style1
    Database Password :  password_style1
    Listening to port 8080
    ```

3. Environment by configMapRef.
    ```sh
    $ cd  K8s-Style-Configmap-and-Secret/

    $ kubectl apply -f configmap-and-secret-style2.yaml
    configmap/style2-configmap created
    ```
    Testing environment on project Go
    ```sh
    $ kubectl apply -f deployment-style2.yaml
    deployment.apps/golang-test created

    $ kubectl get pods
    NAME                           READY   STATUS    RESTARTS   AGE
    golang-test-798fbf86cf-bt6d6   1/1     Running   0          25s
    golang-test-798fbf86cf-j25r5   1/1     Running   0          25s
    golang-test-798fbf86cf-nf48p   1/1     Running   0          25s

    $ kubectl logs golang-test-798fbf86cf-nf48p
    Database Host :  localhost_style2
    Database Name :  fullstack_api_style2
    Database User :  username_style2
    Database Password :  password_style2
    Listening to port 8080
    ```
4. Environment by mountVolume.
   
   Next Chapter ....

### Licence

This work is under [MIT](LICENCE) licence.