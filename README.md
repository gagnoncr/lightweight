## Dependencies

 ```bash
 /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)" 
 ```


### Docker
> **Used these as ref:**
* [clean](https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes)
* [best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
* [builder](https://docs.docker.com/engine/reference/builder/#/cmd)
>
> **List Images:** 
* 1.`docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.ID}}"`
* 2.`docker ps --format "table {{.Image}}\t{{.Ports}} \t{{.Names}}"`
>
> **Helpful Commands**:
 * docker build -t [image-name] . 
 * docker run -p [external]:[interal] -it [image]
 * docker run -d -p [externalport]:[internalport] [image]:[tag]
 * docker run -it [microservice]
 * docker exec -it dsc-exam /bin/bash
 * docker system prune -a 
 
### Minikube
> **link to virtual box:** 
* `brew cask install virtualbox`
* `minikube config set vm-driver virtualbox`

> **Used these as ref:**
* [ingress minikube](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)
>
> **Helpful Commands**:
* minikube start --kubernetes-version=v1.15.0
* eval $(minikube docker-env)
* minikube dashboard
* minikube ssh
* minikube service api_micro --url
* sudo bash -o vi
* cd /var/log/cont
* ls | grep dsc-exam
* sudo head -n7 *.log
    
### Helm
> **Used these as ref:**
* [create your first chart](https://docs.bitnami.com/kubernetes/how-to/create-your-first-helm-chart/)
* [install](https://helm.sh/docs/install/)
* [api dep](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16/)
> **Install**:

        * $ brew install glide
        * $ cd $GOPATH
        * $ mkdir -p src/k8s.io
        * $ cd src/k8s.io
        * $ git clone https://github.com/helm/helm.git // the go mod in here is baf
        * $ go mod tidy
        * $ go clean -modcache
        * $ go mod init
        * $ make bootstrap build

> **Helpful Commands**:  
* helm init 
  * helm create [image]-chart
  * helm install --dry-run --debug ./[image]-chart
  * helm install --dry-run --debug ./[chart] --set service.internalPort=8080
  * helm install
>

# K8s
> **Used these as ref:**
* [cheat-sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
* [get started](https://docs.bitnami.com/kubernetes/get-started-kubernetes/)
* [vol share](https://kubernetes.io/docs/tasks/access-application-cluster/communicate-containers-same-pod-shared-volume/)
* [bare-metal ingress](https://kubernetes.github.io/ingress-nginx/deploy/baremetal/)
* [ingress](https://github.com/nginxinc/kubernetes-ingress)
* [nginx-ingress](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/)
* [Node_service-networking](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport)
* [ingress-controllers](https://kubernetes.io/docs/concepts/services-networking/ingress/#ingress-controllers)
>  
> **Helpful Commands**:    
* brew install kubectl 
* brew link --overwrite kubernetes-cli
* kubectl cluster-info   
* kubectl create -f [name]deploy.yaml (in both go && nginx && api kube dir) _ refactor names && terraform it.
* kubectl run deployment --image=dsc-exam-nginx:latest --port=8080 --label "app=dsc-exam-nginx"
* kubectl get pods -o=custom-columns=NAME:.metadata.name,IP:.status.podIP
* kubectl logs --follow [Pod name]
* kubectl logs -p [Pod name]
* kubectl get deployments --all-namespaces
* kubectl delete -n default deployment <deployname>
* kubectl create -f 
* kubectl get nodes
* kubectl get ep my-nginx
* kubectl exec [Pod] -- printenv | grep SERVICE
* kubectl get services kube-dns --namespace=kube-system
* kubectl run curl --image=radial/busyboxplus:curl -i --tty and nslookup
* kubectl exec -it dsc-exam -c nginx-container -- /bin/bash
 kubectl get svc --all-namespaces -o jsonpath='{range .items[?(@.spec.type=="LoadBalancer")]}{.metadata.name}:{.status.loadBalancer.ingress[0].ip}{"\n"}{end}'
>
> **Simple Deploy**:
* kubectl apply -f https://k8s.io/examples/application/shell-demo.yaml
 - 1.kubectl create -f nginx-deploy.yaml
 - 2.kubectl apply -f nginx-svc.yaml
 - 3.kubectl scale deployment [deployment] --replicas=0; kubectl scale deployment [deployment] --replicas=2;
 - 4.openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /d/tmp/nginx.key -out /d/tmp/nginx.crt -subj "/CN=my-nginx/O=my-nginx"
 - 5.cat /d/tmp/nginx.crt | base64  && cat /d/tmp/nginx.key | base64
 - 6.kubectl apply -f nginx-secret.yaml