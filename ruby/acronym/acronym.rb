module Acronym
  def self.abbreviate(terminology)
    acronym = ""

    terminology.scan(/\b[[:alpha:]]/) do |match|
      acronym += match
    end

    acronym.upcase
  end
end
