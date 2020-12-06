package com.matheuslima.apiworkflow.services;

import com.matheuslima.apiworkflow.domain.WorkFlow;

public interface WorkFlowProducerService {
	
	void send(WorkFlow wf);

}
