# frozen_string_literal: true

COLORS = {
  black: '0',
  brown: '1',
  red: '2',
  orange: '3',
  yellow: '4',
  green: '5',
  blue: '6',
  violet: '7',
  grey: '8',
  white: '9'
}.freeze

module ResistorColorDuo
  def self.value(colors)
    res = colors.map.reduce('') do |r, color|
      r + COLORS[color.to_sym]
    end
    res.to_i
  end
end
