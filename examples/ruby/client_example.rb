$LOAD_PATH.unshift File.expand_path("../", __FILE__)
require_relative './proto/swapi_pb'
require_relative './proto/swapi_services_pb'

stub = Swapi::V1::Starwars::Stub.new("localhost:8080", :this_channel_is_insecure)

begin
  request = Swapi::V1::GetFilmRequest.new(id: 'bunk')
  response = stub.get_film(request)
  puts "Film Title: #{response.film.title}"
rescue StandardError => e
  puts "Code: #{e.code}, ErrorClass: '#{e.class}', Message: #{e.message}"
end
