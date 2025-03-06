require 'spec_helper'

RSpec.describe 'Dashboard API' do
  describe 'GET /dashboard/last_study_session' do
    it 'returns the last study session' do
      response = HTTParty.get("#{BASE_URL}/dashboard/last_study_session")
      expect(response.code).to eq(200)
      expect(JSON.parse(response.body)).to include(
        'id',
        'study_activity_id',
        'group_id',
        'created_at'
      )
    end
  end

  describe 'GET /dashboard/study_progress' do
    it 'returns study progress statistics' do
      response = HTTParty.get("#{BASE_URL}/dashboard/study_progress")
      expect(response.code).to eq(200)
      data = JSON.parse(response.body)
      expect(data).to include(
        'total_words_studied',
        'correct_answers',
        'wrong_answers'
      )
    end
  end

  describe 'GET /dashboard/quick-stats' do
    it 'returns quick statistics' do
      response = HTTParty.get("#{BASE_URL}/dashboard/quick-stats")
      expect(response.code).to eq(200)
      data = JSON.parse(response.body)
      expect(data).to include(
        'total_study_sessions',
        'total_groups',
        'total_words'
      )
    end
  end
end 