CREATE TABLE course (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  desc TEXT NOT NULL,
  start_date TEXT NOT NULL,
  end_date TEXT NOT NULL
);

CREATE TABLE child (
	id INTEGER PRIMARY KEY,

	first_name TEXT NOT NULL,
	last_name TEXT,
	birthday TEXT NOT NULL,

	grade_offset INTEGER NOT NULL,

	family_id INTEGER,
	FOREIGN KEY (family_id) REFERENCES family(id)
);

CREATE TABLE coursejoinchild(
	id INTEGER PRIMARY KEY,

	course_id INTEGER NOT NULL,
	child_id INTEGER NOT NULL,

	FOREIGN KEY (course_id) REFERENCES course(id),
	FOREIGN KEY (child_id) REFERENCES child(id)
);

CREATE TABLE family (
	id INTEGER PRIMARY KEY,

	last_name TEXT NOT NULL,
	main_parent TEXT NOT NULL,
	sec_parent TEXT NOT NULL,

	phone1 TEXT NOT NULL,
	phone2 TEXT NOT NULL,

	phone3 TEXT
);

CREATE TABLE account (
	id INTEGER PRIMARY KEY,
	email TEXT NOT NULL,
	password_hash TEXT NOT NULL,

	default_session_lifetime TEXT NOT NULL,
	priviledge_type TEXT NOT NULL,
	last_updated TEXT NOT NULL
);

CREATE TABLE session (
	id INTEGER PRIMARY KEY,
	token TEXT UNIQUE NOT NULL,
	expiration_datetime TEXT NOT NULL,
	account_id INTEGER NOT NULL,
	FOREIGN KEY (account_id) REFERENCES account(id)
);
