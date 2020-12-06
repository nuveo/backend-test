package com.matheuslima.apiworkflow.resources;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.matheuslima.apiworkflow.domain.dto.WorkFlowDTO;
import com.matheuslima.apiworkflow.services.WorkFlowService;

@RestController
@RequestMapping("/api/v1/")
public class WorkItemConsumerResource {
	
	@Autowired
	private WorkFlowService wfs;
	
	@GetMapping("/workflow")
	public ResponseEntity get() {
		List<WorkFlowDTO> wfd = wfs.findAll();
		return ResponseEntity.ok(wfd);
	}

}

