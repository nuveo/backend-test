/**
 * 
 */
package com.nuveo.backendtest.api.repository;

import com.nuveo.backendtest.api.entity.Workflow;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * @author rsouza
 *
 */

@Repository
public interface WorkflowRepository extends JpaRepository<Workflow, String > {
	
	Workflow findByUuid(String uuid);

}
