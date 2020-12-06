package com.matheuslima.apiworkflow.domain;

import java.io.Serializable;
import java.util.List;
import java.util.UUID;

import javax.persistence.Column;
import javax.persistence.Convert;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.Type;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.matheuslima.apiworkflow.domain.enums.WorkFlowStatus;
import com.matheuslima.apiworkflow.utils.StringListConverter;

import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@ToString
@Entity
@Table(name="WORKFLOW")
public class WorkFlow implements Serializable{
	private static final long serialVersionUID = 1L;
	
	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	@Type(type="org.hibernate.type.UUIDCharType")
	@Column(name="uuid")
	private UUID uuid;
	
	@Column(name="status")
	private WorkFlowStatus wFS;
	
	@JsonProperty("data")
	private String data;
	
    //@ElementCollection
    @Convert(converter = StringListConverter.class)
    @Column(name="workflowSteps")
	private List<String> workflowSteps;
	
}
