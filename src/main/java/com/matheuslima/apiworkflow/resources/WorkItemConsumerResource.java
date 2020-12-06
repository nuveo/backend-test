package com.matheuslima.apiworkflow.resources;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeoutException;

import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;

import org.springframework.amqp.rabbit.annotation.RabbitHandler;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.domain.dto.WorkFlowDTO;
import com.matheuslima.apiworkflow.services.WorkFlowConsumerService;
import com.matheuslima.apiworkflow.services.WorkFlowService;
import com.opencsv.exceptions.CsvDataTypeMismatchException;
import com.opencsv.exceptions.CsvRequiredFieldEmptyException;
import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;
import com.rabbitmq.client.DeliverCallback;

@RestController
@RequestMapping("/api/v1/")
public class WorkItemConsumerResource {

	@Autowired
	private WorkFlowService wfs;

	@Autowired
	private WorkFlowConsumerService wfcs;

	@GetMapping("/workflow")
	public ResponseEntity<List<WorkFlowDTO>> get() {
		List<WorkFlowDTO> wfd = wfs.findAll();
		return ResponseEntity.ok(wfd);
	}

	@Value("${matheuslima.rabbitmq.queue}")
	private String queue;

	@GetMapping("/workflow/consume")
	@RabbitHandler
	public void consume() throws IOException, TimeoutException {
		ConnectionFactory factory = new ConnectionFactory();
		factory.setHost("localhost");
		Connection connection = factory.newConnection();
		Channel channel = connection.createChannel();

		channel.queueDeclare(queue, false, false, false, null);
		System.out.println(" [*] Waiting for messages. To exit press CTRL+C");

		List<WorkFlow> listFileCsv = new ArrayList<>();

		DeliverCallback deliverCallback = (consumerTag, delivery) -> {
			Jsonb jsonb = JsonbBuilder.create();
			try {
				
				listFileCsv.add(jsonb.fromJson(new String(delivery.getBody(), "UTF-8"), WorkFlow.class));
				
			} finally {
				channel.basicAck(delivery.getEnvelope().getDeliveryTag(), false);
			}
			System.out.println("Finished add files");

		};
			channel.basicConsume(queue, true, deliverCallback, consumerTag -> {
			});


		if (!listFileCsv.isEmpty()) {
			try {
				wfcs.writeWorkFlowInFile(listFileCsv);
			} catch (IOException | CsvRequiredFieldEmptyException | CsvDataTypeMismatchException e) {
				e.printStackTrace();
			}
		}
	}

}
