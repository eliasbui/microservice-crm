package com.crm.core.document;

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

/**
 * MongoDB Document for Customer Profile
 * Stores rich customer profile data with flexible schema
 */
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

    @Field("interactions")
    private List<Interaction> interactions;

    @Field("tags")
    private List<String> tags;

    @Field("custom_fields")
    private Map<String, Object> customFields;

    @Field("score")
    private Integer score; // Customer engagement score

    @Field("segment")
    private String segment; // Customer segment (VIP, Regular, etc.)

    @Field("created_at")
    private LocalDateTime createdAt;

    @Field("updated_at")
    private LocalDateTime updatedAt;

    @Data
    @Builder
    @NoArgsConstructor
    @AllArgsConstructor
    public static class CustomerPreferences {
        private String communicationChannel; // email, sms, phone
        private String preferredLanguage;
        private List<String> interests;
        private Map<String, Boolean> notifications;
        private String timezone;
    }

    @Data
    @Builder
    @NoArgsConstructor
    @AllArgsConstructor
    public static class SocialProfile {
        private String platform; // linkedin, twitter, facebook
        private String profileUrl;
        private Integer followers;
        private LocalDateTime lastUpdated;
    }

    @Data
    @Builder
    @NoArgsConstructor
    @AllArgsConstructor
    public static class Interaction {
        private String type; // email, call, meeting, chat
        private String description;
        private LocalDateTime timestamp;
        private String outcome;
        private Map<String, Object> metadata;
    }
}
