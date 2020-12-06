package com.matheuslima.apiworkflow.services;

import java.io.File;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.Stack;
import java.util.UUID;

import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;

import org.springframework.stereotype.Service;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.rabbitmq.client.Channel;
import com.rabbitmq.client.DeliverCallback;

@Service
public class WorkFlowConsumerServiceImpl implements WorkFlowConsumerService {

	@Override
	public void writeWorkFlowInFile(Channel channel, String queue, Stack<WorkFlow> listFileCsv) throws IOException {

		StringBuilder sb = new StringBuilder();
		sb.append("DATA,");
		sb.append("STEPS");
		sb.append('\n');
		
		String fileName = UUID.randomUUID().toString().concat(".csv");
		
		PrintWriter writer = new PrintWriter(new File(fileName));

		writer.write(sb.toString());
		
		System.out.println("done!");

		DeliverCallback deliverCallback = (consumerTag, delivery) -> {
			Jsonb jsonb = JsonbBuilder.create();

			listFileCsv.push(jsonb.fromJson(new String(delivery.getBody(), "UTF-8"), WorkFlow.class));
			writer.write(jsonb.fromJson(new String(delivery.getBody(), "UTF-8"), WorkFlow.class).getData());
			writer.write(jsonb.fromJson(new String(delivery.getBody(), "UTF-8"), WorkFlow.class).getWorkflowSteps().toString());
			writer.write('\n');
			System.out.println("Write Line....");
			writer.flush();

		};
		channel.basicConsume(queue, true, deliverCallback, consumerTag -> {
		});

	}

}
