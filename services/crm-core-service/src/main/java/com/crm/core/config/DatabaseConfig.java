package com.crm.core.config;

import com.datastax.oss.driver.api.core.CqlSession;
import com.mongodb.ConnectionString;
import com.mongodb.MongoClientSettings;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.data.cassandra.config.AbstractCassandraConfiguration;
import org.springframework.data.cassandra.config.CqlSessionFactoryBean;
import org.springframework.data.cassandra.repository.config.EnableCassandraRepositories;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import org.springframework.data.mongodb.config.AbstractMongoClientConfiguration;
import org.springframework.data.mongodb.repository.config.EnableMongoRepositories;

import java.net.InetSocketAddress;

@Configuration
public class DatabaseConfig {

    @Value("${spring.data.mongodb.uri}")
    private String mongoUri;

    @Value("${spring.data.mongodb.database}")
    private String mongoDatabase;

    @Value("${spring.data.cassandra.contact-points}")
    private String cassandraContactPoints;

    @Value("${spring.data.cassandra.port}")
    private int cassandraPort;

    @Value("${spring.data.cassandra.keyspace-name}")
    private String cassandraKeyspace;

    @Value("${spring.data.cassandra.local-datacenter}")
    private String cassandraDatacenter;

    /**
     * PostgreSQL Configuration (Primary Database)
     * Configured in application.yml via spring.datasource
     */
    @Configuration
    @EnableJpaRepositories(
        basePackages = "com.crm.core.repository.jpa",
        entityManagerFactoryRef = "entityManagerFactory",
        transactionManagerRef = "transactionManager"
    )
    public static class PostgresConfig {
        // PostgreSQL is configured via default Spring Boot auto-configuration
    }

    /**
     * MongoDB Configuration (Customer Profiles)
     */
    @Configuration
    @EnableMongoRepositories(basePackages = "com.crm.core.repository.mongo")
    public class MongoConfig extends AbstractMongoClientConfiguration {

        @Override
        protected String getDatabaseName() {
            return mongoDatabase;
        }

        @Override
        @Bean
        @Primary
        public MongoClient mongoClient() {
            ConnectionString connectionString = new ConnectionString(mongoUri);
            MongoClientSettings mongoClientSettings = MongoClientSettings.builder()
                .applyConnectionString(connectionString)
                .build();
            return MongoClients.create(mongoClientSettings);
        }
    }

    /**
     * Cassandra Configuration (Event Logs & Time-Series)
     */
    @Configuration
    @EnableCassandraRepositories(basePackages = "com.crm.core.repository.cassandra")
    public class CassandraConfig extends AbstractCassandraConfiguration {

        @Override
        protected String getKeyspaceName() {
            return cassandraKeyspace;
        }

        @Override
        protected String getContactPoints() {
            return cassandraContactPoints;
        }

        @Override
        protected int getPort() {
            return cassandraPort;
        }

        @Override
        protected String getLocalDataCenter() {
            return cassandraDatacenter;
        }

        @Bean
        @Override
        public CqlSessionFactoryBean cassandraSession() {
            CqlSessionFactoryBean cassandraSession = super.cassandraSession();
            cassandraSession.setKeyspaceName(getKeyspaceName());
            return cassandraSession;
        }
    }
}
