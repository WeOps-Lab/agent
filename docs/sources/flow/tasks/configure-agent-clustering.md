---
aliases:
- /docs/grafana-cloud/agent/flow/tasks/configure-agent-clustering/
- /docs/grafana-cloud/monitor-infrastructure/agent/flow/tasks/configure-agent-clustering/
- /docs/grafana-cloud/monitor-infrastructure/integrations/agent/flow/tasks/configure-agent-clustering/
- /docs/grafana-cloud/send-data/agent/flow/tasks/configure-agent-clustering/
# Previous page aliases for backwards compatibility:
- /docs/grafana-cloud/agent/flow/getting-started/configure-agent-clustering/
- /docs/grafana-cloud/monitor-infrastructure/agent/flow/getting-started/configure-agent-clustering/
- /docs/grafana-cloud/monitor-infrastructure/integrations/agent/flow/getting-started/configure-agent-clustering/
- /docs/grafana-cloud/send-data/agent/flow/getting-started/configure-agent-clustering/
- ../getting-started/configure-agent-clustering/ # /docs/agent/latest/flow/getting-started/configure-agent-clustering/
canonical: https://grafana.com/docs/agent/latest/flow/tasks/configure-agent-clustering/
description: Learn how to configure Grafana Agent clustering in an existing installation
menuTitle: Configure Grafana Agent clustering
title: Configure Grafana Agent clustering in an existing installation
weight: 400
refs:
  ui:
    - pattern: /docs/agent/
      destination: /docs/agent/<AGENT_VERSION>/flow/tasks/debug/#component-detail-page
    - pattern: /docs/grafana-cloud/
      destination: /docs/grafana-cloud/send-data/agent/flow/tasks/debug/#component-detail-page
  install-helm:
    - pattern: /docs/agent/
      destination: /docs/agent/<AGENT_VERSION>/flow/get-started/install/kubernetes/
    - pattern: /docs/grafana-cloud/
      destination: /docs/grafana-cloud/send-data/agent/flow/get-started/install/kubernetes/
  clustering:
    - pattern: /docs/agent/
      destination: /docs/agent/<AGENT_VERSION>/flow/concepts/clustering/
    - pattern: /docs/grafana-cloud/
      destination: /docs/grafana-cloud/send-data/agent/flow/concepts/clustering/
---

# Configure {{% param "PRODUCT_NAME" %}} clustering in an existing installation

You can configure {{< param "PRODUCT_NAME" >}} to run with [clustering](ref:clustering) so that individual {{< param "PRODUCT_ROOT_NAME" >}}s can work together for workload distribution and high availability.

This topic describes how to add clustering to an existing installation.

## Configure {{% param "PRODUCT_NAME" %}} clustering with Helm Chart

This section guides you through enabling clustering when {{< param "PRODUCT_NAME" >}} is installed on Kubernetes using the {{< param "PRODUCT_ROOT_NAME" >}} [Helm chart](ref:install-helm).


### Steps

To configure clustering:

1. Amend your existing `values.yaml` file to add `clustering.enabled=true` inside the `agent` block.

   ```yaml
   agent:
     clustering:
       enabled: true
   ```

1. Upgrade your installation to use the new `values.yaml` file:

   ```bash
   helm upgrade <RELEASE_NAME> -f values.yaml
   ```

   Replace the following:

   - _`<RELEASE_NAME>`_: The name of the installation you chose when you installed the Helm chart.

1. Use the {{< param "PRODUCT_NAME" >}} [UI](ref:ui) to verify the cluster status:

   1. Click **Clustering** in the navigation bar.

   1. Ensure that all expected nodes appear in the resulting table.

