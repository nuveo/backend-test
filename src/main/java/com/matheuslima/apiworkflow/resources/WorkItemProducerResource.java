package com.matheuslima.apiworkflow.resources;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.domain.dto.WorkFlowDTO;
import com.matheuslima.apiworkflow.services.RabbitMQSenderService;
import com.matheuslima.apiworkflow.services.WorkFlowService;

@RestController
@RequestMapping("/api/v1/")
public class WorkItemProducerResource {

	@Autowired
	private WorkFlowService wfs;
	
	@Autowired
	RabbitMQSenderService rabbitMQSender;
	
	@PostMapping("/workflow")
	public ResponseEntity post(@RequestBody WorkFlow wf) { 
		

		//rabbitMQSender.send(wf);
		
		try {
			//WorkFlowDTO wfe = wfs.save(wf);
			return ResponseEntity.ok(wfs.save(wf));
			
		}catch (Exception e) {
			return ResponseEntity.badRequest().build();
		}
	}
}
