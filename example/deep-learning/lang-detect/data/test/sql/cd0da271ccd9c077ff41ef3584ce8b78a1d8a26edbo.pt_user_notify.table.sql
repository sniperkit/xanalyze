﻿CREATE TABLE [dbo].[pt_user_notify] (
    [userID]            NCHAR (35) NOT NULL,
    [projectID]         NCHAR (35) NOT NULL,
    [email_file_new]    BIT        NULL,
    [mobile_file_new]   BIT        NULL,
    [email_file_upd]    BIT        NULL,
    [mobile_file_upd]   BIT        NULL,
    [email_file_com]    BIT        NULL,
    [mobile_file_com]   BIT        NULL,
    [email_issue_new]   BIT        NULL,
    [mobile_issue_new]  BIT        NULL,
    [email_issue_upd]   BIT        NULL,
    [mobile_issue_upd]  BIT        NULL,
    [email_issue_com]   BIT        NULL,
    [mobile_issue_com]  BIT        NULL,
    [email_msg_new]     BIT        NULL,
    [mobile_msg_new]    BIT        NULL,
    [email_msg_upd]     BIT        NULL,
    [mobile_msg_upd]    BIT        NULL,
    [email_msg_com]     BIT        NULL,
    [mobile_msg_com]    BIT        NULL,
    [email_mstone_new]  BIT        NULL,
    [mobile_mstone_new] BIT        NULL,
    [email_mstone_upd]  BIT        NULL,
    [mobile_mstone_upd] BIT        NULL,
    [email_mstone_com]  BIT        NULL,
    [mobile_mstone_com] BIT        NULL,
    [email_todo_new]    BIT        NULL,
    [mobile_todo_new]   BIT        NULL,
    [email_todo_upd]    BIT        NULL,
    [mobile_todo_upd]   BIT        NULL,
    [email_todo_com]    BIT        NULL,
    [mobile_todo_com]   BIT        NULL,
    [email_time_new]    BIT        NULL,
    [mobile_time_new]   BIT        NULL,
    [email_time_upd]    BIT        NULL,
    [mobile_time_upd]   BIT        NULL,
    [email_bill_new]    BIT        NULL,
    [mobile_bill_new]   BIT        NULL,
    [email_bill_upd]    BIT        NULL,
    [mobile_bill_upd]   BIT        NULL,
    [email_bill_paid]   BIT        NULL,
    [mobile_bill_paid]  BIT        NULL
);

