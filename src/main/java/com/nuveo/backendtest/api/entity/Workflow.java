/**
 * 
 */
package com.nuveo.backendtest.api.entity;

import java.util.List;
import java.util.UUID;

import javax.persistence.*;

import javax.validation.constraints.NotEmpty;
import javax.persistence.Id;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.nuveo.backendtest.enums.WorkflowStatus;

/**
 * @author rsouza
 *
 */

@Entity
@Table(name = "tb_workflow")
public class Workflow extends AuditModel {

	/**
	 * 
	 */
	private static final long serialVersionUID = 2067258313899027162L;

	/** Workflow unique indentifier*/
    @Id
	protected String uuid;	
	
	/** workflow Status*/
    @Enumerated
    @Column(columnDefinition = "smallint")
	protected WorkflowStatus status;

	/** JSONB	workflow input */
	@NotEmpty
    @Column(columnDefinition = "text")
	protected String data;

	@OneToMany(mappedBy="workflow")
	@JsonIgnore	
	protected List<Step> steps;

	/**
	 * @return the steps
	 */
	public List<Step> getSteps() {
		return steps;
	}

	/**
	 * @param steps the steps to set
	 */
	public void setSteps(List<Step> steps) {
		this.steps = steps;
	}

	/**
	 * @return the uuid
	 */
	public String getUuid() {
		return uuid;
	}

	/**
	 * @param uuid the uuid to set
	 */
	public void setUuid(String uuid) {
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

	/* (non-Javadoc)
	 * @see java.lang.Object#toString()
	 */
	@Override
	public String toString() {
		return "Workflow [uuid=" + uuid + ", status=" + status + ", data=" + data
				+ "]";
	}

	public Workflow() {
	}
	
	public Workflow(WorkflowStatus pStatus, String pData) {
		super();
		this.uuid = UUID.randomUUID().toString();
		this.status = pStatus;
		this.data = pData;
	}

	public Workflow(String pUUID, WorkflowStatus pStatus, String pData) {
		super();
		this.uuid = pUUID;
		this.status = pStatus;
		this.data = pData;
	}
}
