#!/bin/sh

# shellcheck disable=SC2039
declare -A images=(
  ['k8s.gcr.io/pause:3.1']='kainonly/pause:3.1'
  ['k8s.gcr.io/etcd:3.3.10']='kainonly/etcd:3.3.10'
  ['k8s.gcr.io/coredns:1.3.1']='kainonly/coredns:1.3.1'
  ['k8s.gcr.io/kube-apiserver:v1.15.3']='kainonly/kube-apiserver:v1.15.3'
  ['k8s.gcr.io/kube-controller-manager:v1.15.3']='kainonly/kube-controller-manager:v1.15.3'
  ['k8s.gcr.io/kube-proxy:v1.15.3']='kainonly/kube-proxy:v1.15.3'
  ['k8s.gcr.io/kube-scheduler:v1.15.3']='kainonly/kube-scheduler:v1.15.3'
)

for origin in ${!images[*]}; do
  docker pull ${images[$origin]}
  docker tag ${images[$origin]} ${origin}
  docker rmi ${images[$origin]}
done
