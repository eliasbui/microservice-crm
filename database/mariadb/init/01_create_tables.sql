-- MariaDB Database Schema for Integration Service
-- Stores integration configurations and sync logs

USE crm_integrations;

-- Integration Connections
CREATE TABLE IF NOT EXISTS integration_connections (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    connection_name VARCHAR(100) NOT NULL UNIQUE,
    integration_type VARCHAR(50) NOT NULL,
    provider VARCHAR(50) NOT NULL,
    auth_type ENUM('OAUTH2', 'API_KEY', 'BASIC_AUTH', 'JWT') NOT NULL,
    credentials JSON NOT NULL,
    config JSON,
    is_active BOOLEAN DEFAULT TRUE,
    last_sync_at TIMESTAMP NULL,
    last_error TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_integration_type (integration_type),
    INDEX idx_provider (provider),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Sync Jobs
CREATE TABLE IF NOT EXISTS sync_jobs (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    connection_id BIGINT NOT NULL,
    job_name VARCHAR(255) NOT NULL,
    sync_direction ENUM('IMPORT', 'EXPORT', 'BIDIRECTIONAL') NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    schedule_cron VARCHAR(100),
    status ENUM('PENDING', 'RUNNING', 'COMPLETED', 'FAILED', 'CANCELLED') DEFAULT 'PENDING',
    started_at TIMESTAMP NULL,
    completed_at TIMESTAMP NULL,
    total_records INT DEFAULT 0,
    processed_records INT DEFAULT 0,
    failed_records INT DEFAULT 0,
    error_message TEXT,
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (connection_id) REFERENCES integration_connections(id) ON DELETE CASCADE,
    INDEX idx_connection_status (connection_id, status),
    INDEX idx_status (status),
    INDEX idx_entity_type (entity_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Sync History
CREATE TABLE IF NOT EXISTS sync_history (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    job_id BIGINT NOT NULL,
    entity_id VARCHAR(255) NOT NULL,
    operation ENUM('CREATE', 'UPDATE', 'DELETE') NOT NULL,
    source_data JSON,
    target_data JSON,
    status ENUM('SUCCESS', 'FAILED', 'SKIPPED') NOT NULL,
    error_message TEXT,
    synced_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_id) REFERENCES sync_jobs(id) ON DELETE CASCADE,
    INDEX idx_job_status (job_id, status),
    INDEX idx_entity_id (entity_id),
    INDEX idx_synced_at (synced_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Field Mappings
CREATE TABLE IF NOT EXISTS field_mappings (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    connection_id BIGINT NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    source_field VARCHAR(100) NOT NULL,
    target_field VARCHAR(100) NOT NULL,
    transformation VARCHAR(255),
    is_required BOOLEAN DEFAULT FALSE,
    default_value TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (connection_id) REFERENCES integration_connections(id) ON DELETE CASCADE,
    INDEX idx_connection_entity (connection_id, entity_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Webhook Endpoints
CREATE TABLE IF NOT EXISTS webhook_endpoints (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    endpoint_name VARCHAR(100) NOT NULL UNIQUE,
    url VARCHAR(500) NOT NULL,
    http_method ENUM('GET', 'POST', 'PUT', 'DELETE') DEFAULT 'POST',
    headers JSON,
    event_types JSON NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    retry_strategy JSON,
    last_triggered_at TIMESTAMP NULL,
    failure_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_is_active (is_active),
    INDEX idx_event_types ((CAST(event_types AS CHAR(100) ARRAY)))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Webhook Logs
CREATE TABLE IF NOT EXISTS webhook_logs (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    endpoint_id BIGINT NOT NULL,
    event_type VARCHAR(50) NOT NULL,
    payload JSON NOT NULL,
    response_status INT,
    response_body TEXT,
    duration_ms INT,
    status ENUM('SUCCESS', 'FAILED', 'RETRY') NOT NULL,
    error_message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (endpoint_id) REFERENCES webhook_endpoints(id) ON DELETE CASCADE,
    INDEX idx_endpoint_status (endpoint_id, status),
    INDEX idx_event_type (event_type),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- API Rate Limits
CREATE TABLE IF NOT EXISTS api_rate_limits (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    connection_id BIGINT NOT NULL,
    rate_limit INT NOT NULL,
    time_window_seconds INT NOT NULL,
    current_usage INT DEFAULT 0,
    window_start TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (connection_id) REFERENCES integration_connections(id) ON DELETE CASCADE,
    INDEX idx_connection (connection_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data Transformation Rules
CREATE TABLE IF NOT EXISTS transformation_rules (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    rule_name VARCHAR(100) NOT NULL UNIQUE,
    entity_type VARCHAR(50) NOT NULL,
    transformation_script TEXT NOT NULL,
    script_language ENUM('JAVASCRIPT', 'PYTHON', 'JMESPATH') DEFAULT 'JAVASCRIPT',
    is_active BOOLEAN DEFAULT TRUE,
    test_input JSON,
    test_output JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_entity_type (entity_type),
    INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample integration connections
INSERT INTO integration_connections (connection_name, integration_type, provider, auth_type, credentials, config) VALUES
('Salesforce Integration', 'CRM', 'Salesforce', 'OAUTH2',
 '{"client_id": "xxx", "client_secret": "xxx", "refresh_token": "xxx"}',
 '{"instance_url": "https://example.salesforce.com", "api_version": "v55.0"}'),
('QuickBooks Integration', 'ACCOUNTING', 'QuickBooks', 'OAUTH2',
 '{"client_id": "xxx", "client_secret": "xxx", "refresh_token": "xxx"}',
 '{"company_id": "xxx", "realm_id": "xxx"}'),
('Mailchimp Integration', 'EMAIL_MARKETING', 'Mailchimp', 'API_KEY',
 '{"api_key": "xxx"}',
 '{"datacenter": "us1", "list_id": "xxx"}');
