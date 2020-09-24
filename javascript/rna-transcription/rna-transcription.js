const NUCLEOTIDE_COMPLEMENTS = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};

export const toRna = (dna) => {
  let rna = "";

  dna.split("").map((c) => (rna += NUCLEOTIDE_COMPLEMENTS[c]));

  return rna;
};
