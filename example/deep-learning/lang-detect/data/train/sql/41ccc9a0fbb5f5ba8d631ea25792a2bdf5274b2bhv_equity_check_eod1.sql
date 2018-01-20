--------------------------------------------------------
--  File created - Wednesday-May-28-2014   
--------------------------------------------------------
--------------------------------------------------------
--  DDL for View HV_EQUITY_CHECK_EOD1
--------------------------------------------------------

  CREATE OR REPLACE FORCE VIEW HV_EQUITY_CHECK_EOD1 (SOURCE, NAV_DATE, SECURITY, STK_SEC_ID, AMT, TICKER, S_STK_ID, NSE_CODE, NSE_CLOSE, BSE_CODE, BSE_CLOSE) AS 
  SELECT 'NSE' AS SOURCE,
  TRADED.NAV_DATE,
  TRADED.SECURITY,
  TRADED.STK_SEC_ID,
  TRADED.AMT,
  TRADED.TICKER,
  TRADED.S_STK_ID,
  NSE.NSE_CODE AS NSE_CODE,
  NSE.NSE_CLOSE,
  NULL AS BSE_CODE,
  NULL BSE_CLOSE
FROM
  (SELECT SCHNAVBD.SECURITY,
    SECURITY.STK_SEC_ID,
    -SUM(SCHNAVBD.AMOUNT)/SUM(SCHNAVBD.UNITS) AMT,
    SECURITY.TICKER,
    SCHNAVBD.NAV_DATE,
    SUBSTR(SECURITY.STK_SEC_ID,3,(LENGTH(SECURITY.STK_SEC_ID)-2)) S_STK_ID
  FROM HI_SCHNAVBD_CURR SCHNAVBD,
    HM_SECURITY SECURITY
  WHERE SCHNAVBD.SECURITY     = SECURITY.SECURITY
  AND ( (SCHNAVBD.INVEST_TYPE ='INVEST')
  AND (SCHNAVBD.UNITS        <>'0')
  AND (SECURITY.ASSET_TYPE    ='01')
  AND (SECURITY.SECURITY     <>SECURITY.STK_SEC_ID)
  AND (SCHNAVBD.SCHEME NOT   IN ('HDFCSX','HDFCS+','HDFCNY'))
  AND (SECURITY.RECTYPE       ='L') )
  AND SECURITY.SECURITY      <>'ARSSINF'
    --AND (schnavbd.NAV_DATE=(select cur_date-1 from sysparam where rectype='L'))
  GROUP BY SCHNAVBD.SECURITY,
    SECURITY.STK_SEC_ID,
    SECURITY.TICKER,
    SCHNAVBD.NAV_DATE
  ) TRADED,
  HI_NSE NSE
WHERE NSE.NSE_CODE = TRADED.S_STK_ID(+)
AND NSE.NSE_DT     = TRADED.NAV_DATE
UNION ALL
SELECT 'BSE' AS SOURCE,
  TRADED.NAV_DATE,
  TRADED.SECURITY,
  TRADED.STK_SEC_ID,
  TRADED.AMT,
  TRADED.TICKER,
  TRADED.S_STK_ID,
  NULL AS NSE_CODE,
  NULL NSE_CLOSE,
  BSE.BSE_CODE AS BSE_CODE,
  BSE.BSE_CLOSE
FROM
  (SELECT SCHNAVBD.SECURITY,
    SECURITY.STK_SEC_ID,
    -SUM(SCHNAVBD.AMOUNT)/SUM(SCHNAVBD.UNITS) AMT,
    SECURITY.TICKER,
    SCHNAVBD.NAV_DATE,
    SUBSTR(SECURITY.STK_SEC_ID,3,(LENGTH(SECURITY.STK_SEC_ID)-2)) S_STK_ID
  FROM HI_SCHNAVBD_CURR SCHNAVBD,
    HM_SECURITY SECURITY
  WHERE SCHNAVBD.SECURITY     = SECURITY.SECURITY
  AND ( (SCHNAVBD.INVEST_TYPE ='INVEST')
  AND (SCHNAVBD.UNITS        <>'0')
  AND (SECURITY.ASSET_TYPE    ='01')
  AND (SECURITY.SECURITY     <>SECURITY.STK_SEC_ID)
  AND (SCHNAVBD.SCHEME NOT   IN ('HDFCSX','HDFCS+','HDFCNY'))
  AND (SECURITY.RECTYPE       ='L') )
  AND SECURITY.SECURITY      <>'ARSSINF'
    --AND (schnavbd.NAV_DATE=(select cur_date-1 from sysparam where rectype='L'))
  GROUP BY SCHNAVBD.SECURITY,
    SECURITY.STK_SEC_ID,
    SECURITY.TICKER,
    SCHNAVBD.NAV_DATE
  ) TRADED,
  HI_BSE BSE
WHERE TO_CHAR(BSE.BSE_CODE) = TO_CHAR(TRADED.TICKER)
AND BSE.value_date              = TRADED.NAV_DATE
AND S_STK_ID NOT           IN
  (SELECT NSE.NSE_CODE FROM HI_NSE NSE
  );
