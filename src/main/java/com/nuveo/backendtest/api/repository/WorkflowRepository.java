/**
 * 
 */
package com.nuveo.backendtest.api.repository;

import java.util.UUID;

import com.nuveo.backendtest.api.entity.Workflow;

/**
 * @author rsouza
 *
 */

public interface WorkflowRepository {
	
	Workflow findByUuid(UUID uuid);

}
