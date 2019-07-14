module Complement
  @dna2rna = { 'G' => 'C', 'C' => 'G', 'T' => 'A', 'A' => 'U' }

  def self.of_dna(dna)
    rna = ''

    dna.each_char do |c|
      rna << @dna2rna[c]
    end

    rna
  end
end
