#!/usr/bin/env ruby
# frozen_string_literal: true

$LOAD_PATH.unshift('.')
require 'proto/database_services_pb'

Stub    = Practical::Grpc::V1::Database::Stub
Request = Practical::Grpc::V1::SearchRequest

term = ARGV.first

begin
  stub    = Stub.new('localhost:8080', :this_channel_is_insecure)
  request = Request.new(term: term, max_results: 10)

  stub.search(request).each do |response|
    puts "Term: #{response.matched_term}"
    puts "Rank: #{response.rank}"
    puts "Content: #{response.content}"
    puts
  end
rescue StandardError => e
  code = e.respond_to?(:code) ? e.code : 'Unknown'
  puts "Code: #{code}, Type: '#{e.class}', Message: #{e.message}"
end
