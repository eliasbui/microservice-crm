-- Cassandra Keyspace and Table Creation
-- For CRM Event Logs and Time-Series Data

-- Create Keyspace
CREATE KEYSPACE IF NOT EXISTS crm_keyspace
WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': 3
};

USE crm_keyspace;

-- Customer Event Logs Table (Time-Series Data)
CREATE TABLE IF NOT EXISTS customer_event_logs (
    customer_id bigint,
    event_time timestamp,
    event_id uuid,
    event_type text,
    event_category text,
    user_id text,
    session_id text,
    ip_address text,
    user_agent text,
    page_url text,
    referrer text,
    event_data map<text, text>,
    device_type text,
    browser text,
    os text,
    country text,
    city text,
    PRIMARY KEY ((customer_id), event_time, event_id)
) WITH CLUSTERING ORDER BY (event_time DESC, event_id ASC)
  AND compaction = {'class': 'TimeWindowCompactionStrategy'}
  AND default_time_to_live = 15552000; -- 180 days

-- Analytics Metrics Table
CREATE TABLE IF NOT EXISTS analytics_metrics (
    metric_name text,
    metric_date date,
    metric_hour int,
    metric_value double,
    metric_count bigint,
    dimensions map<text, text>,
    PRIMARY KEY ((metric_name, metric_date), metric_hour)
) WITH CLUSTERING ORDER BY (metric_hour DESC);

-- User Activity Summary
CREATE TABLE IF NOT EXISTS user_activity_summary (
    user_id text,
    activity_date date,
    total_events bigint,
    unique_customers bigint,
    total_revenue decimal,
    activity_by_type map<text, bigint>,
    PRIMARY KEY ((user_id), activity_date)
) WITH CLUSTERING ORDER BY (activity_date DESC);

-- Customer Engagement Scores (Time-Series)
CREATE TABLE IF NOT EXISTS customer_engagement_scores (
    customer_id bigint,
    score_date date,
    engagement_score int,
    email_opens int,
    email_clicks int,
    site_visits int,
    purchases int,
    support_tickets int,
    PRIMARY KEY ((customer_id), score_date)
) WITH CLUSTERING ORDER BY (score_date DESC);

-- Real-Time Dashboard Metrics
CREATE TABLE IF NOT EXISTS dashboard_metrics (
    dashboard_id text,
    metric_timestamp timestamp,
    metric_key text,
    metric_value text,
    PRIMARY KEY ((dashboard_id), metric_timestamp, metric_key)
) WITH CLUSTERING ORDER BY (metric_timestamp DESC)
  AND default_time_to_live = 86400; -- 24 hours

-- Create indexes for common queries
CREATE INDEX IF NOT EXISTS idx_event_type ON customer_event_logs (event_type);
CREATE INDEX IF NOT EXISTS idx_event_category ON customer_event_logs (event_category);
CREATE INDEX IF NOT EXISTS idx_session_id ON customer_event_logs (session_id);
