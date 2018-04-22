$LOAD_PATH.unshift File.expand_path("../proto", __FILE__)
require_relative './proto/swapi_pb'
require_relative './proto/swapi_services_pb'

# Required to use the well known type helpers for proto messages
require 'grpc/google_rpc_status_utils'
require 'google/protobuf/well_known_types'

stub = Swapi::V1::Starwars::Stub.new("localhost:8080", :this_channel_is_insecure)

begin
  request = Swapi::V1::ValidateSpeciesRequest.new(name: 'a')
  response = stub.validate_species(request)
rescue GRPC::InvalidArgument => e
  status = GRPC::GoogleRpcStatusUtils.extract_google_rpc_status(e.to_status)

  status.details.each do |detail|
    error = detail.unpack(Swapi::V1::InvalidKey)
    puts "#{error.key}: #{error.message}"
  end
end
