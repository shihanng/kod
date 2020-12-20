module FlattenArray
  def self.flatten(nested)
    nested.flatten.compact
  end
end
