package com.crm.core.repository;

import com.crm.core.entity.Opportunity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.util.List;

@Repository
public interface OpportunityRepository extends JpaRepository<Opportunity, Long> {

    List<Opportunity> findByCustomerId(Long customerId);

    List<Opportunity> findByStage(Opportunity.OpportunityStage stage);

    List<Opportunity> findByStatus(Opportunity.OpportunityStatus status);

    List<Opportunity> findByAssignedTo(String assignedTo);

    @Query("SELECT o FROM Opportunity o WHERE o.expectedCloseDate BETWEEN :startDate AND :endDate")
    List<Opportunity> findByExpectedCloseDateBetween(@Param("startDate") LocalDate startDate,
                                                      @Param("endDate") LocalDate endDate);

    @Query("SELECT SUM(o.amount) FROM Opportunity o WHERE o.status = :status")
    BigDecimal sumAmountByStatus(@Param("status") Opportunity.OpportunityStatus status);

    @Query("SELECT COUNT(o) FROM Opportunity o WHERE o.stage = :stage")
    long countByStage(@Param("stage") Opportunity.OpportunityStage stage);
}
