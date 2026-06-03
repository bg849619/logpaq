CREATE TABLE IF NOT EXISTS contacts (
    callsign        TEXT NOT NULL,      -- flat, for dedup index
    band            TEXT NOT NULL,      -- flat, for dedup index
    mode            TEXT NOT NULL,      -- flat, for dedup index
    logged_at       INTEGER NOT NULL,   -- flat, for dedup index
    updated_at      INTEGER NOT NULL,  -- for conflict resolution: last writer wins
    sent_exchange   TEXT NOT NULL DEFAULT '{}',   -- JSON blob
    rcvd_exchange   TEXT NOT NULL DEFAULT '{}',   -- JSON blob
    extension_data  TEXT NOT NULL DEFAULT '{}'    -- JSON blob
    PRIMARY KEY (callsign, band, mode)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_dedup
    ON contacts (callsign, band, mode, logged_at / 300);

CREATE TABLE IF NOT EXISTS vclock (
    node_id  TEXT PRIMARY KEY,
    seq      INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS peers (
    node_id      TEXT PRIMARY KEY,
    callsign     TEXT NOT NULL,
    station_name TEXT NOT NULL,
    last_seen    INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS metadata (
    key    TEXT PRIMARY KEY,
    value  TEXT NOT NULL
);