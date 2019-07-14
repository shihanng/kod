const COLORS = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white"
];

export const value = colors => {
  let res = 0;

  colors
    .slice()
    .reverse()
    .forEach((v, i) => {
      let id = COLORS.indexOf(v);

      res += id * 10 ** i;
    });

  return res;
};
