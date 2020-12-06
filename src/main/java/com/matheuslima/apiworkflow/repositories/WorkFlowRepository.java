package com.matheuslima.apiworkflow.repositories;

import java.util.Optional;
import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;

import com.matheuslima.apiworkflow.domain.WorkFlow;

public interface WorkFlowRepository extends JpaRepository<WorkFlow, UUID>{

	Optional<WorkFlow> findByUuid(UUID uuid);

}
