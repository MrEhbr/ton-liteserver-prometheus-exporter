package collector

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/commander-cli/cmd"
)

// LiteServerMetrics holds the parsed metrics from MyTonCtrl output.
type LiteServerMetrics struct {
	// TON Network Status Metrics
	NetworkName         string  `json:"network_name"`
	OnlineValidators    float64 `json:"online_validators"`
	AllValidators       float64 `json:"all_validators"`
	NumberOfShardchains float64 `json:"number_of_shardchains"`
	NewOffers           float64 `json:"new_offers"`
	AllOffers           float64 `json:"all_offers"`
	NewComplaints       float64 `json:"new_complaints"`
	AllComplaints       float64 `json:"all_complaints"`
	ElectionStatus      string  `json:"election_status"`

	// Local Validator Status Metrics
	ValidatorIndex                             float64 `json:"validator_index"`
	AdnlAddress                                string  `json:"adnl_address"`
	WalletAddress                              string  `json:"wallet_address"`
	WalletBalance                              float64 `json:"wallet_balance"`
	MytoncoreStatus                            string  `json:"mytoncore_status"`
	MytoncoreUptimeSeconds                     float64 `json:"mytoncore_uptime_seconds"`
	LocalValidatorStatus                       string  `json:"local_validator_status"`
	LocalValidatorUptimeSeconds                float64 `json:"local_validator_uptime_seconds"`
	LocalValidatorOutOfSyncSeconds             float64 `json:"local_validator_out_of_sync_seconds"`
	LocalValidatorLastStateSerializationBlocks float64 `json:"local_validator_last_state_serialization_blocks"`
	LocalValidatorDatabaseSizeGB               float64 `json:"local_validator_database_size_gb"`
	VersionMytonctrl                           string  `json:"version_mytonctrl"`
	VersionValidator                           string  `json:"version_validator"`

	// TON Network Configuration Metrics
	ConfiguratorAddress        string  `json:"configurator_address"`
	ElectorAddress             string  `json:"elector_address"`
	ValidationPeriodSeconds    float64 `json:"validation_period_seconds"`
	DurationOfElectionsSeconds float64 `json:"duration_of_elections_seconds"`
	HoldPeriodSeconds          float64 `json:"hold_period_seconds"`
	MinimumStakeTONs           float64 `json:"minimum_stake_tons"`
	MaximumStakeTONs           float64 `json:"maximum_stake_tons"`

	// TON Timestamps Metrics
	NetworkLaunchedTimestamp      float64 `json:"network_launched_timestamp"`
	StartValidationCycleTimestamp float64 `json:"start_validation_cycle_timestamp"`
	EndValidationCycleTimestamp   float64 `json:"end_validation_cycle_timestamp"`
	StartElectionsTimestamp       float64 `json:"start_elections_timestamp"`
	EndElectionsTimestamp         float64 `json:"end_elections_timestamp"`
	BeginNextElectionsTimestamp   float64 `json:"begin_next_elections_timestamp"`
}

// Parser encapsulates the parsing logic for MyTonCtrl status output.
type Parser struct{}

// NewParser initializes and returns a new Parser instance.
func NewParser() *Parser {
	return &Parser{}
}

