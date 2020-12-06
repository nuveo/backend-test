package com.matheuslima.apiworkflow.domain.dto;

import java.util.List;
import java.util.UUID;

import org.modelmapper.ModelMapper;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.domain.enums.WorkFlowStatus;

import lombok.Data;

@Data
public class WorkFlowDTO {

	private UUID uuid;
	private WorkFlowStatus wFS;
	private String data;
	private List<String> workflowSteps;
	
	public static WorkFlowDTO create(WorkFlow wf) {
		ModelMapper modelMapper = new ModelMapper();
		return modelMapper.map(wf, WorkFlowDTO.class);
	}
}
