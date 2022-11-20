CREATE TYPE user_role AS ENUM ('super_admin','end_user');

CREATE TABLE "users" (
  "id" varchar(64) PRIMARY KEY,
  "name" varchar(200) NOT NULL DEFAULT '',    
  "email" varchar(512) UNIQUE NOT NULL DEFAULT '',
  "password" varchar(200),
  "role_type" user_role, 
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "stats"(
  "id" varchar(64) PRIMARY KEY,
  "rank" integer NOT NULL DEFAULT 0,    
  "state" varchar(512) UNIQUE NOT NULL DEFAULT '',
  "total_score" varchar(200) NOT NULL DEFAULT '',
  "personal_residencial_safety" integer NOT NULL DEFAULT 0,
  "financial_safety" integer NOT NULL DEFAULT 0,
  "road_safety" integer NOT NULL DEFAULT 0,
  "workplace_safety" integer NOT NULL DEFAULT 0,
  "healthcare_facilities" integer NOT NULL DEFAULT 0,
  "monthly_rent_spare" varchar(200) NOT NULL DEFAULT 0
);


INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (1,46,'Alabama','37.68',36,41,42,37,46,'914.08');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (2,16,'Alaska','54.52',37,16,27,21,2,'1288.42');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (3,20,'Arizona','52.95',33,20,49,9,8,'1208.50');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (4,48,'Arkansas','33.72',47,43,47,47,33,'778.92');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (5,27,'California','48.94',42,22,40,18,19,'1777.67');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (6,37,'Colorado','43.84',44,8,31,43,28,'1404.33');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (7,7,'Connecticut','59.88',1,23,17,27,17,'1315.42');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (8,24,'Delaware','51.36',23,29,25,33,13,'1346.08');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (9,44,'Florida','39.52',34,25,46,44,42,'1329.92');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (10,35,'Georgia','45.27',22,44,39,25,34,'1083.75');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (11,5,'Hawaii','60.36',14,14,20,13,4,'1880.08');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (12,14,'Idaho','54.67',16,18,6,46,15,'901.75');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (13,32,'Illinois','47.68',13,49,18,23,35,'1049.75');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (14,19,'Indiana','52.98',15,37,24,12,29,'889.25');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (15,13,'Iowa','55.29',8,19,5,7,47,'831.17');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (16,39,'Kansas','42.91',29,28,21,34,44,'879.83');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (17,28,'Kentucky','48.73',25,39,35,20,30,'859.25');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (18,50,'Louisiana','31.64',46,48,34,40,49,'869.75');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (19,2,'Maine','66.95',3,13,2,19,1,'1024.67');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (20,18,'Maryland','53.23',9,24,45,14,23,'1505.33');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (21,6,'Massachusetts','60.02',7,5,1,42,11,'1553.75');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (22,26,'Michigan','49.79',43,34,38,15,14,'954.25');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (23,8,'Minnesota','59.81',17,4,3,6,22,'1054.33');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (24,49,'Mississippi','32.52',30,46,50,48,50,'868.42');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (25,43,'Missouri','40.25',41,32,43,24,41,'909.00');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (26,40,'Montana','42.16',45,6,28,38,32,'919.67');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (27,34,'Nebraska','46.58',24,9,9,39,40,'899.83');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (28,31,'Nevada','47.72',48,50,37,17,6,'1245.17');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (29,3,'New Hampshire','63.99',4,1,13,41,3,'1213.50');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (30,15,'New Jersey','54.62',5,30,16,28,25,'1541.08');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (31,25,'New Mexico','49.92',38,38,48,10,12,'901.33');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (32,22,'New York','51.92',10,31,4,30,26,'1431.58');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (33,21,'North Carolina','52.37',12,15,30,4,43,'1008.92');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (34,30,'North Dakota','47.94',20,12,8,49,37,'853.83');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (35,36,'Ohio','45.14',32,47,23,29,24,'865.17');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (36,45,'Oklahoma','39.34',35,45,29,36,45,'847.92');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (37,12,'Oregon','55.48',39,11,12,8,16,'1226.58');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (38,29,'Pennsylvania','48.51',18,35,33,26,27,'1001.42');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (39,10,'Rhode Island','58.99',6,26,19,35,10,'1178.08');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (40,42,'South Carolina','41.03',50,36,44,16,36,'1004.50');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (41,38,'South Dakota','43.38',28,3,32,50,39,'804.83');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (42,41,'Tennessee','41.34',49,40,41,11,38,'979.92');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (43,47,'Texas','36.77',40,42,36,31,48,'1090.83');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (44,4,'Utah','63.48',19,2,7,2,5,'1171.08');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (45,1,'Vermont','68.79',2,10,14,1,7,'905');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (46,11,'Virginia','55.99',11,17,22,3,31,'1351.33');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (47,9,'Washington','59.44',27,7,15,5,9,'1487.58');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (48,33,'West Virginia','47.08',31,33,26,45,18,'786.83');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (49,23,'Wisconsin','51.42',26,21,11,32,20,'943.58');
INSERT INTO public.stats(id, rank, state, total_score, personal_residencial_safety, financial_safety, road_safety, workplace_safety, healthcare_facilities, monthly_rent_spare) VALUES (50,17,'Wyoming','54.04',21,27,10,22,21,'876.92	');

  
 
)