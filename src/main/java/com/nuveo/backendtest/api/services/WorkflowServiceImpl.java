/**
 * 
 */
package com.nuveo.backendtest.api.services;

import java.io.IOException;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import com.amazonaws.services.sqs.model.Message;
import com.nuveo.backendtest.api.entity.Workflow;
import com.nuveo.backendtest.api.repository.WorkflowRepository;
import com.nuveo.backendtest.helper.sqs.SQSReader;
import com.nuveo.backendtest.helper.util.WorkflowPOJOUtil;

/**
 * @author rsouza
 *
 */

@Service
public class WorkflowServiceImpl implements WorkflowService {
	
	@Autowired
	private WorkflowRepository workflowRepository; 
	
	public Workflow create(Workflow workflow) {
		
		return workflowRepository.save(workflow);
	}

	@Override
	public Workflow update(Workflow workflow) {
		return workflowRepository.save(workflow);
	}
	
	@Override
	public Workflow get(String uuid) {
		return workflowRepository.findByUuid(uuid);
	}

    @Value("${sqs.url.testnuveo}")
	private String sqsUrl;
    
	@Override
	public String consume() {
		try {
			
			SQSReader sqsReader = new SQSReader(sqsUrl);
			
			Message msg = sqsReader.receiveMessage();

			Workflow workflowFromQueue = WorkflowPOJOUtil.deserializeMessage(msg.getBody());
			if(workflowFromQueue != null) {

				sqsReader.deleteMessage(msg.getReceiptHandle());
				
				String csv = WorkflowPOJOUtil.JSON2CSV(workflowFromQueue.getData());

				return csv;
			}
			else return null;

		} catch (IOException e) {
			e.printStackTrace();
			return null;			
		}
	}


	@Override
	public List<Workflow> getAll() {
		return workflowRepository.findAll();
	}

}
