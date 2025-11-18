package com.crm.core.cassandra;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.cassandra.core.cql.PrimaryKeyType;
import org.springframework.data.cassandra.core.mapping.Column;
import org.springframework.data.cassandra.core.mapping.PrimaryKeyColumn;
import org.springframework.data.cassandra.core.mapping.Table;

import java.time.Instant;
import java.util.Map;
import java.util.UUID;

/**
 * Cassandra Table for Customer Event Logs
 * Time-series data optimized for write-heavy operations
 */
@Table("customer_event_logs")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CustomerEventLog {

    @PrimaryKeyColumn(name = "customer_id", ordinal = 0, type = PrimaryKeyType.PARTITIONED)
    private Long customerId;

    @PrimaryKeyColumn(name = "event_time", ordinal = 1, type = PrimaryKeyType.CLUSTERED)
    private Instant eventTime;

    @PrimaryKeyColumn(name = "event_id", ordinal = 2, type = PrimaryKeyType.CLUSTERED)
    private UUID eventId;

    @Column("event_type")
    private String eventType; // page_view, button_click, email_open, purchase, etc.

    @Column("event_category")
    private String eventCategory; // engagement, transaction, support

    @Column("user_id")
    private String userId; // User who triggered the event

    @Column("session_id")
    private String sessionId;

    @Column("ip_address")
    private String ipAddress;

    @Column("user_agent")
    private String userAgent;

    @Column("page_url")
    private String pageUrl;

    @Column("referrer")
    private String referrer;

    @Column("event_data")
    private Map<String, String> eventData; // Additional event-specific data

    @Column("device_type")
    private String deviceType; // mobile, desktop, tablet

    @Column("browser")
    private String browser;

    @Column("os")
    private String os;

    @Column("country")
    private String country;

    @Column("city")
    private String city;
}
