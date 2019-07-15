COLORS = %w[
  black
  brown
  red
  orange
  yellow
  green
  blue
  violet
  grey
  white
].freeze

module ResistorColorDuo
  def self.value(colors)
    val = 0

    colors.reverse.each_with_index do |color, i|
      val += COLORS.index(color) * 10**i
    end

    val
  end
end
