
# Using Kubernetes Client-go 

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

Create a new  pod in the namespace 
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

or simply use the command ```kubectl create -f busybox.yml```

A the same time, the watcher sees the new pod and print a log statement   

```
go run main.go -n demo
Starting a Pod watcher in namespace [demo]
2019/08/17 12:43:55  Pod busybox-pod added 
````

Delete the created pod 
````
kubectl delete pod busybox-pod -n demo 
`````

In the watcher logs you can see a log statement like 
````
Starting a Pod watcher in namespace [demo]
.....

2019/08/17 13:11:15  Pod busybox-pod deleted
````
