module Acronym
  def self.abbreviate(terminology)
    splitted = terminology.gsub('-', ' ').split

    acronym = ""

    splitted.each do |word|
      acronym += word[0]
    end

    acronym.upcase
  end
end
