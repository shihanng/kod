
module Complement
  DNA2RNA = { 'G' => 'C', 'C' => 'G', 'T' => 'A', 'A' => 'U' }.freeze

  def self.of_dna(dna)
    dna.gsub(/[GCTA]/, DNA2RNA)
  end
end
