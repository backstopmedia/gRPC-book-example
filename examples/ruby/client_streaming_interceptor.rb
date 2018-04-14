$LOAD_PATH.unshift File.expand_path("../proto", __FILE__)
require_relative './proto/swapi_pb'
require_relative './proto/swapi_services_pb'

class TimingInterceptor < GRPC::ClientInterceptor
  def server_streamer(request: nil, call: nil, method: nil, metadata: nil)
    start = Time.now.to_f
    result = yield
    duration = Time.now.to_f - start
    puts "Call for #{method} took #{duration * 1000}ms"

    result
  end
end

stub = Swapi::V1::Starwars::Stub.new("localhost:8080", :this_channel_is_insecure, interceptors: [ TimingInterceptor.new ])
request = Swapi::V1::ListStarshipActionsRequest.new
response = stub.list_starship_actions(request)

response.each do |action|
  puts "#{action.starship.name} #{action.action}"
end
