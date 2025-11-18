package com.crm.customer.controller;

import com.crm.customer.document.CustomerProfile;
import com.crm.customer.entity.Customer;
import com.crm.customer.service.CustomerService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/customers")
@RequiredArgsConstructor
public class CustomerController {

    private final CustomerService customerService;

    @PostMapping
    public ResponseEntity<Customer> createCustomer(@RequestBody Customer customer) {
        Customer created = customerService.createCustomer(customer);
        return ResponseEntity.status(HttpStatus.CREATED).body(created);
    }

    @GetMapping("/{id}")
    public ResponseEntity<Customer> getCustomer(@PathVariable Long id) {
        return customerService.getCustomerById(id)
            .map(ResponseEntity::ok)
            .orElse(ResponseEntity.notFound().build());
    }

    @GetMapping
    public ResponseEntity<Page<Customer>> getAllCustomers(Pageable pageable) {
        return ResponseEntity.ok(customerService.getAllCustomers(pageable));
    }

    @GetMapping("/search")
    public ResponseEntity<Page<Customer>> searchCustomers(
        @RequestParam String q,
        Pageable pageable
    ) {
        return ResponseEntity.ok(customerService.searchCustomers(q, pageable));
    }

    @PutMapping("/{id}")
    public ResponseEntity<Customer> updateCustomer(
        @PathVariable Long id,
        @RequestBody Customer customer
    ) {
        Customer updated = customerService.updateCustomer(id, customer);
        return ResponseEntity.ok(updated);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteCustomer(@PathVariable Long id) {
        customerService.deleteCustomer(id);
        return ResponseEntity.noContent().build();
    }

    @GetMapping("/{id}/profile")
    public ResponseEntity<CustomerProfile> getCustomerProfile(@PathVariable Long id) {
        return customerService.getCustomerProfile(id)
            .map(ResponseEntity::ok)
            .orElse(ResponseEntity.notFound().build());
    }

    @PutMapping("/{id}/profile")
    public ResponseEntity<CustomerProfile> updateCustomerProfile(
        @PathVariable Long id,
        @RequestBody CustomerProfile profile
    ) {
        CustomerProfile updated = customerService.updateCustomerProfile(id, profile);
        return ResponseEntity.ok(updated);
    }

    @GetMapping("/high-value")
    public ResponseEntity<List<CustomerProfile>> getHighValueCustomers(
        @RequestParam(defaultValue = "10000") Double minValue
    ) {
        return ResponseEntity.ok(customerService.getHighValueCustomers(minValue));
    }

    @GetMapping("/health")
    public ResponseEntity<Object> health() {
        return ResponseEntity.ok(java.util.Map.of(
            "status", "healthy",
            "service", "customer-service",
            "timestamp", java.time.LocalDateTime.now()
        ));
    }
}
