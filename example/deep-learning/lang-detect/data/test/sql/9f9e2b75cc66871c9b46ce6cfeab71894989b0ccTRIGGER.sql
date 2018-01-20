CREATE TRIGGER MOVIMENTATION_SPECS_INS AFTER INSERT ON MOVIMENTATION_SPECS 
	FOR EACH ROW
	CALL DOMOV (NEW.SITE, NEW.MOVIMENTATION_TYPE, NEW.ID_ART, NEW.QUANTITY , 1);
	
CREATE TRIGGER MOVIMENTATION_SPECS_DEL AFTER DELETE ON MOVIMENTATION_SPECS 
	FOR EACH ROW
	CALL DOMOV (OLD.SITE, OLD.MOVIMENTATION_TYPE, OLD.ID_ART, OLD.QUANTITY , -1);
	  
CREATE TRIGGER MOVIMENTATION_SPECS_AF_UPD AFTER UPDATE ON MOVIMENTATION_SPECS 
  	FOR EACH ROW
   	CALL DOUPDATE (OLD.SITE, OLD.MOVIMENTATION_TYPE, OLD.ID_ART, OLD.QUANTITY, NEW.SITE, NEW.MOVIMENTATION_TYPE, NEW.ID_ART, NEW.QUANTITY );
   	
   	
   	
   	
   	