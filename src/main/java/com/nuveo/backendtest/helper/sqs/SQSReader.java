/**
 * 
 */
package com.nuveo.backendtest.helper.sqs;

import com.amazonaws.auth.AWSCredentials;
import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.sqs.AmazonSQSClientBuilder;
import com.amazonaws.services.sqs.model.DeleteMessageRequest;
import com.amazonaws.services.sqs.model.ReceiveMessageRequest;
import com.amazonaws.services.sqs.model.Message;
import com.amazonaws.services.sqs.AmazonSQS;

import java.util.List;

/**
 * @author rsouza
 *
 */

public class SQSReader {

    private static AWSCredentials credentials;

    //TODO: Should use IAM roles here.
    static {
        credentials = new BasicAWSCredentials(
            "AKIAJN6DBGPZS2A5JUSA", 
            "SYP4F3SN3PqKM/ipb0Ste8Jd9FbbNjaTTbklE7Cr"
        );
    }

    private String sqsUrl = "";

    private AmazonSQS sqs = null;
    
    public SQSReader(String pSqsUrl)
    {
    	sqsUrl = pSqsUrl;
    	
        // Set up the client
        sqs = AmazonSQSClientBuilder.standard()
            .withCredentials(new AWSStaticCredentialsProvider(credentials))
            .withRegion(Regions.US_EAST_1)
            .build();
    	
    }
    
    public Message receiveMessage()
    {
    	ReceiveMessageRequest receiveMessageRequest = new ReceiveMessageRequest(sqsUrl)
    			  .withWaitTimeSeconds(10).withMaxNumberOfMessages(1);
    			 
    	List<Message> sqsMessages = sqs.receiveMessage(receiveMessageRequest).getMessages();
    	
    	if(sqsMessages != null && sqsMessages.size() > 0) {
    		
    		return sqsMessages.get(0);
    		
    	} else 
    		return null;
    }
    
    public void deleteMessage(String receiptHandle) {

    	sqs.deleteMessage(new DeleteMessageRequest().withQueueUrl(sqsUrl)
                .withReceiptHandle(receiptHandle));    	
    }

}