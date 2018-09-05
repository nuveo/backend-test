package com.nuveo.backendtest.enums;

public enum WorkflowStatus {

	invalid(0),
	inserted(1),
	consumed(2);
	
	int status = 0;
	
	WorkflowStatus(int pStatus) {
		this.status = pStatus;
	}
}
