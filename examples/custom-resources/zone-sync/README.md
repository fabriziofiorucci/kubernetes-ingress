# Zone-Sync

In this example, we configure the `zone-sync` feature. The feature is available in NGINX Plus.

## Configure NGINX Plus Zone Synchronization and Resolver without TLS

In this step we configure:

- [Zone Synchronization](https://docs.nginx.com/nginx/admin-guide/high-availability/zone_sync/).

Steps:

1. Apply the ConfigMap `nginx-config.yaml`, which contains `zone-sync` data that enables zone synchronization.

    ```console
    kubectl apply -f nginx-config.yaml
    ```

If we don't specify `zone-sync-port` the port number is set to the default value `12345`.
