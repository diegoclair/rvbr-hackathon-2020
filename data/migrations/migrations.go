package migrations

import "github.com/GuiaBolso/darwin"

//Migrations to execute our queries that changes database structure
var (
	Migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table tab_user",
			Script: `CREATE TABLE IF NOT EXISTS health_db.tab_user
				( 	user_id         	INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    				, user_created_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    		    	, user_updated_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					, user_deleted_at 	TIMESTAMP NULL
					, uuid 				CHAR(36) NOT NULL
					, name  			VARCHAR(160) NOT NULL
					, cpf 				VARCHAR(96)  NOT NULL
					, email 			VARCHAR(96)  NOT NULL
					, password 			VARCHAR(96)  NOT NULL
				) ENGINE = InnoDB DEFAULT CHARSET = utf8
			`,
		},
		{
			Version:     2,
			Description: "Inserting datas in table tab_user",
			Script: `
				INSERT INTO health_db.tab_user (uuid, name, cpf, email, password)
					VALUES (
						  '2367cb2d-154c-11eb-9300-0242ac120002'
						, 'Diego Clair'
						, '49305051006'
						, 'diego_rodrigues@healthapp.com'
						, 'teste123'
					),
					(
						  UUID()
					    , 'Thales Assis'
					   	, '95488203079'
					   	, 'thales_assis@healthapp.com'
					   	, 'teste123'
				    )
				;
			`,
		},
		{
			Version:     3,
			Description: "Creating table tab_medicines",
			Script: `CREATE TABLE IF NOT EXISTS health_db.tab_medicines (
				id 			INT NOT NULL AUTO_INCREMENT,
				name 		VARCHAR(100) NOT NULL,
				price 		DECIMAL(7,2) NOT NULL,
				active 		TINYINT(1) NOT NULL DEFAULT 1,
				updated_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				created_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				PRIMARY KEY (id),
				UNIQUE INDEX ID_UNIQUE (id ASC)
			) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     4,
			Description: "Inserting data on table tab_medicines",
			Script: `
				INSERT INTO health_db.tab_medicines 
					(id,name,price)
				VALUES
				(1,'DIPIRONA SODICA 1000 MG COMPRIMIDOS','15.99'),
				(2,'PARACETAMOL 750 MG CP','8.99'),
				(3,'ALLEGRA 180 MG COMPRIMIDOS','3.99'),
				(4,'MARESIS AEROSOL SPRAY NASAL 0,9%','21.99');

			`,
		},
		{
			Version:     5,
			Description: "Creating table tab_doctor",
			Script: `CREATE TABLE IF NOT EXISTS health_db.tab_doctor
				( 	  doctor_id         	INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    				, doctor_created_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    		    	, doctor_updated_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					, doctor_deleted_at 	TIMESTAMP NULL
					, uuid 				CHAR(36) NOT NULL
					, name  			VARCHAR(100) NOT NULL
					, speciality 		VARCHAR(50)  NOT NULL
					, crm			    VARCHAR(10)  NOT NULL
				) ENGINE = InnoDB DEFAULT CHARSET = utf8
			`,
		},
		{
			Version:     6,
			Description: "Inserting datas in table tab_doctor",
			Script: `
				INSERT INTO health_db.tab_doctor (uuid, name, speciality, crm)
					VALUES (
						  UUID()
						, 'Dra. Maria Aparecida Santos'
						, 'Cardiologia'
						, '12345'
					),
					(
						  UUID()
						, 'Dr. Pedro da Silva Santos'
						, 'Oftalmologia'
						, '70123'
					),
					(
						  UUID()
						, 'Dra. Silvana Pereira'
						, 'Ortopedia'
						, '67890'
					),
					(
						  UUID()
					  	, 'Dr. Giovanni Misumi'
					  	, 'Neurologia'
					  	, '12370'
				    ),
				    (
						  UUID()
				  	 	, 'Dra. Patricia Ferreira'
				  		, 'Fonoaudiologia'
				  		, '01020'
			  		)
				;
			`,
		},
		{
			Version:     7,
			Description: "Creating table tab_exam",
			Script: `CREATE TABLE IF NOT EXISTS health_db.tab_exam
				( 	  exam_id         	INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    				, exam_created_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    		    	, exam_updated_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					, exam_deleted_at 	TIMESTAMP NULL
					, uuid 				CHAR(36) NOT NULL
					, name  			VARCHAR(100) NOT NULL
					, type       		VARCHAR(50)  NOT NULL
				) ENGINE = InnoDB DEFAULT CHARSET = utf8
			`,
		},
		{
			Version:     8,
			Description: "Inserting datas in table tab_exam",
			Script: `
				INSERT INTO health_db.tab_exam (uuid, name, type)
					VALUES (
						  UUID()
						, 'Teste Ergometrico'
						, 'Cardiologia'
					),
					(
						  UUID()
						, 'Exame Oftalmologico'
						, 'Oftalmologia'
					),
					(
						  UUID()
						, 'Raio-X'
						, 'Ortopedia'
					),
					(
						  UUID()
					  	, 'Ressonância Magnetica Craniana'
					  	, 'Neurologia'
				    ),
				    (
						  UUID()
				  	 	, 'Audiometria'
				  		, 'Fonoaudiologia'
			  		)
				;
			`,
		},
		{
			Version:     9,
			Description: "Creating table tab_institution",
			Script: `CREATE TABLE IF NOT EXISTS health_db.tab_institution
				( 	  institution_id         	INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    				, institution_created_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    		    	, institution_updated_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					, institution_deleted_at 	TIMESTAMP NULL
					, uuid 				CHAR(36) NOT NULL
					, name  			VARCHAR(100) NOT NULL
					, cnpj       		VARCHAR(16)  NOT NULL
					, phone       		VARCHAR(15)  NOT NULL
					, type       		VARCHAR(50)  NOT NULL
				) ENGINE = InnoDB DEFAULT CHARSET = utf8
			`,
		},
		{
			Version:     10,
			Description: "Inserting datas in table tab_institution",
			Script: `
				INSERT INTO health_db.tab_institution (uuid, name, cnpj, phone, type)
					VALUES (
						  UUID()
						, 'Fluery - São Caetano'
						, '60840055000131'
						, '1131790822'
						, 'Laboratorio'
					),
					(
						  UUID()
						, 'Albert Einstein'
						, '60765823000130'
						, '1121511233'
						, 'Hospital'
					),
					(
						  UUID()
						, 'CDB - Sāo Paulo'
						, '02162577000125'
						, '1159087222'
						, 'Laboratorio'
					),
					(
						  UUID()
					  	, 'Hospital e Maternidade São Luiz - São Caetano'
						, '06047087000139'
						, '1127771100'
						, 'Hospital'
				    )
				;
			`,
		},
		{
			Version:     11,
			Description: "Creating table tab_health_check",
			Script: `CREATE TABLE IF NOT EXISTS health_db.tab_health_check
				( 	  health_check_id         	INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    				, health_check_created_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    		    	, health_check_updated_at 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
					, health_check_deleted_at 	TIMESTAMP NULL
					, uuid 				CHAR(36) NOT NULL

					, user_id INT UNSIGNED NOT NULL
					, CONSTRAINT fk_tab_health_check_user_id
						FOREIGN KEY (user_id)
						REFERENCES tab_user (user_id)
						ON DELETE NO ACTION
						ON UPDATE NO ACTION

					, exam_id INT UNSIGNED NOT NULL
					, CONSTRAINT fk_tab_health_check_exam_id
						FOREIGN KEY (exam_id)
						REFERENCES tab_exam (exam_id)
						ON DELETE NO ACTION
						ON UPDATE NO ACTION

					, doctor_id INT UNSIGNED NOT NULL
					, CONSTRAINT fk_tab_health_check_doctor_id
						FOREIGN KEY (doctor_id)
						REFERENCES tab_doctor (doctor_id)
						ON DELETE NO ACTION
						ON UPDATE NO ACTION

					, institution_id INT UNSIGNED NOT NULL
					, CONSTRAINT fk_tab_health_check_institution_id
						FOREIGN KEY (institution_id)
						REFERENCES tab_institution (institution_id)
						ON DELETE NO ACTION

					, health_check_date TIMESTAMP     NOT NULL
					, note       		VARCHAR(200)  NOT NULL 
				) ENGINE = InnoDB DEFAULT CHARSET = utf8
			`,
		},
		{
			Version:     12,
			Description: "Inserting datas in table tab_health_check",
			Script: `
				INSERT INTO health_db.tab_health_check (uuid, user_id, exam_id, doctor_id, institution_id, health_check_date, note)
					VALUES 
					(
						'5ad2bf31-157b-11eb-952c-0242ac120002'
						, 1
						, 1
						, 1
						, 1
						, '2011-06-22 09:44:35'
						, 'O Paciente apresentou problemas no ventrículo esquerdo.'
					),
					(
						UUID()
						, 1
						, 1
						, 1
						, 2
						, '2012-07-15 13:37:35'
						, 'O Paciente iniciou um processo de tratamento para combater aos altos níveis de gordura na rede sanguínea.'
					),
					(
						UUID()
						, 1
						, 1
						, 1
						, 3
						, '2015-06-22 15:04:35'
						, 'O Paciente está em tratamento com auxilio de profissionais da saúde.'
					),
					(
						UUID()
						, 1
						, 1
						, 1
						, 4
						, '2018-11-28 10:04:35'
						, 'O Paciente nāo apresentou problemas após a realizar todos os exames.'
					)
				;
			`,
		},
	}
)
