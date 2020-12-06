package com.matheuslima.apiworkflow.utils;

import java.util.ArrayList;
import java.util.UUID;
import java.util.Vector;

import javax.annotation.PostConstruct;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.domain.enums.WorkFlowStatus;
import com.matheuslima.apiworkflow.repositories.WorkFlowRepository;

@Component
public class MockData {
	
	@Autowired
	WorkFlowRepository  wfr;
	
	@PostConstruct
	public void saveWorkflow() {
		/*WorkFlow wf1 = new WorkFlow();
		
		Vector<String> steps = new Vector<String>(); 
		steps.add("Propose idea");
		steps.add("Create issues");
		steps.add("Implement issues");
		steps.add("Deploy the project");
		steps.add("Track production");
		
		wf1.setUuid(UUID.randomUUID());
		wf1.setWFS(WorkFlowStatus.INSERTED);
		wf1.setWorkflowSteps(new ArrayList<String>(steps));
		Jsonb jsonb = JsonbBuilder.create();
		String result = jsonb.toJson(wf1);
		wf1.setData(result);
		wfr.save(wf1);*/
	}
	
	public static void main(String[] args) {
		String teste = UUID.randomUUID().toString();
		System.out.println(teste);
	}

}
