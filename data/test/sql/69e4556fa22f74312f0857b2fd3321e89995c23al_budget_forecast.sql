SET DEFINE OFF;
CREATE TABLE L_BUDGET_FORECAST
(
  FISCALYEAR       NUMBER(10),
  MONTH            VARCHAR2(15 BYTE),
  PERIOD           NUMBER(3),
  COMPANY          VARCHAR2(10 BYTE),
  BRANDID          VARCHAR2(10 BYTE),
  CHANNEL          VARCHAR2(10 BYTE),
  DIVISION         VARCHAR2(10 BYTE),
  CLASIFICATION    VARCHAR2(50 BYTE),
  GROUPTYPE        VARCHAR2(100 BYTE),
  AFFILIATION_ID   VARCHAR2(50 BYTE),
  GL_NUMBER        VARCHAR2(50 BYTE),
  GLNAME           VARCHAR2(100 BYTE),
  BUDGETAMOUNT     NUMBER(19,4),
  BUDGET_PAX       NUMBER(10),
  FORECAST_AMOUNT  NUMBER(19,4),
  FORECAST_PAX     NUMBER(10)
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
