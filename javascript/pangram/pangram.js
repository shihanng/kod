const alphabets = [..."abcdefghijklmnopqrstuvwxyz"];

export const isPangram = (sentence) => {
  const normalizedSentence = sentence.toLowerCase();
  return alphabets.every((c) => normalizedSentence.includes(c));
};
