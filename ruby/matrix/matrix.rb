class Matrix
  attr_reader :rows

  def initialize(input)
    @rows = []
    a = input.each_line do |l|
      @rows << l.split.map(&:to_i)
    end
    puts @rows
  end
end
