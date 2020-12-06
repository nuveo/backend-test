package com.matheuslima.apiworkflow.services;

import org.springframework.amqp.core.AmqpTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import com.matheuslima.apiworkflow.domain.WorkFlow;

@Service
public class RabbitMQSenderService {
	
	@Autowired
	private AmqpTemplate rabbitTemplate;
	
	@Value("${matheuslima.rabbitmq.exchange}")
	private String exchange;
	
	@Value("${matheuslima.rabbitmq.routingkey}")
	private String routingkey;	
	
	public void send(WorkFlow wf) {
		rabbitTemplate.convertAndSend(exchange, routingkey, wf);
		System.out.println("Send msg = " + wf);
	    
	}
}