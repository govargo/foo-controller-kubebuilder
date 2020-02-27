# Sample Controller(Foo Controller)

This Controller is developed by Kubebuilder v2.2.0

インプレスR&D NextPublishing「実践入門 Kubernetesカスタムコントローラーへの道」の第七章「Admission Webhook」に掲載したサンプルコード用のリポジトリです。  

https://nextpublishing.jp/book/11389.html

第五章で実装したサンプルは[master](https://github.com/govargo/foo-controller-kubebuilder)に、    
Conversion Webhookのサンプルは[multi_conversion](https://github.com/govargo/foo-controller-kubebuilder/tree/multi_conversion)に用意しています。

## What is this Controller

This Controller reconcile foo custom resource.   
Foo resource owns deployment object.

We define foo.spec.deploymentName & foo.spec.replicas.   
When we apply foo, then deployment owned by foo is created.

If we delete deployment object like 'kubectl delete deployment/\<deployment name\>',    
foo reconcile it and the deployment is created again.

Or if we scale deployment object like 'kubectl scale deployment/\<deployment name\> --replicas 0',       
foo reconcile it and the deployment keeps the foo.spec.replicas.

## How to run

Before you start this operator, you must prepare the kubernetes cluster using like [Minikube](https://github.com/kubernetes/minikube), [Kind](https://github.com/kubernetes-sigs/kind).

We must prepare [cert-manager](https://github.com/jetstack/cert-manager) to deploy admission webhook.

```
$ kubectl create namespace cert-manager
$ kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.11.0/cert-manager.yaml

# Confirm the cert-manager pod's running
$ kubectl get pods -n cert-manager
```

Clone source code

```
$ mkdir -p $GOPATH/src/github.com/govargo
$ cd $GOPATH/src/github.com/govargo
$ git clone https://github.com/govargo/foo-controller-kubebuilder.git
$ cd foo-controller-kubebuilder
```

Run localy

```
$ make install
$ make run
$ kubectl apply -f config/samples/samplecontroller_v1alpha1_foo.yaml
```

Run container as Deployment

```
$ make install
$ make deploy
$ kubectl apply -f config/samples/samplecontroller_v1alpha1_foo.yaml
```

## Reference

This project is inspired by

 * https://github.com/kubernetes/sample-controller
 * https://github.com/jetstack/kubebuilder-sample-controller
