# KuberNeXT Apps

Apps listed below are ready for deployment, all you need to do is executing
`[app]/deploy/k8s/overlays/[env]` with `kubectl apply -k` or registering them to
ArgoCD by referring [this repository](../) with
`[app]/deploy/k8s/overlays/[env]` path setting.

| File          | Description       | Link                |
| ------------- | ----------------- | ------------------- |
| `nginx-test/` | NGINX for testing | [Go](./nginx-test/) |
