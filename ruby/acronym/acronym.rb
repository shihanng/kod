module Acronym
  def self.abbreviate(terminology)
    acronym = terminology.scan(/\b[[:alpha:]]/).join()
    acronym.upcase
  end
end
