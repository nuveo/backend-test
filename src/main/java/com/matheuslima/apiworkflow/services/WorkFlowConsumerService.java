package com.matheuslima.apiworkflow.services;

import java.io.IOException;
import java.util.List;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.opencsv.exceptions.CsvDataTypeMismatchException;
import com.opencsv.exceptions.CsvRequiredFieldEmptyException;

public interface WorkFlowConsumerService {
	
	void writeWorkFlowInFile(List<WorkFlow> wf) throws CsvDataTypeMismatchException, CsvRequiredFieldEmptyException, IOException;
	
}
