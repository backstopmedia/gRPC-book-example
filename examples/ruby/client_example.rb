$LOAD_PATH.unshift File.expand_path("../proto", __FILE__)
require_relative './proto/sfapi_pb'
require_relative './proto/sfapi_services_pb'

stub = Sfapi::V1::Starfriends::Stub.new("localhost:8080", :this_channel_is_insecure)

begin
  request = Sfapi::V1::GetSpeciesRequest.new(id: 'bunk')
  response = stub.get_species(request)
  puts "Film Title: #{response.film.title}"
rescue StandardError => e
  puts "Code: #{e.code}, ErrorClass: '#{e.class}', Message: #{e.message}"
end
