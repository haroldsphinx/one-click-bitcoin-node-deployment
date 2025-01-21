# Prometheus

https://github.com/prometheus-community/helm-charts/tree/main/charts/prometheus

Deploys a prometheus scraper with default scrape configs.

The deployment depends on a kubernetes basic auth secret named `bitcoin-node-exporter-secrets`

## Configuration

Some configuration values must be hardcoded to `values.yaml` before the initial deployment.

Update `values.yaml` setting the `remoteWrite` `url` and `username`:

```yaml
remoteWrite:
  - url: https://<your-prom-url>/api/prom/push
    basic_auth:
      username: <your-numeric-username>
```

For environment-specific changes update `./[development|staging|production]/values.yaml`.

## How to apply changes

Export an `ENVIRONMENT` environment variable and call one of the make targets:

```sh
$ make deploy ENVIRONMENT=development
$ make delete ENVIRONMENT=development
```
