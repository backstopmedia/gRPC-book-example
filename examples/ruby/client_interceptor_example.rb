$LOAD_PATH.unshift File.expand_path("../proto", __FILE__)
require_relative './proto/sfapi_pb'
require_relative './proto/sfapi_services_pb'

class TimingInterceptor < GRPC::ClientInterceptor
  def request_response(request: nil, call: nil, method: nil, metadata: nil)
    start = Time.now.to_f
    yield
    duration = Time.now.to_f - start
    puts "Call for #{method} took #{duration * 1000}ms"
  end
end

stub = Sfapi::V1::Starfriends::Stub.new("localhost:8080", :this_channel_is_insecure, interceptors: [ TimingInterceptor.new ])
request = Sfapi::V1::GetSpeciesRequest.new(id: 'HJFy629NoG')
response = stub.get_species(request)
