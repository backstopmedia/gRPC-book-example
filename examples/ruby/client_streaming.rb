$LOAD_PATH.unshift File.expand_path("../proto", __FILE__)
require_relative './proto/swapi_pb'
require_relative './proto/swapi_services_pb'

stub = Swapi::V1::Starwars::Stub.new("localhost:8080", :this_channel_is_insecure)
request = Swapi::V1::ListStarshipActionsRequest.new
response = stub.list_starship_actions(request)

response.each do |action|
  puts "#{action.starship.name} #{action.action}"
end
