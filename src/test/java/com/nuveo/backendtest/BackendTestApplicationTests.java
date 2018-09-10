package com.nuveo.backendtest;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;
import com.nuveo.backendtest.api.controller.WorkflowController;

import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.boot.test.context.SpringBootTest.WebEnvironment;
import org.springframework.boot.test.web.client.TestRestTemplate;

@RunWith(SpringRunner.class)
@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)
public class BackendTestApplicationTests {

	public BackendTestApplicationTests() {
		
	}
	
	
	@LocalServerPort
	private int port;

	@Autowired
    private WorkflowController controller;

	@Autowired
    private TestRestTemplate restTemplate;
	
	@Test
	public void contextLoads() {
        assertThat(controller).isNotNull();
	}

	@Test
	public void greetingShouldReturnDefaultMessage() throws Exception {
		
		String response = this.restTemplate.getForObject("http://localhost:" + port + "/api/workflow/greetings", String.class); 
		//Object response = this.restTemplate.getForObject("http://localhost:" + port + "/api/workflow/greetings", Object.class); 
		
        assertThat(response).isNotNull();
		
		//assertThat().contains("Hello World");
	}	
}
