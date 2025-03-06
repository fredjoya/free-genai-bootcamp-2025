require 'spec_helper'

RSpec.describe 'Study Sessions API' do
  describe 'GET /study_sessions' do
    it 'returns all study sessions' do
      response = HTTParty.get("#{BASE_URL}/study_sessions")
      expect(response.code).to eq(200)
      sessions = JSON.parse(response.body)
      expect(sessions).to be_an(Array)
      expect(sessions.first).to include(
        'id',
        'study_activity_id',
        'group_id',
        'created_at'
      )
    end
  end

  describe 'POST /study_sessions/:id/review' do
    it 'submits a word review' do
      response = HTTParty.post(
        "#{BASE_URL}/study_sessions/1/review",
        body: { correct: true }.to_json,
        headers: { 'Content-Type' => 'application/json' }
      )
      expect(response.code).to eq(200)
    end
  end
end 