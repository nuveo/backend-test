
create table nuveo.workflow(workflow_id serial not null,
                            uuid uuid not null,
                            status smallint not null,
                            data jsonb,
                            steps text ARRAY not null
                            );

alter table nuveo.workflow
add constraint pk_workflow
primary key (workflow_id);

alter table nuveo.workflow
add constraint uq_workflow_uuid
unique(workflow_id);

alter table nuveo.workflow
add constraint chk_workflow_status
check(status = 0 or status = 1);
