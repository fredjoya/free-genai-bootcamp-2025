require 'httparty'
require 'json'
require 'dotenv'

Dotenv.load

RSpec.configure do |config|
  config.before(:suite) do
    # Wait for the API to be ready
    puts "Waiting for API to be ready..."
    max_attempts = 10
    attempts = 0
    
    while attempts < max_attempts
      begin
        response = HTTParty.get('http://localhost:8080/api/health')
        if response.code == 200
          puts "API is ready!"
          break
        end
      rescue
        attempts += 1
        sleep 1
        print "."
      end
    end
    
    if attempts >= max_attempts
      puts "\nWarning: API server not ready after #{max_attempts} attempts"
      puts "Please make sure the Go server is running on port 8080"
    end
  end
end

BASE_URL = 'http://localhost:8080/api'
