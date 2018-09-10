/**
 * 
 */
package com.nuveo.backendtest.helper.util;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.dataformat.csv.CsvMapper;
import com.fasterxml.jackson.dataformat.csv.CsvSchema;
import com.nuveo.backendtest.api.entity.Workflow;

/**
 * @author rsouza
 *
 */

public class WorkflowPOJOUtil {

	public static Workflow deserializeMessage(String messageBody) throws IOException {
		
		ObjectMapper mapper = new ObjectMapper();
		
		Workflow workflow = mapper.readValue(messageBody, Workflow.class);
		
		return workflow;
	}
	
	public static String serializeWorkflowAsCSV(Workflow workflow)
	{
		try {
			
		    List<Workflow> wkfList = new ArrayList<>();
		    wkfList.add(workflow);
		    
		    List<String> csvList = WorkflowPOJOUtil.convertToString(wkfList);
		    
		    if(csvList != null && csvList.size() > 0) {
		    	return csvList.get(0);
		    }
		    else 
		    	return null;

		} catch (Exception e) {
	    	return null;
		}
	}

	//Got from https://stackoverflow.com/questions/41510496/writing-generic-pojo-to-csv-transformer
	public static <T> List<String> convertToString(List<T> objectList) {

	    if(objectList.isEmpty())
	        return Collections.emptyList();

	    T entry = objectList.get(0);

	    List<String> stringList = new ArrayList<>();
	    char delimiter = ',';
	    char quote = '"';
	    String lineSep = "\n";

	    CsvMapper mapper = new CsvMapper();
	    CsvSchema schema = mapper.schemaFor(entry.getClass());

	    for (T object : objectList) {

	        try {
	            String csv = mapper.writer(schema
	                    .withColumnSeparator(delimiter)
	                    .withQuoteChar(quote)
	                    .withLineSeparator(lineSep)).writeValueAsString(object);

	            stringList.add(csv);
	        } catch (JsonProcessingException e) {
	            System.out.println(e);
	        }
	    }

	    return stringList;
	}	
}