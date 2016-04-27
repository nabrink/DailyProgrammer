def to_palindromic(number)
  unless number.length <= 1 or is_palindromic(number)
    return to_palindromic((number.to_i + (number.reverse).to_i).to_s)
  end
  
  return number
end

def is_palindromic(number)
  return number == number.reverse
end

number = ARGV[0]
puts(to_palindromic(number))
