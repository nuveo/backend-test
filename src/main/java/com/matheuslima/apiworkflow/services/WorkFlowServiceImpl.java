package com.matheuslima.apiworkflow.services;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.List;
import java.util.Optional;
import java.util.UUID;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.util.Assert;

import com.matheuslima.apiworkflow.config.Connect;
import com.matheuslima.apiworkflow.domain.WorkFlow;
import com.matheuslima.apiworkflow.domain.dto.WorkFlowDTO;
import com.matheuslima.apiworkflow.repositories.WorkFlowRepository;

@Service
public class WorkFlowServiceImpl implements WorkFlowService {
	
	@Autowired
	private WorkFlowRepository wfr;

	//consumer
	@Override
	public List<WorkFlowDTO> findAll() {
		return wfr.findAll().parallelStream().map(WorkFlowDTO::create).collect(Collectors.toList());
	}

	//consumer
	@Override
	public Optional<WorkFlowDTO> findByUuid(UUID uuid) {
		Optional<WorkFlow> wf = wfr.findByUuid(uuid);
		return wf.map(WorkFlowDTO::create);
	}

	//producer
	@Override
	public WorkFlowDTO save(WorkFlow wf) {
		Assert.isNull(wf.getUuid(), "It is not possible to insert a null workflow");
		return WorkFlowDTO.create(wfr.save(wf));
	}

	//I preferred to use JDBC :)
	@Override
	public void save(WorkFlowDTO partialUpdate, String uuid) throws ClassNotFoundException, SQLException {
		String sql = "UPDATE WORKFLOW SET status  = 2 WHERE uuid = ?";
		Connection con = Connect.createConnection();
		try {
			PreparedStatement ps = con.prepareStatement(sql);
			ps.setString(1, uuid);
			
			ps.execute();
			ps.close();
			
		} catch(SQLException e) {
			System.out.println("Erro ao tentar atualizar: ".concat(e.getMessage()));
		}
	}

}
