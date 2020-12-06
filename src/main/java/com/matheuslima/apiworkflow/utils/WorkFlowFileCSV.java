package com.matheuslima.apiworkflow.utils;

import java.io.IOException;
import java.io.Writer;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.opencsv.bean.StatefulBeanToCsv;
import com.opencsv.bean.StatefulBeanToCsvBuilder;
import com.opencsv.exceptions.CsvDataTypeMismatchException;
import com.opencsv.exceptions.CsvRequiredFieldEmptyException;

import lombok.Data;

@Data
public class WorkFlowFileCSV {
	 
	public void generateCSV(List<WorkFlow> wf) throws IOException, CsvDataTypeMismatchException, CsvRequiredFieldEmptyException{
		for (WorkFlow workFlow : wf) {
			workFlow.setData("");
		}
        Writer writer = Files.newBufferedWriter(Paths.get("pessoas.csv"));
        StatefulBeanToCsv<WorkFlow> beanToCsv = new StatefulBeanToCsvBuilder(writer).build();
        beanToCsv.write(wf);
        writer.flush();
        writer.close();
	}
}
