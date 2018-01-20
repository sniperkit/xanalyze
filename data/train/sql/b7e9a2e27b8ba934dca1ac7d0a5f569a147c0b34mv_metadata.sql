CREATE TABLE mv_metadata
(
  `code` varchar(255) NOT NULL,
  `type` varchar(255),
  `width` numeric(10),
  `height` numeric(10),
  `field` varchar(255),
  `db` varchar(255),
  `name` varchar(255),
  `outboard` varchar(255),
  `options` varchar(255),
  `attribute` varchar(255),
  `label` varchar(255),
  `help` text,
  `lookup` varchar(255),
  `filter` varchar(255),
  `help_url` varchar(255),
  `pre_filter` varchar(255),
  `lookup_exclude` varchar(255),
  `prepend` varchar(255),
  `append` varchar(255),
  `display_filter` varchar(255),
  `extended` text,
  PRIMARY KEY (`code`)
);
