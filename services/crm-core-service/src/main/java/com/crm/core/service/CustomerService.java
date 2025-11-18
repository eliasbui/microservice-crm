package com.crm.core.service;

import com.crm.core.dto.CustomerDto;
import com.crm.core.entity.Customer;
import com.crm.core.repository.CustomerRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
@Slf4j
@Transactional
public class CustomerService {

    private final CustomerRepository customerRepository;

    @Cacheable(value = "customers", key = "#id")
    public Optional<CustomerDto> getCustomerById(Long id) {
        log.info("Fetching customer with id: {}", id);
        return customerRepository.findById(id).map(this::convertToDto);
    }

    public List<CustomerDto> getAllCustomers() {
        log.info("Fetching all customers");
        return customerRepository.findAll().stream()
                .map(this::convertToDto)
                .collect(Collectors.toList());
    }

    public Page<CustomerDto> searchCustomers(String searchTerm, Pageable pageable) {
        log.info("Searching customers with term: {}", searchTerm);
        return customerRepository.searchCustomers(searchTerm, pageable)
                .map(this::convertToDto);
    }

    public List<CustomerDto> getCustomersByStatus(Customer.CustomerStatus status) {
        log.info("Fetching customers with status: {}", status);
        return customerRepository.findByStatus(status).stream()
                .map(this::convertToDto)
                .collect(Collectors.toList());
    }

    @CacheEvict(value = "customers", allEntries = true)
    public CustomerDto createCustomer(CustomerDto customerDto) {
        log.info("Creating new customer: {}", customerDto.getEmail());

        // Check if customer with email already exists
        if (customerRepository.findByEmail(customerDto.getEmail()).isPresent()) {
            throw new IllegalArgumentException("Customer with email " + customerDto.getEmail() + " already exists");
        }

        Customer customer = convertToEntity(customerDto);
        Customer savedCustomer = customerRepository.save(customer);
        log.info("Customer created with id: {}", savedCustomer.getId());
        return convertToDto(savedCustomer);
    }

    @CacheEvict(value = "customers", key = "#id")
    public Optional<CustomerDto> updateCustomer(Long id, CustomerDto customerDto) {
        log.info("Updating customer with id: {}", id);
        return customerRepository.findById(id).map(customer -> {
            updateCustomerFields(customer, customerDto);
            Customer updatedCustomer = customerRepository.save(customer);
            log.info("Customer updated with id: {}", id);
            return convertToDto(updatedCustomer);
        });
    }

    @CacheEvict(value = "customers", key = "#id")
    public boolean deleteCustomer(Long id) {
        log.info("Deleting customer with id: {}", id);
        return customerRepository.findById(id).map(customer -> {
            customerRepository.delete(customer);
            log.info("Customer deleted with id: {}", id);
            return true;
        }).orElse(false);
    }

    @CacheEvict(value = "customers", key = "#id")
    public Optional<CustomerDto> updateLastContacted(Long id) {
        log.info("Updating last contacted time for customer: {}", id);
        return customerRepository.findById(id).map(customer -> {
            customer.setLastContactedAt(LocalDateTime.now());
            return convertToDto(customerRepository.save(customer));
        });
    }

    private CustomerDto convertToDto(Customer customer) {
        return CustomerDto.builder()
                .id(customer.getId())
                .firstName(customer.getFirstName())
                .lastName(customer.getLastName())
                .email(customer.getEmail())
                .phone(customer.getPhone())
                .company(customer.getCompany())
                .jobTitle(customer.getJobTitle())
                .address(customer.getAddress())
                .city(customer.getCity())
                .state(customer.getState())
                .zipCode(customer.getZipCode())
                .country(customer.getCountry())
                .status(customer.getStatus())
                .notes(customer.getNotes())
                .assignedTo(customer.getAssignedTo())
                .createdAt(customer.getCreatedAt())
                .updatedAt(customer.getUpdatedAt())
                .lastContactedAt(customer.getLastContactedAt())
                .build();
    }

    private Customer convertToEntity(CustomerDto dto) {
        return Customer.builder()
                .firstName(dto.getFirstName())
                .lastName(dto.getLastName())
                .email(dto.getEmail())
                .phone(dto.getPhone())
                .company(dto.getCompany())
                .jobTitle(dto.getJobTitle())
                .address(dto.getAddress())
                .city(dto.getCity())
                .state(dto.getState())
                .zipCode(dto.getZipCode())
                .country(dto.getCountry())
                .status(dto.getStatus() != null ? dto.getStatus() : Customer.CustomerStatus.LEAD)
                .notes(dto.getNotes())
                .assignedTo(dto.getAssignedTo())
                .build();
    }

    private void updateCustomerFields(Customer customer, CustomerDto dto) {
        if (dto.getFirstName() != null) customer.setFirstName(dto.getFirstName());
        if (dto.getLastName() != null) customer.setLastName(dto.getLastName());
        if (dto.getEmail() != null) customer.setEmail(dto.getEmail());
        if (dto.getPhone() != null) customer.setPhone(dto.getPhone());
        if (dto.getCompany() != null) customer.setCompany(dto.getCompany());
        if (dto.getJobTitle() != null) customer.setJobTitle(dto.getJobTitle());
        if (dto.getAddress() != null) customer.setAddress(dto.getAddress());
        if (dto.getCity() != null) customer.setCity(dto.getCity());
        if (dto.getState() != null) customer.setState(dto.getState());
        if (dto.getZipCode() != null) customer.setZipCode(dto.getZipCode());
        if (dto.getCountry() != null) customer.setCountry(dto.getCountry());
        if (dto.getStatus() != null) customer.setStatus(dto.getStatus());
        if (dto.getNotes() != null) customer.setNotes(dto.getNotes());
        if (dto.getAssignedTo() != null) customer.setAssignedTo(dto.getAssignedTo());
    }
}
