package com.crm.core.repository.mongo;

import com.crm.core.document.CustomerProfile;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface CustomerProfileRepository extends MongoRepository<CustomerProfile, String> {

    Optional<CustomerProfile> findByCustomerId(Long customerId);

    List<CustomerProfile> findBySegment(String segment);

    List<CustomerProfile> findByTagsContaining(String tag);

    @Query("{ 'score': { $gte: ?0 } }")
    List<CustomerProfile> findByScoreGreaterThanEqual(Integer minScore);

    @Query("{ 'preferences.communicationChannel': ?0 }")
    List<CustomerProfile> findByPreferredCommunicationChannel(String channel);

    List<CustomerProfile> findByScoreBetween(Integer minScore, Integer maxScore);

    void deleteByCustomerId(Long customerId);
}
