-- Inisialisasi Skema Database PostgreSQL (Sesuai ICD & SRS v0.2)

CREATE TABLE IF NOT EXISTS devices (
    device_id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'offline',
    firmware_version VARCHAR(32),
    ip_address VARCHAR(45),
    wifi_rssi INTEGER,
    free_heap_bytes BIGINT,
    last_seen_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS telemetry_records (
    id BIGSERIAL PRIMARY KEY,
    device_id VARCHAR(64) NOT NULL REFERENCES devices(device_id) ON DELETE CASCADE,
    record_seq BIGINT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    timestamp_quality VARCHAR(32) NOT NULL DEFAULT 'NTP_NETWORK_SYNCED',
    temp_c NUMERIC(5, 2) NOT NULL,
    humidity_pct NUMERIC(5, 2) NOT NULL,
    light_lux NUMERIC(8, 2) NOT NULL,
    soil_raw INTEGER NOT NULL,
    soil_pct NUMERIC(5, 2) NOT NULL, -- Kandungan Air dalam Tanah (%)
    pump_status INTEGER NOT NULL DEFAULT 0,
    pump_duration_sec INTEGER NOT NULL DEFAULT 0,
    wifi_rssi INTEGER,
    free_heap_bytes BIGINT,
    error_code INTEGER NOT NULL DEFAULT 0,
    received_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_device_record_seq UNIQUE (device_id, record_seq)
);

CREATE INDEX IF NOT EXISTS idx_telemetry_device_time ON telemetry_records (device_id, timestamp DESC);
CREATE INDEX IF NOT EXISTS idx_telemetry_time ON telemetry_records (timestamp DESC);

-- Seed initial device
INSERT INTO devices (device_id, name, status, firmware_version, created_at, updated_at)
VALUES ('ESP32S3-POT-01', 'Smart Pot ESP32-S3 Node 1', 'online', '1.0.0', NOW(), NOW())
ON CONFLICT (device_id) DO NOTHING;
