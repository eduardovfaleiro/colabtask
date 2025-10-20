create domain id_uuid_v7 as uuid
	DEFAULT uuidv7()
	NOT NULL;

create table "users" (
	id id_uuid_v7 primary key,
	name varchar(50) not null,
	email varchar(100) not null unique,
	password varchar(50) not null
 );

create table "projects" (
	id id_uuid_v7 primary key,
	title varchar(100) not null,
	description varchar(255) not null
);

create table "tasks" (
	id id_uuid_v7 primary key,
	project_id id_uuid_v7 references "projects" (id),
	title varchar(100) not null,
	description varchar(255),
	due_date timestamptz
);

create table "project_collaborators" (
	user_id id_uuid_v7 references "users" (id),
	project_id id_uuid_v7 references "projects" (id),
	role varchar(20) default 'member'
);

create table "task_assignments" (
	task_id id_uuid_v7 references "tasks" (id),
	user_id id_uuid_v7 references "users" (id),
	assigned_at timestamptz default now()
);
