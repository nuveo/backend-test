package com.matheuslima.apiworkflow.services;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.domain.dto.WorkFlowDTO;

public interface WorkFlowService {

	public List<WorkFlowDTO> findAll();
	
	public Optional<WorkFlowDTO> findByUuid(UUID uuid);
	
	public WorkFlowDTO save(WorkFlow wf);
}
