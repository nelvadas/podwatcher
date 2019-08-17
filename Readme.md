
# Using Kubernetes go-client 

This project shows how to build a custom Pod watcher using kubernetes go-client.
the controller watchs pod creation and deletion in a specific namespace and display a log statement accordingly

````
go run main.go -n demo
Starting a Pod watcher in namespace [demo]
````


Create the demo namespace 
````
kubectl create ns demo
````

Add ad pod in the namespace 
````
$ cat <<EOF  |  kubectl apply -f -
> apiVersion: v1
> kind: Pod
> metadata:
>   name: busybox-pod
>   namespace: demo
> spec:
>   containers:
>   - image: busybox
>     command:
>       - sleep
>       - "3600"
>     imagePullPolicy: IfNotPresent
>     name: busybox-cnt
>   restartPolicy: Always
> EOF
pod/busybox-pod created
````

A the same time, the watcher sees the new pod and print a notification  

go run main.go -n demo
Starting a Pod watcher in namespace [demo]
2019/08/17 12:43:55  Pod busybox-pod added 
````

