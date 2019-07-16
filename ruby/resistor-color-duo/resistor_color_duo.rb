# frozen_string_literal: true

COLORS = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9
}.freeze

module ResistorColorDuo
  def self.value(colors)
    colors.reverse.map.each_with_index.sum do |color, i|
      COLORS[color.to_sym] * 10**i
    end
  end
end
