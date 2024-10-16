package collector

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestParser_ParseOutput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     LiteServerMetrics
		whantErr bool
	}{
		{
			name:     "empty input",
			input:    "",
			want:     LiteServerMetrics{},
			whantErr: true,
		},
		{
			name: "valid input",
			input: `
Welcome to the console. Enter 'help' to display the help menu.
[debug]   16.10.2024, 16:11:47.328 (UTC)  <MainThread>  start GetValidatorStatus function
[debug]   16.10.2024, 16:11:47.334 (UTC)  <MainThread>  start GetConfig32 function
MyTonCtrl> [debug]   16.10.2024, 16:11:47.345 (UTC)  <MainThread>  start GetValidatorWallet function
[debug]   16.10.2024, 16:11:47.345 (UTC)  <MainThread>  start GetLocalWallet function
[debug]   16.10.2024, 16:11:47.345 (UTC)  <MainThread>  start GetWalletFromFile function
[debug]   16.10.2024, 16:11:47.345 (UTC)  <MainThread>  start WalletVersion2Wallet function
[debug]   16.10.2024, 16:11:47.345 (UTC)  <MainThread>  start GetDbSize function
[debug]   16.10.2024, 16:11:47.353 (UTC)  <MainThread>  start GetRootWorkchainEnabledTime function
[debug]   16.10.2024, 16:11:47.353 (UTC)  <MainThread>  start GetConfig function (12)
[debug]   16.10.2024, 16:11:47.361 (UTC)  <MainThread>  start GetConfig34 function
[debug]   16.10.2024, 16:11:47.368 (UTC)  <MainThread>  start GetConfig36 function
[debug]   16.10.2024, 16:11:47.375 (UTC)  <MainThread>  start GetValidatorsLoad function (1729094047, 1729095047)
[debug]   16.10.2024, 16:11:48.578 (UTC)  <MainThread>  start GetConfig function (15)
[debug]   16.10.2024, 16:11:48.585 (UTC)  <MainThread>  start GetConfig function (17)
[debug]   16.10.2024, 16:11:48.593 (UTC)  <MainThread>  start GetFullConfigAddr function
[debug]   16.10.2024, 16:11:48.600 (UTC)  <MainThread>  start GetFullElectorAddr function
[debug]   16.10.2024, 16:11:48.608 (UTC)  <MainThread>  start GetActiveElectionId function
[warning] 16.10.2024, 16:11:48.621 (UTC)  <MainThread>  GetValidatorIndex warning: index not found.
[debug]   16.10.2024, 16:11:48.621 (UTC)  <MainThread>  start GetOffersNumber function
[debug]   16.10.2024, 16:11:48.621 (UTC)  <MainThread>  start GetOffers function
[debug]   16.10.2024, 16:11:48.642 (UTC)  <MainThread>  start GetOffers function
[debug]   16.10.2024, 16:11:48.651 (UTC)  <MainThread>  start GetComplaintsNumber function
[debug]   16.10.2024, 16:11:48.651 (UTC)  <MainThread>  start GetComplaints function
[warning] 16.10.2024, 16:11:48.658 (UTC)  <MainThread>  GetValidatorIndex warning: index not found.
===[ TON network status ]===
Network name: testnet
Number of validators: 23(26)
Number of shardchains: 4
Number of offers: 1(11)
Number of complaints: 2(22)
Election status: open

===[ Node status ]===
ADNL address of local validator: A56F2F60C9309BA2767EF3737A57B4CA1EB4DE66EE288F84ECC615B3EE6C8C81
Public ADNL address of node: 5FEEBBC14F9098F4D216524E9B4D5DA5C56944C792B3CD70FB7BAC25513B5C23
Load average[16]: 0.48, 0.47, 0.44
Network load average (Mbit/s): 25.39, 24.22, 24.79
Memory load: ram:[17.31 Gb, 14.0%], swap:[0.0 Gb, 0.0%]
Disks load average (MB/s): nvme0n1:[0.18, 0.15%], nvme1n1:[0.0, 0.0%], nvme2n1:[0.83, 10.08%]
Mytoncore status: working, 16 days
Local validator status: working, 16 days
Local validator out of sync: 3 s
Local validator last state serialization: 2 blocks ago
Local validator database size: 27.31 Gb, 82.0%
Version mytonctrl: a467af5 (master)
Version validator: 1bef6df (master)


MyTonCtrl> Bye.
`,
			want: LiteServerMetrics{
				NetworkName:                    "testnet",
				OnlineValidators:               23,
				AllValidators:                  26,
				NumberOfShardchains:            4,
				NewOffers:                      1,
				AllOffers:                      11,
				NewComplaints:                  2,
				AllComplaints:                  22,
				ElectionStatus:                 "open",
				AdnlAddress:                    "A56F2F60C9309BA2767EF3737A57B4CA1EB4DE66EE288F84ECC615B3EE6C8C81",
				MytoncoreStatus:                "working",
				MytoncoreUptimeSeconds:         (16 * 24 * time.Hour).Seconds(),
				LocalValidatorStatus:           "working",
				LocalValidatorUptimeSeconds:    (16 * 24 * time.Hour).Seconds(),
				LocalValidatorOutOfSyncSeconds: 3,
				LocalValidatorLastStateSerializationBlocks: 2,
				LocalValidatorDatabaseSizeGB:               27.31,
				VersionMytonctrl:                           "a467af5 (master)",
				VersionValidator:                           "1bef6df (master)",
			},
			whantErr: false,
		},
		{
			name: "another valid input",
			input: `
[debug]
[warning]
[debug]
debugJ
debug]
[debug]
[debug]
start GetActiveElectionId function
===[ TON network status ]===
Network name: testnet
Number of validators: 23(24)
Number of shardchains: 4
Number of offers: 0(0)
Number of complaints: 0(0)
Election status: closed
===[ Node status ]===
Validator index: -1
ADNL address of local validator: D70B2BD8F394EF44DE89B6614DD0D91C672FC28C8A4B7D8C918E2C106A9DF932
Public ADNL address of node: BFE38D4B6B7CAB1FB4396068590C26857A400F93D2F4693208D181B7279B0348
Local validator wallet address: kf_1160NJDnTqgz9AJ6p7EweUP_15xaZw1PK4WPcd48tDbQ1
Local validator wallet balance: 95290.938201014
Load average[16]: 0.47, 0.4, 0.37
Network load average (Mbit/s): 24.61, 23.6, 23.59
Memory load: ram: [62.17 Gb, 47.5%], swap: [0.0 Gb, 0.2%]
Disks load average (MB/s): sr0: [0.0, 0.0%], sr1: [0.0, 0.0%], sr2: [0.0, 0.0%], vda: [0.18, 0.47%], vdb: [0.0, 0.0%], vdc: [0.54, 4.28%]
Mytoncore status: working, 14 days
Local validator status: working, 16 days
Local validator out of sync: 3 s
Local validator last state serialization: 3 blocks ago
Local validator database size: 25.89 Gb, 2.4%
Version mytonctrl: 74536b (master)
Version validator: 0c21ce2 (master)
===[ TON network configuration ]===
Configurator address: -1:5555555555555555555555555555555555555555555555555555555555555555
Elector address: -1:3333333333333333333333333333333333333333333333333333333333333333
Validation period: 7200, Duration of elections: 2400-180, Hold period: 900
Minimum stake: 10000.0, Maximum stake: 5000000.0
===[ TON timestamps ]===
TON network was launched: 15.11.2019 12:44:14 UTC
Start of the validation cycle: 24.09.2024 06:19:55 UTC
End of the validation cycle: 24.09.2024 08:19:55 UTC
Start of elections: 24.09.2024 05:39:55 UTC
End of elections: 24.09.2024 06:16:55 UTC
Beginning of the next elections: 24.09.2024 07:39:55 UTC
`,
			want: LiteServerMetrics{
				NetworkName:                    "testnet",
				OnlineValidators:               23,
				AllValidators:                  24,
				NumberOfShardchains:            4,
				NewOffers:                      0,
				AllOffers:                      0,
				NewComplaints:                  0,
				AllComplaints:                  0,
				ElectionStatus:                 "closed",
				ValidatorIndex:                 -1,
				AdnlAddress:                    "D70B2BD8F394EF44DE89B6614DD0D91C672FC28C8A4B7D8C918E2C106A9DF932",
				WalletAddress:                  "kf_1160NJDnTqgz9AJ6p7EweUP_15xaZw1PK4WPcd48tDbQ1",
				WalletBalance:                  95290.938201014,
				MytoncoreStatus:                "working",
				MytoncoreUptimeSeconds:         (14 * 24 * time.Hour).Seconds(),
				LocalValidatorStatus:           "working",
				LocalValidatorUptimeSeconds:    (16 * 24 * time.Hour).Seconds(),
				LocalValidatorOutOfSyncSeconds: 3,
				LocalValidatorLastStateSerializationBlocks: 3,
				LocalValidatorDatabaseSizeGB:               25.89,
				VersionMytonctrl:                           "74536b (master)",
				VersionValidator:                           "0c21ce2 (master)",
				ConfiguratorAddress:                        "-1:5555555555555555555555555555555555555555555555555555555555555555",
				ElectorAddress:                             "-1:3333333333333333333333333333333333333333333333333333333333333333",
				ValidationPeriodSeconds:                    7200,
				DurationOfElectionsSeconds:                 2400,
				HoldPeriodSeconds:                          900,
				MinimumStakeTONs:                           10000,
				MaximumStakeTONs:                           5000000,
				NetworkLaunchedTimestamp:                   float64(time.Date(2019, 11, 15, 12, 44, 14, 0, time.UTC).Unix()),
				StartValidationCycleTimestamp:              float64(time.Date(2024, 9, 24, 6, 19, 55, 0, time.UTC).Unix()),
				EndValidationCycleTimestamp:                float64(time.Date(2024, 9, 24, 8, 19, 55, 0, time.UTC).Unix()),
				StartElectionsTimestamp:                    float64(time.Date(2024, 9, 24, 5, 39, 55, 0, time.UTC).Unix()),
				EndElectionsTimestamp:                      float64(time.Date(2024, 9, 24, 6, 16, 55, 0, time.UTC).Unix()),
				BeginNextElectionsTimestamp:                float64(time.Date(2024, 9, 24, 7, 39, 55, 0, time.UTC).Unix()),
			},
			whantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewParser()
			got, err := parser.ParseOutput(tt.input)
			if (err != nil) != tt.whantErr {
				t.Errorf("Parser.ParseOutput() error = %v, wantErr %v", err, tt.whantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Parser.ParseOutput() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
