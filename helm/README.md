# OBot Helm Chart

This Helm chart deploys the OBot application to a Kubernetes cluster.

## Prerequisites

- Kubernetes 1.19+
- Helm 3.0+

## Installation

### 1. Create Secret for TELE_TOKEN

Before installing the chart, you need to create a Kubernetes secret with your Telegram bot token:

```bash
kubectl create secret generic obot --from-literal=token=YOUR_TELEGRAM_BOT_TOKEN
```

Replace `YOUR_TELEGRAM_BOT_TOKEN` with your actual Telegram bot token.

### 2. Install the Chart

You can install the chart directly from the GitHub release:

```bash
helm install obot https://github.com/DenisZm/obot/releases/download/v1.1.1/obot-1.1.0.tgz
```

Or install with custom values:

```bash
helm install obot https://github.com/DenisZm/obot/releases/download/v1.1.1/obot-1.1.1.tgz \
  --set image.tag=v1.1.1 \
  --set replicaCount=2
```

### 3. Verify Installation

Check if the deployment is running:

```bash
kubectl get pods -l app.kubernetes.io/name=obot
kubectl get services -l app.kubernetes.io/name=obot
```

## Configuration

The following table lists the configurable parameters of the OBot chart and their default values:

| Parameter | Description | Default |
|-----------|-------------|---------|
| `replicaCount` | Number of replicas | `1` |
| `image.repository` | Container image repository | `ghcr.io/deniszm/obot` |
| `image.tag` | Container image tag | `v1.1.0` |
| `image.arch` | Container image architecture | `amd64` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `secret.name` | Name of the secret containing TELE_TOKEN | `obot` |
| `secret.tokenName` | Environment variable name for the token | `TELE_TOKEN` |
| `secret.tokenKey` | Key in the secret for the token | `token` |
| `securityContext.privileged` | Run container in privileged mode | `true` |

## Uninstalling the Chart

To uninstall/delete the `obot` deployment:

```bash
helm uninstall obot
```

Don't forget to also delete the secret:

```bash
kubectl delete secret obot
```

## Troubleshooting

### Check pod logs

```bash
kubectl logs -l app.kubernetes.io/name=obot
```

### Check pod status

```bash
kubectl describe pods -l app.kubernetes.io/name=obot
```

### Verify secret exists

```bash
kubectl get secret obot
```
