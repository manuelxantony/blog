USE blog

CREATE TABLE post( id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY, title VARCHAR(150) NOT NULL, content TEXT NOT NULL, createdby VARCHAR(50) NOT NULL );