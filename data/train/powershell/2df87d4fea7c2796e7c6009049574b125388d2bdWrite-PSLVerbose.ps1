﻿

<#
.Synopsis
   Writes Verbose messages to various defined targets
.DESCRIPTION
   Writes Verbose log messages to the various defined targets
.PARAMETER Message
    The Log Message to record
.PARAMETER Context
    Where the Log Message was generated,
    usually the name of the function or script that called this cmdlet
.PARAMETER LogName
    The name of the Logger to use to log the message content. Defaults
    to DEFAULT
.EXAMPLE
   Write-PSLVerbose -Message "Logging Verbose Message"
   
   this will Write Waring Strings to the DEFAULT logs targets with a context of NONE
.EXAMPLE
   Write-PSLVerbose -Message "Log This Verbose" -Context 'Add-FooBar' -LogName 'PSLogger'

   this will Write Verbose Strings to the PSLogger Logs targets with a context of Add-FooBar
.INPUTS
   String, String, String
.OUTPUTS
   Void
.COMPONENT
   PoshLogger
#>
function Write-PSLVerbose {
    [CmdletBinding()]
    [OutputType([Void])]
    Param
    (
    [Parameter(
        Mandatory, 
        Position=0
    )]
    [ValidateNotNull()]
    [string]
    $Message,

    [Parameter(
        Position=1
    )]
    [string]
    $Context='NONE',

    [Parameter(
        Position=2
    )]
    [String]
    $LogName='DEFAULT'
    )

    Begin {}

    Process {}

    End {}
}

#Register as Public Function to be exported  
if( Test-Path Variable:PUBLIC ) {
    $PUBLIC.Function+="$(( $MyInvocation.MyCommand.Name -replace '.ps1','' ).Trim() )"
}