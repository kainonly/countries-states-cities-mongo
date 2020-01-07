#!/usr/bin/python3
import subprocess

images = {
    'k8s.gcr.io/pause:3.1': 'kainonly/pause:3.1',
    'k8s.gcr.io/etcd:3.3.10': 'kainonly/etcd:3.3.10',
    'k8s.gcr.io/coredns:1.3.1': 'kainonly/coredns:1.3.1',
    'k8s.gcr.io/kube-apiserver:v1.15.3': 'kainonly/kube-apiserver:v1.15.3',
    'k8s.gcr.io/kube-controller-manager:v1.15.3': 'kainonly/kube-controller-manager:v1.15.3',
    'k8s.gcr.io/kube-proxy:v1.15.3': 'kainonly/kube-proxy:v1.15.3',
    'k8s.gcr.io/kube-scheduler:v1.15.3': 'kainonly/kube-scheduler:v1.15.3'
}

for default, mirror in images.items():
    subprocess.run(['docker', 'pull', mirror])
    subprocess.run(['docker', 'tag', mirror, default])
    subprocess.run(['docker', 'rmi', mirror])
