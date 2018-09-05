/**
 * 
 */
package com.nuveo.backendtest.api.services;

import java.util.UUID;

import com.nuveo.backendtest.api.entity.Workflow;

/**
 * @author rsouza
 *
 */

public interface WorkflowService {
	
	Workflow findByUuid(UUID uuid);

}
