DNA2RNA = { 'G' => 'C', 'C' => 'G', 'T' => 'A', 'A' => 'U' }.freeze

module Complement
  def self.of_dna(dna)
    dna.gsub(/[GCTA]/, DNA2RNA)
  end
end
