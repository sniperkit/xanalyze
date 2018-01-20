--------------------------------------------------------
--  Fichier créé - lundi-mai-26-2014   
--------------------------------------------------------
--------------------------------------------------------
--  DDL for Table INDIVIDUS
--------------------------------------------------------

  CREATE TABLE "JEE"."INDIVIDUS" 
   (	"PRENOM" NVARCHAR2(30), 
	"NOM" NVARCHAR2(30), 
	"AGE" NUMBER, 
	"ID_INDIVIDU" NUMBER
   ) SEGMENT CREATION IMMEDIATE 
  PCTFREE 10 PCTUSED 40 INITRANS 1 MAXTRANS 255 NOCOMPRESS LOGGING
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "SYSTEM" ;
REM INSERTING into JEE.INDIVIDUS
SET DEFINE OFF;
Insert into JEE.INDIVIDUS (PRENOM,NOM,AGE,ID_INDIVIDU) values ('charlotte','henin','20','2');
Insert into JEE.INDIVIDUS (PRENOM,NOM,AGE,ID_INDIVIDU) values ('baptiste','henin','14','3');
--------------------------------------------------------
--  DDL for Index PK_INDIVIDUS
--------------------------------------------------------

  CREATE UNIQUE INDEX "JEE"."PK_INDIVIDUS" ON "JEE"."INDIVIDUS" ("ID_INDIVIDU") 
  PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "SYSTEM" ;
--------------------------------------------------------
--  Constraints for Table INDIVIDUS
--------------------------------------------------------

  ALTER TABLE "JEE"."INDIVIDUS" ADD CONSTRAINT "PK_INDIVIDUS" PRIMARY KEY ("ID_INDIVIDU")
  USING INDEX PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "SYSTEM"  ENABLE;
