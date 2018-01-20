SET DEFINE OFF;
CREATE TABLE D_EXCHANGE_RATES
(
  BASECURRENCY  VARCHAR2(3 BYTE),
  TYPE          VARCHAR2(1 BYTE),
  CODE          VARCHAR2(3 BYTE),
  EXCHRATE      NUMBER(18,6),
  BEGDDATE      DATE,
  ENDDDATE      DATE,
  PRODUCTCODE   VARCHAR2(50 BYTE),
  WHOSTAMP      VARCHAR2(50 BYTE),
  DATESTAMP     DATE
)
TABLESPACE DTW_ADV_TABLES
RESULT_CACHE (MODE DEFAULT)
PCTUSED    0
PCTFREE    10
INITRANS   1
MAXTRANS   255
STORAGE    (
            INITIAL          80K
            NEXT             1M
            MINEXTENTS       1
            MAXEXTENTS       UNLIMITED
            PCTINCREASE      0
            BUFFER_POOL      DEFAULT
            FLASH_CACHE      DEFAULT
            CELL_FLASH_CACHE DEFAULT
           )
LOGGING 
NOCOMPRESS 
NOCACHE
NOPARALLEL
MONITORING;
