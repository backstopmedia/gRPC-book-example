$LOAD_PATH.unshift File.expand_path("../", __FILE__)
require_relative './proto/sfapi_pb'
require_relative './proto/sfapi_services_pb'

# We're going to run 2 RPCs to two different servers
# and choose whichever is fastest
stub1 = Starfriends::Stub.new("localhost:8080", :this_channel_is_insecure)
stub2 = Starfriends::Stub.new("localhost:8081", :this_channel_is_insecure)

request = GetFilmRequest.new(id: '7')
step1 = stub1.get_film(request, return_op: true)
step2 = stub2.get_film(request, return_op: true)

all = [step1, step2]
result = nil
error = nil
threads = []

all.each_with_index do |step, i|
  # we execute each step in its own thread
  threads.push(Thread.new {
    begin
      result = step.execute
      # we have a result, so cancel other
      all.each do |other|
        other.cancel
      end
    rescue => e
      error = e if error.nil?
      # don't cancel: the other one might succeed
    end
  })
end

threads.each do |th|
  th.join
end

if not result.nil? then
    puts result.inspect
else
  puts "failed: #{error.details} (code = #{error.code})"
end
