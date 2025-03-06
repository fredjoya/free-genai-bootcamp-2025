require 'spec_helper'

RSpec.describe 'Groups API' do
  describe 'GET /groups' do
    it 'returns all groups' do
      response = HTTParty.get("#{BASE_URL}/groups")
      expect(response.code).to eq(200)
      groups = JSON.parse(response.body)
      expect(groups).to be_an(Array)
      expect(groups.first).to include(
        'id',
        'name',
        'word_count'
      )
    end
  end

  describe 'GET /groups/:id' do
    it 'returns a specific group' do
      response = HTTParty.get("#{BASE_URL}/groups/1")
      expect(response.code).to eq(200)
      group = JSON.parse(response.body)
      expect(group).to include(
        'id',
        'name',
        'word_count'
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

  describe 'GET /groups/:id/study_sessions' do
    it 'returns study sessions for a group' do
      response = HTTParty.get("#{BASE_URL}/groups/1/study_sessions")
      expect(response.code).to eq(200)
      sessions = JSON.parse(response.body)
      expect(sessions).to be_an(Array)
      expect(sessions.first).to include(
        'id',
        'activity_name',
        'group_name',
        'start_time',
        'end_time',
        'review_items_count'
      )
    end
  end
end 