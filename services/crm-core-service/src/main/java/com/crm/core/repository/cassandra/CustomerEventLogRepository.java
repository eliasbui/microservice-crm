package com.crm.core.repository.cassandra;

import com.crm.core.cassandra.CustomerEventLog;
import org.springframework.data.cassandra.repository.CassandraRepository;
import org.springframework.data.cassandra.repository.Query;
import org.springframework.stereotype.Repository;

import java.time.Instant;
import java.util.List;
import java.util.UUID;

@Repository
public interface CustomerEventLogRepository extends CassandraRepository<CustomerEventLog, UUID> {

    List<CustomerEventLog> findByCustomerId(Long customerId);

    List<CustomerEventLog> findByCustomerIdAndEventTimeBetween(
        Long customerId,
        Instant startTime,
        Instant endTime
    );

    List<CustomerEventLog> findByCustomerIdAndEventType(Long customerId, String eventType);

    @Query("SELECT * FROM customer_event_logs WHERE customer_id = ?0 AND event_time >= ?1 ALLOW FILTERING")
    List<CustomerEventLog> findRecentEvents(Long customerId, Instant since);

    @Query("SELECT * FROM customer_event_logs WHERE customer_id = ?0 AND event_category = ?1 ALLOW FILTERING")
    List<CustomerEventLog> findByCustomerIdAndEventCategory(Long customerId, String eventCategory);
}
