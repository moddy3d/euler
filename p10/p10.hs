-- Problem 10: Summation of primes
--
-- https://projecteuler.net/problem=10
--
-- The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
--
-- Find the sum of all the primes below two million.

primesUnder n = sieve 2 0 n [2..(n - 1)]

sieve p i n xs
    | p * p > n = xs
    | p < n = sieve (xs !! i) (i + 1) n (filter (\x -> (x `mod` p /= 0) || (x == p)) xs)
    | otherwise = xs

main = putStrLn $ show $ sum $ primesUnder 2000000
