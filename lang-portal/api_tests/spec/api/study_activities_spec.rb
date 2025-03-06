require 'spec_helper'

RSpec.describe 'Study Activities API' do
  describe 'GET /api/study_activities/:id' do
    it 'returns a specific study activity' do
      response = HTTParty.get("#{BASE_URL}/api/study_activities/1")
      expect(response.code).to eq(200)
      activity = JSON.parse(response.body)
      expect(activity).to include(
        'id',
        'name',
        'thumbnail_url',
        'description'
      )
    end
  end

  describe 'GET /api/study_activities/:id/study_sessions' do
    it 'returns study sessions for an activity' do
      response = HTTParty.get("#{BASE_URL}/api/study_activities/1/study_sessions")
      expect(response.code).to eq(200)
      result = JSON.parse(response.body)
      expect(result).to include('items', 'pagination')
      expect(result['items'].first).to include(
        'id',
        'activity_name',
        'group_name',
        'start_time',
        'end_time',
        'review_items_count'
      )
      expect(result['pagination']).to include(
        'current_page',
        'total_pages',
        'total_items',
        'items_per_page'
      )
    end
  end

  describe 'POST /study_activities' do
    it 'creates a new study activity' do
      response = HTTParty.post(
        "#{BASE_URL}/study_activities",
        body: {
          group_id: 1,
          study_activity_id: 1
        }.to_json,
        headers: { 'Content-Type' => 'application/json' }
      )
      expect(response.code).to eq(201)
      result = JSON.parse(response.body)
      expect(result).to include(
        'id',
        'group_id'
      )
    end
  end
end 