package com.crm.customer.repository.mongo;

import com.crm.customer.document.CustomerProfile;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface CustomerProfileRepository extends MongoRepository<CustomerProfile, String> {

    Optional<CustomerProfile> findByCustomerId(Long customerId);

    List<CustomerProfile> findByTagsContaining(String tag);

    @Query("{ 'engagement_score': { $gte: ?0 } }")
    List<CustomerProfile> findByEngagementScoreGreaterThan(Integer score);

    @Query("{ 'lifetime_value': { $gte: ?0 } }")
    List<CustomerProfile> findHighValueCustomers(Double minValue);
}
