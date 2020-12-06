package com.matheuslima.apiworkflow.services;

import java.io.IOException;
import java.util.Stack;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.rabbitmq.client.Channel;

public interface WorkFlowConsumerService {
	
	void writeWorkFlowInFile(Channel channel, String queue, Stack<WorkFlow> listFileCsv) throws IOException;
	
}
