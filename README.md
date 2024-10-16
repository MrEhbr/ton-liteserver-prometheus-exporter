# TON LiteServer Prometheus Exporter

![License](https://img.shields.io/github/license/MrEhbr/ton-liteserver-prometheus-exporter)
![Release](https://img.shields.io/github/v/release/MrEhbr/ton-liteserver-prometheus-exporter)
[![Go](https://github.com/MrEhbr/ton-liteserver-prometheus-exporter/actions/workflows/go.yml/badge.svg)](https://github.com/MrEhbr/ton-liteserver-prometheus-exporter/actions/workflows/go.yml)
[![GitHub release](https://img.shields.io/github/release/MrEhbr/ton-liteserver-prometheus-exporter.svg)](https://github.com/MrEhbr/ton-liteserver-prometheus-exporter/releases)
![Made by Alexey Burmistrov](https://img.shields.io/badge/made%20by-Alexey%20Burmistrov-blue.svg?style=flat)

## What It Is

The **TON LiteServer Prometheus Exporter** helps you monitor your TON LiteServer by exposing its metrics to Prometheus. It takes the output from the `mytonctrl status` command and turns it into Prometheus-friendly metrics.

## Install

### Using go

```console
go get -u github.com/MrEhbr/ton-liteserver-prometheus-exporter/cmd/ton-liteserver-prometheus-exporter
```

## Usage

```console
ton-liteserver-prometheus-exporter --port 9100
```

## Metrics

The **TON LiteServer Prometheus Exporter** exposes a variety of metrics to help you monitor the health and performance of your TON LiteServer. Below is a summary of the available metrics:

### TON Network Status Metrics

- **`ton_liteserver_prometheus_exporter_online_validators`**
  - **Description:** Number of online validators.
  
- **`ton_liteserver_prometheus_exporter_all_validators`**
  - **Description:** Total number of validators.
  
- **`ton_liteserver_prometheus_exporter_number_of_shardchains`**
  - **Description:** Number of shardchains.
  
- **`ton_liteserver_prometheus_exporter_new_offers`**
  - **Description:** Number of new offers.
  
- **`ton_liteserver_prometheus_exporter_all_offers`**
  - **Description:** Total number of offers.
  
- **`ton_liteserver_prometheus_exporter_new_complaints`**
  - **Description:** Number of new complaints.
  
- **`ton_liteserver_prometheus_exporter_all_complaints`**
  - **Description:** Total number of complaints.
  
- **`ton_liteserver_prometheus_exporter_election_status`**
  - **Description:** Current election status (e.g., open, closed).
  - **Labels:**
    - `status` – The current status of the election.

### Local Validator Status Metrics

- **`ton_liteserver_prometheus_exporter_validator_index`**
  - **Description:** Index of the local validator.
  
- **`ton_liteserver_prometheus_exporter_local_validator_adnl_address`**
  - **Description:** ADNL address of the local validator.
  - **Labels:**
    - `address` – The ADNL address.
  
- **`ton_liteserver_prometheus_exporter_local_validator_wallet_address`**
  - **Description:** Wallet address of the local validator.
  - **Labels:**
    - `address` – The wallet address.
  
- **`ton_liteserver_prometheus_exporter_local_validator_wallet_balance`**
  - **Description:** Balance of the local validator's wallet.
  
- **`ton_liteserver_prometheus_exporter_mytoncore_status`**
  - **Description:** Status of Mytoncore (e.g., working).
  - **Labels:**
    - `status` – The current status.
  
- **`ton_liteserver_prometheus_exporter_mytoncore_uptime_seconds`**
  - **Description:** Uptime of Mytoncore in seconds.
  
- **`ton_liteserver_prometheus_exporter_local_validator_status`**
  - **Description:** Status of the Local Validator (e.g., working).
  - **Labels:**
    - `status` – The current status.
  
- **`ton_liteserver_prometheus_exporter_local_validator_uptime_seconds`**
  - **Description:** Uptime of the Local Validator in seconds.
  
- **`ton_liteserver_prometheus_exporter_local_validator_out_of_sync_seconds`**
  - **Description:** Time the local validator has been out of sync in seconds.
  
- **`ton_liteserver_prometheus_exporter_local_validator_last_state_serialization_blocks`**
  - **Description:** Number of blocks since the last state serialization.
  
- **`ton_liteserver_prometheus_exporter_local_validator_database_size_gb`**
  - **Description:** Size of the local validator's database in GB.

### TON Network Configuration Metrics

- **`ton_liteserver_prometheus_exporter_configurator_address`**
  - **Description:** Configurator address.
  - **Labels:**
    - `address` – The configurator address.
  
- **`ton_liteserver_prometheus_exporter_elector_address`**
  - **Description:** Elector address.
  - **Labels:**
    - `address` – The elector address.
  
- **`ton_liteserver_prometheus_exporter_validation_period_seconds`**
  - **Description:** Validation period in seconds.
  
- **`ton_liteserver_prometheus_exporter_duration_of_elections_seconds`**
  - **Description:** Duration of elections in seconds.
  
- **`ton_liteserver_prometheus_exporter_hold_period_seconds`**
  - **Description:** Hold period in seconds.
  
- **`ton_liteserver_prometheus_exporter_minimum_stake_tons`**
  - **Description:** Minimum stake required in TONs.
  
- **`ton_liteserver_prometheus_exporter_maximum_stake_tons`**
  - **Description:** Maximum stake allowed in TONs.

### TON Timestamps Metrics

- **`ton_liteserver_prometheus_exporter_network_launched_timestamp`**
  - **Description:** UNIX timestamp when the TON network was launched.
  
- **`ton_liteserver_prometheus_exporter_start_validation_cycle_timestamp`**
  - **Description:** UNIX timestamp for the start of the validation cycle.
  
- **`ton_liteserver_prometheus_exporter_end_validation_cycle_timestamp`**
  - **Description:** UNIX timestamp for the end of the validation cycle.
  
- **`ton_liteserver_prometheus_exporter_start_elections_timestamp`**
  - **Description:** UNIX timestamp for the start of elections.
  
- **`ton_liteserver_prometheus_exporter_end_elections_timestamp`**
  - **Description:** UNIX timestamp for the end of elections.
  
- **`ton_liteserver_prometheus_exporter_begin_next_elections_timestamp`**
  - **Description:** UNIX timestamp for the beginning of the next elections.

### Version Metrics

- **`ton_liteserver_prometheus_exporter_version_mytonctrl`**
  - **Description:** Version of MyTonCtrl.
  - **Labels:**
    - `version` – The version string.
  
- **`ton_liteserver_prometheus_exporter_version_validator`**
  - **Description:** Version of the Validator.
  - **Labels:**
    - `version` – The version string.

---

*For detailed information on each metric and their implementation, refer to the [`collector/metrics.go`](collector/metrics.go).*

### Download releases

<https://github.com/MrEhbr/ton-liteserver-prometheus-exporter/releases>

## License

© 2024 [Alexey Burmistrov]

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE`](LICENSE)). See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: Apache-2.0`
