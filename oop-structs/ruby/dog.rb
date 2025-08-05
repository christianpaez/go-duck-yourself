class Dog
  BEST_DOG_BREEDS = ['Golden Retriever', 'Border Collie'].freeze
  def initialize(name, breed)
    @breed = breed
    @name = name
  end

  def self.description
    puts "This is a class that describes a Dog."
  end

  def is_a_good_dog
    if best_breed?
      puts "#{@name} is the best dog ever, congrats!"
    else
      puts "#{@name} is not the best dog, sorry."
    end
  end

  private

  attr_reader :breed, :name
  def best_breed?
    BEST_DOG_BREEDS.include?(@breed)
  end
end

Dog.description
noah = Dog.new('Noah', 'Golden Retriever')
nina = Dog.new('Nina', 'Border Collie')
pepe = Dog.new('Pepe', 'Pitbull')

noah.is_a_good_dog
nina.is_a_good_dog
pepe.is_a_good_dog

