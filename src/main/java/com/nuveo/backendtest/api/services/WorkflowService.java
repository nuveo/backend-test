/**
 * 
 */
package com.nuveo.backendtest.api.services;

import java.util.List;

import com.nuveo.backendtest.api.entity.Workflow;

/**
 * @author rsouza
 *
 */

public interface WorkflowService {
	
	Workflow create(Workflow workflow);

	String consume();

	Workflow get(String uUID);

	List<Workflow> getAll();

	Workflow update(Workflow workflow);

}
