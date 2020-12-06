package com.matheuslima.apiworkflow.config;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class Connect {

	static final String URL =
			"jdbc:postgresql://localhost:5432/workflow";
	
	static final String USER = "postgres";
	
	static final String PASS = "Hayashii123";
	
	//Set up the bank before using
	public static Connection createConnection() throws ClassNotFoundException, SQLException{
		
		Class.forName("org.postgresql.Driver");
		Connection connect = DriverManager.getConnection(URL, USER, PASS);
		if (connect != null) {
			System.out.println("Connection successful JDBC...");
			return connect;
		}
		return null;
	}
}
