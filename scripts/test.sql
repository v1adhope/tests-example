-- CREATE DATABASE test;

CREATE TABLE members (
  member_id int GENERATED ALWAYS AS IDENTITY,
  first_name varchar(32),
  last_name varchar(32),
  age smallserial,

  CONSTRAINT pk_members_member_id PRIMARY KEY(member_id)
)
