-- 1.4.3
-- * using ib_util.ib_util_malloc library
-- * Fix ReadLnFromTextFile: return by CSTRING(32000) free_it
-- * Type handle of object: integer -> bigint for x64_86 build
-- * Fix declaration:CSTRING(256) --> CSTRING(32000) : WriteToTextFile, WriteLnToTextFile

DECLARE EXTERNAL FUNCTION CreateTextFile
    CSTRING(256)
RETURNS BIGINT BY VALUE
ENTRY_POINT 'CreateTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION ResetTextFile
    BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'ResetTextFile' MODULE_NAME 'afutextfile';


DECLARE EXTERNAL FUNCTION CloseTextFile
    BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'CloseTextFile' MODULE_NAME 'afutextfile';


DECLARE EXTERNAL FUNCTION WriteToTextFile
    BIGINT,CSTRING(32765)
RETURNS INTEGER BY VALUE
ENTRY_POINT 'WriteToTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION WriteLnToTextFile
    BIGINT,CSTRING(32765)
RETURNS INTEGER BY VALUE
ENTRY_POINT 'WriteLnToTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION RewriteTextFile
    BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'RewriteTextFile' MODULE_NAME 'afutextfile';


DECLARE EXTERNAL FUNCTION ReadLnFromTextFile
    BIGINT
RETURNS CSTRING(32765) FREE_IT
ENTRY_POINT 'ReadLnFromTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION EofTextFile
    BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'EofTextFile' MODULE_NAME 'afutextfile';

-- 1.3
DECLARE EXTERNAL FUNCTION FlushTextFile
    BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'FlushTextFile' MODULE_NAME 'afutextfile';

-- 1.4.3
DECLARE EXTERNAL FUNCTION AppendTextFile
    BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'AppendTextFile' MODULE_NAME 'afutextfile';


DECLARE EXTERNAL FUNCTION ReadCharFromTextFile
    BIGINT
RETURNS CSTRING(32765) FREE_IT
ENTRY_POINT 'ReadCharFromTextFile' MODULE_NAME 'afutextfile';


DECLARE EXTERNAL FUNCTION WriteCharToTextFile
    BIGINT,CSTRING(1)
RETURNS INTEGER BY VALUE
ENTRY_POINT 'WriteToTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION WriteInt32ToTextFile
    BIGINT,INTEGER
RETURNS INTEGER BY VALUE
ENTRY_POINT 'WriteInt32ToTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION WriteInt64ToTextFile
    BIGINT,BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'WriteInt64ToTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION ReadInt32FromTextFile
    BIGINT
RETURNS INTEGER BY VALUE
ENTRY_POINT 'ReadInt32FromTextFile' MODULE_NAME 'afutextfile';

DECLARE EXTERNAL FUNCTION ReadInt64FromTextFile
    BIGINT
RETURNS BIGINT BY VALUE
ENTRY_POINT 'ReadInt64FromTextFile' MODULE_NAME 'afutextfile';
