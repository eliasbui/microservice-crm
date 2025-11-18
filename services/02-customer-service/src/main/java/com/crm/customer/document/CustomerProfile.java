package com.crm.customer.document;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;

@Document(collection = "customer_profiles")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CustomerProfile {

    @Id
    private String id;

    @Field("customer_id")
    private Long customerId;

    @Field("preferences")
    private CustomerPreferences preferences;

    @Field("social_profiles")
    private List<SocialProfile> socialProfiles;

    @Field("tags")
    private List<String> tags;

    @Field("custom_fields")
    private Map<String, Object> customFields;

    @Field("engagement_score")
    private Integer engagementScore;

    @Field("lifetime_value")
    private Double lifetimeValue;

    @Field("last_interaction")
    private LocalDateTime lastInteraction;

    @Field("created_at")
    private LocalDateTime createdAt;

    @Field("updated_at")
    private LocalDateTime updatedAt;

    @Data
    @NoArgsConstructor
    @AllArgsConstructor
    public static class CustomerPreferences {
        private String communicationChannel;
        private String language;
        private String timezone;
        private Boolean emailNotifications;
        private Boolean smsNotifications;
        private Boolean pushNotifications;
    }

    @Data
    @NoArgsConstructor
    @AllArgsConstructor
    public static class SocialProfile {
        private String platform;
        private String profileUrl;
        private String username;
    }
}
