import time
import docker
from console_progressbar import ProgressBar


class repository:
    repository = ''
    kubernetes = ''

    def __init__(self, repository, kubernetes):
        self.repository = repository
        self.kubernetes = kubernetes


lists = [
    repository('kainonly/pause:3.1', 'k8s.gcr.io/pause:3.1'),
    repository('kainonly/etcd:3.3.10', 'k8s.gcr.io/etcd:3.3.10'),
    repository('kainonly/coredns:1.3.1', 'k8s.gcr.io/coredns:1.3.1'),
    repository('kainonly/kube-apiserver:v1.15.3', 'k8s.gcr.io/kube-apiserver:v1.15.3'),
    repository('kainonly/kube-controller-manager:v1.15.3', 'k8s.gcr.io/kube-controller-manager:v1.15.3'),
    repository('kainonly/kube-proxy:v1.15.3', 'k8s.gcr.io/kube-proxy:v1.15.3'),
    repository('kainonly/kube-scheduler:v1.15.3', 'k8s.gcr.io/kube-scheduler:v1.15.3'),
]

client = docker.from_env()
pb = ProgressBar(total=100, prefix='Here', suffix='Now', decimals=3, length=50, fill='#', zfill='-')
for i, x in enumerate(lists):
    print('pull image:' + x.repository)
    image = client.images.pull(x.repository)
    print('success image:' + image.id)
    image.tag(x.kubernetes)
    client.images.remove(x.repository)
    pb.print_progress_bar((i + 1) / len(lists) * 100)

print('kubernetes images pull success.')
