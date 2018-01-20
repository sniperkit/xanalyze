﻿using StorageControl.Domain.Model.Entities;
using System.Collections.Generic;

namespace StorageControl.Web.Models.Instruments
{
    public class InstrumentsEditModel
    {
        public InstrumentsEditModel()
        {
            this.Instrument = new Instrument();

            this.Categories = new List<Category>() { new Category() { Id = 0, Name = "Selecione" } };

            this.InstrumentTypes = new List<InstrumentType>() { new InstrumentType() { Id = 0, Name = "Selecione" } };
        }

        public Instrument Instrument { get; set; }

        public List<Category> Categories { get; set; }

        public List<InstrumentType> InstrumentTypes { get; set; }
    }
}