########################################################
#load dependencies
########################################################

import-module functions.psm1

########################################################
#Step1: ask values, show them and create user
########################################################

step1

########################################################
#TODO: Step2: Sync with Office365
########################################################

#Azure AD sync
syncWithO365

########################################################
#TODO: Step3: Office365 settings
########################################################

connectoToMsol
connectToExchangeOnline

    ########################################################
    #TODO: Office365 settings (licenses, groups, ...)
    ########################################################
    askForO365Licenses
    askForO365GroupMembership

    ########################################################
    #TODO: Exchange online settings
    ########################################################

    enableOnlineArchive
    askForProfilePhoto

    ########################################################
    #TODO: SFB settings
    ########################################################

########################################################
#TODO: close Office365 connections
######################################################## 

disconnectFromExchangeOnline


########################################################
#TODO: create Teamviewer account
######################################################## 

########################################################
#TODO: Logging
########################################################