/**
 * 
 */
package com.nuveo.backendtest.api.entity;

import java.util.Arrays;
import java.util.UUID;

import javax.validation.constraints.NotEmpty;

import org.hibernate.validator.constraints.NotBlank;
import org.springframework.data.annotation.Id;

import com.nuveo.backendtest.enums.WorkflowStatus;

/**
 * @author rsouza
 *
 */

public class Workflow {

	/** Workflow unique indentifier*/
	@Id
	protected UUID uuid;	
	
	/** workflow Status*/
	@NotEmpty
	protected WorkflowStatus status;

	/** JSONB	workflow input */
	@NotEmpty
	protected String data;
	
	/** name of workflow steps */
	@NotEmpty
	protected String[] steps;

	/**
	 * @return the uuid
	 */
	public UUID getUuid() {
		return uuid;
	}

	/**
	 * @param uuid the uuid to set
	 */
	public void setUuid(UUID uuid) {
		this.uuid = uuid;
	}

	/**
	 * @return the status
	 */
	public WorkflowStatus getStatus() {
		return status;
	}

	/**
	 * @param status the status to set
	 */
	public void setStatus(WorkflowStatus status) {
		this.status = status;
	}

	/**
	 * @return the data
	 */
	public String getData() {
		return data;
	}

	/**
	 * @param data the data to set
	 */
	public void setData(String data) {
		this.data = data;
	}

	/**
	 * @return the steps
	 */
	public String[] getSteps() {
		return steps;
	}

	/**
	 * @param steps the steps to set
	 */
	public void setSteps(String[] steps) {
		this.steps = steps;
	}

	/* (non-Javadoc)
	 * @see java.lang.Object#toString()
	 */
	@Override
	public String toString() {
		return "Workflow [uuid=" + uuid + ", status=" + status + ", data=" + data + ", steps=" + Arrays.toString(steps)
				+ "]";
	}

	public Workflow(WorkflowStatus pStatus, String pData, String[] pSteps) {
		super();
		this.uuid = UUID.randomUUID();
		this.status = pStatus;
		this.data = pData;
		this.steps = pSteps;
	}

	public Workflow(UUID pUUID, WorkflowStatus pStatus, String pData, String[] pSteps) {
		super();
		this.uuid = pUUID;
		this.status = pStatus;
		this.data = pData;
		this.steps = pSteps;
	}
	
	
	

}
