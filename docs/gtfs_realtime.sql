CREATE SCHEMA IF NOT EXISTS agency_name;

CREATE TABLE IF NOT EXISTS agency_name.vehicle_positions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    trip_id VARCHAR(255),
    route_id VARCHAR(255),
    direction_id INTEGER,
    start_datetime TIMESTAMP WITH TIME ZONE,
    schedule_relationship VARCHAR(255),
    vehicle_id VARCHAR(255),
    vehicle_label VARCHAR(255),
    vehicle_position POINT,
    current_stop_sequence INTEGER,
    stop_id VARCHAR(255),
    timestamp TIMESTAMP WITH TIME ZONE
);
