package com.matheuslima.apiworkflow.domain.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@AllArgsConstructor
public class MenssageQueueDTO {
	
	@Getter @Setter
    private int id;
	
	@Getter @Setter
    private String name;
	
    @Override
    public String toString() {
        return "CustomMessage{" +
                "id=" + id +
                ", name='" + name + '\'' +
                '}';
    }

}
