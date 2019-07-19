# spinnaker-local

## Settings
### Install Halyard
```bash
$ curl -O https://raw.githubusercontent.com/spinnaker/halyard/master/install/macos/InstallHalyard.sh
```

```bash
$ sudo bash ./InstallHalyard.sh
Password:
Please supply a non-root user to run Halyard as: 
```

```bash
$ hal -v
1.21.2-20190701132842
```

### Create Namespace
```bash
$ kubectl config set-context docker-for-desktop/spinnaker --cluster=docker-for-desktop-cluster --user=docker-for-desktop --namespace=spinnaker
Context "docker-for-desktop/spinnaker" created.

$ kubectx kubectx docker-for-desktop/spinnaker
Switched to context "docker-for-desktop/spinnaker".
```

### Install using Helm
```bash
$ helm install --name my-spinnaker stable/spinnaker --timeout 600

NOTES:
1. You will need to create 2 port forwarding tunnels in order to access the Spinnaker UI:
  export DECK_POD=$(kubectl get pods --namespace spinnaker -l "cluster=spin-deck" -o jsonpath="{.items[0].metadata.name}")
  kubectl port-forward --namespace spinnaker $DECK_POD 9000

2. Visit the Spinnaker UI by opening your browser to: http://127.0.0.1:9000

To customize your Spinnaker installation. Create a shell in your Halyard pod:

  kubectl exec --namespace spinnaker -it my-spinnaker-spinnaker-halyard-0 bash

For more info on using Halyard to customize your installation, visit:
  https://www.spinnaker.io/reference/halyard/

For more info on the Kubernetes integration for Spinnaker, visit:
  https://www.spinnaker.io/reference/providers/kubernetes-v2/
```

### Access WebUI
```bash
$ export DECK_POD=$(kubectl get pods --namespace spinnaker -l "cluster=spin-deck" -o jsonpath="{.items[0].metadata.name}")
$ kubectl port-forward $DECK_POD 9000
```

Visit the Spinnaker UI => http://localhost:9000

