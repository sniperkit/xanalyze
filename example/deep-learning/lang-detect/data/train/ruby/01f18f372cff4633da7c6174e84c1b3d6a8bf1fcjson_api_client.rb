require 'faraday'
require 'faraday_middleware'
require 'json'
require 'addressable/uri'

module JsonApiClient
  autoload :Associations, 'json_api_client/associations'
  autoload :Attributes, 'json_api_client/attributes'
  autoload :Connection, 'json_api_client/connection'
  autoload :Errors, 'json_api_client/errors'
  autoload :ErrorCollector, 'json_api_client/error_collector'
  autoload :Formatter, 'json_api_client/formatter'
  autoload :Helpers, 'json_api_client/helpers'
  autoload :Implementation, 'json_api_client/implementation'
  autoload :IncludedData, 'json_api_client/included_data'
  autoload :Linking, 'json_api_client/linking'
  autoload :Relationships, 'json_api_client/relationships'
  autoload :LinkDefinition, 'json_api_client/link_definition'
  autoload :MetaData, 'json_api_client/meta_data'
  autoload :Middleware, 'json_api_client/middleware'
  autoload :Paginating, 'json_api_client/paginating'
  autoload :Parsers, 'json_api_client/parsers'
  autoload :Query, 'json_api_client/query'
  autoload :Resource, 'json_api_client/resource'
  autoload :ResultSet, 'json_api_client/result_set'
  autoload :Schema, 'json_api_client/schema'
  autoload :Utils, 'json_api_client/utils'
  autoload :VERSION, 'json_api_client/version'
end