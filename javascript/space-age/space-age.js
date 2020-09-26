const EARTH_YEARS = {
  earth: 1.0,
  mercury: 0.2408467,
  venus: 0.61519726,
  mars: 1.8808158,
  jupiter: 11.862615,
  saturn: 29.447498,
  uranus: 84.016846,
  neptune: 164.79132,
};

const PER_EARTH_SECONDS = 1 / 31557600;

export const age = (planet, ageInSeconds) => {
  const planetAge = (ageInSeconds * PER_EARTH_SECONDS) / EARTH_YEARS[planet];
  return +planetAge.toFixed(2);
};
