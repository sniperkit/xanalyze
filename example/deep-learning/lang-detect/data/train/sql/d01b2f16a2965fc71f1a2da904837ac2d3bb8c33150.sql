-- WEB-4601 - store validation in reporting
drop table if exists stg_email_validation;

CREATE TABLE stg_email_validation (
  email_address VARCHAR(255) NOT NULL PRIMARY KEY,
  status VARCHAR(255) NOT NULL
);

GRANT SELECT ON stg_EMAIL_VALIDATION TO GROUP READ_ONLY;
GRANT ALL ON stg_EMAIL_VALIDATION TO GROUP READ_WRITE;
GRANT ALL ON stg_EMAIL_VALIDATION TO GROUP SCHEMA_MANAGER;
--
drop table if exists email_validation;

CREATE TABLE email_validation (
  email_address VARCHAR(255) NOT NULL PRIMARY KEY,
  status VARCHAR(255) NOT NULL
);

GRANT SELECT ON EMAIL_VALIDATION TO GROUP READ_ONLY;
GRANT ALL ON EMAIL_VALIDATION TO GROUP READ_WRITE;
GRANT ALL ON EMAIL_VALIDATION TO GROUP SCHEMA_MANAGER;
---

drop table if exists ios_player_device;

CREATE TABLE ios_player_device (
  GAME_TYPE varchar(255) NOT NULL,
  PLAYER_ID decimal(16,2) NOT NULL,
  DEVICE_TOKEN varchar(255) NOT NULL,
  BUNDLE varchar(30) DEFAULT NULL
);

alter table ios_player_device add primary key(bundle,player_id,device_token);

GRANT SELECT ON ios_player_device TO GROUP READ_ONLY;
GRANT ALL ON ios_player_device TO GROUP READ_WRITE;
GRANT ALL ON ios_player_device TO GROUP SCHEMA_MANAGER;
---

drop table if exists SEGMENT_SELECTION;

CREATE TABLE SEGMENT_SELECTION (
  CAMPAIGN_RUN_ID bigint NOT NULL,
  PLAYER_ID decimal(16,2)  NOT NULL,
  PROGRESSIVE_BONUS bigint DEFAULT NULL,
  DISPLAY_NAME varchar(255) DEFAULT NULL,
  CONTENT varchar(10000) DEFAULT NULL
);

alter table segment_selection add PRIMARY KEY (CAMPAIGN_RUN_ID,PLAYER_ID);

GRANT SELECT ON SEGMENT_SELECTION TO GROUP READ_ONLY;
GRANT ALL ON SEGMENT_SELECTION TO GROUP READ_WRITE;
GRANT ALL ON SEGMENT_SELECTION TO GROUP SCHEMA_MANAGER;
---

drop table if exists GCM_PLAYER_DEVICE;

CREATE TABLE GCM_PLAYER_DEVICE (
    PLAYER_ID decimal(16,2) NOT NULL,
    GAME_TYPE varchar(255) NOT NULL,
    REGISTRATION_ID varchar(255) NOT NULL
);

alter table gcm_player_device add PRIMARY KEY (PLAYER_ID, GAME_TYPE, REGISTRATION_ID);
alter table gcm_player_device add unique (PLAYER_ID, REGISTRATION_ID);

GRANT SELECT ON GCM_PLAYER_DEVICE TO GROUP READ_ONLY;
GRANT ALL ON GCM_PLAYER_DEVICE TO GROUP READ_WRITE;
GRANT ALL ON GCM_PLAYER_DEVICE TO GROUP SCHEMA_MANAGER;
---
drop table if exists FACEBOOK_EXCLUSIONS;

CREATE TABLE FACEBOOK_EXCLUSIONS (
    PLAYER_ID decimal(16,2) NOT NULL,
    GAME_TYPE varchar(255) NOT NULL,
    reason varchar(32)
);

alter table FACEBOOK_EXCLUSIONS add PRIMARY KEY (PLAYER_ID, GAME_TYPE);

GRANT SELECT ON FACEBOOK_EXCLUSIONS TO GROUP READ_ONLY;
GRANT ALL ON FACEBOOK_EXCLUSIONS TO GROUP READ_WRITE;
GRANT ALL ON FACEBOOK_EXCLUSIONS TO GROUP SCHEMA_MANAGER;