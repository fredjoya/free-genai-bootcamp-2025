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

  describe 'GET /study_sessions/:id' do
    it 'returns a specific study session' do
      response = HTTParty.get("#{BASE_URL}/study_sessions/1")
      expect(response.code).to eq(200)
      session = JSON.parse(response.body)
      expect(session).to include(
        'id',
        'study_activity_id',
        'group_id',
        'created_at'
      )
    end
  end

  describe 'GET /study_sessions/:id/words' do
    it 'returns words for a study session' do
      response = HTTParty.get("#{BASE_URL}/study_sessions/1/words")
      expect(response.code).to eq(200)
      words = JSON.parse(response.body)
      expect(words).to be_an(Array)
      expect(words.first).to include(
        'id',
        'arabic',
        'transliteration',
        'english',
        'correct_count',
        'wrong_count',
        'review_correct'
      )
    end
  end

  describe 'POST /reset_history' do
    it 'resets study history' do
      response = HTTParty.post("#{BASE_URL}/reset_history")
      expect(response.code).to eq(200)
      result = JSON.parse(response.body)
      expect(result).to include(
        'success',
        'message'
      )
    end
  end

  describe 'POST /full_reset' do
    it 'resets and reseeds the database' do
      response = HTTParty.post("#{BASE_URL}/full_reset")
      expect(response.code).to eq(200)
      result = JSON.parse(response.body)
      expect(result).to include(
        'success',
        'message'
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