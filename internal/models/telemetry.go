package models

import "time"

// TimestampQuality represents the synchronization state of the device clock.
type TimestampQuality string

const (
	QualityNTPNetworkSynced TimestampQuality = "NTP_NETWORK_SYNCED"
	QualityUnsyncedUptime   TimestampQuality = "UNSYNCED_UPTIME"
	QualityHistoricalReplay TimestampQuality = "HISTORICAL_REPLAY"
)

// TelemetryRecord represents an environmental observation matching the SRS/ICD specification.
// Terminology: "Kandungan Air dalam Tanah" (Soil Water Content).
type TelemetryRecord struct {
	ID               int64            `json:"id"`
	DeviceID         string           `json:"device_id"`
	RecordSeq        uint32           `json:"record_seq"`
	Timestamp        time.Time        `json:"timestamp"`
	TimestampQuality TimestampQuality `json:"timestamp_quality"`
	TempC            float64          `json:"temp_c"`
	HumidityPct      float64          `json:"humidity_pct"`
	LightLux         float64          `json:"light_lux"`
	SoilRaw          int32            `json:"soil_raw"`
	SoilPct          float64          `json:"soil_pct"` // Kandungan Air dalam Tanah (%)
	PumpStatus       int32            `json:"pump_status"`
	PumpDurationSec  int32            `json:"pump_duration_sec"`
	WifiRSSI         int32            `json:"wifi_rssi"`
	FreeHeapBytes    uint32           `json:"free_heap_bytes"`
	ErrorCode        int32            `json:"error_code"`
	ReceivedAt       time.Time        `json:"received_at"`
}

// DeviceStatus represents heartbeat / LWT status of an IoT node.
type DeviceStatus struct {
	DeviceID      string    `json:"device_id"`
	Status        string    `json:"status"` // "online" | "offline"
	FirmwareVer   string    `json:"firmware_version,omitempty"`
	IPAddress     string    `json:"ip_address,omitempty"`
	LastSeen      time.Time `json:"last_seen"`
	WifiRSSI      int32     `json:"wifi_rssi"`
	FreeHeapBytes uint32    `json:"free_heap_bytes"`
}

// BackfillBatch represents historical data payload uploaded from SD Card.
type BackfillBatch struct {
	DeviceID   string            `json:"device_id"`
	BatchID    string            `json:"batch_id"`
	TotalItems int               `json:"total_items"`
	Records    []TelemetryRecord `json:"records"`
}

// BackfillACK represents Application ACK sent back to device after DB commit.
type BackfillACK struct {
	DeviceID           string `json:"device_id"`
	AckType            string `json:"ack_type"` // "COMMIT_OK" | "SCHEMA_ERROR"
	CommittedThroughSeq uint32 `json:"committed_through_seq"`
	Status             string `json:"status"`
}
