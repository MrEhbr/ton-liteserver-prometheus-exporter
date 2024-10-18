package collector

import "github.com/prometheus/client_golang/prometheus"

type MetricDef struct {
	desc     *prometheus.Desc
	getValue func(*LiteServerMetrics) (float64, []string)
}

var Metrics = []MetricDef{
	// TON Network Status Metrics
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "online_validators"),
			"Number of online validators",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.OnlineValidators, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "all_validators"),
			"Total number of validators",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.AllValidators, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "number_of_shardchains"),
			"Number of shardchains",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.NumberOfShardchains, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "new_offers"),
			"Number of new offers",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.NewOffers, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "all_offers"),
			"Total number of offers",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.AllOffers, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "new_complaints"),
			"Number of new complaints",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.NewComplaints, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "all_complaints"),
			"Total number of complaints",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.AllComplaints, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "election_status"),
			"Election status (open/closed)",
			[]string{"status"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.ElectionStatus != "" {
				return 1, []string{m.ElectionStatus}
			}
			return 0, []string{"unknown"}
		},
	},

	// Local Validator Status Metrics
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "validator_index"),
			"Index of the local validator",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.ValidatorIndex, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_adnl_address"),
			"ADNL address of the local validator",
			[]string{"address"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.AdnlAddress != "" {
				return 1, []string{m.AdnlAddress}
			}
			return 0, []string{"unknown"}
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_wallet_address"),
			"Local validator wallet address",
			[]string{"address"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.WalletAddress != "" {
				return 1, []string{m.WalletAddress}
			}
			return 0, []string{"unknown"}
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_wallet_balance"),
			"Balance of the local validator wallet",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.WalletBalance, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "mytoncore_status"),
			"Status of Mytoncore",
			[]string{"status"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.MytoncoreStatus != "" {
				return 1, []string{m.MytoncoreStatus}
			}
			return 0, []string{"unknown"}
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "mytoncore_uptime_seconds"),
			"Uptime of Mytoncore in seconds",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.MytoncoreUptimeSeconds, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_status"),
			"Status of Local Validator",
			[]string{"status"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.LocalValidatorStatus != "" {
				return 1, []string{m.LocalValidatorStatus}
			}
			return 0, []string{"unknown"}
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_uptime_seconds"),
			"Uptime of Local Validator in seconds",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.LocalValidatorUptimeSeconds, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_out_of_sync_seconds"),
			"Local validator out of sync in seconds",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.LocalValidatorOutOfSyncSeconds, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_last_state_serialization_blocks"),
			"Number of blocks since last state serialization",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.LocalValidatorLastStateSerializationBlocks, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "local_validator_database_size_gb"),
			"Local validator database size in GB",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.LocalValidatorDatabaseSizeGB, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "version_mytonctrl"),
			"Version of MyTonCtrl",
			[]string{"version"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.VersionMytonctrl != "" {
				return 1, []string{m.VersionMytonctrl}
			}
			return 0, []string{"unknown"}
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "version_validator"),
			"Version of Validator",
			[]string{"version"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.VersionValidator != "" {
				return 1, []string{m.VersionValidator}
			}
			return 0, []string{"unknown"}
		},
	},

	// TON Network Configuration Metrics
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "configurator_address"),
			"Configurator address",
			[]string{"address"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.ConfiguratorAddress != "" {
				return 1, []string{m.ConfiguratorAddress}
			}
			return 0, []string{"unknown"}
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "elector_address"),
			"Elector address",
			[]string{"address"}, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			if m.ElectorAddress != "" {
				return 1, []string{m.ElectorAddress}
			}
			return 0, []string{"unknown"}
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "validation_period_seconds"),
			"Validation period in seconds",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.ValidationPeriodSeconds, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "duration_of_elections_seconds"),
			"Duration of elections in seconds",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.DurationOfElectionsSeconds, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "hold_period_seconds"),
			"Hold period in seconds",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.HoldPeriodSeconds, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "minimum_stake_tons"),
			"Minimum stake in TONs",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.MinimumStakeTONs, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "maximum_stake_tons"),
			"Maximum stake in TONs",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.MaximumStakeTONs, nil
		},
	},

	// TON Timestamps Metrics
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "network_launched_timestamp"),
			"TON network launch timestamp",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.NetworkLaunchedTimestamp, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "start_validation_cycle_timestamp"),
			"Start of the validation cycle timestamp",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.StartValidationCycleTimestamp, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "end_validation_cycle_timestamp"),
			"End of the validation cycle timestamp",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.EndValidationCycleTimestamp, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "start_elections_timestamp"),
			"Start of elections timestamp",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.StartElectionsTimestamp, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "end_elections_timestamp"),
			"End of elections timestamp",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.EndElectionsTimestamp, nil
		},
	},
	{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "begin_next_elections_timestamp"),
			"Beginning of the next elections timestamp",
			nil, nil,
		),
		getValue: func(m *LiteServerMetrics) (float64, []string) {
			return m.BeginNextElectionsTimestamp, nil
		},
	},
}
