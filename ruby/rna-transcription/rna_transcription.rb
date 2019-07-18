module Complement
  @DNA2RNA = { 'G' => 'C', 'C' => 'G', 'T' => 'A', 'A' => 'U' }

  def self.of_dna(dna)
    dna.each_char.sum("") do |c|
      @DNA2RNA[c]
    end
  end
end
