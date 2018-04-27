$LOAD_PATH.unshift File.expand_path("../", __FILE__)
require_relative './proto/sfapi_pb'
require_relative './proto/sfapi_services_pb'

stub = Starfriends::Stub.new("localhost:8080", :this_channel_is_insecure)

# We're going to run 3 RPCs in parallel
request1 = GetFilmRequest.new(id: '4')
step1 = stub.get_film(request1, return_op: true)
request2 = GetFilmRequest.new(id: '5')
step2 = stub.get_film(request2, return_op: true)
request3 = GetFilmRequest.new(id: '6')
step3 = stub.get_film(request3, return_op: true)

all = [step1, step2, step3]
results = [nil, nil, nil]
error = nil
threads = []

all.each_with_index do |step, i|
  # we execute each step in its own thread
  threads.push(Thread.new {
    begin
      results[i] = step.execute
    rescue => e
      error = e if error.nil? 
      # cancel outstanding ops on failure
      all.each do |other|
        other.cancel
      end
    end
  })
end

threads.each do |th|
  th.join
end

if not error.nil? then
  puts "failed: #{error.details} (code = #{error.code})"
else
  results.each do |res|
    puts res.inspect
  end
end
