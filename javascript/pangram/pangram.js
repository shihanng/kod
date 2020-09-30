const alphabets = [..."abcdefghijklmnopqrstuvwxyz"];

export const isPangram = (sentence) => {
  const sentenceSet = new Set(sentence.toLowerCase());
  return alphabets.every((c) => sentenceSet.has(c));
};
