package com.matheuslima.apiworkflow.resources;

import java.sql.SQLException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.domain.dto.WorkFlowDTO;
import com.matheuslima.apiworkflow.services.WorkFlowProducerService;
import com.matheuslima.apiworkflow.services.WorkFlowService;

@RestController
@RequestMapping("/api/v1/")
public class WorkItemProducerResource {

	@Autowired
	private WorkFlowService wfs;
	
	@Autowired
	WorkFlowProducerService rabbitMQSender;
	
	@PostMapping("/workflow")
	public ResponseEntity<WorkFlowDTO> post(@RequestBody WorkFlow wf) { 
		
		try {
			WorkFlowDTO wfd = wfs.save(wf);
			rabbitMQSender.send(wf);
			
			return ResponseEntity.ok(wfd);
			
		}catch (Exception e) {
			return ResponseEntity.badRequest().build();
		}
	}
	
	@PatchMapping(value = "/workflow/{UUID}")
	public ResponseEntity<?> partialUpdateName(
	  @RequestBody WorkFlowDTO partialUpdate, @PathVariable("UUID") String uuid) {
	    
		try {
			wfs.save(partialUpdate, uuid);
		} catch (ClassNotFoundException | SQLException e) {
			e.printStackTrace();
		}
	    return ResponseEntity.ok("resource address updated new status, ".concat(uuid));
	}
}
