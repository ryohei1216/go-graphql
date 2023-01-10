create table if not exists test.users (id varchar(10) not null primary key, name varchar(10) not null);
create table if not exists test.todos (id varchar(10) not null primary key, text varchar(10) not null, user_id varchar(10) not null, foreign key fk_user_id (user_id) references users(id) );
insert into test.users values ('testid', 'miyamoto');
insert into test.todos values ('todo1', 'todo1', 'testid');
select * from test.todos;