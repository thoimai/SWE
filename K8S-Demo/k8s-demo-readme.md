# the process of setting up a local Kubernetes cluster using Minikube and deploying a web application with MongoDB.

Step 1: Install Minikube


Step 2: Start Minikube


Step 3: Verify Minikube is running


Step 4: Install Helm


Step 5: Deploy MongoDB using Helm


Step 6: Create a simple web application


Step 7: Create a Kubernetes deployment for your application


Step 8: Expose your application using a NodePort service

Since Minikube does not support the LoadBalancer service type, you can use a NodePort service to expose your application


Step 9: Access your application

##############



create 4 k8s config files to deploy simple web application with mongodb using kubernetes 


1. configmap.yaml  > mongodb endpoint 

2. secret.yaml > mongodb username and password
 
3. service + deployment > mongodb application with internal service 

4. service + deployment > our own application with external service

+ how to make it an external service
    > change the service type to NodePort

###################
1. Mongo configmap.yaml
###################




# To deploy these files, run the following commands:



kubectl apply -f mongodb-secret.yaml
kubectl apply -f mongodb-deployment.yaml
kubectl apply -f webapp-deployment.yaml

Once the deployment is complete, you can access your web application using the Minikube IP address and the specified node port (e.g., http://<minikube_ip>:30080). To get the Minikube IP address, run:

$ minikube ip

$ minikube start
$ minikube status
$ kubectl get pod
$ kubectl apply -f mongo-config.yaml
$ kubectl apply -f mongo-secret.yaml
