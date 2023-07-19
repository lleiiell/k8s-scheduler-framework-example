# k8s-scheduler-framework-example

This repo is an example for Kubernetes scheduler framework.

## Deploy

```bash
kubectl apply -f deploy/example-scheduler.yaml
```

## Test

### good case

```bash
kubectl apply -f deploy/test-scheduler.yaml
```

Then watch example-scheduler pod logs.

### bad case


```bash
# bad case
kubectl apply -f deploy/test-scheduler-err.yaml
```

Then describe example-scheduler pod.

```log
Events:
  Type     Reason            Age                From               Message
  ----     ------            ----               ----               -------
  Warning  FailedScheduling  20m                example-scheduler  running PreBind plugin "example-sched-plugin": only pods from 'default' namespace are allowed
  Warning  FailedScheduling  19m (x1 over 20m)  example-scheduler  running PreBind plugin "example-sched-plugin": only pods from 'default' namespace are allowed

```

## Issues

### unknown revision v0.0.0

```bash
# k8s version is 1.20.15
bash getK8sMods.sh 1.20.15
# VERSION: 1.20.15
# MODs: k8s.io/api
# MOD: k8s.io/api@kubernetes-1.20.15
# MOD: k8s.io/apiextensions-apiserver@kubernetes-1.20.15
# MOD: k8s.io/apimachinery@kubernetes-1.20.15
# MOD: k8s.io/apiserver@kubernetes-1.20.15
# MOD: k8s.io/cli-runtime@kubernetes-1.20.15
# MOD: k8s.io/client-go@kubernetes-1.20.15
```