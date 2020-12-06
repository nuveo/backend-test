package com.matheuslima.apiworkflow.services;

import java.io.IOException;
import java.util.List;

import org.springframework.stereotype.Service;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.utils.WorkFlowFileCSV;
import com.opencsv.exceptions.CsvDataTypeMismatchException;
import com.opencsv.exceptions.CsvRequiredFieldEmptyException;

@Service
public class WorkFlowConsumerServiceImpl implements WorkFlowConsumerService {

	@Override
	public void writeWorkFlowInFile(List<WorkFlow> wf) throws CsvDataTypeMismatchException, CsvRequiredFieldEmptyException, IOException {
		WorkFlowFileCSV file = new WorkFlowFileCSV();
		file.generateCSV(wf);
	}

}
