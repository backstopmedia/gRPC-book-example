$LOAD_PATH.unshift File.expand_path("../", __FILE__)
require_relative './proto/sfapi_pb'
require_relative './proto/sfapi_services_pb'

stub = Starfriends::Stub.new("localhost:8080", :this_channel_is_insecure)

# Now we can use the stub to make RPCs
request = GetFilmRequest.new(id: '7')
begin
  response = stub.get_film(request)
  puts response.inspect
rescue GRPC::BadStatus => e
  # RPC exceptions have a numeric 'code' field (the gRPC error code)
  # and a string 'details' (the message supplied by the server)
  puts "failed: #{e.details} (code = #{e.code})"
end
