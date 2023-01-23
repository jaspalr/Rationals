class MyRational 
    include Comparable
    attr_accessor :num, :denom
    #Makes a rational function
    def initialize(num,denom)
        raise 'MyRational: denominator cannot be 0' if denom == 0
        @num = num
        @denom = denom
    end
    #Converts a rational a pair
    def pair
        [num,denom].to_a
    end
    #Converts a rational a string
    def to_s
        "#{@num}/#{@denom}"
    end
    #Converts a rational a float
    def to_f
        num.to_f/denom.to_f
    end
    #Checks of a rational is equal to another rational
    def == other
        self.to_f == other.to_f
    end
    #Checks if a rational is less/equal/greater than another rational
    def <=> other
        self.to_f <=> other.to_f
    end
    #Checks of a rational is a integer
    def int?
       (self.to_f) == (num.to_int/denom.to_int)
    end
    #Adds two rationals
    def + other
        newnum = (self.num * other.denom) +  (self.denom * other.num)
        MyRational.new(newnum,(self.denom * other.denom))
    end
    #Multplies two rationals
    def * other
        MyRational.new((self.num * other.num),(self.denom * other.denom))
    end
    #Divides two rationals
    def / other
        MyRational.new((self.num * other.denom),(self.denom * other.num))
    end
    #Inverts a rational
    def invert 
        MyRational.new(denom,num)
    end
    #Converts a rational to lowest terms
    def to_lowest_terms
        ret = MyRational.new(num,denom)
        while ret.num.gcd(ret.denom) > 1 do
            gcd = ret.num.gcd(ret.denom)
            ret = MyRational.new(num/gcd,denom/gcd)
        end
        ret
    end
end
#Returns the harmonic sum of a number
def harmonicSum(n)
    ret = MyRational.new(1,1)
    i = 2
    while i != (n+1) do
        ret += MyRational.new(1,i)
        i +=1
    end
    ret
end