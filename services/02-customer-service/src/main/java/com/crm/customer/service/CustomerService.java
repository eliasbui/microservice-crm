package com.crm.customer.service;

import com.crm.customer.document.CustomerProfile;
import com.crm.customer.entity.Customer;
import com.crm.customer.repository.jpa.CustomerRepository;
import com.crm.customer.repository.mongo.CustomerProfileRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Service
@RequiredArgsConstructor
@Slf4j
public class CustomerService {

    private final CustomerRepository customerRepository;
    private final CustomerProfileRepository profileRepository;

    @Transactional
    public Customer createCustomer(Customer customer) {
        log.info("Creating new customer: {}", customer.getEmail());

        // Save to PostgreSQL
        Customer savedCustomer = customerRepository.save(customer);

        // Create profile in MongoDB
        CustomerProfile profile = CustomerProfile.builder()
            .customerId(savedCustomer.getId())
            .engagementScore(0)
            .lifetimeValue(0.0)
            .createdAt(LocalDateTime.now())
            .updatedAt(LocalDateTime.now())
            .build();

        profileRepository.save(profile);

        log.info("Customer created successfully with ID: {}", savedCustomer.getId());
        return savedCustomer;
    }

    @Transactional(readOnly = true)
    public Optional<Customer> getCustomerById(Long id) {
        return customerRepository.findById(id);
    }

    @Transactional(readOnly = true)
    public Optional<Customer> getCustomerByEmail(String email) {
        return customerRepository.findByEmail(email);
    }

    @Transactional(readOnly = true)
    public Page<Customer> getAllCustomers(Pageable pageable) {
        return customerRepository.findAll(pageable);
    }

    @Transactional(readOnly = true)
    public Page<Customer> searchCustomers(String searchTerm, Pageable pageable) {
        return customerRepository.searchCustomers(searchTerm, pageable);
    }

    @Transactional
    public Customer updateCustomer(Long id, Customer customerDetails) {
        Customer customer = customerRepository.findById(id)
            .orElseThrow(() -> new RuntimeException("Customer not found"));

        customer.setFirstName(customerDetails.getFirstName());
        customer.setLastName(customerDetails.getLastName());
        customer.setEmail(customerDetails.getEmail());
        customer.setPhone(customerDetails.getPhone());
        customer.setCompanyName(customerDetails.getCompanyName());
        customer.setJobTitle(customerDetails.getJobTitle());
        customer.setStatus(customerDetails.getStatus());
        customer.setType(customerDetails.getType());
        customer.setAddress(customerDetails.getAddress());
        customer.setCity(customerDetails.getCity());
        customer.setState(customerDetails.getState());
        customer.setCountry(customerDetails.getCountry());
        customer.setPostalCode(customerDetails.getPostalCode());
        customer.setNotes(customerDetails.getNotes());

        return customerRepository.save(customer);
    }

    @Transactional
    public void deleteCustomer(Long id) {
        customerRepository.deleteById(id);
        profileRepository.findByCustomerId(id).ifPresent(profileRepository::delete);
    }

    @Transactional(readOnly = true)
    public Optional<CustomerProfile> getCustomerProfile(Long customerId) {
        return profileRepository.findByCustomerId(customerId);
    }

    @Transactional
    public CustomerProfile updateCustomerProfile(Long customerId, CustomerProfile profile) {
        CustomerProfile existingProfile = profileRepository.findByCustomerId(customerId)
            .orElseThrow(() -> new RuntimeException("Profile not found"));

        existingProfile.setPreferences(profile.getPreferences());
        existingProfile.setSocialProfiles(profile.getSocialProfiles());
        existingProfile.setTags(profile.getTags());
        existingProfile.setCustomFields(profile.getCustomFields());
        existingProfile.setEngagementScore(profile.getEngagementScore());
        existingProfile.setLifetimeValue(profile.getLifetimeValue());
        existingProfile.setUpdatedAt(LocalDateTime.now());

        return profileRepository.save(existingProfile);
    }

    @Transactional(readOnly = true)
    public List<CustomerProfile> getHighValueCustomers(Double minValue) {
        return profileRepository.findHighValueCustomers(minValue);
    }
}
