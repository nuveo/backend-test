/**
 * 
 */
package com.nuveo.backendtest.helper.util;

import java.io.IOException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.nuveo.backendtest.api.entity.Workflow;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
import org.json.CDL;

/**
 * @author rsouza
 *
 */

public class WorkflowPOJOUtil {

    static String jsonArrayFmt = "{\"jsonObject\": [ %s ]}";

	public static Workflow deserializeMessage(String messageBody) throws IOException {
		
		ObjectMapper mapper = new ObjectMapper();
		
		Workflow workflow = mapper.readValue(messageBody, Workflow.class);
		
		return workflow;
	}
	
	public static String JSON2CSV(String jsonString)
	{
		
		String csv = null;
		
		try {
			JSONObject output = new JSONObject(String.format(jsonArrayFmt, jsonString));
		
		    JSONArray docs = output.getJSONArray("jsonObject");
		
		    csv = CDL.toString(docs);

		} catch (JSONException e) {
		    e.printStackTrace();
		}

		return csv;
	}

}