// Parse executes the 'mytonctrl status' command and parses its output into LightServerMetrics.
func (p *Parser) Parse() (*LiteServerMetrics, error) {
	metrics, err := p.executeCommand()
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

var ansiEscape = regexp.MustCompile(`\x1B[@-_][0-?]*[ -/]*[@-~]`)

// executeCommand runs the 'mytonctrl status' command and parses its output.
func (p *Parser) executeCommand() (*LiteServerMetrics, error) {
	command := cmd.NewCommand("echo 'status' | mytonctrl", cmd.WithInheritedEnvironment(nil))

	if err := command.Execute(); err != nil {
		return nil, err
	}

	if command.ExitCode() != 0 {
		return nil, fmt.Errorf("command %q failed with exit code %d: %s", command.Command, command.ExitCode(), command.Combined())
	}

	metrics, err := p.ParseOutput(command.Stdout())
	if err != nil {
		return nil, fmt.Errorf("error parsing output: %w", err)
	}

	return &metrics, nil
}

// ParseOutput parses the output from 'mytonctrl status' command.
//
//nolint:funlen,gocyclo // This function is long due to the number of fields to parse.
func (p *Parser) ParseOutput(output string) (LiteServerMetrics, error) {
	if output == "" {
		return LiteServerMetrics{}, errors.New("empty input")
	}
	m := LiteServerMetrics{}
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := cleanLine(scanner.Text())

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Check for key prefixes
		switch {
		// TON Network Status
		case strings.HasPrefix(line, "Network name:"):
			m.NetworkName = extractValue(line, "Network name:")
		case strings.HasPrefix(line, "Number of validators:"):
			value := extractValue(line, "Number of validators:")
			m.OnlineValidators, m.AllValidators = parseValidators(value)
		case strings.HasPrefix(line, "Number of shardchains:"):
			m.NumberOfShardchains = parseFloat(extractValue(line, "Number of shardchains:"))
		case strings.HasPrefix(line, "Number of offers:"):
			value := extractValue(line, "Number of offers:")
			m.NewOffers, m.AllOffers = parseOffersOrComplaints(value)
		case strings.HasPrefix(line, "Number of complaints:"):
			value := extractValue(line, "Number of complaints:")
			m.NewComplaints, m.AllComplaints = parseOffersOrComplaints(value)
		case strings.HasPrefix(line, "Election status:"):
			m.ElectionStatus = extractValue(line, "Election status:")

		// Local Validator Status
		case strings.HasPrefix(line, "Validator index:"):
			m.ValidatorIndex = parseFloat(extractValue(line, "Validator index:"))
		case strings.HasPrefix(line, "ADNL address of local validator:"):
			m.AdnlAddress = extractValue(line, "ADNL address of local validator:")
		case strings.HasPrefix(line, "Local validator wallet address:"):
			m.WalletAddress = extractValue(line, "Local validator wallet address:")
		case strings.HasPrefix(line, "Local validator wallet balance:"):
			m.WalletBalance = parseFloat(extractValue(line, "Local validator wallet balance:"))
		case strings.HasPrefix(line, "Mytoncore status:"):
			status, uptime := parseStatusAndUptime(extractValue(line, "Mytoncore status:"))
			m.MytoncoreStatus = status
			m.MytoncoreUptimeSeconds = uptime
		case strings.HasPrefix(line, "Local validator status:"):
			status, uptime := parseStatusAndUptime(extractValue(line, "Local validator status:"))
			m.LocalValidatorStatus = status
			m.LocalValidatorUptimeSeconds = uptime
		case strings.HasPrefix(line, "Local validator out of sync:"):
			m.LocalValidatorOutOfSyncSeconds = parseFloat(extractValue(line, "Local validator out of sync:"))
		case strings.HasPrefix(line, "Local validator last state serialization:"):
			m.LocalValidatorLastStateSerializationBlocks = parseFloat(extractValue(line, "Local validator last state serialization:"))
		case strings.HasPrefix(line, "Local validator database size:"):
			// Assuming the format is "25.89 Gb, 2.4%"
			value := extractValue(line, "Local validator database size:")
			parts := strings.Split(value, ",")
			if len(parts) >= 1 {
				m.LocalValidatorDatabaseSizeGB = parseFloat(parts[0])
			}
		case strings.HasPrefix(line, "Version mytonctrl:"):
			m.VersionMytonctrl = extractValue(line, "Version mytonctrl:")
		case strings.HasPrefix(line, "Version validator:"):
			m.VersionValidator = extractValue(line, "Version validator:")

		// TON Network Configuration
		case strings.HasPrefix(line, "Configurator address:"):
			m.ConfiguratorAddress = extractValue(line, "Configurator address:")
		case strings.HasPrefix(line, "Elector address:"):
			m.ElectorAddress = extractValue(line, "Elector address:")
		case strings.HasPrefix(line, "Validation period:"):
			// Handle "Validation period: 7200, Duration of elections: 2400-180, Hold period: 900"
			pairs := strings.Split(line, ",")
			for _, pair := range pairs {
				pair = strings.TrimSpace(pair)
				switch {
				case strings.HasPrefix(pair, "Validation period:"):
					m.ValidationPeriodSeconds = parseFloat(extractValue(pair, "Validation period:"))
				case strings.HasPrefix(pair, "Duration of elections:"):
					// Assuming format "2400-180"
					subParts := strings.Split(extractValue(pair, "Duration of elections:"), "-")
					if len(subParts) >= 1 {
						m.DurationOfElectionsSeconds = parseFloat(subParts[0])
					}
				case strings.HasPrefix(pair, "Hold period:"):
					m.HoldPeriodSeconds = parseFloat(extractValue(pair, "Hold period:"))
				}
			}
		case strings.Contains(line, "Minimum stake:"):
			// Handle "Minimum stake: 10000.0, Maximum stake: 5000000.0"
			parts := strings.Split(line, ",")
			for _, part := range parts {
				part = strings.TrimSpace(part)
				switch {
				case strings.HasPrefix(part, "Minimum stake:"):
					m.MinimumStakeTONs = parseFloat(extractValue(part, "Minimum stake:"))
				case strings.HasPrefix(part, "Maximum stake:"):
					m.MaximumStakeTONs = parseFloat(extractValue(part, "Maximum stake:"))
				}
			}
		// TON Timestamps
		case strings.HasPrefix(line, "TON network was launched:"):
			m.NetworkLaunchedTimestamp = parseTimestamp(extractValue(line, "TON network was launched:"))
		case strings.HasPrefix(line, "Start of the validation cycle:"):
			m.StartValidationCycleTimestamp = parseTimestamp(extractValue(line, "Start of the validation cycle:"))
		case strings.HasPrefix(line, "End of the validation cycle:"):
			m.EndValidationCycleTimestamp = parseTimestamp(extractValue(line, "End of the validation cycle:"))
		case strings.HasPrefix(line, "Start of elections:"):
			m.StartElectionsTimestamp = parseTimestamp(extractValue(line, "Start of elections:"))
		case strings.HasPrefix(line, "End of elections:"):
			m.EndElectionsTimestamp = parseTimestamp(extractValue(line, "End of elections:"))
		case strings.HasPrefix(line, "Beginning of the next elections:"):
			m.BeginNextElectionsTimestamp = parseTimestamp(extractValue(line, "Beginning of the next elections:"))
		}
	}

	if err := scanner.Err(); err != nil {
		return LiteServerMetrics{}, fmt.Errorf("scanner error: %w", err)
	}

	return m, nil
}

