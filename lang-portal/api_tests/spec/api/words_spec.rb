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
        'english'
      )
    end
  end
end 