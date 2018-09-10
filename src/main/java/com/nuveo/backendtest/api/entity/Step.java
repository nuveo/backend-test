/**
 * 
 */
package com.nuveo.backendtest.api.entity;

import javax.persistence.*;

import javax.validation.constraints.NotEmpty;

import org.hibernate.annotations.OnDelete;
import org.hibernate.annotations.OnDeleteAction;

import com.fasterxml.jackson.annotation.JsonIgnore;

/**
 * @author rsouza
 *
 */

@Entity
@Table(name = "tb_step")
public class Step extends AuditModel {

    /**
	 * 
	 */
	private static final long serialVersionUID = 5064131347090118460L;

	@Id
    @GeneratedValue(generator = "step_generator")
    @SequenceGenerator(
            name = "step_generator",
            sequenceName = "step_sequence",
            initialValue = 1
    )	
	protected Long id;	

	@NotEmpty
    @Column(columnDefinition = "text")
	protected String description;
	
	
	@ManyToOne(fetch = FetchType.LAZY, optional = false)
	@JoinColumn(name = "workflow_UUID", nullable = false)
	@OnDelete(action = OnDeleteAction.CASCADE)
	@JsonIgnore	
	protected Workflow workflow;

	/**
	 * @return the id
	 */
	public Long getId() {
		return id;
	}


	/**
	 * @param id the id to set
	 */
	public void setId(Long id) {
		this.id = id;
	}


	/**
	 * @return the description
	 */
	public String getDescription() {
		return description;
	}


	/**
	 * @param description the description to set
	 */
	public void setDescription(String description) {
		this.description = description;
	}


	/**
	 * @return the workflow
	 */
	public Workflow getWorkflow() {
		return workflow;
	}


	/**
	 * @param workflow the workflow to set
	 */
	public void setWorkflow(Workflow workflow) {
		this.workflow = workflow;
	}
	
	
	

}
