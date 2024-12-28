CREATE TABLE course (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  desc TEXT NOT NULL,
  start_date TEXT,
  end_date TEXT
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
