const NUCLEOTIDE_COMPLEMENTS = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};

export const toRna = (dna) => {
  return dna
    .split("")
    .map((c) => NUCLEOTIDE_COMPLEMENTS[c])
    .join("");
};
