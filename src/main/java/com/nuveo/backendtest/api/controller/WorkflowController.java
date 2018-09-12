/**
 * 
 */
package com.nuveo.backendtest.api.controller;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DuplicateKeyException;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.BindingResult;
import org.springframework.validation.ObjectError;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import com.nuveo.backendtest.api.services.WorkflowService;
import com.nuveo.backendtest.api.entity.Workflow;
import com.nuveo.backendtest.api.response.Response;
import java.util.List;
import java.util.UUID;

/**
 * @author rsouza
 *
 */

@RestController
@CrossOrigin(origins = "*")
public class WorkflowController  {
	
	@Autowired
	private WorkflowService service;
	
	@GetMapping("/api/workflow/greetings")
	public ResponseEntity<Response<String>> greetings() {

		Response<String> response = new Response<String>();
		response.setData("Hello World");
		return ResponseEntity.ok(response);
	}
		
	@PostMapping("/api/workflow")
	public ResponseEntity<Response<Workflow>> create(HttpServletRequest request, @RequestBody Workflow workflow, BindingResult result) {
		
		Response<Workflow> response = new Response<Workflow>();

		try {
			validateWorkflow(workflow, result);
			
			if(result.hasErrors()) {

				result.getAllErrors().forEach(error -> response.getErrors().add(error.getDefaultMessage()));

				return ResponseEntity.badRequest().body(response);
			}
			
			Workflow workflowPersisted = service.create(workflow);
			
			response.setData(workflowPersisted);
			
		} catch (DuplicateKeyException dke) {
			response.getErrors().add("Workflow already inserted");
			return ResponseEntity.badRequest().body(response);
		} catch (Exception ex) {
			return ResponseEntity.badRequest().body(response);
		}
		
		return ResponseEntity.ok(response);
	}

	@GetMapping("/api/workflows")
	public ResponseEntity<Response<List<Workflow>>> getAll() {
		
		Response<List<Workflow>> response = new Response<List<Workflow>>();

		try {
			response.setData(service.getAll());
		} catch (Exception dke) {
			return ResponseEntity.badRequest().body(response);
		}
		
		return ResponseEntity.ok(response);
	}

	public void validateWorkflow(Workflow workflow, BindingResult result) {

		if(workflow.getData() == null) {
			result.addError(new ObjectError("Workflow", "Empty workflow"));
		}

		if(workflow.getUuid() == null || !workflow.getUuid().isEmpty())
		{
			workflow.setUuid(UUID.randomUUID().toString());
		}
	}
	
	
	@PatchMapping("/api/workflow/{UUID}")
	public ResponseEntity<Response<Workflow>> update(@PathVariable String UUID, @RequestBody Workflow workflow, BindingResult result) {
		
		Response<Workflow> response = new Response<Workflow>();

		try {
			validateWorkflow(workflow, result);
			
			if(result.hasErrors()) {

				result.getAllErrors().forEach(error -> response.getErrors().add(error.getDefaultMessage()));

				return ResponseEntity.badRequest().body(response);
			}
			
			Workflow workflowPersisted = service.update(workflow);
			
			response.setData(workflowPersisted);
			
		} catch (Exception dke) {
			return ResponseEntity.badRequest().body(response);
		}
		
		return ResponseEntity.ok(response);
	}

	@GetMapping("/api/workflow/{UUID}")
	public ResponseEntity<Response<Workflow>> get(@PathVariable String UUID) {
		
		Response<Workflow> response = new Response<Workflow>();

		try {

			Workflow workflowFromQueue = service.get(UUID);
			
			response.setData(workflowFromQueue);
			
		} catch (Exception dke) {
			return ResponseEntity.badRequest().body(response);
		}
		
		return ResponseEntity.ok(response);
	}

	@GetMapping("/api/workflow/consume")
	public ResponseEntity<Response<String>> getFromQueue(HttpServletRequest request) {
		
		Response<String> response = new Response<String>();

		try {

			String objSqsAsCSV = service.consume();

			response.setData(objSqsAsCSV);

		} catch (Exception dke) {
			return ResponseEntity.badRequest().body(response);
		}
		
		return ResponseEntity.ok(response);
	}
	
	
}