// extractValue removes the prefix from the line and returns the trimmed value.
func extractValue(line, prefix string) string {
	return strings.TrimSpace(strings.TrimPrefix(line, prefix))
}

func cleanLine(line string) string {
	return ansiEscape.ReplaceAllString(line, "")
}

// parseFloat safely parses a float from string, returns -1 on failure.
func parseFloat(value string) float64 {
	// Handle cases like "123.45 TON" by taking the first field
	fields := strings.Fields(value)
	if len(fields) == 0 {
		return -1
	}
	num, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return -1
	}
	return num
}

// parseTimestamp parses a timestamp string in the format "02.01.2006 15:04:05 UTC" and returns Unix time.
func parseTimestamp(value string) float64 {
	layout := "02.01.2006 15:04:05 UTC"
	t, err := time.Parse(layout, value)
	if err != nil {
		return -1
	}
	return float64(t.Unix())
}

// parseValidators parses the "Number of validators" value.
func parseValidators(value string) (float64, float64) {
	parts := strings.Split(value, "(")
	if len(parts) != 2 {
		return -1, -1
	}
	onlineStr := strings.TrimSpace(parts[0])
	allStr := strings.TrimSuffix(strings.TrimSpace(parts[1]), ")")
	return parseFloat(onlineStr), parseFloat(allStr)
}

// parseOffersOrComplaints parses the "Number of offers/complaints" value.
func parseOffersOrComplaints(value string) (float64, float64) {
	parts := strings.Split(value, "(")
	if len(parts) != 2 {
		return -1, -1
	}
	newStr := strings.TrimSpace(parts[0])
	allStr := strings.TrimSuffix(strings.TrimSpace(parts[1]), ")")
	return parseFloat(newStr), parseFloat(allStr)
}

// parseStatusAndUptime parses status and uptime from a value string.
func parseStatusAndUptime(value string) (string, float64) {
	// Expected format: "working, 14 days"
	parts := strings.SplitN(value, ",", 2)
	if len(parts) != 2 {
		return "unknown", -1
	}
	return strings.TrimSpace(parts[0]), convertUptimeToSeconds(strings.TrimSpace(parts[1]))
}

// convertUptimeToSeconds converts uptime strings like "14 days", "16 days", "3 s" etc., to seconds.
func convertUptimeToSeconds(uptimeStr string) float64 {
	parts := strings.Fields(uptimeStr)
	if len(parts) < 2 {
		return -1
	}
	timeValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1
	}
	unit := strings.ToLower(parts[1])
	switch {
	case strings.Contains(unit, "day"):
		return float64(timeValue * 86400)
	case strings.Contains(unit, "hour"):
		return float64(timeValue * 3600)
	case strings.Contains(unit, "minute"):
		return float64(timeValue * 60)
	case strings.Contains(unit, "second"):
		return float64(timeValue)
	default:
		return -1
	}
}
