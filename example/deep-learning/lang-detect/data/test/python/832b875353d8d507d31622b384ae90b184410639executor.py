#!/usr/bin/python
# (c) Nelen & Schuurmans.  GPL licensed.

from flooding_worker.worker.action_workflow import (
    ActionWorkflow, ActionTaskPublisher, ActionHeartbeat)
from flooding_worker.worker.broker_connection import BrokerConnection
from flooding_worker.worker.message_logging_handler import AMQPMessageHandler
from flooding_worker.models import WorkflowTask
from django.conf import settings

import logging
log = logging.getLogger("flooding.management.start_scenario")


def start_workflow(scenario_id, workflowtemplate_id, log_level='INFO'):
    """
    Opens connection to broker.
    Creates ActionWorkflow object.
    Creates logging handler to send loggings to broker.
    Sets logging handler to ActionWorkflow object.
    Performs workflow.
    Closes connection.
    """
    numeric_level = getattr(logging, log_level.upper(), None)
    if not isinstance(numeric_level, int):
        log.error("Invalid log level: %s" % log_level)
        numeric_level = 10

    broker = BrokerConnection()
    connection = broker.connect_to_broker()
    if connection is None:
        log.error("Could not connect to broker.")
        return

    action = ActionWorkflow(
        connection, scenario_id, workflowtemplate_id)

    logging.handlers.AMQPMessageHandler = AMQPMessageHandler
    broker_handler = logging.handlers.AMQPMessageHandler(action,
                                                         numeric_level)

    action.set_broker_logging_handler(broker_handler)
    status = action.perform_workflow()

    if connection.is_open:
        connection.close()

    return status


def start_task(task_id, log_level='INFO'):
    """
    Publish a message to execute a separate task.
    """
    task = WorkflowTask.objects.get(pk=task_id)
    numeric_level = getattr(logging, log_level.upper(), None)
    if not isinstance(numeric_level, int):
        log.error("Invalid log level: %s" % log_level)
        numeric_level = 10

    broker = BrokerConnection()
    connection = broker.connect_to_broker()
    if connection is None:
        log.error("Could not connect to broker.")
        return

    action = ActionTaskPublisher(connection, task)

    logging.handlers.AMQPMessageHandler = AMQPMessageHandler
    broker_handler = logging.handlers.AMQPMessageHandler(action,
                                                         numeric_level)

    action.set_broker_logging_handler(broker_handler)
    success = action.perform()

    if connection.is_open:
        connection.close()

    return success


def start_heartbeat(log_level='INFO'):
    """
    Publish a message to execute a separate task.
    """
    numeric_level = getattr(logging, log_level.upper(), None)

    broker = BrokerConnection()
    connection = broker.connect_to_broker()
    if connection is None:
        log.error("Could not connect to broker.")
        return

    action = ActionHeartbeat(connection, settings.HEARTBEAT_QUEUES)

    logging.handlers.AMQPMessageHandler = AMQPMessageHandler
    broker_handler = logging.handlers.AMQPMessageHandler(action,
                                                         numeric_level)

    action.set_broker_logging_handler(broker_handler)
    success = action.perform()

    if connection.is_open:
        connection.close()

    return success
