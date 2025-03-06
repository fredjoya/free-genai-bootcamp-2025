require 'spec_helper'

RSpec.describe 'Words API' do
  describe 'GET /words' do
    it 'returns all words' do
      response = HTTParty.get("#{BASE_URL}/words")
      expect(response.code).to eq(200)
      words = JSON.parse(response.body)
      expect(words).to be_an(Array)
      expect(words.first).to include(
        'id',
        'arabic',
        'transliteration',
        'english'
      )
    end
  end

  describe 'GET /words/:id' do
    it 'returns a specific word' do
      response = HTTParty.get("#{BASE_URL}/words/1")
      expect(response.code).to eq(200)
      word = JSON.parse(response.body)
      expect(word).to include(
        'id',
        'arabic',
        'transliteration',
        'english',
        'groups'
      )
    end
  end

  describe 'GET /groups/:id/words' do
    it 'returns words in a group' do
      response = HTTParty.get("#{BASE_URL}/groups/1/words")
      expect(response.code).to eq(200)
      words = JSON.parse(response.body)
      expect(words).to be_an(Array)
      expect(words.first).to include(
        'id',
        'arabic',
        'transliteration',
        'english'
      )
    end
  end

  describe 'POST /study_sessions/:word_id/review' do
    it 'submits a word review' do
      response = HTTParty.post(
        "#{BASE_URL}/study_sessions/1/review",
        body: { correct: true }.to_json,
        headers: { 'Content-Type' => 'application/json' }
      )
      expect(response.code).to eq(200)
      result = JSON.parse(response.body)
      expect(result).to include(
        'success',
        'word_id',
        'study_session_id',
        'correct',
        'created_at'
      )
    end
  end
end 