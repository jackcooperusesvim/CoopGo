CREATE TABLE Class (
	id INTEGER PRIMARY KEY,

	name TEXT UNIQUE NOT NULL,
	desc TEXT NOT NULL,

	start_date TEXT,
	end_date TEXT,
);
CREATE TABLE ClassJoinChild(
	id INTEGER PRIMARY KEY,

	class_id INTEGER NOT NULL,
	child_id INTEGER NOT NULL,

	FOREIGN KEY (class) REFERENCES Class(id)
	FOREIGN KEY (child) REFERENCES Child(id)
);

CREATE TABLE Family (
	id INTEGER PRIMARY KEY,

	last_name TEXT NOT NULL,
	main_parent TEXT NOT NULL,
	sec_parent TEXT NOT NULL,

	phone1 TEXT NOT NULL,
	phone1 TEXT NOT NULL,
	phone2 TEXT NOT NULL,

	phone3 TEXT,
);

CREATE TABLE Child (
	id INTEGER PRIMARY KEY,

	first_name TEXT NOT NULL,
	last_name TEXT,
	birthday TEXT NOT NULL,

	grade_offset INTEGER NOT NULL,

	family_id INTEGER,
	FOREIGN KEY (family_id) REFERENCES Family(id)
);
