// MongoDB Replica Set Initialization
// This script initializes the replica set for high availability

// Wait for MongoDB to be ready
sleep(5000);

rs.initiate({
  _id: "rs0",
  members: [
    { _id: 0, host: "mongodb-primary:27017", priority: 2 },
    { _id: 1, host: "mongodb-secondary1:27017", priority: 1 },
    { _id: 2, host: "mongodb-secondary2:27017", priority: 1 }
  ]
});

// Wait for replica set to be ready
sleep(10000);

// Switch to CRM database
db = db.getSiblingDB('crm_db');

// Create collections with validators

// Customer Profiles Collection
db.createCollection("customer_profiles", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["customer_id", "created_at"],
      properties: {
        customer_id: {
          bsonType: "long",
          description: "Customer ID must be a long and is required"
        },
        preferences: {
          bsonType: "object",
          properties: {
            communicationChannel: { bsonType: "string" },
            preferredLanguage: { bsonType: "string" },
            interests: { bsonType: "array" },
            notifications: { bsonType: "object" },
            timezone: { bsonType: "string" }
          }
        },
        social_profiles: {
          bsonType: "array"
        },
        interactions: {
          bsonType: "array"
        },
        tags: {
          bsonType: "array",
          items: { bsonType: "string" }
        },
        custom_fields: {
          bsonType: "object"
        },
        score: {
          bsonType: "int",
          minimum: 0,
          maximum: 100
        },
        segment: {
          bsonType: "string",
          enum: ["VIP", "Regular", "New", "Inactive", "Churned"]
        },
        created_at: {
          bsonType: "date",
          description: "Creation timestamp is required"
        },
        updated_at: {
          bsonType: "date"
        }
      }
    }
  }
});

// Chat Messages Collection
db.createCollection("chat_messages", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["session_id", "content", "role", "timestamp"],
      properties: {
        session_id: {
          bsonType: "string",
          description: "Session ID is required"
        },
        user_id: {
          bsonType: "string"
        },
        content: {
          bsonType: "string",
          description: "Message content is required"
        },
        role: {
          bsonType: "string",
          enum: ["user", "assistant", "system"],
          description: "Role must be one of: user, assistant, system"
        },
        timestamp: {
          bsonType: "date",
          description: "Timestamp is required"
        },
        metadata: {
          bsonType: "object"
        }
      }
    }
  }
});

// Chat Sessions Collection
db.createCollection("chat_sessions", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["user_id", "started_at"],
      properties: {
        user_id: {
          bsonType: "string",
          description: "User ID is required"
        },
        customer_id: {
          bsonType: "long"
        },
        started_at: {
          bsonType: "date",
          description: "Start timestamp is required"
        },
        ended_at: {
          bsonType: "date"
        },
        is_active: {
          bsonType: "bool"
        }
      }
    }
  }
});

// Notifications Collection
db.createCollection("notifications", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["notification_type", "recipient", "content", "created_at"],
      properties: {
        notification_type: {
          bsonType: "string",
          enum: ["EMAIL", "SMS", "PUSH"],
          description: "Notification type is required"
        },
        recipient: {
          bsonType: "string",
          description: "Recipient is required"
        },
        subject: {
          bsonType: "string"
        },
        content: {
          bsonType: "string",
          description: "Content is required"
        },
        status: {
          bsonType: "string",
          enum: ["PENDING", "SENT", "FAILED", "DELIVERED"]
        },
        metadata: {
          bsonType: "object"
        },
        sent_at: {
          bsonType: "date"
        },
        delivered_at: {
          bsonType: "date"
        },
        created_at: {
          bsonType: "date",
          description: "Creation timestamp is required"
        }
      }
    }
  }
});

// Create indexes for better query performance

// Customer Profiles indexes
db.customer_profiles.createIndex({ "customer_id": 1 }, { unique: true });
db.customer_profiles.createIndex({ "segment": 1 });
db.customer_profiles.createIndex({ "tags": 1 });
db.customer_profiles.createIndex({ "score": -1 });
db.customer_profiles.createIndex({ "preferences.communicationChannel": 1 });
db.customer_profiles.createIndex({ "created_at": -1 });

// Chat Messages indexes
db.chat_messages.createIndex({ "session_id": 1, "timestamp": -1 });
db.chat_messages.createIndex({ "user_id": 1 });
db.chat_messages.createIndex({ "timestamp": -1 });

// Chat Sessions indexes
db.chat_sessions.createIndex({ "user_id": 1, "started_at": -1 });
db.chat_sessions.createIndex({ "customer_id": 1 });
db.chat_sessions.createIndex({ "is_active": 1 });

// Notifications indexes
db.notifications.createIndex({ "recipient": 1, "created_at": -1 });
db.notifications.createIndex({ "status": 1 });
db.notifications.createIndex({ "notification_type": 1 });
db.notifications.createIndex({ "created_at": -1 });

print("MongoDB initialization completed successfully!");
