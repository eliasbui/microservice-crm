#!/bin/bash
# CouchDB Database Initialization Script
# Creates databases and design documents for flexible schema integrations

COUCHDB_URL="http://${COUCHDB_USER:-admin}:${COUCHDB_PASSWORD:-password}@localhost:5984"

# Wait for CouchDB to be ready
echo "Waiting for CouchDB to be ready..."
until curl -f "${COUCHDB_URL}/_up" > /dev/null 2>&1; do
    sleep 2
done
echo "CouchDB is ready!"

# Create databases
echo "Creating databases..."
curl -X PUT "${COUCHDB_URL}/integration_configs"
curl -X PUT "${COUCHDB_URL}/flexible_entities"
curl -X PUT "${COUCHDB_URL}/document_store"
curl -X PUT "${COUCHDB_URL}/audit_logs"

# Create design documents

# Integration Configs - Views
echo "Creating design documents for integration_configs..."
curl -X PUT "${COUCHDB_URL}/integration_configs/_design/queries" -H "Content-Type: application/json" -d '{
  "views": {
    "by_type": {
      "map": "function(doc) { if (doc.integration_type) { emit(doc.integration_type, doc); } }"
    },
    "by_provider": {
      "map": "function(doc) { if (doc.provider) { emit(doc.provider, doc); } }"
    },
    "active_integrations": {
      "map": "function(doc) { if (doc.is_active === true) { emit(doc._id, doc); } }"
    }
  }
}'

# Flexible Entities - Views
echo "Creating design documents for flexible_entities..."
curl -X PUT "${COUCHDB_URL}/flexible_entities/_design/queries" -H "Content-Type: application/json" -d '{
  "views": {
    "by_entity_type": {
      "map": "function(doc) { if (doc.entity_type) { emit(doc.entity_type, doc); } }"
    },
    "by_source_system": {
      "map": "function(doc) { if (doc.source_system) { emit(doc.source_system, doc); } }"
    },
    "by_created_date": {
      "map": "function(doc) { if (doc.created_at) { emit(doc.created_at, doc); } }"
    },
    "by_tags": {
      "map": "function(doc) { if (doc.tags && Array.isArray(doc.tags)) { doc.tags.forEach(function(tag) { emit(tag, doc); }); } }"
    }
  }
}'

# Document Store - Views
echo "Creating design documents for document_store..."
curl -X PUT "${COUCHDB_URL}/document_store/_design/queries" -H "Content-Type: application/json" -d '{
  "views": {
    "by_customer_id": {
      "map": "function(doc) { if (doc.customer_id) { emit(doc.customer_id, doc); } }"
    },
    "by_document_type": {
      "map": "function(doc) { if (doc.document_type) { emit(doc.document_type, doc); } }"
    },
    "recent_documents": {
      "map": "function(doc) { if (doc.uploaded_at) { emit(doc.uploaded_at, doc); } }"
    }
  }
}'

# Audit Logs - Views
echo "Creating design documents for audit_logs..."
curl -X PUT "${COUCHDB_URL}/audit_logs/_design/queries" -H "Content-Type: application/json" -d '{
  "views": {
    "by_user": {
      "map": "function(doc) { if (doc.user_id) { emit(doc.user_id, doc); } }"
    },
    "by_action": {
      "map": "function(doc) { if (doc.action) { emit(doc.action, doc); } }"
    },
    "by_entity": {
      "map": "function(doc) { if (doc.entity_type && doc.entity_id) { emit([doc.entity_type, doc.entity_id], doc); } }"
    },
    "by_timestamp": {
      "map": "function(doc) { if (doc.timestamp) { emit(doc.timestamp, doc); } }"
    }
  }
}'

# Insert sample documents

# Sample Integration Config
echo "Inserting sample integration config..."
curl -X POST "${COUCHDB_URL}/integration_configs" -H "Content-Type: application/json" -d '{
  "integration_type": "REST_API",
  "provider": "Custom ERP",
  "config": {
    "base_url": "https://api.example.com",
    "auth_type": "bearer",
    "endpoints": {
      "customers": "/api/v1/customers",
      "orders": "/api/v1/orders"
    },
    "rate_limit": {
      "requests_per_minute": 100
    }
  },
  "is_active": true,
  "created_at": "'$(date -u +"%Y-%m-%dT%H:%M:%SZ")'"
}'

# Sample Flexible Entity
echo "Inserting sample flexible entity..."
curl -X POST "${COUCHDB_URL}/flexible_entities" -H "Content-Type: application/json" -d '{
  "entity_type": "custom_product",
  "source_system": "legacy_crm",
  "data": {
    "product_name": "Sample Product",
    "custom_field_1": "Custom Value",
    "nested_data": {
      "key1": "value1",
      "key2": "value2"
    }
  },
  "tags": ["imported", "legacy"],
  "created_at": "'$(date -u +"%Y-%m-%dT%H:%M:%SZ")'"
}'

echo "CouchDB initialization completed successfully!"
