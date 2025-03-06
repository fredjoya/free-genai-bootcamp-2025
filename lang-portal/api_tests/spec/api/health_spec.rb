require 'spec_helper'

RSpec.describe 'Health Check API' do
  describe 'GET /health' do
    it 'returns ok status' do
      response = HTTParty.get("#{BASE_URL}/health")
      expect(response.code).to eq(200)
      expect(JSON.parse(response.body)).to include('status' => 'ok')
    end
  end
end
