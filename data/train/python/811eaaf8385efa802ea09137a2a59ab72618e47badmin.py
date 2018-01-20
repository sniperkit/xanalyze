import logging
from shakespeare.lib.base import BaseController, render
from shakespeare import model
from shakespeare import forms
from shakespeare.model import meta
from formalchemy.ext.pylons.admin import FormAlchemyAdminController

log = logging.getLogger(__name__)

class AdminController(BaseController):
    model = model # where your SQLAlchemy mappers are
    forms = forms # module containing FormAlchemy fieldsets definitions
    def Session(self): # Session factory
        return meta.Session

AdminController = FormAlchemyAdminController(AdminController)

