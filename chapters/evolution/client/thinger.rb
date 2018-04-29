#!/usr/bin/env ruby
# frozen_string_literal: true

$LOAD_PATH.unshift('.')
require 'grpc/google_rpc_status_utils'
require 'proto/thing_services_pb'

Stub    = Practical::Grpc::V1::MyService::Stub
Request = Practical::Grpc::V1::UpdateThingRequest
Thing   = Practical::Grpc::V1::Thing

puts Google::Protobuf::FieldMask.method(:initialize).source_location

begin
  stub    = Stub.new('localhost:8080', :this_channel_is_insecure)
  thing   = Thing.new(id: 8_929_212, name: 'New Name')
  request = Request.new(thing: thing, mask: Google::Protobuf::FieldMask.new(paths: %w[thing.id thing.name]))

  resp = stub.update_thing(request)
  puts resp
rescue StandardError => e
  code = e.respond_to?(:code) ? e.code : 'Unknown'

  if code == GRPC::Core::StatusCodes::NOT_FOUND
    status = GRPC::GoogleRpcStatusUtils.extract_google_rpc_status(e.to_status)
    puts Practical::Grpc::V1::ThingNotFoundError.decode(status.details.first.value).reason
    exit 1
  end

  puts "Code: #{code}, Type: '#{e.class}', Message: #{e.message}"
end
