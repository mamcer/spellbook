# kubectl

## setup

```bash
alias k=kubectl
source <(kubectl completion bash)
complete -F __start_kubectl k
```

```bash
# K3s kubeconfig, copied to the home directory to avoid sudo
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown $(id -u):$(id -g) ~/.kube/config
```

```bash
kubectl config set-context --current --namespace=fury-dev
kubectl config view --minify | grep namespace
```
> set a default namespace to skip `-n` on every command

## get

```bash
kubectl get pods -A
kubectl get pods -A -o wide
kubectl get all -n [namespace]
kubectl get svc -A
kubectl get ingress -A
kubectl get pvc -n [namespace]
kubectl get storageclass
kubectl get events -n [namespace] --sort-by='.lastTimestamp'
```
> `-o wide` adds IP, node and restart count

```bash
kubectl top pods -n [namespace]
kubectl top nodes
```
> needs metrics-server, enabled by default on K3s

## daily

```bash
kubectl get pods -n [namespace]
kubectl describe pod [pod] -n [namespace]
kubectl logs -f [pod] -n [namespace]
kubectl exec -it [pod] -n [namespace] -- sh
kubectl apply -f [file].yaml
kubectl rollout restart deployment/[name] -n [namespace]
kubectl port-forward -n [namespace] svc/[service] [local-port]:[remote-port]
```

## logs

```bash
kubectl logs [pod] -n [namespace] -c [container]
kubectl logs [pod] -n [namespace] --previous
kubectl logs -n [namespace] deploy/[name] -f
```
> `-c` for pods with several containers (e.g. the Vault sidecar), `--previous` for the attempt before the last restart

## debug a pod that does not start

```bash
kubectl get pods -n [namespace]
kubectl describe pod [pod] -n [namespace]
kubectl logs [pod] -n [namespace] --previous
kubectl get events -n [namespace] --sort-by='.lastTimestamp'
```
> in that order; in `describe`, read the Events section at the bottom

| status | usual cause |
| --- | --- |
| `Pending` | no node with enough resources, or the PVC cannot bind |
| `ImagePullBackOff` | wrong image name or tag, or no access to the registry |
| `CrashLoopBackOff` | the process crashes on start, check the logs |
| `OOMKilled` | memory limit exceeded, check `resources.limits` |

## apply

```bash
kubectl diff -f deploy.yaml
kubectl apply -f deploy.yaml --dry-run=client -o yaml
kubectl apply -f manifests/
```
> `diff` shows what would change before applying

```bash
kubectl get deployment [name] -n [namespace] -o yaml
```
> the manifest as it is actually running

```bash
kubectl edit configmap [name] -n [namespace]
kubectl rollout restart deployment/[name] -n [namespace]
kubectl rollout status deployment/[name] -n [namespace]
```
> pods do not reload a ConfigMap on their own, restart the deployment

## vault

```bash
kubectl exec -it -n fury-infra vault-0 -- vault status
```
> Vault comes back sealed after every pod restart; first thing to check when something that depends on it starts failing

## traefik

```bash
kubectl get pods -n kube-system | grep traefik
kubectl logs -n kube-system deploy/traefik -f
kubectl get ingressroute -A
```
> on K3s Traefik runs in `kube-system`

## tools

- [k9s](https://k9scli.io/): TUI to browse pods, logs and exec
- [kubectx / kubens](https://github.com/ahmetb/kubectx): switch context and namespace
- [stern](https://github.com/stern/stern): tail logs from several pods at once
